package api

import "github.com/go-chi/chi/v5"

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.HandleSignUpUser)
				r.Post("/login", api.HandleLoginUser)
				r.Post("/logout", api.HandleLogOutUser)
			})
		})
	})

}