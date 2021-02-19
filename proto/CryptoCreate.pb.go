// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: proto/CryptoCreate.proto

package proto

import (
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

//
//Create a new account. After the account is created, the AccountID for it is in the receipt. It can also be retrieved with a GetByKey query. Threshold values can be defined, and records are generated and stored for 25 hours for any transfer that exceeds the thresholds. This account is charged for each record generated, so the thresholds are useful for limiting record generation to happen only for large transactions.
//
//The Key field is the key used to sign transactions for this account. If the account has receiverSigRequired set to true, then all cryptocurrency transfers must be signed by this account's key, both for transfers in and out. If it is false, then only transfers out have to be signed by it. When the account is created, the payer account is charged enough hbars so that the new account will not expire for the next autoRenewPeriod seconds. When it reaches the expiration time, the new account will then be automatically charged to renew for another autoRenewPeriod seconds. If it does not have enough hbars to renew for that long, then the remaining hbars are used to extend its expiration as long as possible. If it is has a zero balance when it expires, then it is deleted. This transaction must be signed by the payer account. If receiverSigRequired is false, then the transaction does not have to be signed by the keys in the keys field. If it is true, then it must be signed by them, in addition to the keys of the payer account.
//An entity (account, file, or smart contract instance) must be created in a particular realm. If the realmID is left null, then a new realm will be created with the given admin key. If a new realm has a null adminKey, then anyone can create/modify/delete entities in that realm. But if an admin key is given, then any transaction to create/modify/delete an entity in that realm must be signed by that key, though anyone can still call functions on smart contract instances that exist in that realm. A realm ceases to exist when everything within it has expired and no longer exists.
//The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in shard 0 and realm 0, with a null key. Future versions of the API will support multiple realms and multiple shards.
type CryptoCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            *Key       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                        // The key that must sign each transfer out of the account. If receiverSigRequired is true, then it must also sign any transfer into the account.
	InitialBalance uint64     `protobuf:"varint,2,opt,name=initialBalance,proto3" json:"initialBalance,omitempty"` // The initial number of tinybars to put into the account
	ProxyAccountID *AccountID `protobuf:"bytes,3,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`  // ID of the account to which this account is proxy staked. If proxyAccountID is null, or is an invalid account, or is an account that isn't a node, then this account is automatically proxy staked to a node chosen by the network, but without earning payments. If the proxyAccountID account refuses to accept proxy staking , or if it is not currently running a node, then it will behave as if proxyAccountID was null.
	// Deprecated: Do not use.
	SendRecordThreshold uint64 `protobuf:"varint,6,opt,name=sendRecordThreshold,proto3" json:"sendRecordThreshold,omitempty"` // [Deprecated]. The threshold amount (in tinybars) for which an account record is created for any send/withdraw transaction
	// Deprecated: Do not use.
	ReceiveRecordThreshold uint64    `protobuf:"varint,7,opt,name=receiveRecordThreshold,proto3" json:"receiveRecordThreshold,omitempty"` // [Deprecated]. The threshold amount (in tinybars) for which an account record is created for any receive/deposit transaction
	ReceiverSigRequired    bool      `protobuf:"varint,8,opt,name=receiverSigRequired,proto3" json:"receiverSigRequired,omitempty"`       // If true, this account's key must sign any transaction depositing into this account (in addition to all withdrawals)
	AutoRenewPeriod        *Duration `protobuf:"bytes,9,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`                // The account is charged to extend its expiration date every this many seconds. If it doesn't have enough balance, it extends as long as possible. If it is empty when it expires, then it is deleted.
	ShardID                *ShardID  `protobuf:"bytes,10,opt,name=shardID,proto3" json:"shardID,omitempty"`                               // The shard in which this account is created
	RealmID                *RealmID  `protobuf:"bytes,11,opt,name=realmID,proto3" json:"realmID,omitempty"`                               // The realm in which this account is created (leave this null to create a new realm)
	NewRealmAdminKey       *Key      `protobuf:"bytes,12,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`             // If realmID is null, then this the admin key for the new realm that will be created
	Memo                   string    `protobuf:"bytes,13,opt,name=memo,proto3" json:"memo,omitempty"`                                     // The memo associated with the account (UTF-8 encoding max 100 bytes)
}

func (x *CryptoCreateTransactionBody) Reset() {
	*x = CryptoCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoCreate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoCreateTransactionBody) ProtoMessage() {}

func (x *CryptoCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoCreate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_CryptoCreate_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoCreateTransactionBody) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *CryptoCreateTransactionBody) GetInitialBalance() uint64 {
	if x != nil {
		return x.InitialBalance
	}
	return 0
}

func (x *CryptoCreateTransactionBody) GetProxyAccountID() *AccountID {
	if x != nil {
		return x.ProxyAccountID
	}
	return nil
}

// Deprecated: Do not use.
func (x *CryptoCreateTransactionBody) GetSendRecordThreshold() uint64 {
	if x != nil {
		return x.SendRecordThreshold
	}
	return 0
}

// Deprecated: Do not use.
func (x *CryptoCreateTransactionBody) GetReceiveRecordThreshold() uint64 {
	if x != nil {
		return x.ReceiveRecordThreshold
	}
	return 0
}

func (x *CryptoCreateTransactionBody) GetReceiverSigRequired() bool {
	if x != nil {
		return x.ReceiverSigRequired
	}
	return false
}

func (x *CryptoCreateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *CryptoCreateTransactionBody) GetShardID() *ShardID {
	if x != nil {
		return x.ShardID
	}
	return nil
}

func (x *CryptoCreateTransactionBody) GetRealmID() *RealmID {
	if x != nil {
		return x.RealmID
	}
	return nil
}

func (x *CryptoCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if x != nil {
		return x.NewRealmAdminKey
	}
	return nil
}

func (x *CryptoCreateTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

var File_proto_CryptoCreate_proto protoreflect.FileDescriptor

var file_proto_CryptoCreate_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9c, 0x04, 0x0a, 0x1b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x34, 0x0a, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x3a, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x18, 0x01, 0x52, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x28,
	0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44, 0x52,
	0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x44, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x49, 0x44, 0x12, 0x36, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x61,
	0x6c, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x42, 0x4b,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_CryptoCreate_proto_rawDescOnce sync.Once
	file_proto_CryptoCreate_proto_rawDescData = file_proto_CryptoCreate_proto_rawDesc
)

func file_proto_CryptoCreate_proto_rawDescGZIP() []byte {
	file_proto_CryptoCreate_proto_rawDescOnce.Do(func() {
		file_proto_CryptoCreate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CryptoCreate_proto_rawDescData)
	})
	return file_proto_CryptoCreate_proto_rawDescData
}

var file_proto_CryptoCreate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_CryptoCreate_proto_goTypes = []interface{}{
	(*CryptoCreateTransactionBody)(nil), // 0: proto.CryptoCreateTransactionBody
	(*Key)(nil),                         // 1: proto.Key
	(*AccountID)(nil),                   // 2: proto.AccountID
	(*Duration)(nil),                    // 3: proto.Duration
	(*ShardID)(nil),                     // 4: proto.ShardID
	(*RealmID)(nil),                     // 5: proto.RealmID
}
var file_proto_CryptoCreate_proto_depIdxs = []int32{
	1, // 0: proto.CryptoCreateTransactionBody.key:type_name -> proto.Key
	2, // 1: proto.CryptoCreateTransactionBody.proxyAccountID:type_name -> proto.AccountID
	3, // 2: proto.CryptoCreateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	4, // 3: proto.CryptoCreateTransactionBody.shardID:type_name -> proto.ShardID
	5, // 4: proto.CryptoCreateTransactionBody.realmID:type_name -> proto.RealmID
	1, // 5: proto.CryptoCreateTransactionBody.newRealmAdminKey:type_name -> proto.Key
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_CryptoCreate_proto_init() }
func file_proto_CryptoCreate_proto_init() {
	if File_proto_CryptoCreate_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_Duration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CryptoCreate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoCreateTransactionBody); i {
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
			RawDescriptor: file_proto_CryptoCreate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CryptoCreate_proto_goTypes,
		DependencyIndexes: file_proto_CryptoCreate_proto_depIdxs,
		MessageInfos:      file_proto_CryptoCreate_proto_msgTypes,
	}.Build()
	File_proto_CryptoCreate_proto = out.File
	file_proto_CryptoCreate_proto_rawDesc = nil
	file_proto_CryptoCreate_proto_goTypes = nil
	file_proto_CryptoCreate_proto_depIdxs = nil
}
