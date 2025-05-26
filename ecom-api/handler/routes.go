package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRouters(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	tokenMaker := handler.TokenMaker

	// wrapped middlewares
	authMiddleware := GetAuthMiddlewareFunc(tokenMaker)
	adminMiddleware := GetAdminMiddlewareFunc(tokenMaker)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.listProducts)
		r.With(adminMiddleware).Post("/", handler.createProduct)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)

			r.Group(func(r chi.Router) {
				r.Use(adminMiddleware)
				r.Patch("/", handler.updateProduct)
				r.Delete("/", handler.deleteProduct)
			})
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/myorder", handler.getOrder)

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", handler.createOrder)
			r.With(adminMiddleware).Get("/", handler.listOrders)

			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", handler.deleteOrder)
			})
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.createUser)
		r.Post("/login", handler.loginUser)

		r.Group(func(r chi.Router) {
			r.Use(adminMiddleware)
			r.Get("/", handler.listUser)
			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", handler.deleteUser)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware)
			r.Patch("/", handler.updateUser)
			r.Post("/logout", handler.logoutUser)
		})

	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/tokens", func(r chi.Router) {
			r.Post("/renew", handler.renewAccessToken)
			r.Post("/revoke", handler.revokeSession)
		})
	})

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
