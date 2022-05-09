// Package vcssql provide SQL queries for vcs package.
// It must not be used outside of vcs, but it can't be internal package
// because otherwise code generation with gorun wouldn't work.
package vcssql

import (
	"io/ioutil"
	sgen "mintter/backend/db/sqlitegen"
	"mintter/backend/db/sqlitegen/qb"
	s "mintter/backend/db/sqliteschema"
)

var _ = generateQueries

//go:generate gorun generateQueries
func generateQueries() error {
	code, err := sgen.CodegenQueries("vcssql",
		qb.MakeQuery(s.Schema, "WorkingCopyReplace", sgen.QueryKindExec,
			"INSERT OR REPLACE INTO", s.WorkingCopy, qb.ListColShort(
				s.WorkingCopyObjectID,
				s.WorkingCopyName,
				s.WorkingCopyVersion,
				s.WorkingCopyData,
				s.WorkingCopyCreateTime,
				s.WorkingCopyUpdateTime,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.WorkingCopyObjectID),
				qb.VarCol(s.WorkingCopyName),
				qb.VarCol(s.WorkingCopyVersion),
				qb.VarCol(s.WorkingCopyData),
				qb.VarCol(s.WorkingCopyCreateTime),
				qb.VarCol(s.WorkingCopyUpdateTime),
			),
		),

		qb.MakeQuery(s.Schema, "WorkingCopyGet", sgen.QueryKindSingle,
			"SELECT", qb.Results(
				qb.ResultCol(s.WorkingCopyData),
				qb.ResultCol(s.WorkingCopyCreateTime),
				qb.ResultCol(s.WorkingCopyUpdateTime),
				qb.ResultCol(s.WorkingCopyVersion),
			), qb.Line,
			"FROM", s.WorkingCopy, qb.Line,
			"WHERE", s.WorkingCopyObjectID, "=", qb.VarCol(s.WorkingCopyObjectID), qb.Line,
			"AND", s.WorkingCopyName, "=", qb.VarCol(s.WorkingCopyName), qb.Line,
			"LIMIT 1",
		),

		qb.MakeQuery(s.Schema, "WorkingCopyDelete", sgen.QueryKindExec,
			"DELETE FROM", s.WorkingCopy, qb.Line,
			"WHERE", s.WorkingCopyObjectID, "=", qb.VarCol(s.WorkingCopyObjectID), qb.Line,
			"AND", s.WorkingCopyName, "=", qb.VarCol(s.WorkingCopyName),
		),

		qb.MakeQuery(s.Schema, "AccountsLookupPK", sgen.QueryKindSingle,
			"SELECT", qb.Results(
				qb.ResultCol(s.AccountsID),
			), qb.Line,
			"FROM", s.Accounts, qb.Line,
			"WHERE", s.AccountsMultihash, "=", qb.VarCol(s.AccountsMultihash),
			"AND", s.AccountsCodec, "=", qb.VarCol(s.AccountsCodec),
		),

		qb.MakeQuery(s.Schema, "ObjectsLookupPK", sgen.QueryKindSingle,
			"SELECT", qb.Results(
				qb.ResultCol(s.ObjectsID),
			), qb.Line,
			"FROM", s.Objects, qb.Line,
			"WHERE", s.ObjectsMultihash, "=", qb.VarCol(s.ObjectsMultihash),
			"AND", s.ObjectsCodec, "=", qb.VarCol(s.ObjectsCodec),
		),

		qb.MakeQuery(s.Schema, "DevicesLookupPK", sgen.QueryKindSingle,
			"SELECT", qb.Results(
				qb.ResultCol(s.DevicesID),
			), qb.Line,
			"FROM", s.Devices, qb.Line,
			"WHERE", s.DevicesMultihash, "=", qb.VarCol(s.DevicesMultihash),
			"AND", s.DevicesCodec, "=", qb.VarCol(s.DevicesCodec),
		),

		qb.MakeQuery(s.Schema, "AccountsInsertPK", sgen.QueryKindSingle,
			"INSERT INTO", s.Accounts, qb.ListColShort(
				s.AccountsMultihash,
				s.AccountsCodec,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.AccountsMultihash),
				qb.VarCol(s.AccountsCodec),
			),
			"RETURNING", qb.Results(qb.ResultCol(s.AccountsID)),
		),

		qb.MakeQuery(s.Schema, "ObjectsInsertPK", sgen.QueryKindSingle,
			"INSERT INTO", s.Objects, qb.ListColShort(
				s.ObjectsMultihash,
				s.ObjectsCodec,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.ObjectsMultihash),
				qb.VarCol(s.ObjectsCodec),
			),
			"RETURNING", qb.Results(qb.ResultCol(s.ObjectsID)),
		),

		qb.MakeQuery(s.Schema, "DevicesInsertPK", sgen.QueryKindSingle,
			"INSERT INTO", s.Devices, qb.ListColShort(
				s.DevicesMultihash,
				s.DevicesCodec,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.DevicesMultihash),
				qb.VarCol(s.DevicesCodec),
			),
			"RETURNING", qb.Results(qb.ResultCol(s.DevicesID)),
		),

		qb.MakeQuery(s.Schema, "ObjectsInsertOrIgnore", sgen.QueryKindExec,
			"INSERT OR IGNORE INTO", s.Objects, qb.ListColShort(
				s.ObjectsMultihash,
				s.ObjectsCodec,
				s.ObjectsAccountID,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.ObjectsMultihash),
				qb.VarCol(s.ObjectsCodec),
				qb.VarCol(s.ObjectsAccountID),
			),
		),

		qb.MakeQuery(s.Schema, "NamedVersionsReplace", sgen.QueryKindExec,
			"INSERT OR REPLACE INTO", s.NamedVersions, qb.ListColShort(
				s.NamedVersionsObjectID,
				s.NamedVersionsAccountID,
				s.NamedVersionsDeviceID,
				s.NamedVersionsName,
				s.NamedVersionsVersion,
			), qb.Line,
			"VALUES", qb.List(
				qb.VarCol(s.NamedVersionsObjectID),
				qb.VarCol(s.NamedVersionsAccountID),
				qb.VarCol(s.NamedVersionsDeviceID),
				qb.VarCol(s.NamedVersionsName),
				qb.VarCol(s.NamedVersionsVersion),
			),
		),

		qb.MakeQuery(s.Schema, "NamedVersionsGet", sgen.QueryKindSingle,
			"SELECT", qb.Results(
				qb.ResultCol(s.NamedVersionsVersion),
			), qb.Line,
			"FROM", s.NamedVersions, qb.Line,
			"WHERE", s.NamedVersionsObjectID, "=", qb.VarCol(s.NamedVersionsObjectID), qb.Line,
			"AND", s.NamedVersionsAccountID, "=", qb.VarCol(s.NamedVersionsAccountID), qb.Line,
			"AND", s.NamedVersionsDeviceID, "=", qb.VarCol(s.NamedVersionsDeviceID), qb.Line,
			"AND", s.NamedVersionsName, "=", qb.VarCol(s.NamedVersionsName), qb.Line,
			"LIMIT 1",
		),

		qb.MakeQuery(s.Schema, "DraftsInsert", sgen.QueryKindExec,
			"INSERT INTO", s.Drafts, qb.ListColShort(
				s.DraftsID,
				s.DraftsTitle,
				s.DraftsSubtitle,
				s.DraftsCreateTime,
				s.DraftsUpdateTime,
			), qb.Line,
			"VALUES", qb.List(
				qb.LookupSubQuery(s.ObjectsID, s.Objects,
					"WHERE", s.ObjectsMultihash, "=", qb.VarCol(s.ObjectsMultihash),
					"AND", s.ObjectsCodec, "=", qb.VarCol(s.ObjectsCodec),
				),
				qb.VarCol(s.DraftsTitle),
				qb.VarCol(s.DraftsSubtitle),
				qb.VarCol(s.DraftsCreateTime),
				qb.VarCol(s.DraftsUpdateTime),
			),
		),

		qb.MakeQuery(s.Schema, "DraftsUpdate", sgen.QueryKindExec,
			"UPDATE", s.Drafts, qb.Line,
			"SET", qb.ListColShort(
				s.DraftsTitle,
				s.DraftsSubtitle,
				s.DraftsUpdateTime,
			), "=", qb.List(
				qb.VarCol(s.DraftsTitle),
				qb.VarCol(s.DraftsSubtitle),
				qb.VarCol(s.DraftsUpdateTime),
			), qb.Line,
			"WHERE", s.DraftsID, "=", qb.LookupSubQuery(s.ObjectsID, s.Objects,
				"WHERE", s.ObjectsMultihash, "=", qb.VarCol(s.ObjectsMultihash),
				"AND", s.ObjectsCodec, "=", qb.VarCol(s.ObjectsCodec),
			),
		),

		qb.MakeQuery(s.Schema, "DraftsList", sgen.QueryKindMany,
			"SELECT", qb.Results(
				qb.ResultCol(s.ObjectsMultihash),
				qb.ResultCol(s.ObjectsCodec),
				qb.ResultCol(s.DraftsTitle),
				qb.ResultCol(s.DraftsSubtitle),
				qb.ResultCol(s.DraftsCreateTime),
				qb.ResultCol(s.DraftsUpdateTime),
			), qb.Line,
			"FROM", s.Drafts, qb.Line,
			"JOIN", s.Objects, "ON", s.ObjectsID, "=", s.DraftsID, qb.Line,
		),
	)

	if err != nil {
		return err
	}

	return ioutil.WriteFile("sqlite_queries.gen.go", code, 0600)
}