package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Kaivv1/blog-aggregator/internal/config"
	"github.com/Kaivv1/blog-aggregator/internal/database"
	feedfollow "github.com/Kaivv1/blog-aggregator/internal/handlers/v1/feed_follow"
	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/feeds"
	"github.com/Kaivv1/blog-aggregator/internal/handlers/v1/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("cannot get port env")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("cannot get port env")
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	v1Router := chi.NewRouter()
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	router.Mount("/v1", v1Router)
	config := config.NewConfig(database.New(db))
	usersRouter := users.NewUsersRouter(config)
	feedsRouter := feeds.NewFeedsRouter(config)
	feedsFollowsRouter := feedfollow.NewFeedFollowRouter(config)

	v1Router.Mount("/users", usersRouter)
	v1Router.Mount("/feeds", feedsRouter)
	v1Router.Mount("/feed_follows", feedsFollowsRouter)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("server running on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
