package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"service/api/pb"
)

type Database interface {
	RegisterClient(client *pb.Client) error
	WriteContent(content *pb.Content, row int64) error
	ReadContent(client *pb.Client, row int64) ([]*pb.Content, error)
}

const url = "postgres://leon.windheuser@localhost:5432/cli2cloud"

type Psql struct {
	conn *pgxpool.Pool
}

// InitPostgres by connection a pool to the cli2cloud database.
// Use a pool for concurrent usage.
func InitPostgres() (*Psql, error) {
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
func (psql *Psql) RegisterClient(client *pb.Client) error {
	queryString := fmt.Sprintf("INSERT INTO users (id, encrypted, created) "+
		"VALUES ('%s', false, now());", client.Id)

	_, err := psql.runQueryNoReturn(queryString)
	if err != nil {
		log.Println("couldn't insert new user", err)
		return err
	}

	return nil
}

// WriteContent Write new content to the database with its respected row
func (psql *Psql) WriteContent(content *pb.Content, row int64) error {
	queryString := fmt.Sprintf("INSERT INTO cli_storage (userid, content, row) "+
		"VALUES ('%s', '%s', %d);", content.Client.Id, content.Payload, row)

	_, err := psql.runQueryNoReturn(queryString)
	if err != nil {
		log.Println("couldn't insert new content.", err)
		return err
	}

	return nil
}

// ReadContent Return all rows of the client which are newer (greater) than the given row.
func (psql *Psql) ReadContent(client *pb.Client, row int64) ([]*pb.Content, error) {
	queryString := fmt.Sprintf("SELECT content, row FROM cli_storage "+
		"WHERE userid = '%s' AND row > %d;", client.Id, row)

	rows, err := psql.conn.Query(context.Background(), queryString)
	if err != nil {
		log.Println("couldn't query contents", err)
		return nil, err
	}

	var contents []*pb.Content

	for rows.Next() {
		var curr pb.Content

		if err := rows.Scan(&curr.Payload, &curr.Row); err != nil {
			log.Println("couldn't marshall row into content", err)
		}

		contents = append(contents, &curr)
	}

	return contents, nil
}

func (psql *Psql) runQueryNoReturn(query string) (int64, error) {
	cmdTag, err := psql.conn.Exec(context.Background(), query)
	if err != nil {
		return -1, err
	}

	return cmdTag.RowsAffected(), nil
}
