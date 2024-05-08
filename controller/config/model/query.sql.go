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
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type AddPluginParams struct {
	Name              string
	Description       sql.NullString
	Path              string
	Executablepath    string
	Version           string
	Homepage          sql.NullString
	Registry          string
	Config            string
	Checksumdir       sql.NullString
	Dev               int64
	Author            sql.NullString
	Tablename         string
	Issharedextension int64
}

func (q *Queries) AddPlugin(ctx context.Context, arg AddPluginParams) error {
	_, err := q.db.ExecContext(ctx, addPlugin,
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
        pluginName,
        registry,
        config
    )
VALUES
    (?, ?, ?, ?)
`

type AddProfileParams struct {
	Name       string
	Pluginname string
	Registry   string
	Config     string
}

func (q *Queries) AddProfile(ctx context.Context, arg AddProfileParams) error {
	_, err := q.db.ExecContext(ctx, addProfile,
		arg.Name,
		arg.Pluginname,
		arg.Registry,
		arg.Config,
	)
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
	Url              string
	Lastupdated      int64
	Checksumregistry string
	Registryjson     string
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
    name, description, path, executablepath, version, homepage, registry, config, checksumdir, dev, author, tablename, issharedextension
FROM
    plugin_installed
WHERE
    name = ?
    AND registry = ?
`

type GetPluginParams struct {
	Name     string
	Registry string
}

func (q *Queries) GetPlugin(ctx context.Context, arg GetPluginParams) (PluginInstalled, error) {
	row := q.db.QueryRowContext(ctx, getPlugin, arg.Name, arg.Registry)
	var i PluginInstalled
	err := row.Scan(
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
    name, description, path, executablepath, version, homepage, registry, config, checksumdir, dev, author, tablename, issharedextension
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
    name, pluginname, registry, config
FROM
    profile
WHERE
    name = ?
    AND pluginName = ?
    AND registry = ?
`

type GetProfileParams struct {
	Name       string
	Pluginname string
	Registry   string
}

func (q *Queries) GetProfile(ctx context.Context, arg GetProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, arg.Name, arg.Pluginname, arg.Registry)
	var i Profile
	err := row.Scan(
		&i.Name,
		&i.Pluginname,
		&i.Registry,
		&i.Config,
	)
	return i, err
}

const getProfiles = `-- name: GetProfiles :many
SELECT
    name, pluginname, registry, config
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
		if err := rows.Scan(
			&i.Name,
			&i.Pluginname,
			&i.Registry,
			&i.Config,
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

const getProfilesOfPlugin = `-- name: GetProfilesOfPlugin :many
SELECT
    name, pluginname, registry, config
FROM
    profile
WHERE
    pluginName = ?
    AND registry = ?
`

type GetProfilesOfPluginParams struct {
	Pluginname string
	Registry   string
}

func (q *Queries) GetProfilesOfPlugin(ctx context.Context, arg GetProfilesOfPluginParams) ([]Profile, error) {
	rows, err := q.db.QueryContext(ctx, getProfilesOfPlugin, arg.Pluginname, arg.Registry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.Name,
			&i.Pluginname,
			&i.Registry,
			&i.Config,
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

const updateRegistry = `-- name: UpdateRegistry :exec
UPDATE
    registry
SET
    url = ?,
    lastUpdated = ?,
    checksumRegistry = ?,
    registryJSON = ?
WHERE
    name = ?
`

type UpdateRegistryParams struct {
	Url              string
	Lastupdated      int64
	Checksumregistry string
	Registryjson     string
	Name             string
}

// --------------------------------------------------------------------------
//
//	Updates
//
// --------------------------------------------------------------------------
func (q *Queries) UpdateRegistry(ctx context.Context, arg UpdateRegistryParams) error {
	_, err := q.db.ExecContext(ctx, updateRegistry,
		arg.Url,
		arg.Lastupdated,
		arg.Checksumregistry,
		arg.Registryjson,
		arg.Name,
	)
	return err
}
