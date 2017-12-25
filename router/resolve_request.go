package router

import (
	"net/http"
	"regexp"
	"strings"
)

var pathRgx *regexp.Regexp

func init() {
	pathRgx = regexp.MustCompile("/+")
}

type RequestResolver struct {
	Router *Router
}

type RRequest struct {
	Parts []string
	Route *Route
}

func NewRequestResolver(router *Router) *RequestResolver {
	rr := new(RequestResolver)
	rr.Router = router
	return rr
}

func (r *RequestResolver) Resolve(req *http.Request) *RRequest {
	rr := r.initRRequest(req)
	rr.Route = nil

	for _, route := range r.Router.routes {
		if len(rr.Parts) != len(route.Segments) {
			continue
		}

		if route.Match(rr) {
			rr.Route = route
			break
		}
	}

	return rr
}

func (r *RequestResolver) initRRequest(req *http.Request) *RRequest {
	cleanPath := pathRgx.ReplaceAllString(req.URL.Path, "/")
	cleanPath = strings.Trim(cleanPath, "/")
	parts := strings.Split(cleanPath, "/")

	rr := new(RRequest)
	rr.Parts = parts
	return rr
}
