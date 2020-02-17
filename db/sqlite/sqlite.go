package sqlite

import (
	"database/sql"
	"mtworldview/coords"
	"mtworldview/db"

	_ "github.com/mattn/go-sqlite3"
)


type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func convertRows(pos int64, data []byte) *db.Block {
	c := coords.PlainToCoord(pos)
	return &db.Block{Pos: c, Data: data}
}

func (db *Sqlite3Accessor) GetBlock(pos *coords.MapBlockCoords) (*db.Block, error) {
	ppos := coords.CoordToPlain(pos)

	rows, err := db.db.Query(getBlockQuery, ppos)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var pos int64
		var data []byte

		err = rows.Scan(&pos, &data)
		if err != nil {
			return nil, err
		}

		mb := convertRows(pos, data)
		return mb, nil
	}

	return nil, nil
}

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename+"?mode=ro&_timeout=2000")
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
