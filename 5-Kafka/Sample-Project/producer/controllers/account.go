package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
	ShowBalance(c *fiber.Ctx) error
	ShowTransactions(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService}
}

func (obj accountController) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	id, err := obj.accountService.OpenAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "open account success",
		"id":      id,
	})
}

func (obj accountController) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "deposit fund",
	})
}

func (obj accountController) WithdrawFund(c *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "withdraw fund",
	})
}

func (obj accountController) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "close account success",
	})
}

func (obj accountController) ShowBalance(c *fiber.Ctx) error {
	command := commands.ShowBalanceCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.ShowBalance(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check balance success",
	})
}

func (obj accountController) ShowTransactions(c *fiber.Ctx) error {
	command := commands.ShowTransactionsCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = obj.accountService.ShowTransactions(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "check transactions",
	})
}
