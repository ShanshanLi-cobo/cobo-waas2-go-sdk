/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EstimatedUtxoFeeSlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimatedUtxoFeeSlow{}

// EstimatedUtxoFeeSlow struct for EstimatedUtxoFeeSlow
type EstimatedUtxoFeeSlow struct {
	// The fee rate in sat/vByte. The fee rate represents the satoshis you are willing to pay for each byte of data that your transaction will consume on the blockchain.
	FeeRate string `json:"fee_rate"`
	// The fee that you need to pay for the transaction.
	FeeAmount string `json:"fee_amount"`
}

type _EstimatedUtxoFeeSlow EstimatedUtxoFeeSlow

// NewEstimatedUtxoFeeSlow instantiates a new EstimatedUtxoFeeSlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedUtxoFeeSlow(feeRate string, feeAmount string) *EstimatedUtxoFeeSlow {
	this := EstimatedUtxoFeeSlow{}
	this.FeeRate = feeRate
	this.FeeAmount = feeAmount
	return &this
}

// NewEstimatedUtxoFeeSlowWithDefaults instantiates a new EstimatedUtxoFeeSlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedUtxoFeeSlowWithDefaults() *EstimatedUtxoFeeSlow {
	this := EstimatedUtxoFeeSlow{}
	return &this
}

// GetFeeRate returns the FeeRate field value
func (o *EstimatedUtxoFeeSlow) GetFeeRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value
// and a boolean to check if the value has been set.
func (o *EstimatedUtxoFeeSlow) GetFeeRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRate, true
}

// SetFeeRate sets field value
func (o *EstimatedUtxoFeeSlow) SetFeeRate(v string) {
	o.FeeRate = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *EstimatedUtxoFeeSlow) GetFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *EstimatedUtxoFeeSlow) GetFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *EstimatedUtxoFeeSlow) SetFeeAmount(v string) {
	o.FeeAmount = v
}

func (o EstimatedUtxoFeeSlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimatedUtxoFeeSlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fee_rate"] = o.FeeRate
	toSerialize["fee_amount"] = o.FeeAmount
	return toSerialize, nil
}

func (o *EstimatedUtxoFeeSlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_rate",
		"fee_amount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEstimatedUtxoFeeSlow := _EstimatedUtxoFeeSlow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimatedUtxoFeeSlow)

	if err != nil {
		return err
	}

	*o = EstimatedUtxoFeeSlow(varEstimatedUtxoFeeSlow)

	return err
}

type NullableEstimatedUtxoFeeSlow struct {
	value *EstimatedUtxoFeeSlow
	isSet bool
}

func (v NullableEstimatedUtxoFeeSlow) Get() *EstimatedUtxoFeeSlow {
	return v.value
}

func (v *NullableEstimatedUtxoFeeSlow) Set(val *EstimatedUtxoFeeSlow) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedUtxoFeeSlow) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedUtxoFeeSlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedUtxoFeeSlow(val *EstimatedUtxoFeeSlow) *NullableEstimatedUtxoFeeSlow {
	return &NullableEstimatedUtxoFeeSlow{value: val, isSet: true}
}

func (v NullableEstimatedUtxoFeeSlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedUtxoFeeSlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

