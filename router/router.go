package router

import (
	"openthailand/controllers/api"
	"openthailand/route"

	"github.com/gorilla/mux"
)

func RegisterRouting(r *mux.Router) {
	route.Router(r).Register(
		api.Routes(),
	)
}
