// currency-converter/internal/service/conversion_service.go
package service

import (
	"context"
	"fmt"
	"math/rand"

	"currency-converter/internal/domain"
	"currency-converter/internal/repository"
)

// ConversionService define a interface para a lógica de negócio de conversão.
type ConversionService interface {
	Convert(ctx context.Context, req domain.ConversionRequest) (*domain.Conversion, error)
}

type conversionService struct {
	repo repository.ConversionRepository
}

// NewConversionService cria uma nova instância do serviço.
// Note que ele recebe a INTERFACE do repositório, não a implementação.
// Isso é a chave do baixo acoplamento!
func NewConversionService(repo repository.ConversionRepository) ConversionService {
	return &conversionService{repo: repo}
}

func (s *conversionService) Convert(ctx context.Context, req domain.ConversionRequest) (*domain.Conversion, error) {
	// Passo 1: Buscar a taxa de conversão.
	// SIMULAÇÃO: Em uma aplicação real, aqui você chamaria uma API externa (ex: Banco Central).
	rate, err := s.fetchConversionRateFromExternalAPI(req.From, req.To)
	if err != nil {
		return nil, fmt.Errorf("could not get conversion rate: %w", err)
	}

	// Passo 2: Calcular o valor convertido.
	convertedAmount := req.Amount * rate

	// Passo 3: Criar o objeto de domínio para salvar.
	conversion := &domain.Conversion{
		FromCurrency:    req.From,
		ToCurrency:      req.To,
		InitialAmount:   req.Amount,
		ConvertedAmount: convertedAmount,
		Rate:            rate,
	}

	// Passo 4: Persistir o resultado usando o repositório.
	err = s.repo.Save(ctx, conversion)
	if err != nil {
		return nil, fmt.Errorf("could not save conversion: %w", err)
	}

	return conversion, nil
}

// fetchConversionRateFromExternalAPI é uma simulação de uma chamada a um serviço externo.
func (s *conversionService) fetchConversionRateFromExternalAPI(from, to string) (float64, error) {
	fmt.Printf("SIMULATING: fetching rate for %s to %s\n", from, to)
	// Para este exemplo, vamos retornar uma taxa aleatória para simular a variação.
	// Exemplo: BRL para USD -> algo em torno de 0.18 - 0.20
	// Exemplo: USD para BRL -> algo em torno de 5.0 - 5.5
	if from == "USD" && to == "BRL" {
		return 5.0 + rand.Float64()*(0.5), nil
	}
	if from == "BRL" && to == "USD" {
		return 0.18 + rand.Float64()*(0.02), nil
	}
	// Retornar um erro para pares não suportados na simulação
	return 0, fmt.Errorf("currency pair not supported in this simulation")
}
