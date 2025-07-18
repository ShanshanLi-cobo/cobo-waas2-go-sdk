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

// checks if the TokenizationAllowlistActivationParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationAllowlistActivationParams{}

// TokenizationAllowlistActivationParams struct for TokenizationAllowlistActivationParams
type TokenizationAllowlistActivationParams struct {
	Source TokenizationTokenOperationSource `json:"source"`
	// Whether to activate the allowlist feature for the token.
	Activation bool `json:"activation"`
}

type _TokenizationAllowlistActivationParams TokenizationAllowlistActivationParams

// NewTokenizationAllowlistActivationParams instantiates a new TokenizationAllowlistActivationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationAllowlistActivationParams(source TokenizationTokenOperationSource, activation bool) *TokenizationAllowlistActivationParams {
	this := TokenizationAllowlistActivationParams{}
	this.Source = source
	this.Activation = activation
	return &this
}

// NewTokenizationAllowlistActivationParamsWithDefaults instantiates a new TokenizationAllowlistActivationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationAllowlistActivationParamsWithDefaults() *TokenizationAllowlistActivationParams {
	this := TokenizationAllowlistActivationParams{}
	return &this
}

// GetSource returns the Source field value
func (o *TokenizationAllowlistActivationParams) GetSource() TokenizationTokenOperationSource {
	if o == nil {
		var ret TokenizationTokenOperationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TokenizationAllowlistActivationParams) GetSourceOk() (*TokenizationTokenOperationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TokenizationAllowlistActivationParams) SetSource(v TokenizationTokenOperationSource) {
	o.Source = v
}

// GetActivation returns the Activation field value
func (o *TokenizationAllowlistActivationParams) GetActivation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Activation
}

// GetActivationOk returns a tuple with the Activation field value
// and a boolean to check if the value has been set.
func (o *TokenizationAllowlistActivationParams) GetActivationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Activation, true
}

// SetActivation sets field value
func (o *TokenizationAllowlistActivationParams) SetActivation(v bool) {
	o.Activation = v
}

func (o TokenizationAllowlistActivationParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationAllowlistActivationParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["activation"] = o.Activation
	return toSerialize, nil
}

func (o *TokenizationAllowlistActivationParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source",
		"activation",
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

	varTokenizationAllowlistActivationParams := _TokenizationAllowlistActivationParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationAllowlistActivationParams)

	if err != nil {
		return err
	}

	*o = TokenizationAllowlistActivationParams(varTokenizationAllowlistActivationParams)

	return err
}

type NullableTokenizationAllowlistActivationParams struct {
	value *TokenizationAllowlistActivationParams
	isSet bool
}

func (v NullableTokenizationAllowlistActivationParams) Get() *TokenizationAllowlistActivationParams {
	return v.value
}

func (v *NullableTokenizationAllowlistActivationParams) Set(val *TokenizationAllowlistActivationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationAllowlistActivationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationAllowlistActivationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationAllowlistActivationParams(val *TokenizationAllowlistActivationParams) *NullableTokenizationAllowlistActivationParams {
	return &NullableTokenizationAllowlistActivationParams{value: val, isSet: true}
}

func (v NullableTokenizationAllowlistActivationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationAllowlistActivationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


