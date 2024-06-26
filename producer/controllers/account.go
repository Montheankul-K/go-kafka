package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"producer/commands"
	"producer/services"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService: accountService}
}

func (h accountController) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	id, err := h.accountService.OpenAccount(command)
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

func (h accountController) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = h.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "deposit fund success",
	})
}

func (h accountController) WithdrawFund(c *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = h.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "withdraw fund success",
	})
}

func (h accountController) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = h.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(fiber.Map{
		"message": "close account success",
	})
}
