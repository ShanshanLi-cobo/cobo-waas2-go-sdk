/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the SOLBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SOLBase{}

// SOLBase struct for SOLBase
type SOLBase struct {
	// A fixed fee charged per signature. The default is 5,000 lamports per signature.
	BaseFee *string `json:"base_fee,omitempty"`
	// The rent fee charged by the network to store non–rent-exempt accounts on-chain. It is deducted periodically until the account maintains the minimum balance required for rent exemption.
	RentAmount *string `json:"rent_amount,omitempty"`
}

// NewSOLBase instantiates a new SOLBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSOLBase() *SOLBase {
	this := SOLBase{}
	return &this
}

// NewSOLBaseWithDefaults instantiates a new SOLBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSOLBaseWithDefaults() *SOLBase {
	this := SOLBase{}
	return &this
}

// GetBaseFee returns the BaseFee field value if set, zero value otherwise.
func (o *SOLBase) GetBaseFee() string {
	if o == nil || IsNil(o.BaseFee) {
		var ret string
		return ret
	}
	return *o.BaseFee
}

// GetBaseFeeOk returns a tuple with the BaseFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SOLBase) GetBaseFeeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseFee) {
		return nil, false
	}
	return o.BaseFee, true
}

// HasBaseFee returns a boolean if a field has been set.
func (o *SOLBase) HasBaseFee() bool {
	if o != nil && !IsNil(o.BaseFee) {
		return true
	}

	return false
}

// SetBaseFee gets a reference to the given string and assigns it to the BaseFee field.
func (o *SOLBase) SetBaseFee(v string) {
	o.BaseFee = &v
}

// GetRentAmount returns the RentAmount field value if set, zero value otherwise.
func (o *SOLBase) GetRentAmount() string {
	if o == nil || IsNil(o.RentAmount) {
		var ret string
		return ret
	}
	return *o.RentAmount
}

// GetRentAmountOk returns a tuple with the RentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SOLBase) GetRentAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RentAmount) {
		return nil, false
	}
	return o.RentAmount, true
}

// HasRentAmount returns a boolean if a field has been set.
func (o *SOLBase) HasRentAmount() bool {
	if o != nil && !IsNil(o.RentAmount) {
		return true
	}

	return false
}

// SetRentAmount gets a reference to the given string and assigns it to the RentAmount field.
func (o *SOLBase) SetRentAmount(v string) {
	o.RentAmount = &v
}

func (o SOLBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SOLBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseFee) {
		toSerialize["base_fee"] = o.BaseFee
	}
	if !IsNil(o.RentAmount) {
		toSerialize["rent_amount"] = o.RentAmount
	}
	return toSerialize, nil
}

type NullableSOLBase struct {
	value *SOLBase
	isSet bool
}

func (v NullableSOLBase) Get() *SOLBase {
	return v.value
}

func (v *NullableSOLBase) Set(val *SOLBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSOLBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSOLBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSOLBase(val *SOLBase) *NullableSOLBase {
	return &NullableSOLBase{value: val, isSet: true}
}

func (v NullableSOLBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSOLBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


