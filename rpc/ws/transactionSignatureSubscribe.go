package ws

import (
	"context"
	"fmt"

	sjson "encoding/json"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var ErrInvalidParams = fmt.Errorf("invalid params")

type TransactionSubscribeMethodProvider string

const (
	TransactionSubscribeMethodProviderHelius TransactionSubscribeMethodProvider = "helius"
	TransactionSubscribeMethodProviderTriton TransactionSubscribeMethodProvider = "triton"
)

type TransactionSignatureSubscription struct {
	sub *Subscription
}

type TransactionSignatureResult struct {
	Transaction struct {
		Meta struct {
			LogMessages       []string           `json:"logMessages"`
			PreTokenBalances  []rpc.TokenBalance `json:"preTokenBalances"`
			PostTokenBalances []rpc.TokenBalance `json:"postTokenBalances"`
		} `json:"meta"`
	} `json:"transaction"`
	Signature solana.Signature `json:"signature"`
	Slot      uint64           `json:"slot"`
}

// TransactionSignatureResultG is a generic version of TransactionSignatureResult that allows for either regular or raw signature types
type TransactionSignatureResultG[S solana.Signature | solana.RawSolanaSignature, P solana.PublicKey | solana.RawPublicKey] struct {
	Transaction struct {
		Meta struct {
			LogMessages sjson.RawMessage `json:"logMessages"`
			// PreTokenBalances  []rpc.TokenBalanceG[P] `json:"preTokenBalances"`
			// PostTokenBalances []rpc.TokenBalanceG[P] `json:"postTokenBalances"`
		} `json:"meta"`
	} `json:"transaction"`
	Signature S      `json:"signature"`
	Slot      uint64 `json:"slot"`
}

// TransactionSignatureSubscribe subscribes to a transaction signature. Only Helius rpc nodes support this method.
func (cl *Client) TransactionSignatureSubscribe(
	accountInclude []string,
	accountRequired []string,
	methodProvider TransactionSubscribeMethodProvider,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscription, error) {
	params := make([]any, 0, 3)
	param := rpc.M{}
	switch methodProvider {
	case TransactionSubscribeMethodProviderHelius:
		if len(accountInclude) > 0 {
			param["accountInclude"] = accountInclude
		}
		if len(accountRequired) > 0 {
			param["accountRequired"] = accountRequired
		}
	case TransactionSubscribeMethodProviderTriton:
		accountsParam := rpc.M{}
		if len(accountInclude) > 0 {
			accountsParam["include"] = accountInclude
		}
		if len(accountRequired) > 0 {
			accountsParam["required"] = accountRequired
		}
		param["accounts"] = accountsParam
	default:
		return nil, ErrInvalidParams
	}
	if len(param) == 0 {
		return nil, ErrInvalidParams
	}
	param["vote"] = false
	param["failed"] = false
	params = append(params, param)
	return cl.transactionSignatureSubscribe(
		params,
		commitment,
	)
}

// TransactionSignatureSubscribeG is a generic version of TransactionSignatureSubscribe that allows for either regular or raw signature/pubkey types
func TransactionSignatureSubscribeG[S solana.Signature | solana.RawSolanaSignature, P solana.PublicKey | solana.RawPublicKey](
	cl *Client,
	accountInclude []string,
	accountRequired []string,
	methodProvider TransactionSubscribeMethodProvider,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscriptionG[S, P], error) {
	params := make([]any, 0, 3)
	param := rpc.M{}
	switch methodProvider {
	case TransactionSubscribeMethodProviderHelius:
		if len(accountInclude) > 0 {
			param["accountInclude"] = accountInclude
		}
		if len(accountRequired) > 0 {
			param["accountRequired"] = accountRequired
		}
	case TransactionSubscribeMethodProviderTriton:
		accountsParam := rpc.M{}
		if len(accountInclude) > 0 {
			accountsParam["include"] = accountInclude
		}
		if len(accountRequired) > 0 {
			accountsParam["required"] = accountRequired
		}
		param["accounts"] = accountsParam
	default:
		return nil, ErrInvalidParams
	}
	if len(param) == 0 {
		return nil, ErrInvalidParams
	}
	param["vote"] = false
	param["failed"] = false
	params = append(params, param)
	return transactionSignatureSubscribeG[S, P](
		cl,
		params,
		commitment,
	)
}

func (cl *Client) transactionSignatureSubscribe(
	params []any,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscription, error) {
	conf := map[string]interface{}{}
	conf["transactionDetails"] = "full"
	conf["maxSupportedTransactionVersion"] = 0
	if commitment != "" {
		conf["commitment"] = commitment
	}
	genSub, err := cl.subscribe(
		params,
		conf,
		"transactionSubscribe",
		"transactionUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res TransactionSignatureResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)

	if err != nil {
		return nil, err
	}
	return &TransactionSignatureSubscription{
		sub: genSub,
	}, nil
}

// transactionSignatureSubscribeG is a generic version of transactionSignatureSubscribe that allows for either regular or raw signature/pubkey types
func transactionSignatureSubscribeG[S solana.Signature | solana.RawSolanaSignature, P solana.PublicKey | solana.RawPublicKey](
	cl *Client,
	params []any,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscriptionG[S, P], error) {
	conf := map[string]interface{}{}
	conf["transactionDetails"] = "full"
	conf["maxSupportedTransactionVersion"] = 0
	if commitment != "" {
		conf["commitment"] = commitment
	}
	genSub, err := cl.subscribe(
		params,
		conf,
		"transactionSubscribe",
		"transactionUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res TransactionSignatureResultG[S, P]
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)

	if err != nil {
		return nil, err
	}
	return &TransactionSignatureSubscriptionG[S, P]{
		sub: genSub,
	}, nil
}

func (sw *TransactionSignatureSubscription) Recv(ctx context.Context) (*TransactionSignatureResult, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case d, ok := <-sw.sub.stream:
		if !ok {
			return nil, ErrSubscriptionClosed
		}
		return d.(*TransactionSignatureResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *TransactionSignatureSubscription) Err() <-chan error {
	return sw.sub.err
}

func (sw *TransactionSignatureSubscription) Response() <-chan *TransactionSignatureResult {
	typedChan := make(chan *TransactionSignatureResult, 1)
	go func(ch chan *TransactionSignatureResult) {
		// TODO: will this subscription yield more than one result?
		d, ok := <-sw.sub.stream
		if !ok {
			return
		}
		ch <- d.(*TransactionSignatureResult)
	}(typedChan)
	return typedChan
}

func (sw *TransactionSignatureSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}

// TransactionSignatureSubscriptionG is a generic version of TransactionSignatureSubscription that allows for either regular or raw signature/pubkey types
type TransactionSignatureSubscriptionG[S solana.Signature | solana.RawSolanaSignature, P solana.PublicKey | solana.RawPublicKey] struct {
	sub *Subscription
}

func (sw *TransactionSignatureSubscriptionG[S, P]) Recv(ctx context.Context) (*TransactionSignatureResultG[S, P], error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case d, ok := <-sw.sub.stream:
		if !ok {
			return nil, ErrSubscriptionClosed
		}
		return d.(*TransactionSignatureResultG[S, P]), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *TransactionSignatureSubscriptionG[S, P]) Err() <-chan error {
	return sw.sub.err
}

func (sw *TransactionSignatureSubscriptionG[S, P]) Response() <-chan *TransactionSignatureResultG[S, P] {
	typedChan := make(chan *TransactionSignatureResultG[S, P], 1)
	go func(ch chan *TransactionSignatureResultG[S, P]) {
		// TODO: will this subscription yield more than one result?
		d, ok := <-sw.sub.stream
		if !ok {
			return
		}
		ch <- d.(*TransactionSignatureResultG[S, P])
	}(typedChan)
	return typedChan
}

func (sw *TransactionSignatureSubscriptionG[S, P]) Unsubscribe() {
	sw.sub.Unsubscribe()
}
