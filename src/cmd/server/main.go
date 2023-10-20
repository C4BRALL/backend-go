package main

import (
	"fmt"
	"log"

	_ "github.com/backend/docs"
	"github.com/backend/src/configs"
	"github.com/backend/src/internal/entity"
	"github.com/backend/src/internal/infra/database"
	"github.com/backend/src/internal/infra/web"
	"github.com/backend/src/internal/infra/web/webserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			Backend GO
//	@version		1.0
//	@description	This is a sample server e-commerce server.

//	@contact.name	João Cabral
//	@contact.url	https://github.com/C4BRALL
//	@contact.email	cabral047dev@gmail.com

// @host		localhost:9874
func main() {
	config, err := configs.LoadConfig("../../../")
	if err != nil {
		panic(err)
	}
	fmt.Println("db name:", config.DbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Seller{}, &entity.Store{})

	webserver := webserver.NewWebServer(config.WebServerPort)
	sellerDB := database.NewSellerRepository(db)
	webSellerHandler := web.NewWebSellerHandler(sellerDB)
	webserver.AddHandler("/seller", webSellerHandler.Create)
	log.Printf("Http server running at http://localhost%s", config.WebServerPort)
	webserver.Start()

	// sellerHandler := handlers.NewSellerHandler(sellerDB)
	// StoreDB := database.NewStore(db)
	// StoreHandler := handlers.NewStoreHandler(StoreDB)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.WithValue("jwt", config.TokenAuth))
	// r.Use(middleware.WithValue("tokenExpiresIn", config.JWTExpiresIn))

	// r.Route("/seller", func(r chi.Router) {
	// 	r.Post("/", sellerHandler.CreateSeller)
	// 	r.Post("/signin", sellerHandler.GetJWT)
	// 	r.Get("/{email}", sellerHandler.GetSeller)
	// 	r.Put("/{email}", sellerHandler.UpdateSeller)
	// 	r.Delete("/{document}", sellerHandler.DeleteSeller)
	// 	r.Get("/all", sellerHandler.GetSellers)
	// })

	// r.Route("/store", func(r chi.Router) {
	// 	r.Use(jwtauth.Verifier(config.TokenAuth))
	// 	r.Use(jwtauth.Authenticator)
	// 	r.Post("/", StoreHandler.NewStore)
	// 	r.Get("/", StoreHandler.GetStore)
	// 	r.Put("/{id}", StoreHandler.UpdateStore)
	// 	r.Delete("/{id}", StoreHandler.DeleteStore)
	// 	r.Get("/all", StoreHandler.GetStores)
	// })

	// r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", config.WebServerPort))))

	// log.Printf("Http server running at http://localhost%s", config.WebServerPort)

	// err = http.ListenAndServe(config.WebServerPort, r)
	// if err != nil {
	// 	panic(err.Error())
	// }
}
