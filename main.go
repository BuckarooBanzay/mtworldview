package main

import (
	"fmt"
	"mtworldview/app"
	"mtworldview/params"
	"mtworldview/web"
	"runtime"

	"github.com/sirupsen/logrus"
)

//go:generate sh -c "go run github.com/mjibson/esc -o vfs/static.go -prefix='static/' -pkg vfs static"

func main() {
	//Parse command line

	p := params.Parse()

	if p.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if p.Help {
		params.PrintHelp()
		return
	}

	if p.Version {
		fmt.Print("Mt world view version: ")
		fmt.Println(app.Version)
		fmt.Print("OS: ")
		fmt.Println(runtime.GOOS)
		fmt.Print("Architecture: ")
		fmt.Println(runtime.GOARCH)
		return
	}

	//parse Config
	cfg, err := app.ParseConfig(app.ConfigFile)
	if err != nil {
		panic(err)
	}

	//write back config with all values
	err = cfg.Save()
	if err != nil {
		panic(err)
	}

	//exit after creating the config
	if p.CreateConfig {
		return
	}

	//setup app context
	ctx := app.Setup(p, cfg)

	//Start http server
	web.Serve(ctx)

}
