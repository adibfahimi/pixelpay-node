package main

import (
	"flag"
	"log"

	"github.com/adibfahimi/pixelpay-node/routes"
	"github.com/adibfahimi/pixelpay-node/routes/blocks"
	"github.com/adibfahimi/pixelpay-node/routes/mine"
	"github.com/adibfahimi/pixelpay-node/routes/txs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("p", "3000", "port to run the node on")

	flag.Parse()

	log.Println("blockchain started")

	app := fiber.New()

	app.Get("/", routes.IndexRoute)

	app.Get("/blocks", blocks.GetAllBlocksRoute)
	app.Get("/blocks/:hash", blocks.GetBlockByHashRoute)

	app.Get("/txs/:hash", txs.GetTxsByHashRoute)
	app.Post("/txs", txs.AddTxRoute)

	app.Get("/mine", mine.GetBlockToMineRoute)
	app.Post("/mine", mine.AddBlockRoute)

	app.Get("/balance/:address", routes.GetBalanceByAddressRoute)

	log.Fatal(app.Listen(":" + *port))
}
