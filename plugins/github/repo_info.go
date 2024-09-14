package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/go-github/v63/github"
	"github.com/julien040/anyquery/rpc"
)

// A constructor to create a new table instance
// This function is called everytime a new connection is made to the plugin
//
// It should return a new table instance, the database schema and if there is an error
func repoInfoCreator(args rpc.TableCreatorArgs) (rpc.Table, *rpc.DatabaseSchema, error) {
	// Request a connection
	client, token, err := getClient(args)
	if err != nil {
		return nil, nil, err
	}

	// Open the database
	db, err := openDatabase("repo_info", token)
	if err != nil {
		return nil, nil, err
	}

	return &githubRepoByIdTable{
			client: client,
			db:     db,
		}, &rpc.DatabaseSchema{
			Columns: repositorySchema,
		}, nil
}

type githubRepoByIdTable struct {
	client *github.Client
	db     *badger.DB
}

type githubRepoByIdCursor struct {
	client *github.Client
	pageID int
	db     *badger.DB
}

// Return a slice of rows that will be returned to Anyquery and filtered.
// The second return value is true if the cursor has no more rows to return
//
// The constraints are used for optimization purposes to "pre-filter" the rows
// If the rows returned don't match the constraints, it's not an issue. Anyquery will filter them out
func (t *githubRepoByIdCursor) Query(constraints rpc.QueryConstraint) ([][]interface{}, bool, error) {
	repository := retrieveArgString(constraints, 0)
	repoInfo := strings.Split(repository, "/")
	repoName := ""
	repoOwner := ""
	if len(repoInfo) == 2 {
		repoOwner = repoInfo[0]
		repoName = repoInfo[1]
	}

	// Retrieve the repositories
	rows := [][]interface{}{}
	cacheKey := fmt.Sprintf("repositories-%d-%s-%s", t.pageID, repoOwner, repoName)

	// Try to load the cache
	err := loadCache(t.db, cacheKey, &rows)
	if err == nil {
		t.pageID++
		return rows, len(rows) == 0, nil
	}

	// Otherwise, fetch the repositories
	repo, _, err := t.client.Repositories.Get(context.Background(), repoOwner, repoName)

	if err != nil {
		return nil, true, fmt.Errorf("failed to list repository info: %w", err)
	}

	// Serialize the custom properties
	rows = convertRepoInfo(repo, rows)

	// Save the cache
	err = saveCache(t.db, cacheKey, rows)
	if err != nil {
		log.Printf("Failed to save cache: %v", err)
	}

	t.pageID++

	return rows, len(rows) == 0, nil
}

func convertRepoInfo(repo *github.Repository, rows [][]interface{}) [][]interface{} {
	rows = append(rows, []interface{}{
		repo.GetID(),
		repo.GetNodeID(),
		repo.GetOwner().GetLogin(),
		repo.GetName(),
		repo.GetFullName(),
		repo.GetDescription(),
		repo.GetHomepage(),
		repo.GetDefaultBranch(),
		repo.GetCreatedAt().Format(time.RFC3339),
		repo.GetPushedAt().Format(time.RFC3339),
		repo.GetUpdatedAt().Format(time.RFC3339),
		repo.GetHTMLURL(),
		repo.GetCloneURL(),
		repo.GetGitURL(),
		repo.GetMirrorURL(),
		repo.GetSSHURL(),
		repo.GetLanguage(),
		repo.GetFork(),
		repo.GetForksCount(),
		repo.GetNetworkCount(),
		repo.GetOpenIssuesCount(),
		repo.GetStargazersCount(),
		repo.GetSubscribersCount(),
		repo.GetSize(),
		repo.GetAllowRebaseMerge(),
		repo.GetAllowUpdateBranch(),
		repo.GetAllowSquashMerge(),
		repo.GetAllowMergeCommit(),
		repo.GetAllowAutoMerge(),
		repo.GetAllowForking(),
		repo.GetDeleteBranchOnMerge(),
		repo.Topics,
		serializeJSON(repo.GetCustomProperties()),
		repo.GetArchived(),
		repo.GetDisabled(),
		repo.GetVisibility(),
	})
	return rows
}

// Create a new cursor that will be used to read rows
func (t *githubRepoByIdTable) CreateReader() rpc.ReaderInterface {
	return &githubCursor{
		client: t.client,
		pageID: 1,
		db:     t.db,
	}
}

// A slice of rows to insert
func (t *githubRepoByIdTable) Insert(rows [][]interface{}) error {
	return nil
}

// A slice of rows to update
// The first element of each row is the primary key
// while the rest are the values to update
// The primary key is therefore present twice
func (t *githubRepoByIdTable) Update(rows [][]interface{}) error {
	return nil
}

// A slice of primary keys to delete
func (t *githubRepoByIdTable) Delete(primaryKeys []interface{}) error {
	return nil
}

// A destructor to clean up resources
func (t *githubRepoByIdTable) Close() error {
	return nil
}
