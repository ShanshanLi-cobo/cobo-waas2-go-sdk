/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TokenizationBurnTokenParamsBurnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationBurnTokenParamsBurnsInner{}

// TokenizationBurnTokenParamsBurnsInner struct for TokenizationBurnTokenParamsBurnsInner
type TokenizationBurnTokenParamsBurnsInner struct {
	// The amount of tokens to burn from this source.
	Amount string `json:"amount"`
	// The address to burn tokens from for this source.
	FromAddress string `json:"from_address"`
}

type _TokenizationBurnTokenParamsBurnsInner TokenizationBurnTokenParamsBurnsInner

// NewTokenizationBurnTokenParamsBurnsInner instantiates a new TokenizationBurnTokenParamsBurnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationBurnTokenParamsBurnsInner(amount string, fromAddress string) *TokenizationBurnTokenParamsBurnsInner {
	this := TokenizationBurnTokenParamsBurnsInner{}
	this.Amount = amount
	this.FromAddress = fromAddress
	return &this
}

// NewTokenizationBurnTokenParamsBurnsInnerWithDefaults instantiates a new TokenizationBurnTokenParamsBurnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationBurnTokenParamsBurnsInnerWithDefaults() *TokenizationBurnTokenParamsBurnsInner {
	this := TokenizationBurnTokenParamsBurnsInner{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TokenizationBurnTokenParamsBurnsInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenizationBurnTokenParamsBurnsInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenizationBurnTokenParamsBurnsInner) SetAmount(v string) {
	o.Amount = v
}

// GetFromAddress returns the FromAddress field value
func (o *TokenizationBurnTokenParamsBurnsInner) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *TokenizationBurnTokenParamsBurnsInner) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *TokenizationBurnTokenParamsBurnsInner) SetFromAddress(v string) {
	o.FromAddress = v
}

func (o TokenizationBurnTokenParamsBurnsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationBurnTokenParamsBurnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["from_address"] = o.FromAddress
	return toSerialize, nil
}

func (o *TokenizationBurnTokenParamsBurnsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"from_address",
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

	varTokenizationBurnTokenParamsBurnsInner := _TokenizationBurnTokenParamsBurnsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationBurnTokenParamsBurnsInner)

	if err != nil {
		return err
	}

	*o = TokenizationBurnTokenParamsBurnsInner(varTokenizationBurnTokenParamsBurnsInner)

	return err
}

type NullableTokenizationBurnTokenParamsBurnsInner struct {
	value *TokenizationBurnTokenParamsBurnsInner
	isSet bool
}

func (v NullableTokenizationBurnTokenParamsBurnsInner) Get() *TokenizationBurnTokenParamsBurnsInner {
	return v.value
}

func (v *NullableTokenizationBurnTokenParamsBurnsInner) Set(val *TokenizationBurnTokenParamsBurnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationBurnTokenParamsBurnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationBurnTokenParamsBurnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationBurnTokenParamsBurnsInner(val *TokenizationBurnTokenParamsBurnsInner) *NullableTokenizationBurnTokenParamsBurnsInner {
	return &NullableTokenizationBurnTokenParamsBurnsInner{value: val, isSet: true}
}

func (v NullableTokenizationBurnTokenParamsBurnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationBurnTokenParamsBurnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


