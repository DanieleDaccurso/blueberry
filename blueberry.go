package blueberry

import (
	"fmt"
	"net/http"

	"github.com/danieledaccurso/blueberry/glue"
	"github.com/danieledaccurso/blueberry/router"
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
	router    *router.Router
	AppConfig *AppConfig
}

func New() *App {
	a := new(App)
	a.router = router.NewRouter()
	a.AppConfig = new(AppConfig)
	return a
}

func CreateDefault() *App {
	a := New()
	a.EnableCORS()
	a.EnableJSONResponse()
	a.EnableParams()
	return a
}

func (a *App) EnableJSONResponse() *App {
	a.router.AppendPostRequestEvent(new(glue.RenderResponseEvent))
	return a
}

func (a *App) EnableCORS() *App {
	a.router.AppendPreRequestEvent(new(glue.CORSEvent))
	return a
}

func (a *App) EnableParams() *App {
	a.router.AddInjector(new(glue.ParameterInjector))
	return a
}

func (a *App) Router() *router.Router {
	return a.router
}

func (a *App) ListenAndServe() {
	if !a.AppConfig.Server.TLS.Enabled {
		err := http.ListenAndServe(a.AppConfig.Server.ListenHost, a.Router)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
