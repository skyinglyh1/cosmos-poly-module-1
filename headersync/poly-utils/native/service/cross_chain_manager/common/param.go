/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package common

import (
	"fmt"
	"github.com/polynetwork/cosmos-poly-module/headersync/poly-utils/common"
)

type MakeTxParam struct {
	TxHash              []byte
	CrossChainID        []byte
	FromContractAddress []byte
	ToChainID           uint64
	ToContractAddress   []byte
	Method              string
	Args                []byte
}

func (this *MakeTxParam) Serialization(sink *common.ZeroCopySink) {
	sink.WriteVarBytes(this.TxHash)
	sink.WriteVarBytes(this.CrossChainID)
	sink.WriteVarBytes(this.FromContractAddress)
	sink.WriteUint64(this.ToChainID)
	sink.WriteVarBytes(this.ToContractAddress)
	sink.WriteVarBytes([]byte(this.Method))
	sink.WriteVarBytes(this.Args)
}

func (this *MakeTxParam) Deserialization(source *common.ZeroCopySource) error {
	txHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize txHash error")
	}
	crossChainID, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize crossChainID error")
	}
	fromContractAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize fromContractAddress error")
	}
	toChainID, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize toChainID error")
	}
	toContractAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize toContractAddress error")
	}
	method, eof := source.NextString()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize method error")
	}
	args, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MakeTxParam deserialize args error")
	}

	this.TxHash = txHash
	this.CrossChainID = crossChainID
	this.FromContractAddress = fromContractAddress
	this.ToChainID = toChainID
	this.ToContractAddress = toContractAddress
	this.Method = method
	this.Args = args
	return nil
}

type ToMerkleValue struct {
	TxHash      []byte
	FromChainID uint64
	MakeTxParam *MakeTxParam
}

func (this *ToMerkleValue) Serialization(sink *common.ZeroCopySink) {
	sink.WriteVarBytes(this.TxHash)
	sink.WriteUint64(this.FromChainID)
	this.MakeTxParam.Serialization(sink)
}

func (this *ToMerkleValue) Deserialization(source *common.ZeroCopySource) error {
	txHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("MerkleValue deserialize txHash error")
	}
	fromChainID, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("MerkleValue deserialize fromChainID error")
	}

	makeTxParam := new(MakeTxParam)
	err := makeTxParam.Deserialization(source)
	if err != nil {
		return fmt.Errorf("MerkleValue deserialize makeTxParam error:%s", err)
	}

	this.TxHash = txHash
	this.FromChainID = fromChainID
	this.MakeTxParam = makeTxParam
	return nil
}
