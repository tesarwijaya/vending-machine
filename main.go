package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tesarwijaya/vending-machine/vending"
)

func inputCoinStock(v *vending.Vending, value int, amount int) {
	for i := 0; i < amount; i++ {
		v.CoinStock(&vending.Coin{Value: value})
	}
}

func main() {
	v := vending.InitVending()

	// Stock coin for change
	inputCoinStock(v, 10, 10)
	inputCoinStock(v, 50, 10)
	inputCoinStock(v, 100, 10)
	inputCoinStock(v, 500, 5)

	// Stock initial product
	v.ItemStock(&vending.Item{Name: "Canned coffee", Price: 120, Slot: "1"})
	v.ItemStock(&vending.Item{Name: "Canned coffee", Price: 120, Slot: "1"})
	v.ItemStock(&vending.Item{Name: "Canned coffee", Price: 120, Slot: "1"})

	v.ItemStock(&vending.Item{Name: "Water PET bottle", Price: 100, Slot: "2"})
	v.ItemStock(&vending.Item{Name: "Water PET bottle", Price: 100, Slot: "2"})
	v.ItemStock(&vending.Item{Name: "Water PET bottle", Price: 100, Slot: "2"})
	v.ItemStock(&vending.Item{Name: "Water PET bottle", Price: 100, Slot: "2"})

	v.ItemStock(&vending.Item{Name: "Sport drinks", Price: 150, Slot: "3"})
	v.ItemStock(&vending.Item{Name: "Sport drinks", Price: 150, Slot: "3"})
	v.ItemStock(&vending.Item{Name: "Sport drinks", Price: 150, Slot: "3"})
	v.ItemStock(&vending.Item{Name: "Sport drinks", Price: 150, Slot: "3"})
	v.ItemStock(&vending.Item{Name: "Sport drinks", Price: 150, Slot: "3"})

	v.Print()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		v.InputMap(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
