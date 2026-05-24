package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler
	// for i := len(middlewares) - 1; i >= 0; i-- {
	// 	middleware := middlewares[i]
	// 	n = middleware(n)
	// }

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		h = globalMiddleware(h)
	}

	return h

}
