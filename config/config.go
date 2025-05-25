package config

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)
fun LoadEnv(){
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found, using default environment variables")
	}
}