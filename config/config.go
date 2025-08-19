// currency-converter/config/config.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // 1. Importar a biblioteca
)

type Config struct {
	DBConnectionString string
	ServerPort         string
}

func Load() *Config {
	// 2. Chamar godotenv.Load() no início da função
	// Ele vai procurar por um arquivo .env na raiz e carregar as variáveis dele
	// para o ambiente da aplicação.
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables from OS")
	}

	// 3. Ler as variáveis de ambiente usando os.Getenv()
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}

	return &Config{
		DBConnectionString: connStr,
		ServerPort:         port,
	}
}
