package hedera

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScheduleCreateTransaction_Execute(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	tx, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		SetNodeAccountIDs([]AccountID{{Account: 3}}).
		FreezeWith(client)
	assert.NoError(t, err)

	scheduleTx := tx.Schedule()

	frozen, err := scheduleTx.SetNodeAccountIDs([]AccountID{{Account: 3}}).SetAdminKey(newKey.PublicKey()).FreezeWith(client)
	assert.NoError(t, err)

	frozen = frozen.Sign(newKey)

	txID, err := frozen.Execute(client)
	assert.NoError(t, err)

	receipt, err := txID.GetReceipt(client)
	assert.NoError(t, err)

	//println(receipt.)

	_, err = NewScheduleInfoQuery().
		SetScheduleID(*receipt.ScheduleID).
		SetMaxQueryPayment(NewHbar(2)).
		Execute(client)
	assert.NoError(t, err)

	//println(info.)

	tx2, err := NewScheduleDeleteTransaction().
		SetScheduleID(*receipt.ScheduleID).
		FreezeWith(client)
	assert.NoError(t, err)

	txID, err = tx2.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = txID.GetReceipt(client)
	assert.NoError(t, err)
}

func TestScheduleCreateTransaction_SetTransaction_Execute(t *testing.T) {
	client := newTestClient(t)

	newKey, err := GeneratePrivateKey()
	assert.NoError(t, err)

	newBalance := NewHbar(1)

	assert.Equal(t, HbarUnits.Hbar.numberOfTinybar(), newBalance.tinybar)

	tx, err := NewAccountCreateTransaction().
		SetKey(newKey.PublicKey()).
		SetMaxTransactionFee(NewHbar(2)).
		SetInitialBalance(newBalance).
		SetNodeAccountIDs([]AccountID{{Account: 3}}).
		FreezeWith(client)
	assert.NoError(t, err)

	scheduleTx, err := NewScheduleCreateTransaction().SetAdminKey(newKey.PublicKey()).SetTransaction(tx.Transaction).FreezeWith(client)
	assert.NoError(t, err)

	scheduleTx = scheduleTx.Sign(newKey)

	txID, err := scheduleTx.Execute(client)
	assert.NoError(t, err)

	receipt, err := txID.GetReceipt(client)
	assert.NoError(t, err)

	//println(receipt.)

	_, err = NewScheduleInfoQuery().
		SetScheduleID(*receipt.ScheduleID).
		SetMaxQueryPayment(NewHbar(2)).
		Execute(client)
	assert.NoError(t, err)

	//println(info.)

	tx2, err := NewScheduleDeleteTransaction().
		SetScheduleID(*receipt.ScheduleID).
		FreezeWith(client)
	assert.NoError(t, err)

	txID, err = tx2.
		Sign(newKey).
		Execute(client)
	assert.NoError(t, err)

	_, err = txID.GetReceipt(client)
	assert.NoError(t, err)
}
