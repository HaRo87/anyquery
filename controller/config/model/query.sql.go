// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package model

import (
	"context"
	"database/sql"
)

const addAlias = `-- name: AddAlias :exec
INSERT INTO
    alias (tableName, alias)
VALUES
    (?, ?)
`

type AddAliasParams struct {
	Tablename sql.NullString
	Alias     sql.NullString
}

func (q *Queries) AddAlias(ctx context.Context, arg AddAliasParams) error {
	_, err := q.db.ExecContext(ctx, addAlias, arg.Tablename, arg.Alias)
	return err
}

const addPlugin = `-- name: AddPlugin :exec
INSERT INTO
    plugin_installed (
        id,
        name,
        description,
        path,
        executablePath,
        version,
        homepage,
        registry,
        config,
        checksumDir,
        dev,
        author,
        tablename,
        isSharedExtension
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type AddPluginParams struct {
	ID                sql.NullString
	Name              sql.NullString
	Description       sql.NullString
	Path              sql.NullString
	Executablepath    sql.NullString
	Version           sql.NullString
	Homepage          sql.NullString
	Registry          sql.NullString
	Config            sql.NullString
	Checksumdir       sql.NullString
	Dev               sql.NullInt64
	Author            sql.NullString
	Tablename         sql.NullString
	Issharedextension sql.NullInt64
}

func (q *Queries) AddPlugin(ctx context.Context, arg AddPluginParams) error {
	_, err := q.db.ExecContext(ctx, addPlugin,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Path,
		arg.Executablepath,
		arg.Version,
		arg.Homepage,
		arg.Registry,
		arg.Config,
		arg.Checksumdir,
		arg.Dev,
		arg.Author,
		arg.Tablename,
		arg.Issharedextension,
	)
	return err
}

const addProfile = `-- name: AddProfile :exec
INSERT INTO
    profile (
        name,
        pluginId,
        config
    )
VALUES
    (?, ?, ?)
`

type AddProfileParams struct {
	Name     sql.NullString
	Pluginid sql.NullString
	Config   sql.NullString
}

func (q *Queries) AddProfile(ctx context.Context, arg AddProfileParams) error {
	_, err := q.db.ExecContext(ctx, addProfile, arg.Name, arg.Pluginid, arg.Config)
	return err
}

const addRegistry = `-- name: AddRegistry :exec
INSERT INTO
    registry (
        name,
        url,
        lastUpdated,
        checksumRegistry,
        registryJSON
    )
VALUES
    (?, ?, ?, ?, ?)
`

type AddRegistryParams struct {
	Name             string
	Url              sql.NullString
	Lastupdated      sql.NullInt64
	Checksumregistry sql.NullString
	Registryjson     sql.NullString
}

func (q *Queries) AddRegistry(ctx context.Context, arg AddRegistryParams) error {
	_, err := q.db.ExecContext(ctx, addRegistry,
		arg.Name,
		arg.Url,
		arg.Lastupdated,
		arg.Checksumregistry,
		arg.Registryjson,
	)
	return err
}

const getAlias = `-- name: GetAlias :one
SELECT
    tablename, alias
FROM
    alias
WHERE
    tableName = ?
`

func (q *Queries) GetAlias(ctx context.Context, tablename sql.NullString) (Alias, error) {
	row := q.db.QueryRowContext(ctx, getAlias, tablename)
	var i Alias
	err := row.Scan(&i.Tablename, &i.Alias)
	return i, err
}

const getAliases = `-- name: GetAliases :many
SELECT
    tablename, alias
FROM
    alias
`

func (q *Queries) GetAliases(ctx context.Context) ([]Alias, error) {
	rows, err := q.db.QueryContext(ctx, getAliases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Alias
	for rows.Next() {
		var i Alias
		if err := rows.Scan(&i.Tablename, &i.Alias); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlugin = `-- name: GetPlugin :one
SELECT
    id, name, description, path, executablepath, version, homepage, registry, config, checksumdir, dev, author, tablename, issharedextension
FROM
    plugin_installed
WHERE
    id = ?
`

func (q *Queries) GetPlugin(ctx context.Context, id sql.NullString) (PluginInstalled, error) {
	row := q.db.QueryRowContext(ctx, getPlugin, id)
	var i PluginInstalled
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Path,
		&i.Executablepath,
		&i.Version,
		&i.Homepage,
		&i.Registry,
		&i.Config,
		&i.Checksumdir,
		&i.Dev,
		&i.Author,
		&i.Tablename,
		&i.Issharedextension,
	)
	return i, err
}

const getPlugins = `-- name: GetPlugins :many
SELECT
    id, name, description, path, executablepath, version, homepage, registry, config, checksumdir, dev, author, tablename, issharedextension
FROM
    plugin_installed
`

func (q *Queries) GetPlugins(ctx context.Context) ([]PluginInstalled, error) {
	rows, err := q.db.QueryContext(ctx, getPlugins)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PluginInstalled
	for rows.Next() {
		var i PluginInstalled
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Path,
			&i.Executablepath,
			&i.Version,
			&i.Homepage,
			&i.Registry,
			&i.Config,
			&i.Checksumdir,
			&i.Dev,
			&i.Author,
			&i.Tablename,
			&i.Issharedextension,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfile = `-- name: GetProfile :one
SELECT
    name, pluginid, config
FROM
    profile
WHERE
    name = ?
    AND pluginId = ?
`

type GetProfileParams struct {
	Name     sql.NullString
	Pluginid sql.NullString
}

func (q *Queries) GetProfile(ctx context.Context, arg GetProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, arg.Name, arg.Pluginid)
	var i Profile
	err := row.Scan(&i.Name, &i.Pluginid, &i.Config)
	return i, err
}

const getProfiles = `-- name: GetProfiles :many
SELECT
    name, pluginid, config
FROM
    profile
`

func (q *Queries) GetProfiles(ctx context.Context) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(&i.Name, &i.Pluginid, &i.Config); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfilesOfPlugin = `-- name: GetProfilesOfPlugin :many
SELECT
    name, pluginid, config
FROM
    profile
WHERE
    pluginId = ?
`

func (q *Queries) GetProfilesOfPlugin(ctx context.Context, pluginid sql.NullString) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, getProfilesOfPlugin, pluginid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(&i.Name, &i.Pluginid, &i.Config); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRegistries = `-- name: GetRegistries :many
SELECT
    name, url, lastupdated, checksumregistry, registryjson
FROM
    registry
`

func (q *Queries) GetRegistries(ctx context.Context) ([]Registry, error) {
	rows, err := q.db.QueryContext(ctx, getRegistries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Registry
	for rows.Next() {
		var i Registry
		if err := rows.Scan(
			&i.Name,
			&i.Url,
			&i.Lastupdated,
			&i.Checksumregistry,
			&i.Registryjson,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRegistry = `-- name: GetRegistry :one
SELECT
    name, url, lastupdated, checksumregistry, registryjson
FROM
    registry
WHERE
    name = ?
`

func (q *Queries) GetRegistry(ctx context.Context, name string) (Registry, error) {
	row := q.db.QueryRowContext(ctx, getRegistry, name)
	var i Registry
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.Lastupdated,
		&i.Checksumregistry,
		&i.Registryjson,
	)
	return i, err
}
