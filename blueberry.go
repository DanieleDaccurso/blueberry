package blueberry

import (
	"fmt"
	"github.com/danieledaccurso/blueberry/glue"
	"github.com/danieledaccurso/blueberry/router"
	"net/http"
)

type AppConfig struct {
	Server struct {
		TLS struct {
			Enabled  bool
			CertFile string
			KeyFile  string
		}
		ListenHost string
	}
}

type App struct {
	Router    *router.Router
	AppConfig *AppConfig
}

func New() *App {
	a := new(App)
	a.Router = router.NewRouter()
	a.AppConfig = new(AppConfig)
	return a
}

func (a *App) WithJsonResponse() *App {
	a.Router.AppendPostRequestEvent(new(glue.RenderResponseEvent))
	return a
}

func (a *App) WithCORSEnabled() *App {
	a.Router.AppendPreRequestEvent(new(glue.CORSEvent))
	return a
}

func (a *App) ListenAndServe() {
	if !a.AppConfig.Server.TLS.Enabled {
		err := http.ListenAndServe(a.AppConfig.Server.ListenHost, a.Router)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
