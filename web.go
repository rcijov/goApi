package main
 
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "lib"
)
 
 
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/people", lib.GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", lib.GetPersonEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", lib.CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people/{id}", lib.DeletePersonEndpoint).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))


}