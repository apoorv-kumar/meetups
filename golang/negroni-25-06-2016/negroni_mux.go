package main

import(
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "net/http"
)

func SimpleUserHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Hello %v\n", vars["id"])
}

func EmailUserHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Hello %v. You're logged in with your %v account\n", vars["id"], vars["domain"])
}

func main() {
    mux := mux.NewRouter()
    mux.HandleFunc("/user/{id:[a-zA-Z0-9]+}@{domain:[a-z]+}.com", EmailUserHandler).Methods("GET")
    mux.HandleFunc("/user/{id}", SimpleUserHandler).Methods("GET")
    n := negroni.Classic()
    n.UseHandler(mux)
    n.Run(":3001")
}