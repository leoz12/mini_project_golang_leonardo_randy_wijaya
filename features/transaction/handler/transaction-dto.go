package transactionHandler

import (
	"mini_project/features/transaction"
	"time"
)

type CreateRequest struct {
	GameId   string `json:"gameId" form:"gameId"`
	Quantity int    `json:"quantity" form:"quantity"`
}

type TransactionResponse struct {
	Id              string
	UserId          string
	GameId          string
	GameName        string
	GameDescription string
	Price           float32
	Quantity        int
	Discount        float32
	TotalPrice      float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CoreToResponse(data transaction.Core) TransactionResponse {
	return TransactionResponse{
		Id:              data.Id,
		UserId:          data.UserId,
		GameId:          data.GameId,
		GameName:        data.GameName,
		GameDescription: data.GameDescription,
		Price:           data.Price,
		Quantity:        data.Quantity,
		Discount:        data.Discount,
		TotalPrice:      (data.Price * (100 - data.Discount) / 100) * float32(data.Quantity),
	}
}
