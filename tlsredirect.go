package tlsredirect

import (
	"net/http"

	"github.com/mholt/caddy/caddy/setup"
	"github.com/mholt/caddy/middleware"
)

type handler struct {
	next middleware.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Header.Get("X-Forwarded-Proto") == "http" {
		http.Redirect(w, r, "https://"+r.Host+r.URL.Path, 301)
		return 0, nil
	}
	return h.next.ServeHTTP(w, r)
}

func Setup(c *setup.Controller) (middleware.Middleware, error) {
	return func(next middleware.Handler) middleware.Handler {
		return &handler{next: next}
	}, nil
}
