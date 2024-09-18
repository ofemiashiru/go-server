package main

import (
	"acme/api"
	"acme/postgres"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello Learner")
}

// Function to stop any Cors errors
func CorsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(writer, request)
	})
}

func main() {
	// Use .env
	godotenv.Load()
	postgrespass := os.Getenv("POSTGRESSPASS")

	// Postgress connection
	connectionString := fmt.Sprintf("user=postgres dbname=acme password=%s host=localhost sslmode=disable", postgrespass)
	if err := postgres.InitDB(connectionString); err != nil {
		fmt.Println("Error initialising databse:", err)
		return
	}

	defer postgres.DB.Close()

	//set up our multiplexer - we then use the router.HandleFunc to handle
	router := http.NewServeMux()

	// router.HandlFunc can handle multiple routes
	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", api.GetUsers)
	router.HandleFunc("POST /api/users", api.CreateUser)
	router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", api.DeleteSingleUser)
	router.HandleFunc("PUT /api/users/{id}", api.UpdateSingleUser)

	// Start Server here
	fmt.Println("Server listening on port 8080")

	// make sure to pass "router" to ListenAndServe
	err := http.ListenAndServe(":8080", CorsMiddleWare(router))

	if err != nil {
		fmt.Println("Error starting", err)
	}
}
