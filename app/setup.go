package app

import (
	"mtworldview/colormapping"
	"mtworldview/db/postgres"
	"mtworldview/db/sqlite"
	"mtworldview/mapblockaccessor"
	"mtworldview/params"
	"mtworldview/worldconfig"
	"time"

	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"errors"
)

func Setup(p params.ParamsType, cfg *Config) *App {
	a := App{}
	a.Params = p
	a.Config = cfg

	//Parse world config
	a.Worldconfig = worldconfig.Parse("world.mt")
	logrus.WithFields(logrus.Fields{"version": Version}).Info("Starting mtworldview")

	var err error

	switch a.Worldconfig[worldconfig.CONFIG_BACKEND] {
	case worldconfig.BACKEND_SQLITE3:
		a.Blockdb, err = sqlite.New("map.sqlite")
		if err != nil {
			panic(err)
		}

	case worldconfig.BACKEND_POSTGRES:
		a.Blockdb, err = postgres.New(a.Worldconfig[worldconfig.CONFIG_PSQL_CONNECTION])
		if err != nil {
			panic(err)
		}

	default:
		panic(errors.New("map-backend not supported: " + a.Worldconfig[worldconfig.CONFIG_BACKEND]))
	}

	//mapblock accessor
	expireDuration, err := time.ParseDuration(cfg.MapBlockAccessorCfg.Expiretime)
	if err != nil {
		panic(err)
	}

	purgeDuration, err := time.ParseDuration(cfg.MapBlockAccessorCfg.Purgetime)
	if err != nil {
		panic(err)
	}

	// mapblock accessor
	a.MapBlockAccessor = mapblockaccessor.NewMapBlockAccessor(
		a.Blockdb,
		expireDuration, purgeDuration,
		cfg.MapBlockAccessorCfg.MaxItems)

	//color mapping
	a.Colormapping = colormapping.NewColorMapping()

	colorfiles := []string{
		//https://daconcepts.com/vanessa/hobbies/minetest/colors.txt
		"/colors/vanessa.txt",
		"/colors/advtrains.txt",
		"/colors/scifi_nodes.txt",
		"/colors/mcl2_colors.txt",
		"/colors/miles.txt",
		"/colors/custom.txt",
	}

	for _, colorfile := range colorfiles {
		_, err := a.Colormapping.LoadVFSColors(false, colorfile)
		if err != nil {
			panic(err.Error() + " file:" + colorfile)
		}
	}

	//load provided colors, if available
	info, err := os.Stat("colors.txt")
	if info != nil && err == nil {
		logrus.WithFields(logrus.Fields{"filename": "colors.txt"}).Info("Loading colors from filesystem")

		data, err := ioutil.ReadFile("colors.txt")
		if err != nil {
			panic(err)
		}

		count, err := a.Colormapping.LoadBytes(data)
		if err != nil {
			panic(err)
		}

		logrus.WithFields(logrus.Fields{"count": count}).Info("Loaded custom colors")

	}
	return &a
}
