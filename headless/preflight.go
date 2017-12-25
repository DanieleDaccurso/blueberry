package headless

import "github.com/danieledaccurso/blueberry/router"

type PreflightEvent struct {
}

func (ev *PreflightEvent) Exec(ctx *router.PostMatchEventContext) {

}
