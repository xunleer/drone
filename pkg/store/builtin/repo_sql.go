package builtin

// DO NOT EDIT
// code generated by go:generate

import (
	"database/sql"
	"encoding/json"

	. "github.com/drone/drone/pkg/types"
)

var _ = json.Marshal

// generic database interface, matching both *sql.Db and *sql.Tx
type repoDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func getRepo(db repoDB, query string, args ...interface{}) (*Repo, error) {
	row := db.QueryRow(query, args...)
	return scanRepo(row)
}

func getRepos(db repoDB, query string, args ...interface{}) ([]*Repo, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRepos(rows)
}

func createRepo(db repoDB, query string, v *Repo) error {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 bool
	var v10 bool
	var v11 int64
	var v12 string
	var v13 string
	var v14 bool
	var v15 bool
	var v16 bool
	var v17 []byte
	v0 = v.UserID
	v1 = v.Owner
	v2 = v.Name
	v3 = v.FullName
	v4 = v.Avatar
	v5 = v.Self
	v6 = v.Link
	v7 = v.Clone
	v8 = v.Branch
	v9 = v.Private
	v10 = v.Trusted
	v11 = v.Timeout
	if v.Keys != nil {
		v12 = v.Keys.Public
		v13 = v.Keys.Private
	}
	if v.Hooks != nil {
		v14 = v.Hooks.PullRequest
		v15 = v.Hooks.Push
		v16 = v.Hooks.Tags
	}
	v17, _ = json.Marshal(v.Params)

	res, err := db.Exec(query,
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
	)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func updateRepo(db repoDB, query string, v *Repo) error {
	var v0 int64
	var v1 int64
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 bool
	var v11 bool
	var v12 int64
	var v13 string
	var v14 string
	var v15 bool
	var v16 bool
	var v17 bool
	var v18 []byte
	v0 = v.ID
	v1 = v.UserID
	v2 = v.Owner
	v3 = v.Name
	v4 = v.FullName
	v5 = v.Avatar
	v6 = v.Self
	v7 = v.Link
	v8 = v.Clone
	v9 = v.Branch
	v10 = v.Private
	v11 = v.Trusted
	v12 = v.Timeout
	if v.Keys != nil {
		v13 = v.Keys.Public
		v14 = v.Keys.Private
	}
	if v.Hooks != nil {
		v15 = v.Hooks.PullRequest
		v16 = v.Hooks.Push
		v17 = v.Hooks.Tags
	}
	v18, _ = json.Marshal(v.Params)

	_, err := db.Exec(query,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
		&v18,
		&v0,
	)
	return err
}

func scanRepo(row *sql.Row) (*Repo, error) {
	var v0 int64
	var v1 int64
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string
	var v9 string
	var v10 bool
	var v11 bool
	var v12 int64
	var v13 string
	var v14 string
	var v15 bool
	var v16 bool
	var v17 bool
	var v18 []byte

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
		&v16,
		&v17,
		&v18,
	)
	if err != nil {
		return nil, err
	}

	v := &Repo{}
	v.ID = v0
	v.UserID = v1
	v.Owner = v2
	v.Name = v3
	v.FullName = v4
	v.Avatar = v5
	v.Self = v6
	v.Link = v7
	v.Clone = v8
	v.Branch = v9
	v.Private = v10
	v.Trusted = v11
	v.Timeout = v12
	v.Keys = &Keypair{}
	v.Keys.Public = v13
	v.Keys.Private = v14
	v.Hooks = &Hooks{}
	v.Hooks.PullRequest = v15
	v.Hooks.Push = v16
	v.Hooks.Tags = v17
	json.Unmarshal(v18, &v.Params)

	return v, nil
}

func scanRepos(rows *sql.Rows) ([]*Repo, error) {
	var err error
	var vv []*Repo
	for rows.Next() {
		var v0 int64
		var v1 int64
		var v2 string
		var v3 string
		var v4 string
		var v5 string
		var v6 string
		var v7 string
		var v8 string
		var v9 string
		var v10 bool
		var v11 bool
		var v12 int64
		var v13 string
		var v14 string
		var v15 bool
		var v16 bool
		var v17 bool
		var v18 []byte
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
			&v9,
			&v10,
			&v11,
			&v12,
			&v13,
			&v14,
			&v15,
			&v16,
			&v17,
			&v18,
		)
		if err != nil {
			return vv, err
		}

		v := &Repo{}
		v.ID = v0
		v.UserID = v1
		v.Owner = v2
		v.Name = v3
		v.FullName = v4
		v.Avatar = v5
		v.Self = v6
		v.Link = v7
		v.Clone = v8
		v.Branch = v9
		v.Private = v10
		v.Trusted = v11
		v.Timeout = v12
		v.Keys = &Keypair{}
		v.Keys.Public = v13
		v.Keys.Private = v14
		v.Hooks = &Hooks{}
		v.Hooks.PullRequest = v15
		v.Hooks.Push = v16
		v.Hooks.Tags = v17
		json.Unmarshal(v18, &v.Params)
		vv = append(vv, v)
	}
	return vv, rows.Err()
}

const stmtRepoSelectList = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
`

const stmtRepoSelectRange = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
LIMIT ? OFFSET ?
`

const stmtRepoSelect = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
WHERE repo_id = ?
`

const stmtRepoSelectRepoUserId = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
WHERE repo_user_id = ?
`

const stmtRepoSelectRepoOwnerName = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
WHERE repo_owner = ?
AND repo_name = ?
`

const stmtRepoSelectRepoFullName = `
SELECT
 repo_id
,repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
FROM repos
WHERE repo_full_name = ?
`

const stmtRepoSelectCount = `
SELECT count(1)
FROM repos
`

const stmtRepoInsert = `
INSERT INTO repos (
 repo_user_id
,repo_owner
,repo_name
,repo_full_name
,repo_avatar
,repo_self
,repo_link
,repo_clone
,repo_branch
,repo_private
,repo_trusted
,repo_timeout
,repo_keys_public
,repo_keys_private
,repo_hooks_pull_request
,repo_hooks_push
,repo_hooks_tags
,repo_params
) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);
`

const stmtRepoUpdate = `
UPDATE repos SET
 repo_user_id = ?
,repo_owner = ?
,repo_name = ?
,repo_full_name = ?
,repo_avatar = ?
,repo_self = ?
,repo_link = ?
,repo_clone = ?
,repo_branch = ?
,repo_private = ?
,repo_trusted = ?
,repo_timeout = ?
,repo_keys_public = ?
,repo_keys_private = ?
,repo_hooks_pull_request = ?
,repo_hooks_push = ?
,repo_hooks_tags = ?
,repo_params = ?
WHERE repo_id = ?
`

const stmtRepoDelete = `
DELETE FROM repos
WHERE repo_id = ?
`

const stmtRepoTable = `
CREATE TABLE IF NOT EXISTS repos (
 repo_id		INTEGER PRIMARY KEY AUTOINCREMENT
,repo_user_id		INTEGER
,repo_owner		VARCHAR
,repo_name		VARCHAR
,repo_full_name		VARCHAR
,repo_avatar		VARCHAR
,repo_self		VARCHAR
,repo_link		VARCHAR
,repo_clone		VARCHAR
,repo_branch		VARCHAR
,repo_private		BOOLEAN
,repo_trusted		BOOLEAN
,repo_timeout		INTEGER
,repo_keys_public	VARCHAR
,repo_keys_private	VARCHAR
,repo_hooks_pull_request BOOLEAN
,repo_hooks_push	BOOLEAN
,repo_hooks_tags	BOOLEAN
,repo_params		BLOB
);
`

const stmtRepoRepoUserIdIndex = `
CREATE INDEX IF NOT EXISTS ix_repo_user_id ON repos (repo_user_id);
`

const stmtRepoRepoOwnerNameIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS ux_repo_owner_name ON repos (repo_owner,repo_name);
`

const stmtRepoRepoFullNameIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS ux_repo_full_name ON repos (repo_full_name);
`