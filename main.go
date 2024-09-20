package main

import (
	"acme/api"
	"acme/config"
	"acme/db/postgres"
	"acme/repository/user"
	"acme/service"
	"fmt"
	"io"
	"net/http"
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
	//Load config change to .Postgress or .inmemory
	config := config.LoadDatabaseConfig( /*".env.inmemory"*/ )

	var userRepo user.UserRepository

	switch config.Type {
	case "postgres":
		connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)

		db, err := postgres.PostgresConnection(connectionString)
		if err != nil {
			panic(err)
		}

		userRepo = user.NewPostgresUserRepository(db.DB)

	case "inmemory":
		//for in-memory, we don't need db connection details as the repository itself does this
		userRepo = user.NewInMemoryRepository()

	default:
		fmt.Errorf("unsupported database type: %s", config.Type)
	}

	// Initialize services
	userService := service.NewUserService(userRepo)
	userAPI := api.NewUserAPI(userService)

	//set up our multiplexer - we then use the router.HandleFunc to handle
	router := http.NewServeMux()

	// router.HandlFunc can handle multiple routes
	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", userAPI.GetUsers)
	router.HandleFunc("POST /api/users", userAPI.CreateUser)
	router.HandleFunc("GET /api/users/{id}", userAPI.GetSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", userAPI.DeleteSingleUser)
	router.HandleFunc("PUT /api/users/{id}", userAPI.UpdateSingleUser)

	// Start Server here
	fmt.Println("Server listening on port 8080")

	// make sure to pass "router" to ListenAndServe
	err := http.ListenAndServe(":8080", CorsMiddleWare(router))

	if err != nil {
		fmt.Println("Error starting", err)
	}
}
