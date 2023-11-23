package main

import (
	"database/sql"
	"fmt"

	"github.com/HeitorFM2/GoIntensive.go/internal/infra/database"
	"github.com/HeitorFM2/GoIntensive.go/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1",
		Price: 10.0,
		Tax:   0.1,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
