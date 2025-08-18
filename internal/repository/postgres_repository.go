// currency-converter/internal/repository/postgres_repository.go
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"currency-converter/internal/domain"
)

// ConversionRepository define a interface para as operações de persistência.
// Usar uma interface aqui permite "mockar" o repositório em testes.
type ConversionRepository interface {
	Save(ctx context.Context, conversion *domain.Conversion) error
}

// postgresRepository é a implementação concreta usando PostgreSQL.
type postgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository cria uma nova instância do repositório.
func NewPostgresRepository(db *sql.DB) ConversionRepository {
	return &postgresRepository{db: db}
}

// Save insere um novo registro de conversão no banco de dados.
func (r *postgresRepository) Save(ctx context.Context, c *domain.Conversion) error {
	query := `INSERT INTO conversion_history (from_currency, to_currency, initial_amount, converted_amount, rate)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query, c.FromCurrency, c.ToCurrency, c.InitialAmount, c.ConvertedAmount, c.Rate).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save conversion: %w", err)
	}
	return nil
}
