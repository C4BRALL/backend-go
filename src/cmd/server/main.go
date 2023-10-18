package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/backend/docs"
	"github.com/backend/src/configs"
	"github.com/backend/src/internal/entity"
	database "github.com/backend/src/internal/infra/db"
	"github.com/backend/src/internal/infra/webServer/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			Backend GO
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:6415
// @BasePath	/v1
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

	sellerDB := database.NewSeller(db)
	sellerHandler := handlers.NewSellerHandler(sellerDB)
	StoreDB := database.NewStore(db)
	StoreHandler := handlers.NewStoreHandler(StoreDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("tokenExpiresIn", config.JWTExpiresIn))

	r.Route("/seller", func(r chi.Router) {
		r.Post("/", sellerHandler.CreateSeller)
		r.Post("/signin", sellerHandler.GetJWT)
		r.Get("/{email}", sellerHandler.GetSeller)
		r.Put("/{email}", sellerHandler.UpdateSeller)
		r.Delete("/{document}", sellerHandler.DeleteSeller)
		r.Get("/all", sellerHandler.GetSellers)
	})

	r.Route("/store", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", StoreHandler.NewStore)
		r.Get("/", StoreHandler.GetStore)
		r.Put("/{id}", StoreHandler.UpdateStore)
		r.Delete("/{id}", StoreHandler.DeleteStore)
		r.Get("/all", StoreHandler.GetStores)
	})

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:6415/swagger/doc.json")))

	log.Printf("Http server running at http://localhost%s", config.WebServerPort)

	err = http.ListenAndServe(config.WebServerPort, r)
	if err != nil {
		panic(err.Error())
	}
}
