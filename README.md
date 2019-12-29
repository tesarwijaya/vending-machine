# Vending Machine

## How to run

### Development

`go run main.go`

### Build Production

`CGO_ENABLED=0 go build -o vending-machine`

## Available Command

Command (1)
	Command name     : Insert coins
	Command number   : 1
	Argument         : int, coin types (any of 10, 50, 100, 500)
For example: “1 500” (Insert 500 JPY coin)

Command (2)
	Command name     : Choose item to purchase
	Command number   : 2
	Argument         : int, item types (any of item numbers)
For example: “2 1” (1: Choose Canned coffee)

Command (3)
	Command name     : Get items
	Command number   : 3
	Argument         : None
For example: “3” (Get items)

Command (4)
	Command name     : Return coins
	Command number   : 4
	Argument         : None
For example: “4” (Pull Return lever)

Command (5)
	Command name     : Get returned coins
	Command number   : 5
	Argument         : None
For example: “5” (Get returned coins)