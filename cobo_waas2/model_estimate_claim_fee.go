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

// checks if the EstimateClaimFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimateClaimFee{}

// EstimateClaimFee struct for EstimateClaimFee
type EstimateClaimFee struct {
	ActivityType ActivityType `json:"activity_type"`
	// The ID of the staking position. You can retrieve a list of staking positions by calling [List staking positions](/v2/api-references/stakings/list-staking-positions).
	StakingId *string `json:"staking_id,omitempty"`
	Fee *TransactionRequestFee `json:"fee,omitempty"`
}

type _EstimateClaimFee EstimateClaimFee

// NewEstimateClaimFee instantiates a new EstimateClaimFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateClaimFee(activityType ActivityType) *EstimateClaimFee {
	this := EstimateClaimFee{}
	this.ActivityType = activityType
	return &this
}

// NewEstimateClaimFeeWithDefaults instantiates a new EstimateClaimFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateClaimFeeWithDefaults() *EstimateClaimFee {
	this := EstimateClaimFee{}
	return &this
}

// GetActivityType returns the ActivityType field value
func (o *EstimateClaimFee) GetActivityType() ActivityType {
	if o == nil {
		var ret ActivityType
		return ret
	}

	return o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value
// and a boolean to check if the value has been set.
func (o *EstimateClaimFee) GetActivityTypeOk() (*ActivityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivityType, true
}

// SetActivityType sets field value
func (o *EstimateClaimFee) SetActivityType(v ActivityType) {
	o.ActivityType = v
}

// GetStakingId returns the StakingId field value if set, zero value otherwise.
func (o *EstimateClaimFee) GetStakingId() string {
	if o == nil || IsNil(o.StakingId) {
		var ret string
		return ret
	}
	return *o.StakingId
}

// GetStakingIdOk returns a tuple with the StakingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateClaimFee) GetStakingIdOk() (*string, bool) {
	if o == nil || IsNil(o.StakingId) {
		return nil, false
	}
	return o.StakingId, true
}

// HasStakingId returns a boolean if a field has been set.
func (o *EstimateClaimFee) HasStakingId() bool {
	if o != nil && !IsNil(o.StakingId) {
		return true
	}

	return false
}

// SetStakingId gets a reference to the given string and assigns it to the StakingId field.
func (o *EstimateClaimFee) SetStakingId(v string) {
	o.StakingId = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *EstimateClaimFee) GetFee() TransactionRequestFee {
	if o == nil || IsNil(o.Fee) {
		var ret TransactionRequestFee
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateClaimFee) GetFeeOk() (*TransactionRequestFee, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *EstimateClaimFee) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given TransactionRequestFee and assigns it to the Fee field.
func (o *EstimateClaimFee) SetFee(v TransactionRequestFee) {
	o.Fee = &v
}

func (o EstimateClaimFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimateClaimFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activity_type"] = o.ActivityType
	if !IsNil(o.StakingId) {
		toSerialize["staking_id"] = o.StakingId
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	return toSerialize, nil
}

func (o *EstimateClaimFee) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activity_type",
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

	varEstimateClaimFee := _EstimateClaimFee{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEstimateClaimFee)

	if err != nil {
		return err
	}

	*o = EstimateClaimFee(varEstimateClaimFee)

	return err
}

type NullableEstimateClaimFee struct {
	value *EstimateClaimFee
	isSet bool
}

func (v NullableEstimateClaimFee) Get() *EstimateClaimFee {
	return v.value
}

func (v *NullableEstimateClaimFee) Set(val *EstimateClaimFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateClaimFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateClaimFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateClaimFee(val *EstimateClaimFee) *NullableEstimateClaimFee {
	return &NullableEstimateClaimFee{value: val, isSet: true}
}

func (v NullableEstimateClaimFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateClaimFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

