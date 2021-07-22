package main

import (
	"fmt"
	"github.com/DABronskikh/bgo-3_03.1/pkg/card"
	"time"
)

func main() {
	master := &card.Card{
		Id:       1,
		Issuer:   "MasterCard",
		Balance:  50_000_00,
		Currency: "RUB",
		Number:   "0001",
		Icon:     "MasterCardIcon",
	}

	var transactionsCard []card.Transaction

	transactionsCard = append(transactionsCard, card.Transaction{
		Id:     1,
		Amount: -789_00,
		Date:   time.Now().Unix(),
		MCC:    "5411",
		Type:   "payment",
		Status: "in progress",
	})
	transactionsCard = append(transactionsCard, card.Transaction{
		Id:     2,
		Amount: 1000_00,
		Date:   time.Now().Unix(),
		MCC:    "",
		Type:   "top up",
		Status: "completed",
	})

	master.Transactions = transactionsCard

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println(category)

	category = card.TranslateMCC(master.Transactions[1].MCC)
	fmt.Println(category)
}
