package main

import (
	"encoding/json"
	"fmt"

	"github.com/julien040/anyquery/rpc"
)

// A constructor to create a new table instance
// This function is called everytime a new connection is made to the plugin
//
// It should return a new table instance, the database schema and if there is an error
func albumTableCreator(args rpc.TableCreatorArgs) (rpc.Table, *rpc.DatabaseSchema, error) {
	refreshToken, ok := args.UserConfig["token"]
	if !ok {
		return nil, nil, fmt.Errorf("token is missing")
	}
	clientID, ok := args.UserConfig["client_id"]
	if !ok {
		return nil, nil, fmt.Errorf("client_id is missing")
	}
	clientSecret, ok := args.UserConfig["client_secret"]
	if !ok {
		return nil, nil, fmt.Errorf("client_secret is missing")
	}

	accessToken, err := getAccessToken(refreshToken.(string), clientID.(string), clientSecret.(string))
	if err != nil {
		return nil, nil, err
	}

	return &albumTable{
			accessToken: accessToken,
		}, &rpc.DatabaseSchema{
			PrimaryKey: 0,
			Columns: []rpc.DatabaseSchemaColumn{
				{
					Name:        "id",
					Type:        rpc.ColumnTypeString,
					IsParameter: true,
					IsRequired:  true,
				},
				{
					Name: "album_type",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "total_tracks_album",
					Type: rpc.ColumnTypeInt,
				},
				{
					Name: "href",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "album_name",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "release_date",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "artist_name",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "copyright",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "album_popularity",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_name",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_duration_ms",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_disc_number",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_explicit",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_href",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_artists",
					Type: rpc.ColumnTypeString,
				},
				{
					Name: "track_number",
					Type: rpc.ColumnTypeString,
				},
			},
		}, nil
}

type albumTable struct {
	accessToken string
}

type albumCursor struct {
	accessToken string
}

// Return a slice of rows that will be returned to Anyquery and filtered.
// The second return value is true if the cursor has no more rows to return
//
// The constraints are used for optimization purposes to "pre-filter" the rows
// If the rows returned don't match the constraints, it's not an issue. Anyquery will filter them out
func (t *albumCursor) Query(constraints rpc.QueryConstraint) ([][]interface{}, bool, error) {
	albumID := ""
	var err error
	var ok bool
	for _, constraint := range constraints.Columns {
		if constraint.ColumnID == 0 {
			if albumID, ok = constraint.Value.(string); !ok {
				return nil, true, fmt.Errorf("id is a string, got %T", constraint.Value)
			}
			break
		}
	}

	if albumID == "" {
		return nil, true, fmt.Errorf("id is empty")
	}

	endpoint := "https://api.spotify.com/v1/albums/" + albumID + "?limit=50"

	var data albumAPI
	res, err := restyClient.R().SetHeader("Authorization", "Bearer "+t.accessToken).SetResult(&data).Get(endpoint)
	if err != nil {
		return nil, true, err
	}

	if res.StatusCode() != 200 {
		return nil, true, fmt.Errorf("failed to get album: %s", res.String())
	}

	copyright := ""
	for _, c := range data.Copyright {
		copyright = copyright + c.Text + " "
	}

	albumArtists := []string{}
	for _, artist := range data.Artists {
		albumArtists = append(albumArtists, artist.Name)
	}

	albumArtistsJSON, err := json.Marshal(albumArtists)
	if err != nil {
		return nil, true, err
	}

	var rows [][]interface{}
	for _, track := range data.Tracks.Items {
		artists := []string{}
		for _, artist := range track.Artists {
			artists = append(artists, artist.Name)
		}

		jsonArtists, err := json.Marshal(artists)
		if err != nil {
			return nil, true, err
		}

		rows = append(rows, []interface{}{
			data.ID,
			data.AlbumType,
			data.TotalTracks,
			data.Href,
			data.Name,
			data.ReleaseDate,
			string(albumArtistsJSON),
			copyright,
			data.Popularity,
			track.Name,
			track.DurationMs,
			track.DiscNumber,
			track.Explicit,
			track.Href,
			string(jsonArtists),
			track.TrackNumber,
		})
	}

	return rows, true, nil
}

// Create a new cursor that will be used to read rows
func (t *albumTable) CreateReader() rpc.ReaderInterface {
	return &albumCursor{
		accessToken: t.accessToken,
	}
}

// A slice of rows to insert
func (t *albumTable) Insert(rows [][]interface{}) error {
	return nil
}

// A slice of rows to update
// The first element of each row is the primary key
// while the rest are the values to update
// The primary key is therefore present twice
func (t *albumTable) Update(rows [][]interface{}) error {
	return nil
}

// A slice of primary keys to delete
func (t *albumTable) Delete(primaryKeys []interface{}) error {
	return nil
}

// A destructor to clean up resources
func (t *albumTable) Close() error {
	return nil
}
