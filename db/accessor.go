package db

import (
	"mtworldview/coords"
)

type Block struct {
	Pos   *coords.MapBlockCoords
	Data  []byte
}

type DBAccessor interface {
	GetBlock(pos *coords.MapBlockCoords) (*Block, error)
}
