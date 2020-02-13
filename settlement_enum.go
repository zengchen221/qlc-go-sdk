// Code generated by go-enum
// DO NOT EDIT!

package qlcchain

import (
	"fmt"
	"strings"
)

const (
	// ContractStatusActiveStage1 is a ContractStatus of type ActiveStage1
	ContractStatusActiveStage1 ContractStatus = iota
	// ContractStatusActived is a ContractStatus of type Actived
	ContractStatusActived
	// ContractStatusDestroyStage1 is a ContractStatus of type DestroyStage1
	ContractStatusDestroyStage1
	// ContractStatusDestroyed is a ContractStatus of type Destroyed
	ContractStatusDestroyed
)

const _ContractStatusName = "ActiveStage1ActivedDestroyStage1Destroyed"

var _ContractStatusNames = []string{
	_ContractStatusName[0:12],
	_ContractStatusName[12:19],
	_ContractStatusName[19:32],
	_ContractStatusName[32:41],
}

// ContractStatusNames returns a list of possible string values of ContractStatus.
func ContractStatusNames() []string {
	tmp := make([]string, len(_ContractStatusNames))
	copy(tmp, _ContractStatusNames)
	return tmp
}

var _ContractStatusMap = map[ContractStatus]string{
	0: _ContractStatusName[0:12],
	1: _ContractStatusName[12:19],
	2: _ContractStatusName[19:32],
	3: _ContractStatusName[32:41],
}

// String implements the Stringer interface.
func (x ContractStatus) String() string {
	if str, ok := _ContractStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ContractStatus(%d)", x)
}

var _ContractStatusValue = map[string]ContractStatus{
	_ContractStatusName[0:12]:  0,
	_ContractStatusName[12:19]: 1,
	_ContractStatusName[19:32]: 2,
	_ContractStatusName[32:41]: 3,
}

// ParseContractStatus attempts to convert a string to a ContractStatus
func ParseContractStatus(name string) (ContractStatus, error) {
	if x, ok := _ContractStatusValue[name]; ok {
		return x, nil
	}
	return ContractStatus(0), fmt.Errorf("%s is not a valid ContractStatus, try [%s]", name, strings.Join(_ContractStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x *ContractStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ContractStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseContractStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// DLRStatusDelivered is a DLRStatus of type Delivered
	DLRStatusDelivered DLRStatus = iota
	// DLRStatusUnknown is a DLRStatus of type Unknown
	DLRStatusUnknown
	// DLRStatusUndelivered is a DLRStatus of type Undelivered
	DLRStatusUndelivered
	// DLRStatusEmpty is a DLRStatus of type Empty
	DLRStatusEmpty
)

const _DLRStatusName = "DeliveredUnknownUndeliveredEmpty"

var _DLRStatusNames = []string{
	_DLRStatusName[0:9],
	_DLRStatusName[9:16],
	_DLRStatusName[16:27],
	_DLRStatusName[27:32],
}

// DLRStatusNames returns a list of possible string values of DLRStatus.
func DLRStatusNames() []string {
	tmp := make([]string, len(_DLRStatusNames))
	copy(tmp, _DLRStatusNames)
	return tmp
}

var _DLRStatusMap = map[DLRStatus]string{
	0: _DLRStatusName[0:9],
	1: _DLRStatusName[9:16],
	2: _DLRStatusName[16:27],
	3: _DLRStatusName[27:32],
}

// String implements the Stringer interface.
func (x DLRStatus) String() string {
	if str, ok := _DLRStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("DLRStatus(%d)", x)
}

var _DLRStatusValue = map[string]DLRStatus{
	_DLRStatusName[0:9]:   0,
	_DLRStatusName[9:16]:  1,
	_DLRStatusName[16:27]: 2,
	_DLRStatusName[27:32]: 3,
}

// ParseDLRStatus attempts to convert a string to a DLRStatus
func ParseDLRStatus(name string) (DLRStatus, error) {
	if x, ok := _DLRStatusValue[name]; ok {
		return x, nil
	}
	return DLRStatus(0), fmt.Errorf("%s is not a valid DLRStatus, try [%s]", name, strings.Join(_DLRStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x *DLRStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *DLRStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseDLRStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// SendingStatusSend is a SendingStatus of type Send
	SendingStatusSend SendingStatus = iota
	// SendingStatusError is a SendingStatus of type Error
	SendingStatusError
	// SendingStatusEmpty is a SendingStatus of type Empty
	SendingStatusEmpty
)

const _SendingStatusName = "SendErrorEmpty"

var _SendingStatusNames = []string{
	_SendingStatusName[0:4],
	_SendingStatusName[4:9],
	_SendingStatusName[9:14],
}

// SendingStatusNames returns a list of possible string values of SendingStatus.
func SendingStatusNames() []string {
	tmp := make([]string, len(_SendingStatusNames))
	copy(tmp, _SendingStatusNames)
	return tmp
}

var _SendingStatusMap = map[SendingStatus]string{
	0: _SendingStatusName[0:4],
	1: _SendingStatusName[4:9],
	2: _SendingStatusName[9:14],
}

// String implements the Stringer interface.
func (x SendingStatus) String() string {
	if str, ok := _SendingStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("SendingStatus(%d)", x)
}

var _SendingStatusValue = map[string]SendingStatus{
	_SendingStatusName[0:4]:  0,
	_SendingStatusName[4:9]:  1,
	_SendingStatusName[9:14]: 2,
}

// ParseSendingStatus attempts to convert a string to a SendingStatus
func ParseSendingStatus(name string) (SendingStatus, error) {
	if x, ok := _SendingStatusValue[name]; ok {
		return x, nil
	}
	return SendingStatus(0), fmt.Errorf("%s is not a valid SendingStatus, try [%s]", name, strings.Join(_SendingStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x *SendingStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *SendingStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseSendingStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// SettlementStatusUnknown is a SettlementStatus of type Unknown
	SettlementStatusUnknown SettlementStatus = iota
	// SettlementStatusStage1 is a SettlementStatus of type Stage1
	SettlementStatusStage1
	// SettlementStatusSuccess is a SettlementStatus of type Success
	SettlementStatusSuccess
	// SettlementStatusFailure is a SettlementStatus of type Failure
	SettlementStatusFailure
	// SettlementStatusMissing is a SettlementStatus of type Missing
	SettlementStatusMissing
	// SettlementStatusDuplicate is a SettlementStatus of type Duplicate
	SettlementStatusDuplicate
)

const _SettlementStatusName = "unknownstage1successfailuremissingduplicate"

var _SettlementStatusNames = []string{
	_SettlementStatusName[0:7],
	_SettlementStatusName[7:13],
	_SettlementStatusName[13:20],
	_SettlementStatusName[20:27],
	_SettlementStatusName[27:34],
	_SettlementStatusName[34:43],
}

// SettlementStatusNames returns a list of possible string values of SettlementStatus.
func SettlementStatusNames() []string {
	tmp := make([]string, len(_SettlementStatusNames))
	copy(tmp, _SettlementStatusNames)
	return tmp
}

var _SettlementStatusMap = map[SettlementStatus]string{
	0: _SettlementStatusName[0:7],
	1: _SettlementStatusName[7:13],
	2: _SettlementStatusName[13:20],
	3: _SettlementStatusName[20:27],
	4: _SettlementStatusName[27:34],
	5: _SettlementStatusName[34:43],
}

// String implements the Stringer interface.
func (x SettlementStatus) String() string {
	if str, ok := _SettlementStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("SettlementStatus(%d)", x)
}

var _SettlementStatusValue = map[string]SettlementStatus{
	_SettlementStatusName[0:7]:   0,
	_SettlementStatusName[7:13]:  1,
	_SettlementStatusName[13:20]: 2,
	_SettlementStatusName[20:27]: 3,
	_SettlementStatusName[27:34]: 4,
	_SettlementStatusName[34:43]: 5,
}

// ParseSettlementStatus attempts to convert a string to a SettlementStatus
func ParseSettlementStatus(name string) (SettlementStatus, error) {
	if x, ok := _SettlementStatusValue[name]; ok {
		return x, nil
	}
	return SettlementStatus(0), fmt.Errorf("%s is not a valid SettlementStatus, try [%s]", name, strings.Join(_SettlementStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x *SettlementStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *SettlementStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseSettlementStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
