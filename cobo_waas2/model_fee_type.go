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

// FeeType The fee model. Possible values include: - `Fixed`: The fixed fee model.  - `EVM_EIP_1559`: The EIP-1559 fee model. - `EVM_Legacy`: The legacy fee model. - `UTXO`: The fee model used in UTXO-based blockchains, such as Bitcoin. - `SOL`: The fee model used in Solana. - `FIL`: The fee model used in Filecoin.   Each fee model requires a different set of properties. Switch between the above tabs for details.  To learn more about the fee models, refer to [Fee models](https://www.cobo.com/developers/v2/guides/transactions/estimate-fees#fee-models). 
type FeeType string

// List of FeeType
const (
	FEETYPE_FIXED FeeType = "Fixed"
	FEETYPE_EVM_EIP_1559 FeeType = "EVM_EIP_1559"
	FEETYPE_EVM_LEGACY FeeType = "EVM_Legacy"
	FEETYPE_UTXO FeeType = "UTXO"
	FEETYPE_SOL FeeType = "SOL"
	FEETYPE_FIL FeeType = "FIL"
)

// All allowed values of FeeType enum
var AllowedFeeTypeEnumValues = []FeeType{
	"Fixed",
	"EVM_EIP_1559",
	"EVM_Legacy",
	"UTXO",
	"SOL",
	"FIL",
}

func (v *FeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeeType(value)
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
	*v = FeeType("unknown")
	return nil
}

// NewFeeTypeFromValue returns a pointer to a valid FeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeTypeFromValue(v string) (*FeeType, error) {
	ev := FeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeeType: valid values are %v", v, AllowedFeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeeType) IsValid() bool {
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeeType value
func (v FeeType) Ptr() *FeeType {
	return &v
}

type NullableFeeType struct {
	value *FeeType
	isSet bool
}

func (v NullableFeeType) Get() *FeeType {
	return v.value
}

func (v *NullableFeeType) Set(val *FeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeType(val *FeeType) *NullableFeeType {
	return &NullableFeeType{value: val, isSet: true}
}

func (v NullableFeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

