package http

import (
	"fmt"
	"libraryes/database"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(HTTPHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: HTTPHandlers,
	}
}

func (h *HTTPServer) StartServer() error{
	if err := godotenv.Load(); err != nil {
		fmt.Println("error read env file")
	}
	db, err := database.NewPostgresDB(database.Config{
				Host: os.Getenv("DB_HOST"),
				Port: os.Getenv("DB_PORT"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				DBName: os.Getenv("DB_NAME"),
				SSLMode: "disable",
				})
				if err != nil {
					fmt.Print(err)
					return err
				}

				defer db.Close()
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(h.httpHandlers.HandleCreateBook(db))
	router.Path("/books/{title}").Methods("GET").HandlerFunc(h.httpHandlers.HandleGetBook)
	router.Path("/books").Methods("GET").Queries("readed", "true").HandlerFunc(h.httpHandlers.HandleGetReadedBook)
	router.Path("/books").Methods("GET").Queries("readed", "false").HandlerFunc(h.httpHandlers.HandleGetUnReadedBook)
	router.Path("/books").Methods("GET").Queries("author", "{author}").HandlerFunc(h.httpHandlers.HandleListBookAuthor)
	router.Path("/books").Methods("GET").HandlerFunc(h.httpHandlers.HandleGetAllBook)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(h.httpHandlers.HandleReadBook)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(h.httpHandlers.HandleDeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		return err
	}
	return nil
}