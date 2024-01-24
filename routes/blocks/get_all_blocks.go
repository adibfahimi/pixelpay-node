package blocks

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func GetAllBlocksRoute(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "List of blocks",
		"data":    core.BC.Chain,
	})
}
