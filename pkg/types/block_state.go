package types

import (
	"bytes"
	"encoding/json"

	"github.com/qlcchain/qlc-go-sdk/pkg/util"
)

//go:generate msgp
type StateBlock struct {
	Type           BlockType `msg:"type" json:"type"`
	Token          Hash      `msg:"token,extension" json:"token"`
	Address        Address   `msg:"address,extension" json:"address"`
	Balance        Balance   `msg:"balance,extension" json:"balance"`
	Vote           Balance   `msg:"vote,extension" json:"vote"`
	Network        Balance   `msg:"network,extension" json:"network"`
	Storage        Balance   `msg:"storage,extension" json:"storage"`
	Oracle         Balance   `msg:"oracle,extension" json:"oracle"`
	Previous       Hash      `msg:"previous,extension" json:"previous"`
	Link           Hash      `msg:"link,extension" json:"link"`
	Sender         []byte    `msg:"sender" json:"sender,omitempty"`
	Receiver       []byte    `msg:"receiver" json:"receiver,omitempty"`
	Message        Hash      `msg:"message,extension" json:"message,omitempty"`
	Data           []byte    `msg:"data" json:"data,omitempty"`
	PoVHeight      uint64    `msg:"povHeight" json:"povHeight"`
	Timestamp      int64     `msg:"timestamp" json:"timestamp"`
	Extra          Hash      `msg:"extra,extension" json:"extra,omitempty"`
	Representative Address   `msg:"representative,extension" json:"representative"`

	PrivateFrom    string   `msg:"priFrom,omitempty" json:"privateFrom,omitempty"`
	PrivateFor     []string `msg:"priFor,omitempty" json:"privateFor,omitempty"`
	PrivateGroupID string   `msg:"priGid,omitempty" json:"privateGroupID,omitempty"`

	Work      Work      `msg:"work,extension" json:"work"`
	Signature Signature `msg:"signature,extension" json:"signature"`
}

func (b *StateBlock) BuildHashData() []byte {
	buf := new(bytes.Buffer)

	// fields for public txs
	buf.WriteByte(byte(b.Type))
	buf.Write(b.Token[:])
	buf.Write(b.Address[:])
	buf.Write(b.Balance.Bytes())
	buf.Write(b.Vote.Bytes())
	buf.Write(b.Network.Bytes())
	buf.Write(b.Storage.Bytes())
	buf.Write(b.Oracle.Bytes())
	buf.Write(b.Previous[:])
	buf.Write(b.Link[:])
	buf.Write(b.Sender)
	buf.Write(b.Receiver)
	buf.Write(b.Message[:])
	buf.Write(b.Data)
	buf.Write(util.BE_Int2Bytes(b.Timestamp))
	buf.Write(util.BE_Uint64ToBytes(b.PoVHeight))
	buf.Write(b.Extra[:])
	buf.Write(b.Representative[:])

	// additional fields for private txs
	if len(b.PrivateFrom) > 0 {
		buf.WriteString(b.PrivateFrom)
	}
	if len(b.PrivateFor) > 0 {
		for _, pf := range b.PrivateFor {
			if len(pf) > 0 {
				buf.WriteString(pf)
			}
		}
	}
	if len(b.PrivateGroupID) > 0 {
		buf.WriteString(b.PrivateGroupID)
	}

	return buf.Bytes()
}

func (b *StateBlock) GetHash() Hash {
	data := b.BuildHashData()
	hash := HashData(data)
	return hash
}

func (b *StateBlock) GetHashWithoutPrivacy() Hash {
	t := []byte{byte(b.Type)}
	hash, _ := HashBytes(t, b.Token[:], b.Address[:], b.Balance.Bytes(), b.Vote.Bytes(), b.Network.Bytes(),
		b.Storage.Bytes(), b.Oracle.Bytes(), b.Previous[:], b.Link[:], b.Sender, b.Receiver, b.Message[:], b.Data,
		util.BE_Int2Bytes(b.Timestamp), util.BE_Uint64ToBytes(b.PoVHeight),
		b.Extra[:], b.Representative[:])
	return hash
}

func (b *StateBlock) GetType() BlockType {
	return b.Type
}

func (b *StateBlock) GetToken() Hash {
	return b.Token
}

func (b *StateBlock) GetAddress() Address {
	return b.Address
}

func (b *StateBlock) GetPrevious() Hash {
	return b.Previous
}

func (b *StateBlock) GetBalance() Balance {
	return b.Balance
}

func (b *StateBlock) GetVote() Balance {
	if b.Vote.Int == nil {
		return ZeroBalance
	}
	return b.Vote
}

func (b *StateBlock) GetOracle() Balance {
	if b.Oracle.Int == nil {
		return ZeroBalance
	}
	return b.Oracle
}

func (b *StateBlock) GetNetwork() Balance {
	if b.Network.Int == nil {
		return ZeroBalance
	}
	return b.Network
}

func (b *StateBlock) GetStorage() Balance {
	if b.Storage.Int == nil {
		return ZeroBalance
	}
	return b.Storage
}

func (b *StateBlock) GetData() []byte {
	return b.Data
}

func (b *StateBlock) GetLink() Hash {
	return b.Link
}

func (b *StateBlock) GetSignature() Signature {
	return b.Signature
}

func (b *StateBlock) GetWork() Work {
	return b.Work
}

func (b *StateBlock) GetExtra() Hash {
	return b.Extra
}

func (b *StateBlock) GetRepresentative() Address {
	return b.Representative
}

func (b *StateBlock) GetReceiver() []byte {
	return b.Receiver
}

func (b *StateBlock) GetSender() []byte {
	return b.Sender
}

func (b *StateBlock) GetMessage() Hash {
	return b.Message
}

func (b *StateBlock) GetTimestamp() int64 {
	return b.Timestamp
}

func (b *StateBlock) TotalBalance() Balance {
	balance := b.Balance
	if b.Vote.Int != nil {
		balance = balance.Add(b.Vote)
	}
	if b.Network.Int != nil {
		balance = balance.Add(b.Network)
	}
	if b.Oracle.Int != nil {
		balance = balance.Add(b.Oracle)
	}
	if b.Storage.Int != nil {
		balance = balance.Add(b.Storage)
	}
	return balance
}

func (b *StateBlock) IsOpen() bool {
	return b.Previous.IsZero()
}

func (b *StateBlock) Root() Hash {
	if b.IsOpen() {
		return b.Address.ToHash()
	}
	return b.Previous
}

func (b *StateBlock) Parent() Hash {
	if b.IsOpen() {
		return b.Link
	}
	return b.Previous
}

func (b *StateBlock) Size() int {
	return b.Msgsize()
}

func (b *StateBlock) IsValid() bool {
	if b.IsOpen() {
		return b.Work.IsValid(Hash(b.Address))
	}
	return b.Work.IsValid(b.Previous)
}

func (b *StateBlock) Serialize() ([]byte, error) {
	return b.MarshalMsg(nil)
}

func (b *StateBlock) Deserialize(text []byte) error {
	_, err := b.UnmarshalMsg(text)
	if err != nil {
		return err
	}
	return nil
}

func (b *StateBlock) String() string {
	bytes, _ := json.Marshal(b)
	return string(bytes)
}

func (b *StateBlock) IsReceiveBlock() bool {
	return b.Type == Receive || b.Type == Open || b.Type == ContractReward
}

func (b *StateBlock) IsSendBlock() bool {
	return b.Type == Send || b.Type == ContractSend
}

func (b *StateBlock) IsContractBlock() bool {
	return b.Type == ContractReward || b.Type == ContractSend || b.Type == ContractRefund || b.Type == ContractError
}

func (b *StateBlock) Clone() *StateBlock {
	clone := StateBlock{}
	bytes, _ := b.Serialize()
	_ = clone.Deserialize(bytes)
	return &clone
}

func (b *StateBlock) IsPrivate() bool {
	if len(b.PrivateFrom) > 0 {
		return true
	}
	return false
}

type StateBlockList []*StateBlock

func (bs *StateBlockList) Serialize() ([]byte, error) {
	return bs.MarshalMsg(nil)
}

func (bs *StateBlockList) Deserialize(text []byte) error {
	_, err := bs.UnmarshalMsg(text)
	if err != nil {
		return err
	}
	return nil
}

//
////go:generate msgp
//type BlockExtra struct {
//	KeyHash Hash    `msg:"key,extension" json:"key"`
//	Abi     []byte  `msg:"abi" json:"abi"`
//	Issuer  Address `msg:"issuer,extension" json:"issuer"`
//}
