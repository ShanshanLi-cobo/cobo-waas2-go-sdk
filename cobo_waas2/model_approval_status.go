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

// ApprovalStatus The approval status. Possible values include:    - `Pending`: The approval is pending.   - `Completed`: The approval is completed.    - `Failed`: The approval is rejected.   - `Revoked`: The approval is revoked.  
type ApprovalStatus string

// List of ApprovalStatus
const (
	APPROVALSTATUS_PENDING ApprovalStatus = "Pending"
	APPROVALSTATUS_COMPLETED ApprovalStatus = "Completed"
	APPROVALSTATUS_FAILED ApprovalStatus = "Failed"
	APPROVALSTATUS_REVOKED ApprovalStatus = "Revoked"
)

// All allowed values of ApprovalStatus enum
var AllowedApprovalStatusEnumValues = []ApprovalStatus{
	"Pending",
	"Completed",
	"Failed",
	"Revoked",
}

func (v *ApprovalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApprovalStatus(value)
	for _, existing := range AllowedApprovalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = ApprovalStatus("unknown")
	return nil
}

// NewApprovalStatusFromValue returns a pointer to a valid ApprovalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApprovalStatusFromValue(v string) (*ApprovalStatus, error) {
	ev := ApprovalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApprovalStatus: valid values are %v", v, AllowedApprovalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApprovalStatus) IsValid() bool {
	for _, existing := range AllowedApprovalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApprovalStatus value
func (v ApprovalStatus) Ptr() *ApprovalStatus {
	return &v
}

type NullableApprovalStatus struct {
	value *ApprovalStatus
	isSet bool
}

func (v NullableApprovalStatus) Get() *ApprovalStatus {
	return v.value
}

func (v *NullableApprovalStatus) Set(val *ApprovalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalStatus(val *ApprovalStatus) *NullableApprovalStatus {
	return &NullableApprovalStatus{value: val, isSet: true}
}

func (v NullableApprovalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

