// Copyright 2016 The go-simplechain Authors
// This file is part of the go-simplechain library.
//
// The go-simplechain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-simplechain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-simplechain library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/simplechain-org/go-simplechain/common"
)

type ConfirmedMakerEvent struct {
	Txs []*CrossTransaction
}

type NewTakerEvent struct {
	Takers []*CrossTransactionModifier
}

type ConfirmedTakerEvent struct {
	Txs []*ReceptTransaction
}

type SignedCtxEvent struct { // pool event
	Tx       *CrossTransactionWithSignatures
	CallBack func(cws *CrossTransactionWithSignatures, invalidSigIndex ...int)
}

type NewFinishEvent struct {
	Finishes []*CrossTransactionModifier
}

type ConfirmedFinishEvent struct {
	Finishes []*CrossTransactionModifier
}

type NewAnchorEvent struct {
	ChainInfo []*RemoteChainInfo
}

type ModType uint8

const (
	Normal = ModType(iota)
	Remote
	Reorg
)

type CrossTransactionModifier struct {
	Type          ModType
	ID            common.Hash
	Status        CtxStatus
	AtBlockNumber uint64
}

type CrossBlockEvent struct {
	Number          *big.Int
	ConfirmedMaker  ConfirmedMakerEvent
	NewTaker        NewTakerEvent
	ConfirmedTaker  ConfirmedTakerEvent
	NewFinish       NewFinishEvent
	ConfirmedFinish ConfirmedFinishEvent
	NewAnchor       NewAnchorEvent
	ReorgTaker      NewTakerEvent
	ReorgFinish     NewFinishEvent
}

func (e CrossBlockEvent) IsEmpty() bool {
	return len(e.ConfirmedMaker.Txs)|len(e.ConfirmedTaker.Txs)|
		len(e.ConfirmedFinish.Finishes)|len(e.NewTaker.Takers)|
		len(e.NewFinish.Finishes)|len(e.NewAnchor.ChainInfo)|
		len(e.ReorgTaker.Takers)|len(e.ReorgFinish.Finishes) == 0
}
