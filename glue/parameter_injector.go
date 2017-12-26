package glue

import (
	"strings"

	"github.com/danieledaccurso/blueberry/router"
)

type Params map[string]string

func (p Params) Get(t string) string {
	if val, ok := p[t]; ok {
		return val
	}
	return ""
}

type ParameterInjector struct{}

func (p *ParameterInjector) Supports(ctx *router.InjectorContext) bool {
	return ctx.DataType == "glue.Params"
}

func (p *ParameterInjector) Get(ctx *router.InjectorContext) interface{} {
	params := make(Params)

	for index, segment := range ctx.Route.Segments {
		if segment.CatchAll {
			params[strings.Replace(segment.Value, ":", "", 1)] = ctx.RRequest.Parts[index]
		}
	}

	return params
}
