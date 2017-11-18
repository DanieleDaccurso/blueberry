package glue

import (
	"github.com/danieledaccurso/blueberry/response"
	"github.com/danieledaccurso/blueberry/wrouter"
)

type RenderResponseEvent struct {
}

func (r *RenderResponseEvent) Exec(ctx *wrouter.PostRequestEventContext) {
	if len(ctx.Values) == 0 || ctx.Values[0].Type().String() != "*response.JsonResponse" {
		return
	}

	ctx.Values[0].Interface().(*response.JsonResponse).Render(ctx.ResponseWriter)
}
