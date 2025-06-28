package router
import(
"github.com/gorilla/mux"

)


func Router() *mux.Router {
	router := mux.NewRouter()
r.HandleFunc("/", controllers.SetString).Methods("GET")
	r.HandleFunc("/api/json", controllers.GetJSON).Methods("GET")
	r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{ID}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{ID}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{ID}", controllers.DeleteBook).Methods("DELETE")

	return router
}