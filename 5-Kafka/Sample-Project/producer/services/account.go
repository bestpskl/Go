package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) (id string, err error)
	DepositFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
	ShowBalance(command commands.ShowBalanceCommand) error
	ShowTransactions(command commands.ShowTransactionsCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer: eventProducer}
}

func (obj accountService) OpenAccount(command commands.OpenAccountCommand) (id string, err error) {

	if command.AccountHolder == "" || command.AccountType == 0 || command.Balance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:            uuid.NewString(),
		AccountHolder: command.AccountHolder,
		AccountType:   command.AccountType,
		Balance:       command.Balance,
	}

	log.Printf("%#v", event)
	return event.ID, obj.eventProducer.Produce(event)
}

func (obj accountService) DepositFund(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) WithdrawFund(command commands.WithdrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) CloseAccount(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.CloseAccountEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) ShowBalance(command commands.ShowBalanceCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.ShowBalanceEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) ShowTransactions(command commands.ShowTransactionsCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.ShowTransactionsEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}
