package mine

import (
	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func GetBlockToMineRoute(c *fiber.Ctx) error {
	if len(core.BC.PendingBlock.Txs) == 0 {
		return c.JSON(fiber.Map{
			"message": "no pending tx",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"message": "block",
		"data": fiber.Map{
			"block":  core.BC.PendingBlock,
			"reward": core.BC.MiningReward,
		},
	})
}
