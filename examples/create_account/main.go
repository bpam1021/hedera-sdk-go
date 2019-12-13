package main

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	client, err := hedera.NewClient(
		// Node ID
		hedera.AccountID{Account: 3},
		// Node Address
		"0.testnet.hedera.com:50211",
	)

	if err != nil {
		panic(err)
	}

	operatorPrivateKey, err := hedera.Ed25519PrivateKeyFromString("302e020100300506032b657004220420db484b828e64b2d8f12ce3c0a0e93a0b8cce7af1bb8f39c97732394482538e10")

	if err != nil {
		panic(err)
	}

	client.SetOperator(
		// Operator Account ID
		hedera.AccountID{Account: 2},
		// Operator Private Key
		operatorPrivateKey,
	)

	newKey, err := hedera.GenerateEd25519PrivateKey()

	if err != nil {
		panic(err)
	}

	newPublicKey := newKey.PublicKey()

	tx, err := hedera.NewAccountCreateTransaction(client).
		SetKey(newPublicKey).
		SetInitialBalance(1000).
		SetMaxTransactionFee(10000000).
		Build()

	if err != nil {
		panic(err)
	}

	receipt, err := tx.ExecuteForReceipt()

	if err != nil {
		panic(err)
	}

	newAccountID := receipt.AccountID()

	fmt.Println(newAccountID.String())
}