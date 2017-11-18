package blueberry

import (
	"github.com/danieledaccurso/blueberry/glue"
	"github.com/danieledaccurso/blueberry/wrouter"
)

func DefaultRouter() *wrouter.Router {
	r := wrouter.NewRouter()
	r.AppendPostRequestEvent(new(glue.RenderResponseEvent))
	return r
}
