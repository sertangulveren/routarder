package routarder

import (
	"net/http"
	"testing"
)

var noneHandler = func(_ http.ResponseWriter, _ *http.Request) {return}

func TestPrefix(t *testing.T) {
	r := New()
	if r.prefix != "/" {
		t.Errorf("expected %s, got %s", "/", r.prefix)
	}
}

func TestGroupPrefix(t *testing.T) {
	r := New()                  // prefix = /
	api := r.Group("/api")      // prefix = /api
	v1 := api.Group("/v1")      // prefix = /api/v1
	posts := v1.Group("/posts") // prefix = /api/v1/posts

	if posts.prefix != "/api/v1/posts" {
		t.Errorf("expected %s, got %s", "/api/v1/posts", posts.prefix)
	}
}

func TestRouteList(t *testing.T) {
	r := New() // prefix = /
	r.HandleFunc("/hi", noneHandler)
	r.HandleFunc("/dude", noneHandler)

	for _, route := range r.routes {
		if route != "/hi" && route != "/dude" {
			t.Fatalf("route must be equal to %s or %s", "/hi", "/dude")
		}
	}
}

func TestChildrenSize(t *testing.T) {
	r := New()

	apiGroup := r.Group("/api")
	apiGroup.HandleFunc("/posts", noneHandler)

	if len(r.children) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(r.children))
	}
}

func TestChildrenPrefix(t *testing.T) {
	r := New()

	_ = r.Group("/api")

	if r.children[0].prefix != "/api" {
		t.Errorf("expected %s, got %s", "/api", r.children[0].prefix)
	}
}

func TestSelfRoutes(t *testing.T) {
}

func TestSelfRoutesValue(t *testing.T) {
}

