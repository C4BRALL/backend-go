package postgres

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Repository struct {
	Db  *pgx.Conn
	cfg *configs
	// TokenManager auth.TokenFactory
}

func NewRepository() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return
	}

	// Caminho completo para o arquivo .env na raiz do projeto
	caminhoEnv := filepath.Join(dir, "../../../../.env")

	fmt.Println(caminhoEnv)

	err = godotenv.Load(caminhoEnv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar o arquivo .env: %v\n", err)
		os.Exit(1)
	}

	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse database URL: %v\n", err)
		os.Exit(1)
	}

	// Configurar TLS (remova estas linhas se você não estiver usando TLS)
	config.TLSConfig = &tls.Config{
		InsecureSkipVerify: false, // Defina como false para verificar certificados SSL
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Connected'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
