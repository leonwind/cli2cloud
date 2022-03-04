package storage

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Row struct {
	Content string
	Line    int64
}

type Database interface {
	RegisterClient(clientId string, encrypted bool, salt *string, iv *string) error
	GetClientById(clientId string) (bool, *string, *string, error)
	WriteContent(clientId string, row Row) error
	ReadContent(clientId string, line int64) ([]Row, error)
}

type Psql struct {
	conn *pgxpool.Pool
}

// InitPostgres by connecting a pool to the cli2cloud database.
// Use a pool for concurrent usage.
func InitPostgres(url string) (*Psql, error) {
	poolConn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	psql := Psql{
		conn: poolConn,
	}
	return &psql, nil
}

// RegisterClient Register a new Client in the database and store the ID, if encrypted, and timestamp.
// Encryption is yet not implemented
func (psql *Psql) RegisterClient(clientId string, encrypted bool, salt *string, iv *string) error {
	queryString := "INSERT INTO clients (id, encrypted, salt, iv, created) " +
		"VALUES ($1, $2, $3, $4, now());"

	_, err := psql.runQueryNoReturn(queryString, clientId, encrypted, ptrToNullString(salt), ptrToNullString(iv))
	if err != nil {
		return err
	}

	return nil
}

// GetClientById fetches if the client, based on its Id, encrypted the data, and if so then returns
// the salt and the IV required for the decryption.
func (psql *Psql) GetClientById(clientId string) (bool, *string, *string, error) {
	queryString := "SELECT encrypted, salt, iv FROM clients " +
		"WHERE id = $1;"

	queriedRow, err := psql.conn.Query(context.Background(), queryString, clientId)
	if err != nil {
		return false, nil, nil, err
	}

	var encrypted bool
	var salt sql.NullString
	var iv sql.NullString

	for queriedRow.Next() {
		if err := queriedRow.Scan(&encrypted, &salt, &iv); err != nil {
			return false, nil, nil, err
		}
	}
	return encrypted, nullStringToPtr(salt), nullStringToPtr(iv), nil
}

// WriteContent Write new content to the database with its respected row
func (psql *Psql) WriteContent(clientId string, row Row) error {
	queryString := "INSERT INTO cli_storage (clientId, content, line) " +
		"VALUES ($1, $2, $3);"

	_, err := psql.runQueryNoReturn(queryString, clientId, row.Content, row.Line)

	if err != nil {
		return err
	}

	return nil
}

// ReadContent Return all rows of the client which are newer (greater) than the given row.
func (psql *Psql) ReadContent(clientId string, line int64) ([]Row, error) {
	queryString := "SELECT content, line FROM cli_storage " +
		"WHERE clientId = $1 AND line >= $2 " +
		"ORDER BY line " +
		"LIMIT 100;"

	queriedRows, err := psql.conn.Query(context.Background(), queryString, clientId, line)
	if err != nil {
		return nil, err
	}

	var rows []Row

	for queriedRows.Next() {
		var curr Row

		if err := queriedRows.Scan(&curr.Content, &curr.Line); err != nil {
			return nil, err
		}

		rows = append(rows, curr)
	}

	return rows, nil
}

// runQueryNoReturns executes a query with its arguments and returns the number of lines it affected
func (psql *Psql) runQueryNoReturn(query string, arguments ...interface{}) (int64, error) {
	cmdTag, err := psql.conn.Exec(context.Background(), query, arguments...)
	if err != nil {
		return -1, err
	}

	return cmdTag.RowsAffected(), nil
}

// ptrToNullString checks if the given string is nil or empty and returns a NullString.
// Pgx then stores the string as NULL in postgres.
func ptrToNullString(s *string) sql.NullString {
	if s == nil || len(*s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

// Returns nil if the NullString is null, else the pointer to the string value
func nullStringToPtr(s sql.NullString) *string {
	if s.Valid {
		return &(s.String)
	}
	return nil
}
