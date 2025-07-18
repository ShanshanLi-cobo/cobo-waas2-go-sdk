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

// checks if the TokenizationUpdateAllowlistAddressesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenizationUpdateAllowlistAddressesRequest{}

// TokenizationUpdateAllowlistAddressesRequest struct for TokenizationUpdateAllowlistAddressesRequest
type TokenizationUpdateAllowlistAddressesRequest struct {
	Action TokenizationUpdateAddressAction `json:"action"`
	Source TokenizationTokenOperationSource `json:"source"`
	// A list of addresses to manage. For 'add' operations, notes can be provided. For 'remove' operations, notes are ignored.
	Addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner `json:"addresses"`
	// The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator.
	AppInitiator *string `json:"app_initiator,omitempty"`
	Fee TransactionRequestFee `json:"fee"`
}

type _TokenizationUpdateAllowlistAddressesRequest TokenizationUpdateAllowlistAddressesRequest

// NewTokenizationUpdateAllowlistAddressesRequest instantiates a new TokenizationUpdateAllowlistAddressesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizationUpdateAllowlistAddressesRequest(action TokenizationUpdateAddressAction, source TokenizationTokenOperationSource, addresses []TokenizationUpdateAllowlistAddressesParamsAddressesInner, fee TransactionRequestFee) *TokenizationUpdateAllowlistAddressesRequest {
	this := TokenizationUpdateAllowlistAddressesRequest{}
	this.Action = action
	this.Source = source
	this.Addresses = addresses
	this.Fee = fee
	return &this
}

// NewTokenizationUpdateAllowlistAddressesRequestWithDefaults instantiates a new TokenizationUpdateAllowlistAddressesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizationUpdateAllowlistAddressesRequestWithDefaults() *TokenizationUpdateAllowlistAddressesRequest {
	this := TokenizationUpdateAllowlistAddressesRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *TokenizationUpdateAllowlistAddressesRequest) GetAction() TokenizationUpdateAddressAction {
	if o == nil {
		var ret TokenizationUpdateAddressAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetActionOk() (*TokenizationUpdateAddressAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *TokenizationUpdateAllowlistAddressesRequest) SetAction(v TokenizationUpdateAddressAction) {
	o.Action = v
}

// GetSource returns the Source field value
func (o *TokenizationUpdateAllowlistAddressesRequest) GetSource() TokenizationTokenOperationSource {
	if o == nil {
		var ret TokenizationTokenOperationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TokenizationUpdateAllowlistAddressesRequest) SetSource(v TokenizationTokenOperationSource) {
	o.Source = v
}

// GetAddresses returns the Addresses field value
func (o *TokenizationUpdateAllowlistAddressesRequest) GetAddresses() []TokenizationUpdateAllowlistAddressesParamsAddressesInner {
	if o == nil {
		var ret []TokenizationUpdateAllowlistAddressesParamsAddressesInner
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetAddressesOk() ([]TokenizationUpdateAllowlistAddressesParamsAddressesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *TokenizationUpdateAllowlistAddressesRequest) SetAddresses(v []TokenizationUpdateAllowlistAddressesParamsAddressesInner) {
	o.Addresses = v
}

// GetAppInitiator returns the AppInitiator field value if set, zero value otherwise.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetAppInitiator() string {
	if o == nil || IsNil(o.AppInitiator) {
		var ret string
		return ret
	}
	return *o.AppInitiator
}

// GetAppInitiatorOk returns a tuple with the AppInitiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetAppInitiatorOk() (*string, bool) {
	if o == nil || IsNil(o.AppInitiator) {
		return nil, false
	}
	return o.AppInitiator, true
}

// HasAppInitiator returns a boolean if a field has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) HasAppInitiator() bool {
	if o != nil && !IsNil(o.AppInitiator) {
		return true
	}

	return false
}

// SetAppInitiator gets a reference to the given string and assigns it to the AppInitiator field.
func (o *TokenizationUpdateAllowlistAddressesRequest) SetAppInitiator(v string) {
	o.AppInitiator = &v
}

// GetFee returns the Fee field value
func (o *TokenizationUpdateAllowlistAddressesRequest) GetFee() TransactionRequestFee {
	if o == nil {
		var ret TransactionRequestFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *TokenizationUpdateAllowlistAddressesRequest) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *TokenizationUpdateAllowlistAddressesRequest) SetFee(v TransactionRequestFee) {
	o.Fee = v
}

func (o TokenizationUpdateAllowlistAddressesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizationUpdateAllowlistAddressesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["source"] = o.Source
	toSerialize["addresses"] = o.Addresses
	if !IsNil(o.AppInitiator) {
		toSerialize["app_initiator"] = o.AppInitiator
	}
	toSerialize["fee"] = o.Fee
	return toSerialize, nil
}

func (o *TokenizationUpdateAllowlistAddressesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"source",
		"addresses",
		"fee",
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

	varTokenizationUpdateAllowlistAddressesRequest := _TokenizationUpdateAllowlistAddressesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenizationUpdateAllowlistAddressesRequest)

	if err != nil {
		return err
	}

	*o = TokenizationUpdateAllowlistAddressesRequest(varTokenizationUpdateAllowlistAddressesRequest)

	return err
}

type NullableTokenizationUpdateAllowlistAddressesRequest struct {
	value *TokenizationUpdateAllowlistAddressesRequest
	isSet bool
}

func (v NullableTokenizationUpdateAllowlistAddressesRequest) Get() *TokenizationUpdateAllowlistAddressesRequest {
	return v.value
}

func (v *NullableTokenizationUpdateAllowlistAddressesRequest) Set(val *TokenizationUpdateAllowlistAddressesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizationUpdateAllowlistAddressesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizationUpdateAllowlistAddressesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenizationUpdateAllowlistAddressesRequest(val *TokenizationUpdateAllowlistAddressesRequest) *NullableTokenizationUpdateAllowlistAddressesRequest {
	return &NullableTokenizationUpdateAllowlistAddressesRequest{value: val, isSet: true}
}

func (v NullableTokenizationUpdateAllowlistAddressesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizationUpdateAllowlistAddressesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


