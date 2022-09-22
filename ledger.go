package ledger

import (
	"errors"
)

func initLedger() []map[string]int {
	var ledger []map[string]int
	genesisBlock := make(map[string]int)
	ledger = append(ledger, genesisBlock)
	return ledger
}

func set(ledger []map[string]int, uid string, balance int) error {
	if balance < 0 {
		return errors.New("user balance cannot be negative")
	}
	height := len(ledger) - 1
	ledger[height][uid] = balance
	return nil
}

func increment(ledger []map[string]int) []map[string]int {
	latestBlock := ledger[len(ledger)-1]
	newBlock := make(map[string]int, len(latestBlock))
	for k, v := range latestBlock {
		newBlock[k] = v
	}
	ledger = append(ledger, newBlock)
	return ledger
}

func get(ledger []map[string]int, uid string, height int) (int, error) {
	if height > len(ledger)-1 {
		return -1, errors.New("this block height does not exist")
	}
	//fmt.Println("Block", height, "\n", uid, ":", ledger[height][uid])
	return ledger[height][uid], nil
}

func tx(ledger []map[string]int, uid_source string, uid_dest string, amount int) error {
	if amount < 0 {
		return errors.New("tx amount cannot be negative")
	}
	height := len(ledger) - 1
	if amount > ledger[height][uid_source] {
		return errors.New("tx sender does not have enough balance")
	}
	ledger[height][uid_source] -= amount
	ledger[height][uid_dest] += amount
	return nil
}
