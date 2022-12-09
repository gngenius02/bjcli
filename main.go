package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrayConv(str []string) (out []int) {
	out = make([]int, len(str))
	for i, v := range str {
		temp, _ := strconv.ParseInt(v, 10, 8)
		out[i] = int(temp)
	}
	return out
}

func main() {
	argv := os.Args
	argc := len(argv)
	command := argv[0]
	if argc < 3 {
		fmt.Printf("Usage:\n\t bjcli \"<player hand>\" <dealer card> [number of hands Default:1]\n")
		fmt.Printf("\n<player hand>:\t\t comma separated value of cards \n\t\t\t A: 1\n\t\t\t J-K: 10\n")
		fmt.Printf("\n<dealer card>:\t\t value of dealer card")
		fmt.Printf("\n[numebr of hands]:\t should always be 1 except after a split will be 2 regardless of which hand is being asked.\n\n")
		fmt.Printf("Examples:\n\t%s \"1,6\" 1 \t - This shows player A6 vs dealer A\n\n\t%s \"1,1\" 1 2\t - This shows player AA vs dealer A after a split\n", command, command)
		os.Exit(1)
	}

	playerHands := strings.Split(argv[1], ",")
	dealerCard, _ := strconv.ParseInt(argv[2], 10, 8)
	var handsCount int = 1
	if argc > 3 {
		temp, _ := strconv.ParseInt(argv[3], 10, 8)
		handsCount = int(temp)
	}
	b := NewBJ(arrayConv(playerHands), int(dealerCard), handsCount)

	result := b.BlackjackLogic()

	fmt.Printf("Player: %v\tDealer: %d\t\n\n", arrayConv(playerHands), int(dealerCard))
	fmt.Printf("result: %s\n", result)
}
