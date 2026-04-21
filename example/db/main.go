package main

// chat gpt apresenta:

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// connection string TCP
	dsn := "postgresql://admin:admin123@localhost:8080/appdb"
	// contexto base (controle de timeout em sistemas reais)
	ctx := context.Background()
	// cria pool de conexões
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	defer pool.Close()

	// valida conexão de forma explícita (boa prática)
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("banco não respondeu: %v", err)
	}
	_, err = pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS users ()")
	if err != nil {
		log.Fatalf("erro ao executar: %v", err)
	}
	fmt.Println("conectado no PostgreSQL com sucesso")
	time.Sleep(2 * time.Second)
}
