package mine

import (
	"time"

	"github.com/adibfahimi/pixelpay-node/core"
	"github.com/gofiber/fiber/v2"
)

func AddBlockRoute(c *fiber.Ctx) error {
	if len(core.BC.PendingBlock.Txs) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "no pending tx",
			"data":    nil,
		})
	}

	var block core.Block
	if err := c.BodyParser(&block); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body",
			"data":    nil,
		})
	}

	if block.PreviousHash != core.BC.GetLatestBlock().Hash {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid block",
			"data":    nil,
		})
	}

	if err := block.IsValid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	core.BC.Chain = append(core.BC.Chain, block)

	lastBlock := core.BC.GetLatestBlock()
	core.BC.PendingBlock = core.Block{
		Hash:         "",
		PreviousHash: lastBlock.Hash,
		MerkleRoot:   "",
		Txs:          []core.Tx{},
		Index:        lastBlock.Index + 1,
		Timestamp:    uint(time.Now().Unix()),
		Nonce:        0,
		Difficulty:   core.BC.Difficulty,
	}

	return c.JSON(fiber.Map{
		"message": "Block added successfully",
		"data":    block,
	})
}
