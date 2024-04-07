// Package plugin provides interfaces and methods to write a plugin
// that communicates with the main program
//
// We recommend that you read the official docs rather than this
// autogenerated documentation unless for the API reference.
package plugin

// InternalExchangeInterface is an interface that defines the methods
// that a plugin must implement to communicate with the main program
//
// This part should be handled by the plugin library and should not be
// implemented by the user
type InternalExchangeInterface interface {
	// Initialize is a method that is called when the plugin is initialized
	//
	// It is called once when the plugin is loaded and is used by the main
	// program to infer the schema of the tables
	Initialize(tableIndex int, config PluginConfig) (DatabaseSchema, error)

	// Query is a method that returns rows for a given SELECT query
	//
	// Constraints are passed as arguments for optimization purposes
	// However, the plugin is free to ignore them because
	// the main program will filter the results to match the constraints
	//
	// The return value is a 2D slice of interface{} where each row is a slice
	// and each element in the row is an interface{} representing the value.
	// The order and type of the values should match the schema of the table
	Query(tableIndex int, cursorIndex int, constraint QueryConstraint) ([][]interface{}, error)
}

// PluginConfig is a struct that holds the configuration for the plugin
//
// It is mostly used to specify user-defined configuration
// and is passed to the plugin during initialization
type PluginConfig map[string]string

// PluginManifest is a struct that holds the metadata of the plugin
//
// It is often represented as a JSON file in the plugin directory
type PluginManifest struct {
	Name        string
	Version     string
	Author      string
	Description string

	UserConfig []PluginConfigField
}

type PluginConfigField struct {
	Name     string
	Required bool
}

// DatabaseSchema holds the schema of the database
//
// It must stay the same throughout the lifetime of the plugin
// and for every cursor opened.
//
// One and only field must be the primary key. If you don't have a primary key,
// you can generate a unique key. The primary key must be unique for each row.
// It is used to update and delete rows.
type DatabaseSchema struct {
	// ... other fields
	HandleLimit  bool
	HandleOffset bool
}

// ColumnType is an enum that represents the type of a column
type ColumnType int8

const (
	// ColumnTypeInt represents an INTEGER column
	ColumnTypeInt ColumnType = iota
	// ColumnTypeFloat represents a REAL column
	ColumnTypeFloat
	// ColumnTypeString represents a TEXT column
	ColumnTypeString
	// ColumnTypeBlob represents a BLOB column
	ColumnTypeBlob
)

type DatabaseSchemaColumn struct {
	Name string
	Type ColumnType
}

// QueryConstraint is a struct that holds the constraints for a SELECT query
//
// It specifies the WHERE conditions in the Columns field,
// the LIMIT and OFFSET in the Limit and Offset fields,
// and the ORDER BY clause in the OrderBy field
type QueryConstraint struct {
	Columns []ColumnConstraint
	Limit   int
	Offset  int
	OrderBy string
}

type ColumnConstraint struct {
	ColumnID int
	Operator string
	Value    interface{}
}
