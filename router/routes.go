package router

import (
	"log"
	"reflect"
	"strings"
)

type Segment struct {
	Value    string
	CatchAll bool
}

// Route represents one callable route
type Route struct {
	// Methods contains a slice of strings, identifying the allowed HTTP methods for the current route
	AllowedMethods []string
	// RMethod contains the reflect.Method of the Controller's action to be called on the current route
	RMethod reflect.Value
	// Path contains the path for the current route
	Path string
	// Ordered slice of segments
	Segments []Segment
}

func newRoute(path string, function interface{}) *Route {
	rf := reflect.ValueOf(function)
	if !strings.Contains(rf.Type().String(), "func(") {
		log.Fatal("Trying to add non-function to route match")
	}

	r := new(Route)
	r.RMethod = rf
	r.Path = path
	r.initRoute()
	return r
}

func (r *Route) Match(rr *RRequest) bool {
	if len(rr.Parts) != len(r.Segments) {
		return false
	}

	for index, part := range rr.Parts {
		if r.Segments[index].Value != part &&
			!r.Segments[index].CatchAll {
			return false
		}
	}

	return true
}

func (r *Route) Methods(methods ...string) {
	for _, method := range methods {
		r.AllowedMethods = append(r.AllowedMethods, method)
	}
}

func (r *Route) initRoute() {
	r.Path = pathRgx.ReplaceAllString(r.Path, "/")
	r.Path = strings.Trim(r.Path, "/")
	parts := strings.Split(r.Path, "/")

	for _, part := range parts {
		catchAll := false
		if part != "" && part[0] == ':' {
			catchAll = true
		}
		r.Segments = append(r.Segments, Segment{part, catchAll})
	}
}
