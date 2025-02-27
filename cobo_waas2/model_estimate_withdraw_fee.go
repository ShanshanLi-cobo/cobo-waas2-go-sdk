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

// checks if the EstimateWithdrawFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimateWithdrawFee{}

// EstimateWithdrawFee struct for EstimateWithdrawFee
type EstimateWithdrawFee struct {
	ActivityType ActivityType `json:"activity_type"`
	// The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization.
	RequestId *string `json:"request_id,omitempty"`
	// The ID of the corresponding staking position.
	StakingId string `json:"staking_id"`
	// The amount to withdraw.
	Amount *string `json:"amount,omitempty"`
	Fee TransactionRequestFee `json:"fee"`
}

type _EstimateWithdrawFee EstimateWithdrawFee

// NewEstimateWithdrawFee instantiates a new EstimateWithdrawFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateWithdrawFee(activityType ActivityType, stakingId string, fee TransactionRequestFee) *EstimateWithdrawFee {
	this := EstimateWithdrawFee{}
	this.ActivityType = activityType
	this.StakingId = stakingId
	this.Fee = fee
	return &this
}

// NewEstimateWithdrawFeeWithDefaults instantiates a new EstimateWithdrawFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateWithdrawFeeWithDefaults() *EstimateWithdrawFee {
	this := EstimateWithdrawFee{}
	return &this
}

// GetActivityType returns the ActivityType field value
func (o *EstimateWithdrawFee) GetActivityType() ActivityType {
	if o == nil {
		var ret ActivityType
		return ret
	}

	return o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value
// and a boolean to check if the value has been set.
func (o *EstimateWithdrawFee) GetActivityTypeOk() (*ActivityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityType, true
}

// SetActivityType sets field value
func (o *EstimateWithdrawFee) SetActivityType(v ActivityType) {
	o.ActivityType = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *EstimateWithdrawFee) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateWithdrawFee) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *EstimateWithdrawFee) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *EstimateWithdrawFee) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStakingId returns the StakingId field value
func (o *EstimateWithdrawFee) GetStakingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value
// and a boolean to check if the value has been set.
func (o *EstimateWithdrawFee) GetStakingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakingId, true
}

// SetStakingId sets field value
func (o *EstimateWithdrawFee) SetStakingId(v string) {
	o.StakingId = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *EstimateWithdrawFee) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateWithdrawFee) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *EstimateWithdrawFee) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *EstimateWithdrawFee) SetAmount(v string) {
	o.Amount = &v
}

// GetFee returns the Fee field value
func (o *EstimateWithdrawFee) GetFee() TransactionRequestFee {
	if o == nil {
		var ret TransactionRequestFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *EstimateWithdrawFee) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *EstimateWithdrawFee) SetFee(v TransactionRequestFee) {
	o.Fee = v
}

func (o EstimateWithdrawFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimateWithdrawFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activity_type"] = o.ActivityType
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["staking_id"] = o.StakingId
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["fee"] = o.Fee
	return toSerialize, nil
}

func (o *EstimateWithdrawFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activity_type",
		"staking_id",
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

	varEstimateWithdrawFee := _EstimateWithdrawFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimateWithdrawFee)

	if err != nil {
		return err
	}

	*o = EstimateWithdrawFee(varEstimateWithdrawFee)

	return err
}

type NullableEstimateWithdrawFee struct {
	value *EstimateWithdrawFee
	isSet bool
}

func (v NullableEstimateWithdrawFee) Get() *EstimateWithdrawFee {
	return v.value
}

func (v *NullableEstimateWithdrawFee) Set(val *EstimateWithdrawFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateWithdrawFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateWithdrawFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateWithdrawFee(val *EstimateWithdrawFee) *NullableEstimateWithdrawFee {
	return &NullableEstimateWithdrawFee{value: val, isSet: true}
}

func (v NullableEstimateWithdrawFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateWithdrawFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


