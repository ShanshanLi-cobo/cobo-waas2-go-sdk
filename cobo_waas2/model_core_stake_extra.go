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

// checks if the CoreStakeExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreStakeExtra{}

// CoreStakeExtra struct for CoreStakeExtra
type CoreStakeExtra struct {
	PoolType StakingPoolType `json:"pool_type"`
	// The Unix timestamp (in seconds) when the staking position will be unlocked and available for withdrawal.
	Timelock int32 `json:"timelock"`
	// The change address on the Bitcoin chain. If not provided, the source wallet's address will be used as the change address.
	ChangeAddress *string `json:"change_address,omitempty"`
	// The validator's EVM address.
	ValidatorAddress string `json:"validator_address"`
	// The EVM address to receive staking rewards.
	RewardAddress string `json:"reward_address"`
}

type _CoreStakeExtra CoreStakeExtra

// NewCoreStakeExtra instantiates a new CoreStakeExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreStakeExtra(poolType StakingPoolType, timelock int32, validatorAddress string, rewardAddress string) *CoreStakeExtra {
	this := CoreStakeExtra{}
	this.PoolType = poolType
	this.Timelock = timelock
	this.ValidatorAddress = validatorAddress
	this.RewardAddress = rewardAddress
	return &this
}

// NewCoreStakeExtraWithDefaults instantiates a new CoreStakeExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreStakeExtraWithDefaults() *CoreStakeExtra {
	this := CoreStakeExtra{}
	return &this
}

// GetPoolType returns the PoolType field value
func (o *CoreStakeExtra) GetPoolType() StakingPoolType {
	if o == nil {
		var ret StakingPoolType
		return ret
	}

	return o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value
// and a boolean to check if the value has been set.
func (o *CoreStakeExtra) GetPoolTypeOk() (*StakingPoolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolType, true
}

// SetPoolType sets field value
func (o *CoreStakeExtra) SetPoolType(v StakingPoolType) {
	o.PoolType = v
}

// GetTimelock returns the Timelock field value
func (o *CoreStakeExtra) GetTimelock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timelock
}

// GetTimelockOk returns a tuple with the Timelock field value
// and a boolean to check if the value has been set.
func (o *CoreStakeExtra) GetTimelockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timelock, true
}

// SetTimelock sets field value
func (o *CoreStakeExtra) SetTimelock(v int32) {
	o.Timelock = v
}

// GetChangeAddress returns the ChangeAddress field value if set, zero value otherwise.
func (o *CoreStakeExtra) GetChangeAddress() string {
	if o == nil || IsNil(o.ChangeAddress) {
		var ret string
		return ret
	}
	return *o.ChangeAddress
}

// GetChangeAddressOk returns a tuple with the ChangeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreStakeExtra) GetChangeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeAddress) {
		return nil, false
	}
	return o.ChangeAddress, true
}

// HasChangeAddress returns a boolean if a field has been set.
func (o *CoreStakeExtra) HasChangeAddress() bool {
	if o != nil && !IsNil(o.ChangeAddress) {
		return true
	}

	return false
}

// SetChangeAddress gets a reference to the given string and assigns it to the ChangeAddress field.
func (o *CoreStakeExtra) SetChangeAddress(v string) {
	o.ChangeAddress = &v
}

// GetValidatorAddress returns the ValidatorAddress field value
func (o *CoreStakeExtra) GetValidatorAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorAddress
}

// GetValidatorAddressOk returns a tuple with the ValidatorAddress field value
// and a boolean to check if the value has been set.
func (o *CoreStakeExtra) GetValidatorAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorAddress, true
}

// SetValidatorAddress sets field value
func (o *CoreStakeExtra) SetValidatorAddress(v string) {
	o.ValidatorAddress = v
}

// GetRewardAddress returns the RewardAddress field value
func (o *CoreStakeExtra) GetRewardAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardAddress
}

// GetRewardAddressOk returns a tuple with the RewardAddress field value
// and a boolean to check if the value has been set.
func (o *CoreStakeExtra) GetRewardAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardAddress, true
}

// SetRewardAddress sets field value
func (o *CoreStakeExtra) SetRewardAddress(v string) {
	o.RewardAddress = v
}

func (o CoreStakeExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreStakeExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_type"] = o.PoolType
	toSerialize["timelock"] = o.Timelock
	if !IsNil(o.ChangeAddress) {
		toSerialize["change_address"] = o.ChangeAddress
	}
	toSerialize["validator_address"] = o.ValidatorAddress
	toSerialize["reward_address"] = o.RewardAddress
	return toSerialize, nil
}

func (o *CoreStakeExtra) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pool_type",
		"timelock",
		"validator_address",
		"reward_address",
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

	varCoreStakeExtra := _CoreStakeExtra{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreStakeExtra)

	if err != nil {
		return err
	}

	*o = CoreStakeExtra(varCoreStakeExtra)

	return err
}

type NullableCoreStakeExtra struct {
	value *CoreStakeExtra
	isSet bool
}

func (v NullableCoreStakeExtra) Get() *CoreStakeExtra {
	return v.value
}

func (v *NullableCoreStakeExtra) Set(val *CoreStakeExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreStakeExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreStakeExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreStakeExtra(val *CoreStakeExtra) *NullableCoreStakeExtra {
	return &NullableCoreStakeExtra{value: val, isSet: true}
}

func (v NullableCoreStakeExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreStakeExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

