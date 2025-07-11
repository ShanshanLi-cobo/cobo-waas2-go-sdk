/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// SettleStatus The current status of settlement. - `Pending`: The settlement has been created and is awaiting processing. - `Processing`: The settlement is being processed. - `Completed`: The funds have been successfully deposited into the bank account. - `PartiallyCompleted`: Some settlement transactions have been completed successfully, while others have failed. - `Failed`: The settlement could not be completed due to an error. 
type SettleStatus string

// List of SettleStatus
const (
	SETTLESTATUS_PENDING SettleStatus = "Pending"
	SETTLESTATUS_PROCESSING SettleStatus = "Processing"
	SETTLESTATUS_COMPLETED SettleStatus = "Completed"
	SETTLESTATUS_PARTIALLY_COMPLETED SettleStatus = "PartiallyCompleted"
	SETTLESTATUS_FAILED SettleStatus = "Failed"
)

// All allowed values of SettleStatus enum
var AllowedSettleStatusEnumValues = []SettleStatus{
	"Pending",
	"Processing",
	"Completed",
	"PartiallyCompleted",
	"Failed",
}

func (v *SettleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SettleStatus(value)
	for _, existing := range AllowedSettleStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = SettleStatus("unknown")
	return nil
}

// NewSettleStatusFromValue returns a pointer to a valid SettleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSettleStatusFromValue(v string) (*SettleStatus, error) {
	ev := SettleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SettleStatus: valid values are %v", v, AllowedSettleStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SettleStatus) IsValid() bool {
	for _, existing := range AllowedSettleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SettleStatus value
func (v SettleStatus) Ptr() *SettleStatus {
	return &v
}

type NullableSettleStatus struct {
	value *SettleStatus
	isSet bool
}

func (v NullableSettleStatus) Get() *SettleStatus {
	return v.value
}

func (v *NullableSettleStatus) Set(val *SettleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSettleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSettleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettleStatus(val *SettleStatus) *NullableSettleStatus {
	return &NullableSettleStatus{value: val, isSet: true}
}

func (v NullableSettleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

