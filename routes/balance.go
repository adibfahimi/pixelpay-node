package routes

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func GetBalanceByAddressRoute(c *fiber.Ctx) error {
	address := c.Params("address")

	balance := uint(0)

	for _, block := range core.BC.Chain {
		for _, tx := range block.Txs {
			if tx.From == address {
				balance -= tx.Amount
			}

			if tx.To == address {
				balance += tx.Amount
			}
		}
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"amount": balance,
		},
	})
}
