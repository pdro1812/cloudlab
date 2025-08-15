package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // Pacote para criar o servidor web
	"os"

	// O _ significa que estamos importando o pacote por seus efeitos colaterais
	// (ele se registra como um driver de banco de dados), mas não vamos usá-lo diretamente no código.
	_ "github.com/lib/pq"
)

func main() {
	// --- Conexão com o Banco de Dados ---
	// Lê as variáveis de ambiente que o Docker Compose nos forneceu.
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Monta a string de conexão (connection string)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abre uma conexão com o banco de dados.
	// sql.Open não estabelece a conexão imediatamente, apenas prepara.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// log.Fatal irá imprimir o erro e encerrar a aplicação.
		log.Fatal("Erro ao preparar a conexão com o banco de dados: ", err)
	}
	// Garante que a conexão será fechada quando a função main terminar.
	defer db.Close()

	// --- Servidor Web ---
	// Criamos uma "rota" ou "endpoint". Quando alguém acessar "/", a função handler será executada.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Dentro desta função, vamos testar a conexão com o banco de dados de verdade.
		// db.Ping() força uma conexão com o banco e retorna um erro se falhar.
		err := db.Ping()
		if err != nil {
			log.Println("Erro ao conectar com o banco de dados:", err)
			// Retorna um erro HTTP 500 (Internal Server Error) se a conexão falhar.
			http.Error(w, "Erro: Não foi possível conectar ao banco de dados.", http.StatusInternalServerError)
			return
		}

		// Se a conexão for bem-sucedida, escreve uma mensagem de sucesso.
		fmt.Fprintln(w, "Conexão com o banco de dados PostgreSQL bem-sucedida!")
	})

	// Informa no console que o servidor está iniciando.
	log.Println("Servidor iniciado na porta :8080")
	log.Println("Acesse http://localhost:8080 para testar a conexão com o banco.")

	// Inicia o servidor web na porta 8080 e fica escutando por requisições.
	// Se houver um erro ao iniciar o servidor, o programa será encerrado.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
