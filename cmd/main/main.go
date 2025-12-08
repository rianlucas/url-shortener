package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rianlucas/url-shortener/config"
	"github.com/rianlucas/url-shortener/internal/database"
	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/handler"
	"github.com/rianlucas/url-shortener/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

	db := client.Database("url-shortener")
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

	clickRepository := repositories.NewClickAnalyticsRepository(ctx, db)
	clickService := service.NewClickService(ctx, clickRepository)

	urlHandler := handler.NewUrlHandler(urlService, clickService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/{shortCode}", urlHandler.FindByShortCode)
		r.Post("/", urlHandler.Create)
		r.Get("/qr-code/{shortCode}", urlHandler.ShowQrCode)
	})

	log.Println("✅ Server configured successfully!")
	log.Println("📋 Available routes:")
	log.Println("   POST /               - Create short URL")
	log.Println("   GET  /{shortCode}    - Redirect to long URL")
	log.Println("   GET  /qr-code/{shortCode} - Generate QR Code")
	log.Println("")
	log.Println("🌐 Server listening on http://localhost:8000")

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		println("error", err)
		return
	}
}
