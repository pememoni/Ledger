package ledger

import (
	_ "errors"
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	ledger := initLedger()
	check := make(map[string]int)
	assert.Equal(t, ledger[0], check, "genesis block initiation: FAILED")
}

func TestSet(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	assert.Equal(t, ledger[0]["Alice"], 100, "writing user balances on chain: FAILED")
}

func TestSetNegative(t *testing.T) {
	ledger := initLedger()
	err := set(ledger, "Alice", -100)
	if err == nil {
		t.Error(err, "Throwing error for writing negative user balances: Failed")
	}
}

func TestIncrement(t *testing.T) {
	ledger := initLedger()
	ledger = increment(ledger)
	ledger = increment(ledger)
	assert.Equal(t, len(ledger), 3, "Adding new blocks to the ledger: FAILED")
}

func TestGet(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	balanceAlice, _ := get(ledger, "Alice", 0)

	balanceBob, _ := get(ledger, "Bob", 0)

	assert.Equal(t, balanceAlice, 100, "Reading users balances from the current block: FAILED")
	assert.Equal(t, balanceBob, 50, "Reading users balances from the current block: FAILED")
}

func TestGetNonExistent(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	balanceStranger, _ := get(ledger, "Stranger", 0)
	assert.Equal(t, balanceStranger, 0, "Setting non-existent user balance to zero: FAILED")
}

func TestGetFutureBlock(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	_, err := get(ledger, "Alice", 1)
	if err == nil {
		t.Error(err, "Throwing error for reading future blocks: Failed")
	}
}

func TestUpdateBalance(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	ledger = increment(ledger)
	_ = set(ledger, "Alice", 70)

	assert.Equal(t, ledger[1]["Alice"], 70, "Updating users balance in the current block: FAILED")
	assert.Equal(t, ledger[0]["Alice"], 100, "Preserving users balance in previous blocks while updating current balance: FAILED")

}

func TestTx(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	_ = tx(ledger, "Alice", "Bob", 25)

	AliceBalance, _ := get(ledger, "Alice", 0)
	BobBalance, _ := get(ledger, "Bob", 0)

	assert.Equal(t, AliceBalance, 75, "Updating sender balance after transaction: FAILED")
	assert.Equal(t, BobBalance, 75, "Updating receiver balance after transaction: FAILED")

}

func TestInvalidTx(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	err := tx(ledger, "Alice", "Bob", -25)
	if err == nil {
		t.Error(err, "Throwing error for negative amount in a tx: Failed")
	}
	err = tx(ledger, "Alice", "Bob", 101)
	if err == nil {
		t.Error(err, "Throwing error for spending more than balance: Failed")
	}

}

func TestDemo(t *testing.T) {
	ledger := initLedger()
	_ = set(ledger, "Alice", 100)
	_ = set(ledger, "Bob", 50)
	ledger = increment(ledger)
	_ = set(ledger, "Alice", 70)
	_ = set(ledger, "Charlie", 30)
	ledger = increment(ledger)
	_ = set(ledger, "Bob", 30)
	_ = set(ledger, "Daniel", 20)

	Alice0, _ := get(ledger, "Alice", 0)
	Bob0, _ := get(ledger, "Bob", 0)
	Alice1, _ := get(ledger, "Alice", 1)
	Bob1, _ := get(ledger, "Bob", 1)
	Charlie1, _ := get(ledger, "Charlie", 1)
	Alice2, _ := get(ledger, "Alice", 2)
	Bob2, _ := get(ledger, "Bob", 2)
	Charlie2, _ := get(ledger, "Charlie", 2)
	Daniel2, _ := get(ledger, "Daniel", 2)
	_ = tx(ledger, "Alice", "Daniel", 30)
	DanielAfterTx, _ := get(ledger, "Daniel", 2)
	AliceAfterTx, _ := get(ledger, "Alice", 2)

	assert.Equal(t, Alice0, 100, "Updating Alice balance in demo: FAILED")
	assert.Equal(t, Alice1, 70, "Updating Alice balance in demo: FAILED")
	assert.Equal(t, Alice2, 70, "Updating Alice balance in demo: FAILED")
	assert.Equal(t, Bob0, 50, "Updating Bob balance in demo: FAILED")
	assert.Equal(t, Bob1, 50, "Updating Bob balance in demo: FAILED")
	assert.Equal(t, Bob2, 30, "Updating Bob balance in demo: FAILED")
	assert.Equal(t, Charlie1, 30, "Updating Charlie balance in demo: FAILED")
	assert.Equal(t, Charlie2, 30, "Updating Charlie balance in demo: FAILED")
	assert.Equal(t, Daniel2, 20, "Updating Daniel balance in demo: FAILED")
	assert.Equal(t, DanielAfterTx, 50, "Updating Daniel balance after transaction in demo: FAILED")
	assert.Equal(t, AliceAfterTx, 40, "Updating Alice balance after transaction in demo: FAILED")

}
