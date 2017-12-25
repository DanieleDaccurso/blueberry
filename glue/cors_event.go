package glue

import "github.com/danieledaccurso/blueberry/router"

type CORSEvent struct{}

func (c *CORSEvent) Exec(ctx *router.PreRequestEventContext) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
}
