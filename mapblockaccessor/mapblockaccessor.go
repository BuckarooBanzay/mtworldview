package mapblockaccessor

import (
	"fmt"
	"mtworldview/coords"
	"mtworldview/db"

	"time"

	cache "github.com/patrickmn/go-cache"
)

type MapBlockAccessor struct {
	accessor   db.DBAccessor
	blockcache *cache.Cache
	maxcount   int
}

func getKey(pos *coords.MapBlockCoords) string {
	return fmt.Sprintf("Coord %d/%d/%d", pos.X, pos.Y, pos.Z)
}

func NewMapBlockAccessor(accessor db.DBAccessor, expiretime, purgetime time.Duration, maxcount int) *MapBlockAccessor {
	blockcache := cache.New(expiretime, purgetime)

	return &MapBlockAccessor{
		accessor:   accessor,
		blockcache: blockcache,
		maxcount:   maxcount,
	}
}
