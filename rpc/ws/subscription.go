// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ws

import (
	"context"
	"sync"
)

type Subscription struct {
	req               *request
	subID             uint64
	stream            chan result
	err               chan error
	closeFunc         func(err error)
	mu                sync.Mutex
	closed            bool
	unsubscribeMethod string
	decoderFunc       decoderFunc
}

type decoderFunc func([]byte) (any, error)

func newSubscription(
	req *request,
	closeFunc func(err error),
	unsubscribeMethod string,
	decoderFunc decoderFunc,
) *Subscription {
	return &Subscription{
		req:               req,
		subID:             0,
		stream:            make(chan result, 20000),
		err:               make(chan error, 100),
		closeFunc:         closeFunc,
		unsubscribeMethod: unsubscribeMethod,
		decoderFunc:       decoderFunc,
	}
}

func (s *Subscription) Recv(ctx context.Context) (any, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case d, ok := <-s.stream:
		if !ok {
			return nil, ErrSubscriptionClosed
		}
		return d, nil
	case err, ok := <-s.err:
		if !ok {
			return nil, ErrSubscriptionClosed
		}
		return nil, err
	}
}

func (s *Subscription) Unsubscribe() {
	s.closeFunc(ErrSubscriptionClosed)
}
