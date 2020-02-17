package postgres

import (
	"database/sql"
	"mtworldview/coords"
	"mtworldview/db"

	_ "github.com/lib/pq"
)

type PostgresAccessor struct {
	db *sql.DB
}

func convertRows(posx, posy, posz int, data []byte) *db.Block {
	c := coords.NewMapBlockCoords(posx, posy, posz)
	return &db.Block{Pos: c, Data: data}
}

func (this *PostgresAccessor) GetBlock(pos *coords.MapBlockCoords) (*db.Block, error) {
	rows, err := this.db.Query(getBlockQuery, pos.X, pos.Y, pos.Z)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var posx, posy, posz int
		var data []byte

		err = rows.Scan(&posx, &posy, &posz, &data)
		if err != nil {
			return nil, err
		}

		mb := convertRows(posx, posy, posz, data)
		return mb, nil
	}

	return nil, nil
}

func New(connStr string) (*PostgresAccessor, error) {
	db, err := sql.Open("postgres", connStr+" sslmode=disable")
	if err != nil {
		return nil, err
	}

	sq := &PostgresAccessor{db: db}
	return sq, nil
}
