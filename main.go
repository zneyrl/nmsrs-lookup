package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"gopkg.in/unrolled/render.v1"
)

func layout(l string) *render.Render {
	return render.New(render.Options{
		Directory:  "views",
		Layout:     "layouts/" + l,
		Extensions: []string{".gohtml"},
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	layout("main").HTML(w, http.StatusOK, "home/index", map[string]interface{}{
		"title": "Home",
	})
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	layout("main").HTML(w, http.StatusOK, "home/welcome", map[string]interface{}{
		"title": "Welcome",
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/login", map[string]interface{}{
		"title": "Login",
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/register", map[string]interface{}{
		"title": "Register",
	})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	layout("search").HTML(w, http.StatusOK, "search/index", map[string]interface{}{
		"title": "Search",
	})
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	layout("search").HTML(w, http.StatusOK, "search/results", map[string]interface{}{
		"title": "Results",
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/welcome", welcomeHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/register", registerHandler).Methods("GET")
	r.HandleFunc("/search", searchHandler).Methods("GET")
	r.HandleFunc("/results", resultsHandler).Methods("GET")

	h := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(h).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}
