package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRouters(handler *handler) *chi.Mux {
	r = chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.listProducts)
		r.Post("/", handler.createProduct)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)
			r.Patch("/", handler.updateProduct)
			r.Delete("/", handler.deleteProduct)
		})
	})

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", handler.listOrders)
		r.Post("/", handler.createOrder)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getOrder)
			r.Delete("/", handler.deleteOrder)
		})
	})

	return r
}

func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
