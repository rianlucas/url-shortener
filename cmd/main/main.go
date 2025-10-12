package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/handler"
	"github.com/rianlucas/url-shortener/internal/service"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbUrl := viper.Get("DB_URL").(string)

	client, err := mongo.Connect(options.Client().ApplyURI(dbUrl))
	ctx := context.Background()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	urlRepository := repositories.CreateUrlRepository(ctx, client)
	urlService := service.CreateUrlService(urlRepository)
	urlHandler := handler.CreateUrlHandler(urlService)

	http.HandleFunc("/", urlHandler.Create)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		println("error", err)
		return
	}
}
