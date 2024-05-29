package rpc

import (
	"os"
	"os/exec"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
)

func TestRPCPlugin(t *testing.T) {
	// Build the normal plugin
	os.Mkdir("_test", 0755)
	err := exec.Command("go", "build", "-o", "_test/plugin.out", "../test/normalplugin.go").Run()
	require.NoError(t, err, "The plugin should be built without errors")

	var client *InternalClient

	pool := NewConnectionPool()

	logger := hclog.Default()
	if testing.Verbose() {
		logger.SetLevel(hclog.Debug)
	}

	client, err = pool.NewClient("_test/plugin.out", hclog.Default())
	if err != nil {
		t.Fatal("Could not create a new client", err)
	}

	defer pool.CloseConnection("_test/plugin.out", 0)

	t.Run("Create a connection to the plugin", func(t *testing.T) {
		client, err = pool.NewClient("_test/plugin.out", hclog.Default())
		require.NoError(t, err, "The plugin should be created without errors")
	})

	t.Run("Initialize the plugin", func(t *testing.T) {
		schema, err := client.Plugin.Initialize(0, 0, nil)
		require.NoError(t, err, "The plugin should be initialized without errors")
		require.Equal(t, DatabaseSchema{
			Columns: []DatabaseSchemaColumn{
				{
					Name:        "id",
					Type:        ColumnTypeInt,
					IsParameter: false,
				},
				{
					Name:        "name",
					Type:        ColumnTypeString,
					IsParameter: false,
				},
			},
			PrimaryKey:   -1,
			HandleOffset: false,
		}, schema, "The schema should be correct")
	})

	t.Run("Query the plugin", func(t *testing.T) {
		rows, noMoreRows, err := client.Plugin.Query(0, 0, 0, QueryConstraint{})
		require.NoError(t, err, "The plugin should be queried without errors")
		require.Equal(t, [][]interface{}{
			{1, "hello"},
			{2, "world"},
		}, rows, "The rows should be correct")
		require.True(t, noMoreRows, "The noMoreRows should be true")
	})

}
