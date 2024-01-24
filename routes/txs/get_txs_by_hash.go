package txs

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func GetTxsByHashRoute(c *fiber.Ctx) error {
	hash := c.Params("hash")

	for _, block := range core.BC.Chain {
		for _, tx := range block.Txs {
			if tx.Hash == hash {
				return c.JSON(fiber.Map{
					"message": "tx found",
					"data":    tx,
				})
			}
		}
	}

	return c.JSON(fiber.Map{
		"message": "tx not found",
		"data":    nil,
	})
}
