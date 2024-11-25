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
	dbClient.DB.AutoMigrate(&user.ConfigUser{})
	crypto := crypto.NewCrypto()

	userRepository := user.NewUserRepository(dbClient.DB)
	userService := user.NewUserService(
		userRepository,
		crypto,
		config.AccessTokenSecret,
	)
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
		c.Set("crypto", crypto)
		c.Set("secret", config.AccessTokenSecret)
		c.Next()
	})

	routes.Use(middleware.Cors())
	user.AddUserRouter(routes, userHandler)
	auth.AddAuthRouter(routes, authHandler)
	adult.AddRoutesAdult(routes, *adultHandler)
	stats.AddStatsRouter(routes, statsHandler)

	fmt.Printf("API server listening at: %s\n\n", config.HttpPort)
	routes.Run(config.HttpPort)
}
