package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type TransferTransaction struct {
	TransactionBuilder
	pb             *proto.CryptoTransferTransactionBody
	tokenIdIndexes map[string]int
}

// NewTransferTransaction creates a TransferTransaction builder which can be
// used to construct and execute a Token Transfers Transaction.
//
// Deprecated: Use `TransferTransaction` instead
func NewTransferTransaction() TransferTransaction {
	pb := &proto.CryptoTransferTransactionBody{
		Transfers: &proto.TransferList{
			AccountAmounts: []*proto.AccountAmount{},
		},
		TokenTransfers: make([]*proto.TokenTransferList, 0),
	}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_CryptoTransfer{CryptoTransfer: pb}

	builder := TransferTransaction{inner, pb, make(map[string]int)}

	return builder
}

// AddSender adds an account and the amount of hbar (as a positive value) to be sent from the sender. If any sender
// account fails to have a sufficient balance to do the withdrawal, then the entire transaction fails, and none of those
// transfers occur, though the transaction fee is still charged.
func (builder TransferTransaction) AddHbarSender(id AccountID, amount Hbar) TransferTransaction {
	return builder.AddHbarTransfer(id, amount.negated())
}

// AddRecipient adds a recipient account and the amount of hbar to be received from the sender(s).
func (builder TransferTransaction) AddHbarRecipient(id AccountID, amount Hbar) TransferTransaction {
	return builder.AddHbarTransfer(id, amount)
}

// AddTransfer adds the accountID to the internal accounts list and the amounts to the internal amounts list. Each
// negative amount is withdrawn from the corresponding account (a sender), and each positive one is added to the
// corresponding account (a receiver). The amounts list must sum to zero and there can be a maximum of 10 transfers.
//
// AddSender and AddRecipient are provided as convenience wrappers around AddTransfer.
func (builder TransferTransaction) AddHbarTransfer(id AccountID, amount Hbar) TransferTransaction {
	builder.pb.Transfers.AccountAmounts = append(builder.pb.Transfers.AccountAmounts, &proto.AccountAmount{
		AccountID: id.toProto(),
		Amount:    amount.AsTinybar(),
	})

	return builder
}

func (builder TransferTransaction) AddTokenSender(tokenID TokenID, accountID AccountID, amount uint64) TransferTransaction {
	return builder.AddTokenTransfers(tokenID, accountID, -int64(amount))
}

func (builder TransferTransaction) AddTokenRecipient(tokenID TokenID, accountID AccountID, amount uint64) TransferTransaction {
	return builder.AddTokenTransfers(tokenID, accountID, int64(amount))
}

func (builder TransferTransaction) AddTokenTransfers(tokenID TokenID, accountID AccountID, amount int64) TransferTransaction {
	index, ok := builder.tokenIdIndexes[tokenID.String()]

	if !ok {
		index := len(builder.pb.TokenTransfers)
		builder.tokenIdIndexes[tokenID.String()] = index

		builder.pb.TokenTransfers = append(builder.pb.TokenTransfers, &proto.TokenTransferList{
			Token:     tokenID.toProto(),
			Transfers: make([]*proto.AccountAmount, 0),
		})

	}

	builder.pb.TokenTransfers[index].Transfers = append(builder.pb.TokenTransfers[index].Transfers, &proto.AccountAmount{
		AccountID: accountID.toProto(),
		Amount:    amount,
	})

	return builder
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

// SetMaxTransactionFee sets the max transaction fee for this Transaction.
func (builder TransferTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) TransferTransaction {
	return TransferTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb, builder.tokenIdIndexes}
}

// SetTransactionMemo sets the memo for this Transaction.
func (builder TransferTransaction) SetTransactionMemo(memo string) TransferTransaction {
	return TransferTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb, builder.tokenIdIndexes}
}

// SetTransactionValidDuration sets the valid duration for this Transaction.
func (builder TransferTransaction) SetTransactionValidDuration(validDuration time.Duration) TransferTransaction {
	return TransferTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb, builder.tokenIdIndexes}
}

// SetTransactionID sets the TransactionID for this Transaction.
func (builder TransferTransaction) SetTransactionID(transactionID TransactionID) TransferTransaction {
	return TransferTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb, builder.tokenIdIndexes}
}

// SetNodeAccountID sets the node AccountID for this Transaction.
func (builder TransferTransaction) SetNodeAccountID(nodeAccountID AccountID) TransferTransaction {
	return TransferTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb, builder.tokenIdIndexes}
}