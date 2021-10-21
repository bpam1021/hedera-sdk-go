// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/stream/account_balance_file.proto

//-
// ‌
// Hedera Network Services Protobuf
// ​
// Copyright (C) 2018 - 2021 Hedera Hashgraph, LLC
// ​
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
// ‍

package stream

import (
	proto "github.com/hashgraph/hedera-sdk-go/v2/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenUnitBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// A unique token id
	TokenId *proto.TokenID `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	//*
	// Number of transferable units of the identified token. For token of type FUNGIBLE_COMMON -
	// balance in the smallest denomination. For token of type NON_FUNGIBLE_UNIQUE - the number of
	// NFTs held by the account
	Balance uint64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *TokenUnitBalance) Reset() {
	*x = TokenUnitBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stream_account_balance_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenUnitBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenUnitBalance) ProtoMessage() {}

func (x *TokenUnitBalance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_account_balance_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenUnitBalance.ProtoReflect.Descriptor instead.
func (*TokenUnitBalance) Descriptor() ([]byte, []int) {
	return file_proto_stream_account_balance_file_proto_rawDescGZIP(), []int{0}
}

func (x *TokenUnitBalance) GetTokenId() *proto.TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *TokenUnitBalance) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

//*
// Includes all currency balances (both hbar and token) of a single account in the ledger.
type SingleAccountBalances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The account
	AccountID *proto.AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	//*
	// The account's hbar balance
	HbarBalance uint64 `protobuf:"varint,2,opt,name=hbarBalance,proto3" json:"hbarBalance,omitempty"`
	//*
	// The list of the account's token balances
	TokenUnitBalances []*TokenUnitBalance `protobuf:"bytes,3,rep,name=tokenUnitBalances,proto3" json:"tokenUnitBalances,omitempty"`
}

func (x *SingleAccountBalances) Reset() {
	*x = SingleAccountBalances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stream_account_balance_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleAccountBalances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleAccountBalances) ProtoMessage() {}

func (x *SingleAccountBalances) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_account_balance_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleAccountBalances.ProtoReflect.Descriptor instead.
func (*SingleAccountBalances) Descriptor() ([]byte, []int) {
	return file_proto_stream_account_balance_file_proto_rawDescGZIP(), []int{1}
}

func (x *SingleAccountBalances) GetAccountID() *proto.AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *SingleAccountBalances) GetHbarBalance() uint64 {
	if x != nil {
		return x.HbarBalance
	}
	return 0
}

func (x *SingleAccountBalances) GetTokenUnitBalances() []*TokenUnitBalance {
	if x != nil {
		return x.TokenUnitBalances
	}
	return nil
}

//*
// Includes all currency balances (both hbar and token) of all accounts in the ledger.
type AllAccountBalances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// An instant in consensus time
	ConsensusTimestamp *proto.Timestamp `protobuf:"bytes,1,opt,name=consensusTimestamp,proto3" json:"consensusTimestamp,omitempty"`
	//*
	// The list of account balances for all accounts, after handling all transactions with consensus
	// timestamp up to and including the above instant
	AllAccounts []*SingleAccountBalances `protobuf:"bytes,2,rep,name=allAccounts,proto3" json:"allAccounts,omitempty"`
}

func (x *AllAccountBalances) Reset() {
	*x = AllAccountBalances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stream_account_balance_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllAccountBalances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllAccountBalances) ProtoMessage() {}

func (x *AllAccountBalances) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_account_balance_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllAccountBalances.ProtoReflect.Descriptor instead.
func (*AllAccountBalances) Descriptor() ([]byte, []int) {
	return file_proto_stream_account_balance_file_proto_rawDescGZIP(), []int{2}
}

func (x *AllAccountBalances) GetConsensusTimestamp() *proto.Timestamp {
	if x != nil {
		return x.ConsensusTimestamp
	}
	return nil
}

func (x *AllAccountBalances) GetAllAccounts() []*SingleAccountBalances {
	if x != nil {
		return x.AllAccounts
	}
	return nil
}

var File_proto_stream_account_balance_file_proto protoreflect.FileDescriptor

var file_proto_stream_account_balance_file_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55,
	0x6e, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xc6,
	0x01, 0x0a, 0x15, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x62, 0x61, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x68,
	0x62, 0x61, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x40,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x54, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x53, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x01, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_stream_account_balance_file_proto_rawDescOnce sync.Once
	file_proto_stream_account_balance_file_proto_rawDescData = file_proto_stream_account_balance_file_proto_rawDesc
)

func file_proto_stream_account_balance_file_proto_rawDescGZIP() []byte {
	file_proto_stream_account_balance_file_proto_rawDescOnce.Do(func() {
		file_proto_stream_account_balance_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stream_account_balance_file_proto_rawDescData)
	})
	return file_proto_stream_account_balance_file_proto_rawDescData
}

var file_proto_stream_account_balance_file_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_stream_account_balance_file_proto_goTypes = []interface{}{
	(*TokenUnitBalance)(nil),      // 0: com.hedera.stream.api.proto.TokenUnitBalance
	(*SingleAccountBalances)(nil), // 1: com.hedera.stream.api.proto.SingleAccountBalances
	(*AllAccountBalances)(nil),    // 2: com.hedera.stream.api.proto.AllAccountBalances
	(*proto.TokenID)(nil),         // 3: proto.TokenID
	(*proto.AccountID)(nil),       // 4: proto.AccountID
	(*proto.Timestamp)(nil),       // 5: proto.Timestamp
}
var file_proto_stream_account_balance_file_proto_depIdxs = []int32{
	3, // 0: com.hedera.stream.api.proto.TokenUnitBalance.tokenId:type_name -> proto.TokenID
	4, // 1: com.hedera.stream.api.proto.SingleAccountBalances.accountID:type_name -> proto.AccountID
	0, // 2: com.hedera.stream.api.proto.SingleAccountBalances.tokenUnitBalances:type_name -> com.hedera.stream.api.proto.TokenUnitBalance
	5, // 3: com.hedera.stream.api.proto.AllAccountBalances.consensusTimestamp:type_name -> proto.Timestamp
	1, // 4: com.hedera.stream.api.proto.AllAccountBalances.allAccounts:type_name -> com.hedera.stream.api.proto.SingleAccountBalances
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_stream_account_balance_file_proto_init() }
func file_proto_stream_account_balance_file_proto_init() {
	if File_proto_stream_account_balance_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stream_account_balance_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenUnitBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stream_account_balance_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleAccountBalances); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_stream_account_balance_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllAccountBalances); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_stream_account_balance_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_stream_account_balance_file_proto_goTypes,
		DependencyIndexes: file_proto_stream_account_balance_file_proto_depIdxs,
		MessageInfos:      file_proto_stream_account_balance_file_proto_msgTypes,
	}.Build()
	File_proto_stream_account_balance_file_proto = out.File
	file_proto_stream_account_balance_file_proto_rawDesc = nil
	file_proto_stream_account_balance_file_proto_goTypes = nil
	file_proto_stream_account_balance_file_proto_depIdxs = nil
}