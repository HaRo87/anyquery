// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package model

import (
	"database/sql"
)

type Alias struct {
	Tablename sql.NullString
	Alias     sql.NullString
}

type PluginInstalled struct {
	ID             sql.NullString
	Name           sql.NullString
	Description    sql.NullString
	Path           sql.NullString
	Executablepath sql.NullString
	Version        sql.NullString
	Homepage       sql.NullString
	Registry       sql.NullString
	Config         sql.NullString
	Checksumdir    sql.NullString
	Dev            sql.NullInt64
}

type Profile struct {
	Name     sql.NullString
	Pluginid sql.NullString
	Config   sql.NullString
}

type Registry struct {
	Name             string
	Url              sql.NullString
	Lastupdated      sql.NullInt64
	Checksumregistry sql.NullString
	Registryjson     sql.NullString
}