package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo: accountRepo}
}

func (obj accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {

	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := &events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.Balance,
		}
		err = obj.accountRepo.SaveAccount(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

		accountTransaction := repositories.AccountTransaction{
			ID:              uuid.New().String(),
			AccountID:       event.ID,
			TransactionType: "deposit",
			Amount:          event.Balance,
			CreateAt:        time.Now().Add(time.Hour * time.Duration(7)),
		}
		err = obj.accountRepo.CreateTransaction(accountTransaction)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := obj.accountRepo.FindAccountByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance += event.Amount
		err = obj.accountRepo.SaveAccount(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

		accountTransaction := repositories.AccountTransaction{
			ID:              uuid.New().String(),
			AccountID:       event.ID,
			TransactionType: "deposit",
			Amount:          event.Amount,
			CreateAt:        time.Now().Add(time.Hour * time.Duration(7)),
		}
		err = obj.accountRepo.CreateTransaction(accountTransaction)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := obj.accountRepo.FindAccountByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance -= event.Amount
		err = obj.accountRepo.SaveAccount(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

		accountTransaction := repositories.AccountTransaction{
			ID:              uuid.New().String(),
			AccountID:       event.ID,
			TransactionType: "withdraw",
			Amount:          event.Amount,
			CreateAt:        time.Now().Add(time.Hour * time.Duration(7)),
		}
		err = obj.accountRepo.CreateTransaction(accountTransaction)
		if err != nil {
			log.Println(err)
			return
		}

	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := &events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		err = obj.accountRepo.DeleteAccount(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

	case reflect.TypeOf(events.ShowBalanceEvent{}).Name():
		event := &events.ShowBalanceEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := obj.accountRepo.FindAccountByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		accountid := bankAccount.ID
		balance := bankAccount.Balance
		log.Printf("[%v] %#v\n", topic, event)
		log.Printf("Your balance (AccountID : %v) is %v", accountid, balance)

	case reflect.TypeOf(events.ShowTransactionsEvent{}).Name():
		event := &events.ShowTransactionsEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		accountTransactions, err := obj.accountRepo.FindTransactionsByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		accountid := accountTransactions[1].AccountID
		log.Printf("[%v] %#v\n", topic, event)
		log.Printf("List of transactions (AccountID : %v)\n", accountid)
		for _, v := range accountTransactions {
			log.Println(v)
		}

	default:
		log.Println("no event handler")
	}
}
