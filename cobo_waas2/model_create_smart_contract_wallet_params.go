/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"fmt"
)

// CreateSmartContractWalletParams - struct for CreateSmartContractWalletParams
type CreateSmartContractWalletParams struct {
	CreateSafeWalletParams *CreateSafeWalletParams
}

// CreateSafeWalletParamsAsCreateSmartContractWalletParams is a convenience function that returns CreateSafeWalletParams wrapped in CreateSmartContractWalletParams
func CreateSafeWalletParamsAsCreateSmartContractWalletParams(v *CreateSafeWalletParams) CreateSmartContractWalletParams {
	return CreateSmartContractWalletParams{
		CreateSafeWalletParams: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateSmartContractWalletParams) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'Safe'
	if jsonDict["smart_contract_wallet_type"] == "Safe" {
		// try to unmarshal JSON data into CreateSafeWalletParams
		err = json.Unmarshal(data, &dst.CreateSafeWalletParams)
		if err == nil {
			return nil // data stored in dst.CreateSafeWalletParams, return on the first match
		} else {
			dst.CreateSafeWalletParams = nil
			return fmt.Errorf("failed to unmarshal CreateSmartContractWalletParams as CreateSafeWalletParams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CreateSafeWalletParams'
	if jsonDict["smart_contract_wallet_type"] == "CreateSafeWalletParams" {
		// try to unmarshal JSON data into CreateSafeWalletParams
		err = json.Unmarshal(data, &dst.CreateSafeWalletParams)
		if err == nil {
			return nil // data stored in dst.CreateSafeWalletParams, return on the first match
		} else {
			dst.CreateSafeWalletParams = nil
			return fmt.Errorf("failed to unmarshal CreateSmartContractWalletParams as CreateSafeWalletParams: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSmartContractWalletParams) MarshalJSON() ([]byte, error) {
	if src.CreateSafeWalletParams != nil {
		return json.Marshal(&src.CreateSafeWalletParams)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateSmartContractWalletParams) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateSafeWalletParams != nil {
		return obj.CreateSafeWalletParams
	}

	// all schemas are nil
	return nil
}

type NullableCreateSmartContractWalletParams struct {
	value *CreateSmartContractWalletParams
	isSet bool
}

func (v NullableCreateSmartContractWalletParams) Get() *CreateSmartContractWalletParams {
	return v.value
}

func (v *NullableCreateSmartContractWalletParams) Set(val *CreateSmartContractWalletParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSmartContractWalletParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSmartContractWalletParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSmartContractWalletParams(val *CreateSmartContractWalletParams) *NullableCreateSmartContractWalletParams {
	return &NullableCreateSmartContractWalletParams{value: val, isSet: true}
}

func (v NullableCreateSmartContractWalletParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSmartContractWalletParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

