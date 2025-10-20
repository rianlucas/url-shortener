package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rianlucas/url-shortener/config"
	"github.com/rianlucas/url-shortener/internal/database"
	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/handler"
	"github.com/rianlucas/url-shortener/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"net/http"
	"time"
)

func main() {

	limiter := time.Tick(100 * time.Millisecond)

	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	dbUrl := conf.DbConn

	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}

	db := client.Database("url_shortener")
	err = database.CreateUrlIndexes(db)
	if err != nil {
		fmt.Println("Error creating indexes:", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	urlRepository := repositories.NewUrlRepository(ctx, db)
	urlService := service.NewUrlService(urlRepository)
	urlHandler := handler.NewUrlHandler(urlService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-limiter:
			if r.Method == "POST" {
				urlHandler.Create(w, r)
			} else if r.Method == "GET" {
				urlHandler.FindByShortCode(w, r)
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}

		default:
			w.WriteHeader(http.StatusTooManyRequests)
			w.Header().Set("Content-Type", "Application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"success": false,
				"error":   "too many requests",
			})
		}

	})

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		println("error", err)
		return
	}
}
