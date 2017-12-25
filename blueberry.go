package blueberry

import (
	"github.com/danieledaccurso/blueberry/router"
	"github.com/danieledaccurso/blueberry/glue"
	"github.com/danieledaccurso/blueberry/headless"
)

type App struct {
	Router *router.Router
}

func New() *App {
	a := new(App)
	a.Router = router.NewRouter()
	return a
}

func (a *App) WithJsonResponse() *App {
	a.Router.AppendPostRequestEvent(new(glue.RenderResponseEvent))
	return a
}

func (a *App) WithCORSEnables() *App {
	a.Router.AppendPreRequestEvent(new(headless.CORSEvent))
	return a
}