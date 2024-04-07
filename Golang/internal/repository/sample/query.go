package sample

import "github.com/api-sekejap/internal/repository/base"

const (
	tableDefinition = "channels"

	// Select statements.
	querySelect     = `SELECT id, package_id, name, link, asset_url, description,` + base.GetBaseAttrQuery + `FROM ` + tableDefinition
	querySelectByID = querySelect + ` WHERE id = ?`

	// Insert statements.
	queryInsert = `INSERT INTO ` + tableDefinition +
		`(package_id, name, link, asset_url, description, is_active, created_by, updated_by) ` +
		`VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
)
