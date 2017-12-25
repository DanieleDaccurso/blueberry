package router

import (
	"github.com/danieledaccurso/blueberry/clitable"
	"github.com/danieledaccurso/blueberry/events"
	"io"
	"net/http"
	"reflect"
)

// Router represents one implementation of the http.Handler interface
type Router struct {
	routes    []*Route
	injectors []Injector

	// Eventcollections
	// See: github.com/danieledaccurso/blueberry/events
	preRequest  *events.EventCollection
	postRequest *events.EventCollection
	postMatch   *events.EventCollection

	RequestResolver *RequestResolver
}

type RValues map[string]string

// Create a new Router
func NewRouter() *Router {
	r := new(Router)
	r.preRequest = new(events.EventCollection)
	r.postRequest = new(events.EventCollection)
	r.postMatch = new(events.EventCollection)
	r.RequestResolver = NewRequestResolver(r)
	return r
}

// ServeHTTP satisfies the http.Handler interface, so that this Router can be used as the
// second parameter of http.ListenAndServe
func (r *Router) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	// Execute PreRequestEvents if any
	if r.preRequest.Len() != 0 {
		ctx := createPreRequestEventContext(h, w)
		events.DispatchEvents(r.preRequest, ctx)
	}

	rreq := r.findRequestRoute(h)
	if r.postMatch.Len() != 0 {
		ctx := createPostMatchEventContext(h, w, rreq)
		events.DispatchEvents(r.postMatch, ctx)
	}

	if rreq.Route == nil {
		// @TODO: Implement check if an ErrorController exists
		// Define Specifications for ErrorController (ex: StatusNotFoundAction??)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	r.callRoute(rreq, w, h)
}

// PrintRoutes will print out all routes to io.Writer
func (r *Router) PrintRoutes(writer io.Writer) {
	t := clitable.New()
	t.AddRow("ID", "METHODS", "PATH", "CONTROLLER", "METHOD")
	t.Fprint(writer)
}

func (r *Router) NewRoute(path string, fn interface{}) *Route {
	route := newRoute(path, fn)
	r.routes = append(r.routes, route)
	return route
}

// AddInjector will append an Injector to the list of Injectors in this router.
// It will automatically be the last executed injector as the first injector to support a certain
// type, will be the one who serves the value.
func (r *Router) AddInjector(in Injector) {
	r.injectors = append(r.injectors, in)
}

// AddPreRequestEvent will add a PreRequestEvent with a given priority. It will return an error
// if the selected priority is already taken.
func (r *Router) AddPreRequestEvent(ev PreRequestEvent, priority uint) {
	r.preRequest.AddEvent(ev, priority)
}

// AddPostRequestEvent will add a PostRequestEvent with a given priority. It will return an error
// if the selected priority is already taken.
func (r *Router) AddPostRequestEvent(ev PostRequestEvent, priority uint) {
	r.postRequest.AddEvent(ev, priority)
}

// AppendPreRequestEvent will append a PreRequestEvent at the end of the current PreRequestEvent
// queue. If you want to set a priority, see AddPreRequestEvent
func (r *Router) AppendPreRequestEvent(ev PreRequestEvent) {
	r.preRequest.AppendEvent(ev)
}

// AppendPostRequestEvent will append a PreRequestEvent at the end of the current AddPostRequestEvent
// queue. If you want to set a priority, see AddPostRequestEvent
func (r *Router) AppendPostRequestEvent(ev PostRequestEvent) {
	r.postRequest.AppendEvent(ev)
}

func (r *Router) AppendPostMatchEvent(ev PostMatchEvent) {
	r.postMatch.AppendEvent(ev)
}

func (r *Router) findRequestRoute(h *http.Request) *RRequest {
	return r.RequestResolver.Resolve(h)
}

func (r *Router) callRoute(rreq *RRequest, w http.ResponseWriter, h *http.Request) {
	values := make([]reflect.Value, 0)
	route := rreq.Route

	ctx := createInjectorContext(h, route, r, w, rreq)

	// argument resolving switch is only called, if a method has more than one argument
	if route.RMethod.Type().NumIn() > 0 {
		// Call controller method and inject arguents by reflection
		for i := 0; i < route.RMethod.Type().NumIn(); i++ {
			arg := route.RMethod.Type().In(i)
			switch arg.String() {
			case "http.ResponseWriter":
				values = append(values, reflect.ValueOf(w))
			case "*http.Request":
				values = append(values, reflect.ValueOf(h))
			default:
				values = append(values, reflect.ValueOf(r.inject(arg.Name(), arg.String(), ctx)))
			}
		}
	}

	ret := route.RMethod.Call(values)

	// Execute PostRequest events
	if r.postRequest.Len() != 0 {
		ctx := createPostRequestEventContext(h, w, ret)
		events.DispatchEvents(r.postRequest, ctx)
	}
}

func (r *Router) inject(e string, t string, ctx *InjectorContext) interface{} {
	if len(r.injectors) != 0 {
		for _, injector := range r.injectors {
			if injector.Supports(t, e) {
				return injector.Get(ctx)
			}
		}
	}

	return nil
}
