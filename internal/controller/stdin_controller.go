package controller

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/elct9620/demo-stdio-go-plugin/internal/usecase"
)

type Stdin struct {
	encodeUsecase *usecase.EncodeProduct
}

func NewStdin(encodeUsecase *usecase.EncodeProduct) *Stdin {
	return &Stdin{
		encodeUsecase: encodeUsecase,
	}
}

func (s *Stdin) Handle(ctx context.Context, reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	items := make([]usecase.EncodeInputItem, 0)

	fmt.Println("Please input product name and price, separated by space")
	fmt.Println("Example: iPhone 1000")
	fmt.Println("To exit, press Ctrl + D")
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}

		values := strings.SplitN(input, " ", 2)
		if len(values) != 2 {
			fmt.Println("Invalid input")
			fmt.Print("> ")
			continue
		}

		name := values[0]
		price, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Printf("Invalid price: %v\n", err)
			fmt.Print("> ")
			continue
		}

		items = append(items, usecase.EncodeInputItem{
			Name:  name,
			Price: price,
		})

		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading input: %v", err)
	}

	return s.encodeUsecase.Execute(ctx, &usecase.EncodeInput{
		Items: items,
	})
}
