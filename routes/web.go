package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/controllers/applicant"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/dashboard"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
	"github.com/zneyrl/nmsrs-lookup/controllers/reports"
	"github.com/zneyrl/nmsrs-lookup/controllers/search"
	"github.com/zneyrl/nmsrs-lookup/controllers/user"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
)

func Web() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").Methods("GET").HandlerFunc(home.Index)
	r.Path("/welcome").Methods("GET").HandlerFunc(home.Welcome)

	login := r.Path("/login").Subrouter()
	login.Methods("GET").HandlerFunc(auth.ShowLoginForm)
	login.Methods("POST").HandlerFunc(auth.Login)

	// register := r.Path("/register").Subrouter()
	// register.Methods("GET").HandlerFunc(auth.ShowRegisterForm)
	// register.Methods("POST").HandlerFunc(auth.Register)

	r.Path("/dashboard").Methods("GET").Handler(middlewares.Secure(dashboard.Index))
	r.Path("/dashboard/overview").Methods("GET").Handler(middlewares.Secure(dashboard.Overview))

	r.Path("/users").Methods("GET").Handler(middlewares.Secure(user.Index))
	r.Path("/users/create").Methods("GET").Handler(middlewares.Secure(user.Create))
	r.Path("/users").Methods("POST").Handler(middlewares.Secure(user.Store))
	r.Path("/users/ids").Methods("POST").Handler(middlewares.Secure(user.DestroyMany))
	r.Path("/users/{id}").Methods("GET").Handler(middlewares.Secure(user.Show))
	r.Path("/users/{id}/edit").Methods("GET").Handler(middlewares.Secure(user.Edit))
	r.Path("/users/{id}").Methods("PUT").Handler(middlewares.Secure(user.Update))
	r.Path("/users/{id}").Methods("DELETE").Handler(middlewares.Secure(user.Destroy))
	r.Path("/users/{id}/reset-password").Methods("POST").Handler(middlewares.Secure(user.ResetPassword))

	r.Path("/applicants").Methods("GET").Handler(middlewares.Secure(applicant.Index))
	r.Path("/reports").Methods("GET").Handler(middlewares.Secure(reports.Index))

	r.Path("/search").Methods("GET").Handler(middlewares.Secure(search.Index))
	r.Path("/results").Methods("GET").Handler(middlewares.Secure(search.Results))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return r
}
