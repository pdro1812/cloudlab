// currency-converter/cmd/api/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq" // Driver do PostgreSQL

	"currency-converter/config"
	"currency-converter/internal/api"
	"currency-converter/internal/repository"
	"currency-converter/internal/service"
)

func main() {
	// 1. Carregar configurações
	cfg := config.Load()

	// 2. Conectar ao banco de dados
	db, err := sql.Open("postgres", cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	defer db.Close()

	// Verifica se a conexão com o banco de dados está viva
	if err := db.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}
	log.Println("Successfully connected to the database!")

	// 3. Injeção de Dependência (montando as camadas)
	// A ordem é de "dentro para fora": Repository -> Service -> Handler
	conversionRepo := repository.NewPostgresRepository(db)
	conversionSvc := service.NewConversionService(conversionRepo)
	conversionHandler := api.NewConversionHandler(conversionSvc)

	// 4. Configurar o roteador HTTP (estamos usando Chi)
	r := chi.NewRouter()
	r.Use(middleware.Logger)    // Middleware para logar as requisições
	r.Use(middleware.Recoverer) // Middleware para recuperar de panics

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Registra as rotas específicas do nosso handler
	conversionHandler.RegisterRoutes(r)

	// 5. Iniciar o servidor
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
