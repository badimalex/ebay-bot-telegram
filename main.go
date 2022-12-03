package main

import (
	"fmt"

	"github.com/badimalex/ebay-bot-telegram/ebay"
	"github.com/badimalex/ebay-bot-telegram/telegram"
)

func main() {

	fmt.Println("Start")

	ebay.InitSearch();

	telegram.SendMessage("Hello world");
}

// @todo plan

// Add .env with TOKEN
// ADD logic for chat id, fetch and send messages from DB
// LOGIC
//  Parse products
//  Send message with product id to db, and send it to recepient

// Bot asks user about what to search
// - "Tell what to search"
// - User: "Hohner Verdi II"

// Bot sends to user first Batch with products and store it to DB => [http:...1, http://2, ....]
// Cron gows to ebay and parse new added products
// if product with this id is not in messages for this user (user_id, product_id)
// Cron job sends to user message with new product

// Finish.
