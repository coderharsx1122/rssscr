package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coderharsx1122/rssscr/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load env file
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found")
	}

	DB_URL := os.Getenv("DB_URL")

	if DB_URL == "" {
		log.Fatal("DB URL not found")
	}

	// db connection
	con, err := sql.Open("postgres", DB_URL)

	if err != nil {
		log.Fatal("Can't connect to database")
	}
	queries := database.New(con)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	// cors config
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// v1 router
	v1Router := chi.NewRouter()

	// v1Router.HandleFunc("/ready",handleReq) // will handle function for any method
	v1Router.Get("/ready", handleReq) // will handle function for only get request
	v1Router.Get("/err", handleErr)
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.getUserHandler)

	// mount v1Router at /v1
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("server starting on port %v", port)

	// start server
	er := server.ListenAndServe()

	if er != nil {
		log.Fatal(er)
	}

	fmt.Println("PORT:", port)

}
