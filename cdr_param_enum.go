// Code generated by go-enum
// DO NOT EDIT!

package qlcchain

import (
	"fmt"
	"strings"
)

const (
	// DLRStatusDelivered is a DLRStatus of type Delivered
	DLRStatusDelivered DLRStatus = iota
	// DLRStatusRejected is a DLRStatus of type Rejected
	DLRStatusRejected
	// DLRStatusUnknown is a DLRStatus of type Unknown
	DLRStatusUnknown
	// DLRStatusUndelivered is a DLRStatus of type Undelivered
	DLRStatusUndelivered
	// DLRStatusEmpty is a DLRStatus of type Empty
	DLRStatusEmpty
)

const _DLRStatusName = "DeliveredRejectedUnknownUndeliveredEmpty"

var _DLRStatusNames = []string{
	_DLRStatusName[0:9],
	_DLRStatusName[9:17],
	_DLRStatusName[17:24],
	_DLRStatusName[24:35],
	_DLRStatusName[35:40],
}

// DLRStatusNames returns a list of possible string values of DLRStatus.
func DLRStatusNames() []string {
	tmp := make([]string, len(_DLRStatusNames))
	copy(tmp, _DLRStatusNames)
	return tmp
}

var _DLRStatusMap = map[DLRStatus]string{
	0: _DLRStatusName[0:9],
	1: _DLRStatusName[9:17],
	2: _DLRStatusName[17:24],
	3: _DLRStatusName[24:35],
	4: _DLRStatusName[35:40],
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
	_DLRStatusName[9:17]:  1,
	_DLRStatusName[17:24]: 2,
	_DLRStatusName[24:35]: 3,
	_DLRStatusName[35:40]: 4,
}

// ParseDLRStatus attempts to convert a string to a DLRStatus
func ParseDLRStatus(name string) (DLRStatus, error) {
	if x, ok := _DLRStatusValue[name]; ok {
		return x, nil
	}
	return DLRStatus(0), fmt.Errorf("%s is not a valid DLRStatus, try [%s]", name, strings.Join(_DLRStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x DLRStatus) MarshalText() ([]byte, error) {
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
	// SendingStatusSent is a SendingStatus of type Sent
	SendingStatusSent SendingStatus = iota
	// SendingStatusError is a SendingStatus of type Error
	SendingStatusError
	// SendingStatusEmpty is a SendingStatus of type Empty
	SendingStatusEmpty
)

const _SendingStatusName = "SentErrorEmpty"

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
func (x SendingStatus) MarshalText() ([]byte, error) {
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