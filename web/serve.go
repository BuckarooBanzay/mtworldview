package web

import (
	"mtworldview/app"
	"mtworldview/vfs"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func Serve(ctx *app.App) {
	fields := logrus.Fields{
		"port":   ctx.Config.Port,
		"webdev": ctx.Config.Webdev,
	}
	logrus.WithFields(fields).Info("Starting http server")

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(vfs.FS(ctx.Config.Webdev)))

	mux.Handle("/api/config", &ConfigHandler{ctx: ctx})
	mux.Handle("/api/colormapping", &ColorMappingHandler{ctx: ctx})
	mux.Handle("/api/viewblock/", &ViewMapblockHandler{ctx: ctx})

	if ctx.Config.EnablePrometheus {
		mux.Handle("/metrics", promhttp.Handler())
	}

	ws := NewWS(ctx)
	mux.Handle("/api/ws", ws)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI

		if len(uri) >= 3 {
			suffix := uri[len(uri)-3:]

			switch suffix {
			case "css":
				w.Header().Set("Content-Type", "text/css")
			case ".js":
				w.Header().Set("Content-Type", "application/javascript")
			case "png":
				w.Header().Set("Content-Type", "image/png")
			}
		}
		mux.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(ctx.Config.Port), nil)
	if err != nil {
		panic(err)
	}
}
