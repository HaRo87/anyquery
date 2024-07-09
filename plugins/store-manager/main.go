package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
	"github.com/pelletier/go-toml"
	flag "github.com/spf13/pflag"
	"golang.org/x/term"
)

const registryURL = "https://registry.anyquery.dev"

type UserConfig struct {
	Name        string `toml:"name" json:"name"`
	Description string `toml:"description" json:"description"`
	Required    bool   `toml:"required" json:"required"`
	Type        string `toml:"type" json:"type"`
}

type File struct {
	Platform       string `toml:"platform" json:"platform"`
	Directory      string `toml:"directory" json:"directory"`
	ExecutablePath string `toml:"executablePath" json:"executablePath"`
}

type Plugin struct {
	Name                   string `toml:"name" json:"name"`
	Version                string `toml:"version" json:"version"`
	Description            string `toml:"description" json:"description"`
	Author                 string `toml:"author" json:"author"`
	License                string `toml:"license" json:"license"`
	Repository             string `toml:"repository" json:"repository"`
	Homepage               string `toml:"homepage" json:"homepage"`
	Type                   string `toml:"type" json:"type"`
	MinimumAnyqueryVersion string `toml:"minimumAnyqueryVersion" json:"minimumAnyqueryVersion"`

	Tables []string `toml:"tables" json:"tables"`

	UserConfig []UserConfig `toml:"userConfig" json:"userConfig"`

	File []File `toml:"file" json:"file"`

	// Should only be populated by the server
	Versions    []string `json:"versions"`
	ID          string   `json:"id"`
	PageContent string   `json:"pageContent"`
}

func main() {
	var user, password string
	flag.StringVarP(&user, "user", "u", "", "User")

	// Parse the flags
	var configurationFile string
	flag.StringVarP(&configurationFile, "config", "c", "", "Configuration file")

	var packageName string
	flag.StringVarP(&packageName, "package", "p", "", "Package name")

	flag.Parse()

	if user == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Username: ")
		text, err := reader.ReadString('\n')
		user = text
		if err != io.EOF && err != nil {
			panic(err)
		}
	}

	password = os.Getenv("ANYQUERY_PASSWORD")

	// Request the password
	if password == "" {
		fmt.Print("Password: ")
		rawPass, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		password = string(rawPass)
	}

	urlAuth, err := url.Parse(registryURL + "/api/admins/auth-with-password?fields=*")
	if err != nil {
		panic(err)
	}

	rawBody := map[string]string{
		"identity": user,
		"password": password,
	}
	marshalled, err := json.Marshal(rawBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", urlAuth.String(), bytes.NewBuffer(marshalled))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	token, err := script.Do(req).JQ(".token").Replace("\"", "").String()
	if err != nil {
		panic(err)
	}
	token = strings.Trim(token, "\" \n")

	if configurationFile == "" || packageName == "" {
		fmt.Println("Configuration file and package name are required")
		flag.Usage()
		return
	}

	// Load the configuration
	plugin := Plugin{}
	rawContent, err := os.ReadFile(configurationFile)
	if err != nil {
		panic(err)
	}
	toml.Unmarshal(rawContent, &plugin)

	// Set the current directory relative to the configuration file
	// This is required to load the files

	err = os.Chdir(filepath.Dir(configurationFile))
	if err != nil {
		panic(fmt.Errorf("error changing directory: %w", err))
	}
	ids := []string{}

	fmt.Println("Uploading files")

	for _, file := range plugin.File {
		fileName := fmt.Sprintf("%s_%s_%s.zip", plugin.Name, file.Platform, plugin.Version)
		id, err := uploadFile(file.Platform, file.Directory, file.ExecutablePath, fileName, token)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
		fmt.Println("Uploaded file with id: ", id)
	}

	versionId, err := uploadVersion(plugin, ids, token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploaded version %s (%s)\n", plugin.Version, versionId)

	queryParams := url.Values{}
	queryParams.Add("filter", fmt.Sprintf("(name='%s')", plugin.Name))
	queryParams.Add("perPage", "1")
	queryParams.Add("expand", "versions")

	// Get the plugin
	req, err = http.NewRequest("GET", registryURL+"/api/collections/plugin/records"+"?"+queryParams.Encode(), nil)
	if err != nil {
		panic(err)
	}
	req.Header["Authorization"] = []string{"Bearer " + token}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	contentReq, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(contentReq))
		panic(fmt.Errorf("status code: %d", resp.StatusCode))
	}

	// Get the length of items
	length, err := script.Echo(string(contentReq)).JQ(".totalItems").String()
	if err != nil {
		panic(err)
	}
	length = strings.Trim(length, "\" \n")
	if length != "1" {
		fmt.Println(string(contentReq), length)
		panic(fmt.Errorf("Plugin %s not found", plugin.Name))
	}

	itemValue, err := script.Echo(string(contentReq)).JQ(".items[0]").String()
	if err != nil {
		panic(err)
	}

	remotePlugin := Plugin{}
	err = json.Unmarshal([]byte(itemValue), &remotePlugin)
	if err != nil {
		panic(err)
	}

	// Add the version to the plugin
	remotePlugin.Versions = append(remotePlugin.Versions, versionId)

	// Read the README.md and set the content to pageContent
	readmeContent, err := os.ReadFile("README.md")
	if err != nil {
		// Leave the pageContent as is
	} else {
		remotePlugin.PageContent = string(readmeContent)
	}

	// Set the other fields
	remotePlugin.Author = plugin.Author
	remotePlugin.Description = plugin.Description
	remotePlugin.Homepage = plugin.Homepage
	remotePlugin.License = plugin.License
	remotePlugin.Repository = plugin.Repository
	remotePlugin.Type = plugin.Type

	// Update the plugin
	urlUpdate := registryURL + "/api/collections/plugin/records/" + remotePlugin.ID
	marshalled, err = json.Marshal(remotePlugin)
	if err != nil {
		panic(err)
	}

	req, err = http.NewRequest("PATCH", urlUpdate, bytes.NewBuffer(marshalled))
	if err != nil {
		panic(err)
	}
	req.Header["Authorization"] = []string{"Bearer " + token}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	contentReq, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(contentReq))
		panic(fmt.Errorf("status code: %d", resp.StatusCode))
	}

	fmt.Printf("Updated plugin %s\n", plugin.Name)

}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func uploadVersion(plugin Plugin, ids []string, token string) (string, error) {
	fmt.Println("Uploading version")
	postUrl := registryURL + "/api/collections/pluginVersion/records"

	// Compute the version ID. It'a 15 character long string
	versionID := strings.Builder{}
	versionID.WriteString(strings.ReplaceAll(plugin.Version, ".", ""))
	// Add the first character of the package name until the length is 15 substraction the length of the version
	i := 0
	currentLen := len(versionID.String())
	for i < 15-currentLen {
		if i < len(plugin.Name) {
			// If the character is non-alphanumeric, skip it
			if !('a' <= plugin.Name[i] && plugin.Name[i] <= 'z') {
				continue
			}
			versionID.WriteByte(plugin.Name[i])
		} else {
			// Add a random character if the package name is shorter than 15
			versionID.WriteByte(alphabet[rand.IntN(len(alphabet))])
		}
		i++
	}

	if len(versionID.String()) != 15 {
		panic("Version ID is not 15 characters long " + versionID.String() + " " + string(len(versionID.String())))
	}

	rawBody := map[string]interface{}{
		"id":              versionID.String(),
		"version":         plugin.Version,
		"minimum_version": plugin.MinimumAnyqueryVersion,
		"files":           ids,
		"user_config":     plugin.UserConfig,
		"tables":          plugin.Tables,
	}

	marshalled, err := json.Marshal(rawBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(marshalled))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header["Authorization"] = []string{"Bearer " + token}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// Parse the json response and get the id
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(response)
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return response["id"].(string), nil
}

func uploadFile(platform string, filePath string, executablePath string, pluginZipName string, token string) (string, error) {
	postUrl := registryURL + "/api/collections/pluginFile/records"

	// Zip the filepath
	fs := os.DirFS(filePath)

	body := bytes.Buffer{}
	multiPartWriter := multipart.NewWriter(&body)
	fileWriter, err := multiPartWriter.CreateFormFile("file", pluginZipName)
	if err != nil {
		return "", err
	}

	fileWriterBuffer := &bytes.Buffer{}

	zipWriter := zip.NewWriter(fileWriterBuffer)
	err = zipWriter.AddFS(fs)
	if err != nil {
		return "", err
	}
	zipWriter.Close()

	_, err = io.Copy(fileWriter, fileWriterBuffer)
	if err != nil {
		return "", err
	}

	// Add the platform
	err = multiPartWriter.WriteField("platform", platform)
	if err != nil {
		return "", err
	}

	// Add the executable path
	err = multiPartWriter.WriteField("path", executablePath)
	if err != nil {
		return "", err
	}

	// Compute the checksum of the file
	checksum := sha256.Sum256(fileWriterBuffer.Bytes())
	err = multiPartWriter.WriteField("hash", fmt.Sprintf("%x", checksum))
	if err != nil {
		return "", err
	}

	err = multiPartWriter.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", postUrl, &body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	req.Header["Authorization"] = []string{"Bearer " + token}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// Parse the json response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error uploading file: %+v\n", response)
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return response["id"].(string), nil

}
