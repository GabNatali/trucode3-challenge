package main

import (
	"fmt"
	"log"

	// "sync"

	"github.com/GabNatali/trucode3-challenge-final/cli"
	"github.com/gin-gonic/gin"

	// "github.com/GabNatali/trucode3-challenge-final/utils"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	stats "github.com/GabNatali/trucode3-challenge-final/internal/Stats"
	"github.com/GabNatali/trucode3-challenge-final/internal/adult"
	"github.com/GabNatali/trucode3-challenge-final/internal/auth"
	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
	"github.com/GabNatali/trucode3-challenge-final/internal/database"
	"github.com/GabNatali/trucode3-challenge-final/internal/user"
	"github.com/GabNatali/trucode3-challenge-final/middleware"
)

func main() {

	parser := cli.NewParser()

	config, err := parser.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbClient := database.NewService(config.DatabaseURL)

	err = dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	dbClient.DB.AutoMigrate(&adult.Adult{})
	dbClient.DB.AutoMigrate(&user.User{})
	crypto := crypto.NewCrypto()

	// linesChan := make(chan []adult.Adult, 4)
	// var wg sync.WaitGroup

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for batch := range linesChan {
	// 			if err := dbClient.DB.Create(&batch).Error; err != nil {
	// 				log.Printf("Failed to insert remaining batch: %v", err)
	// 			}
	// 		}
	// 	}()
	// }

	// path := "../data/source.data"

	// err = utils.ReadFile(path, linesChan)

	// if err != nil {
	// 	log.Fatalf("Error leyendo el archivo: %v", err)
	// }

	// close(linesChan)
	// wg.Wait()

	userRepository := user.NewUserRepository(dbClient.DB)
	userService := user.NewUserService(userRepository, crypto)
	userHandler := user.NewUserHandler(userService)

	authService := auth.NewAuthService(
		auth.AuthServiceOpts{
			UserRepository: userRepository,
			Crypto:         crypto,
			Config:         config.AccessTokenSecret,
		})
	authHandler := auth.NewAuthHandler(authService)

	adultRepository := adult.NewAdultRepository(dbClient.DB)
	adultService := adult.NewAdultService(adultRepository)
	adultHandler := adult.NewUserHandler(adultService)

	statsRepository := stats.NewRepository(dbClient.DB)
	statsHandler := stats.NewHandler(statsRepository)

	routes := gin.Default()

	routes.Use(func(c *gin.Context) {
		c.Set("authService", authService)
		c.Next()
	})

	routes.Use(middleware.Cors())
	user.AddUserRouter(routes, userHandler)
	auth.AddAuthRouter(routes, authHandler)
	adult.AddRoutesAdult(routes, *adultHandler)
	stats.AddStatsRouter(routes, statsHandler)

	fmt.Printf("API server listening at: %s\n\n", ":3333")
	log.Fatal(routes.Run(":3333"))
}
