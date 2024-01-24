package txs

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func AddTxRoute(c *fiber.Ctx) error {
	var tx core.Tx
	if err := c.BodyParser(&tx); err != nil {
		return c.JSON(fiber.Map{
			"message": "invalid tx",
			"data":    nil,
		})
	}

	if err := tx.IsValid(); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	core.BC.PendingBlock.Txs = append(core.BC.PendingBlock.Txs, tx)

	return c.JSON(fiber.Map{
		"message": "tx added",
		"data":    tx,
	})
}
