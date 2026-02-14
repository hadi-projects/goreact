package main

import (
	"fmt"

	"github.com/hadi-projects/go-react-starter/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("App Name: ", cfg.AppName)
	fmt.Println("App Port: ", cfg.AppPort)
	fmt.Println("App Env: ", cfg.APPEnv)
	fmt.Println("DB Host: ", cfg.DBHost)
	fmt.Println("DB Port: ", cfg.DBPort)
	fmt.Println("DB Username: ", cfg.DBUserName)
	fmt.Println("DB Password: ", cfg.DBPassword)
	fmt.Println("DB Name: ", cfg.DBName)
	fmt.Println("Redis Host: ", cfg.RedisHost)
	fmt.Println("Redis Port: ", cfg.RedisPort)
	fmt.Println("Redis Password: ", cfg.RedisPassword)
	fmt.Println("Redis DB: ", cfg.RedisDB)
	fmt.Println("CORS Allowed Origins: ", cfg.CORSAllowedOrigins)
	fmt.Println("CORS Allowed Methods: ", cfg.CORSAllowedMethods)
	fmt.Println("CORS Allowed Headers: ", cfg.CORSAllowedHeaders)
	fmt.Println("CORS Max Age: ", cfg.CORSMaxAge)
	fmt.Println("CORS Exposed Headers: ", cfg.CORSExposedHeaders)
	fmt.Println("CORS Allow Credentials: ", cfg.CORSAllowCredentials)
	fmt.Println("JWT Secret: ", cfg.JwtSecret)
	fmt.Println("JWT Issuer: ", cfg.JwtIssuer)
	fmt.Println("JWT Access Expiration Time: ", cfg.JwtAccessExpirationTime)
	fmt.Println("Rate Limit RPS: ", cfg.RateLimitRps)
	fmt.Println("Rate Limit Burst: ", cfg.RateLimitBurst)
	fmt.Println("Request Time Out: ", cfg.RequestTimeOut)
	fmt.Println("API Key: ", cfg.APIKey)
	fmt.Println("BCrypt Cost: ", cfg.BCryptCost)
	fmt.Println("Admin Email: ", cfg.AdminEmail)
	fmt.Println("Admin Password: ", cfg.AdminPassword)
}
