package app

import (
	"mtworldview/colormapping"
	"mtworldview/db"
	"mtworldview/mapblockaccessor"
	"mtworldview/params"
)

type App struct {
	Params      params.ParamsType
	Config      *Config
	Worldconfig map[string]string

	Blockdb  db.DBAccessor

	MapBlockAccessor *mapblockaccessor.MapBlockAccessor
	Colormapping     *colormapping.ColorMapping
}
