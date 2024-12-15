package usecase

import (
	"context"

	"github.com/elct9620/demo-stdio-go-plugin/internal/entity"
)

type EncodeInputItem struct {
	Name  string
	Price int
}

type EncodeInput struct {
	Items []EncodeInputItem
}

type Encoder interface {
	Encode(products []*entity.Product) ([]byte, error)
}

type Presenter interface {
	Render([]byte) error
}

type EncodeProduct struct {
	encoder   Encoder
	presenter Presenter
}

func NewEncodeProduct(encoder Encoder, presenter Presenter) *EncodeProduct {
	return &EncodeProduct{
		encoder:   encoder,
		presenter: presenter,
	}
}

func (e *EncodeProduct) Execute(ctx context.Context, input *EncodeInput) error {
	products := make([]*entity.Product, 0, len(input.Items))
	for _, item := range input.Items {
		products = append(products, entity.NewProduct(item.Name, item.Price))
	}

	res, err := e.encoder.Encode(products)
	if err != nil {
		return err
	}

	return e.presenter.Render(res)
}
