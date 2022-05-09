// Code generated by sqlitegen. DO NOT EDIT.

package vcssql

import (
	"errors"
	"fmt"

	"crawshaw.io/sqlite"
	"mintter/backend/db/sqlitegen"
)

var _ = errors.New

func WorkingCopyReplace(conn *sqlite.Conn, workingCopyObjectID int, workingCopyName string, workingCopyVersion string, workingCopyData []byte, workingCopyCreateTime int, workingCopyUpdateTime int) error {
	const query = `INSERT OR REPLACE INTO working_copy (object_id, name, version, data, create_time, update_time)
VALUES (:workingCopyObjectID, :workingCopyName, :workingCopyVersion, :workingCopyData, :workingCopyCreateTime, :workingCopyUpdateTime)`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetInt(":workingCopyObjectID", workingCopyObjectID)
		stmt.SetText(":workingCopyName", workingCopyName)
		stmt.SetText(":workingCopyVersion", workingCopyVersion)
		stmt.SetBytes(":workingCopyData", workingCopyData)
		stmt.SetInt(":workingCopyCreateTime", workingCopyCreateTime)
		stmt.SetInt(":workingCopyUpdateTime", workingCopyUpdateTime)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: WorkingCopyReplace: %w", err)
	}

	return err
}

type WorkingCopyGetResult struct {
	WorkingCopyData       []byte
	WorkingCopyCreateTime int
	WorkingCopyUpdateTime int
	WorkingCopyVersion    string
}

func WorkingCopyGet(conn *sqlite.Conn, workingCopyObjectID int, workingCopyName string) (WorkingCopyGetResult, error) {
	const query = `SELECT working_copy.data, working_copy.create_time, working_copy.update_time, working_copy.version
FROM working_copy
WHERE working_copy.object_id = :workingCopyObjectID
AND working_copy.name = :workingCopyName
LIMIT 1`

	var out WorkingCopyGetResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetInt(":workingCopyObjectID", workingCopyObjectID)
		stmt.SetText(":workingCopyName", workingCopyName)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("WorkingCopyGet: more than one result return for a single-kind query")
		}

		out.WorkingCopyData = stmt.ColumnBytes(0)
		out.WorkingCopyCreateTime = stmt.ColumnInt(1)
		out.WorkingCopyUpdateTime = stmt.ColumnInt(2)
		out.WorkingCopyVersion = stmt.ColumnText(3)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: WorkingCopyGet: %w", err)
	}

	return out, err
}

func WorkingCopyDelete(conn *sqlite.Conn, workingCopyObjectID int, workingCopyName string) error {
	const query = `DELETE FROM working_copy
WHERE working_copy.object_id = :workingCopyObjectID
AND working_copy.name = :workingCopyName`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetInt(":workingCopyObjectID", workingCopyObjectID)
		stmt.SetText(":workingCopyName", workingCopyName)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: WorkingCopyDelete: %w", err)
	}

	return err
}

type AccountsLookupPKResult struct {
	AccountsID int
}

func AccountsLookupPK(conn *sqlite.Conn, accountsMultihash []byte, accountsCodec int) (AccountsLookupPKResult, error) {
	const query = `SELECT accounts.id
FROM accounts
WHERE accounts.multihash = :accountsMultihash AND accounts.codec = :accountsCodec`

	var out AccountsLookupPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":accountsMultihash", accountsMultihash)
		stmt.SetInt(":accountsCodec", accountsCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("AccountsLookupPK: more than one result return for a single-kind query")
		}

		out.AccountsID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: AccountsLookupPK: %w", err)
	}

	return out, err
}

type ObjectsLookupPKResult struct {
	ObjectsID int
}

func ObjectsLookupPK(conn *sqlite.Conn, objectsMultihash []byte, objectsCodec int) (ObjectsLookupPKResult, error) {
	const query = `SELECT objects.id
FROM objects
WHERE objects.multihash = :objectsMultihash AND objects.codec = :objectsCodec`

	var out ObjectsLookupPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":objectsMultihash", objectsMultihash)
		stmt.SetInt(":objectsCodec", objectsCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("ObjectsLookupPK: more than one result return for a single-kind query")
		}

		out.ObjectsID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: ObjectsLookupPK: %w", err)
	}

	return out, err
}

type DevicesLookupPKResult struct {
	DevicesID int
}

func DevicesLookupPK(conn *sqlite.Conn, devicesMultihash []byte, devicesCodec int) (DevicesLookupPKResult, error) {
	const query = `SELECT devices.id
FROM devices
WHERE devices.multihash = :devicesMultihash AND devices.codec = :devicesCodec`

	var out DevicesLookupPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":devicesMultihash", devicesMultihash)
		stmt.SetInt(":devicesCodec", devicesCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("DevicesLookupPK: more than one result return for a single-kind query")
		}

		out.DevicesID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: DevicesLookupPK: %w", err)
	}

	return out, err
}

type AccountsInsertPKResult struct {
	AccountsID int
}

func AccountsInsertPK(conn *sqlite.Conn, accountsMultihash []byte, accountsCodec int) (AccountsInsertPKResult, error) {
	const query = `INSERT INTO accounts (multihash, codec)
VALUES (:accountsMultihash, :accountsCodec) RETURNING accounts.id`

	var out AccountsInsertPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":accountsMultihash", accountsMultihash)
		stmt.SetInt(":accountsCodec", accountsCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("AccountsInsertPK: more than one result return for a single-kind query")
		}

		out.AccountsID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: AccountsInsertPK: %w", err)
	}

	return out, err
}

type ObjectsInsertPKResult struct {
	ObjectsID int
}

func ObjectsInsertPK(conn *sqlite.Conn, objectsMultihash []byte, objectsCodec int) (ObjectsInsertPKResult, error) {
	const query = `INSERT INTO objects (multihash, codec)
VALUES (:objectsMultihash, :objectsCodec) RETURNING objects.id`

	var out ObjectsInsertPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":objectsMultihash", objectsMultihash)
		stmt.SetInt(":objectsCodec", objectsCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("ObjectsInsertPK: more than one result return for a single-kind query")
		}

		out.ObjectsID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: ObjectsInsertPK: %w", err)
	}

	return out, err
}

type DevicesInsertPKResult struct {
	DevicesID int
}

func DevicesInsertPK(conn *sqlite.Conn, devicesMultihash []byte, devicesCodec int) (DevicesInsertPKResult, error) {
	const query = `INSERT INTO devices (multihash, codec)
VALUES (:devicesMultihash, :devicesCodec) RETURNING devices.id`

	var out DevicesInsertPKResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":devicesMultihash", devicesMultihash)
		stmt.SetInt(":devicesCodec", devicesCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("DevicesInsertPK: more than one result return for a single-kind query")
		}

		out.DevicesID = stmt.ColumnInt(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: DevicesInsertPK: %w", err)
	}

	return out, err
}

func ObjectsInsertOrIgnore(conn *sqlite.Conn, objectsMultihash []byte, objectsCodec int, objectsAccountID int) error {
	const query = `INSERT OR IGNORE INTO objects (multihash, codec, account_id)
VALUES (:objectsMultihash, :objectsCodec, :objectsAccountID)`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":objectsMultihash", objectsMultihash)
		stmt.SetInt(":objectsCodec", objectsCodec)
		stmt.SetInt(":objectsAccountID", objectsAccountID)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: ObjectsInsertOrIgnore: %w", err)
	}

	return err
}

func NamedVersionsReplace(conn *sqlite.Conn, namedVersionsObjectID int, namedVersionsAccountID int, namedVersionsDeviceID int, namedVersionsName string, namedVersionsVersion string) error {
	const query = `INSERT OR REPLACE INTO named_versions (object_id, account_id, device_id, name, version)
VALUES (:namedVersionsObjectID, :namedVersionsAccountID, :namedVersionsDeviceID, :namedVersionsName, :namedVersionsVersion)`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetInt(":namedVersionsObjectID", namedVersionsObjectID)
		stmt.SetInt(":namedVersionsAccountID", namedVersionsAccountID)
		stmt.SetInt(":namedVersionsDeviceID", namedVersionsDeviceID)
		stmt.SetText(":namedVersionsName", namedVersionsName)
		stmt.SetText(":namedVersionsVersion", namedVersionsVersion)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: NamedVersionsReplace: %w", err)
	}

	return err
}

type NamedVersionsGetResult struct {
	NamedVersionsVersion string
}

func NamedVersionsGet(conn *sqlite.Conn, namedVersionsObjectID int, namedVersionsAccountID int, namedVersionsDeviceID int, namedVersionsName string) (NamedVersionsGetResult, error) {
	const query = `SELECT named_versions.version
FROM named_versions
WHERE named_versions.object_id = :namedVersionsObjectID
AND named_versions.account_id = :namedVersionsAccountID
AND named_versions.device_id = :namedVersionsDeviceID
AND named_versions.name = :namedVersionsName
LIMIT 1`

	var out NamedVersionsGetResult

	before := func(stmt *sqlite.Stmt) {
		stmt.SetInt(":namedVersionsObjectID", namedVersionsObjectID)
		stmt.SetInt(":namedVersionsAccountID", namedVersionsAccountID)
		stmt.SetInt(":namedVersionsDeviceID", namedVersionsDeviceID)
		stmt.SetText(":namedVersionsName", namedVersionsName)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		if i > 1 {
			return errors.New("NamedVersionsGet: more than one result return for a single-kind query")
		}

		out.NamedVersionsVersion = stmt.ColumnText(0)
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: NamedVersionsGet: %w", err)
	}

	return out, err
}

func DraftsInsert(conn *sqlite.Conn, objectsMultihash []byte, objectsCodec int, draftsTitle string, draftsSubtitle string, draftsCreateTime int, draftsUpdateTime int) error {
	const query = `INSERT INTO drafts (id, title, subtitle, create_time, update_time)
VALUES (COALESCE((SELECT objects.id FROM objects WHERE objects.multihash = :objectsMultihash AND objects.codec = :objectsCodec LIMIT 1), -1000), :draftsTitle, :draftsSubtitle, :draftsCreateTime, :draftsUpdateTime)`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetBytes(":objectsMultihash", objectsMultihash)
		stmt.SetInt(":objectsCodec", objectsCodec)
		stmt.SetText(":draftsTitle", draftsTitle)
		stmt.SetText(":draftsSubtitle", draftsSubtitle)
		stmt.SetInt(":draftsCreateTime", draftsCreateTime)
		stmt.SetInt(":draftsUpdateTime", draftsUpdateTime)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: DraftsInsert: %w", err)
	}

	return err
}

func DraftsUpdate(conn *sqlite.Conn, draftsTitle string, draftsSubtitle string, draftsUpdateTime int, objectsMultihash []byte, objectsCodec int) error {
	const query = `UPDATE drafts
SET (title, subtitle, update_time) = (:draftsTitle, :draftsSubtitle, :draftsUpdateTime)
WHERE drafts.id = COALESCE((SELECT objects.id FROM objects WHERE objects.multihash = :objectsMultihash AND objects.codec = :objectsCodec LIMIT 1), -1000)`

	before := func(stmt *sqlite.Stmt) {
		stmt.SetText(":draftsTitle", draftsTitle)
		stmt.SetText(":draftsSubtitle", draftsSubtitle)
		stmt.SetInt(":draftsUpdateTime", draftsUpdateTime)
		stmt.SetBytes(":objectsMultihash", objectsMultihash)
		stmt.SetInt(":objectsCodec", objectsCodec)
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: DraftsUpdate: %w", err)
	}

	return err
}

type DraftsListResult struct {
	ObjectsMultihash []byte
	ObjectsCodec     int
	DraftsTitle      string
	DraftsSubtitle   string
	DraftsCreateTime int
	DraftsUpdateTime int
}

func DraftsList(conn *sqlite.Conn) ([]DraftsListResult, error) {
	const query = `SELECT objects.multihash, objects.codec, drafts.title, drafts.subtitle, drafts.create_time, drafts.update_time
FROM drafts
JOIN objects ON objects.id = drafts.id
`

	var out []DraftsListResult

	before := func(stmt *sqlite.Stmt) {
	}

	onStep := func(i int, stmt *sqlite.Stmt) error {
		out = append(out, DraftsListResult{
			ObjectsMultihash: stmt.ColumnBytes(0),
			ObjectsCodec:     stmt.ColumnInt(1),
			DraftsTitle:      stmt.ColumnText(2),
			DraftsSubtitle:   stmt.ColumnText(3),
			DraftsCreateTime: stmt.ColumnInt(4),
			DraftsUpdateTime: stmt.ColumnInt(5),
		})

		return nil
	}

	err := sqlitegen.ExecStmt(conn, query, before, onStep)
	if err != nil {
		err = fmt.Errorf("failed query: DraftsList: %w", err)
	}

	return out, err
}