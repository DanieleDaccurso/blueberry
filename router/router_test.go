package router

import (
	"testing"
	"net/http"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	router.NewRoute("/user/:userid/info", func() {})
	res := router.RequestResolver

	req1, _ := http.NewRequest("GET", "/user/12/info", nil)
	if rr := res.Resolve(req1); rr.Route == nil {
		t.Error("Failed to match req1")
	}

	req2, _ := http.NewRequest("POST", "/user/12/info", nil)
	if rr := res.Resolve(req2); rr.Route == nil {
		t.Error("Failed to match req2")
	}

	req3, _ := http.NewRequest("POST", "/user/12/info/", nil)
	if rr := res.Resolve(req3); rr.Route == nil {
		t.Error("Failed to match req3")
	}

	req4, _ := http.NewRequest("POST", "/user//12/info/", nil)
	if rr := res.Resolve(req4); rr.Route == nil {
		t.Error("Failed to match req4")
	}

	req5, _ := http.NewRequest("POST", "/auth/12/info/", nil)
	if rr := res.Resolve(req5); rr.Route != nil {
		t.Error("Failed to mismatch req5")
	}

	req6, _ := http.NewRequest("POST", "/user/12/info/2", nil)
	if rr := res.Resolve(req6); rr.Route != nil {
		t.Error("Failed to mismatch req6")
	}
}
