package blocks

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func GetBlockByHashRoute(c *fiber.Ctx) error {
	hash := c.Params("hash")

	for _, block := range core.BC.Chain {
		if block.Hash == hash {
			return c.JSON(fiber.Map{
				"message": "Block found",
				"data":    block,
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Block not found",
		"data":    nil,
	})
}
