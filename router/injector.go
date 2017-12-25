package router

import "net/http"

type InjectorContext struct {
	Request        *http.Request
	Route          *Route
	Router         *Router
	ResponseWriter http.ResponseWriter
	RRequest       *RRequest
}

func createInjectorContext(request *http.Request, route *Route, router *Router, rwriter http.ResponseWriter, rreq *RRequest) *InjectorContext {
	return &InjectorContext{
		Request:        request,
		Route:          route,
		Router:         router,
		ResponseWriter: rwriter,
		RRequest:       rreq,
	}
}

// Injector is the default API to provide values in controller actions. The implementation is executed on every request
// which requires a value to be injected. If the implementation returns true on support, Get will be executed.
type Injector interface {
	Supports(string, string) bool
	Get(*InjectorContext) interface{}
}
