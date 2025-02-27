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

// checks if the Stakings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stakings{}

// Stakings The information about a staking position.
type Stakings struct {
	// The ID of the staking position.
	Id string `json:"id"`
	// The staker's wallet ID.
	WalletId string `json:"wallet_id"`
	// The staker's wallet address.
	Address string `json:"address"`
	// The details about the staking amount.
	Amounts []AmountDetailsInner `json:"amounts"`
	PoolId StakingPoolId `json:"pool_id"`
	// The token ID.
	TokenId string `json:"token_id"`
	// The information about the staking rewards.
	RewardsInfo map[string]interface{} `json:"rewards_info,omitempty"`
	// The time when the staking position was created.
	CreatedTimestamp int64 `json:"created_timestamp"`
	// The time when the staking position was last updated.
	UpdatedTimestamp int64 `json:"updated_timestamp"`
	ValidatorInfo BabylonValidator `json:"validator_info"`
	Extra *StakingsExtra `json:"extra,omitempty"`
}

type _Stakings Stakings

// NewStakings instantiates a new Stakings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStakings(id string, walletId string, address string, amounts []AmountDetailsInner, poolId StakingPoolId, tokenId string, createdTimestamp int64, updatedTimestamp int64, validatorInfo BabylonValidator) *Stakings {
	this := Stakings{}
	this.Id = id
	this.WalletId = walletId
	this.Address = address
	this.Amounts = amounts
	this.PoolId = poolId
	this.TokenId = tokenId
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	this.ValidatorInfo = validatorInfo
	return &this
}

// NewStakingsWithDefaults instantiates a new Stakings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStakingsWithDefaults() *Stakings {
	this := Stakings{}
	return &this
}

// GetId returns the Id field value
func (o *Stakings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Stakings) SetId(v string) {
	o.Id = v
}

// GetWalletId returns the WalletId field value
func (o *Stakings) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *Stakings) SetWalletId(v string) {
	o.WalletId = v
}

// GetAddress returns the Address field value
func (o *Stakings) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Stakings) SetAddress(v string) {
	o.Address = v
}

// GetAmounts returns the Amounts field value
func (o *Stakings) GetAmounts() []AmountDetailsInner {
	if o == nil {
		var ret []AmountDetailsInner
		return ret
	}

	return o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetAmountsOk() ([]AmountDetailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amounts, true
}

// SetAmounts sets field value
func (o *Stakings) SetAmounts(v []AmountDetailsInner) {
	o.Amounts = v
}

// GetPoolId returns the PoolId field value
func (o *Stakings) GetPoolId() StakingPoolId {
	if o == nil {
		var ret StakingPoolId
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetPoolIdOk() (*StakingPoolId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *Stakings) SetPoolId(v StakingPoolId) {
	o.PoolId = v
}

// GetTokenId returns the TokenId field value
func (o *Stakings) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *Stakings) SetTokenId(v string) {
	o.TokenId = v
}

// GetRewardsInfo returns the RewardsInfo field value if set, zero value otherwise.
func (o *Stakings) GetRewardsInfo() map[string]interface{} {
	if o == nil || IsNil(o.RewardsInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.RewardsInfo
}

// GetRewardsInfoOk returns a tuple with the RewardsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stakings) GetRewardsInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RewardsInfo) {
		return map[string]interface{}{}, false
	}
	return o.RewardsInfo, true
}

// HasRewardsInfo returns a boolean if a field has been set.
func (o *Stakings) HasRewardsInfo() bool {
	if o != nil && !IsNil(o.RewardsInfo) {
		return true
	}

	return false
}

// SetRewardsInfo gets a reference to the given map[string]interface{} and assigns it to the RewardsInfo field.
func (o *Stakings) SetRewardsInfo(v map[string]interface{}) {
	o.RewardsInfo = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *Stakings) GetCreatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *Stakings) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *Stakings) GetUpdatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *Stakings) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = v
}

// GetValidatorInfo returns the ValidatorInfo field value
func (o *Stakings) GetValidatorInfo() BabylonValidator {
	if o == nil {
		var ret BabylonValidator
		return ret
	}

	return o.ValidatorInfo
}

// GetValidatorInfoOk returns a tuple with the ValidatorInfo field value
// and a boolean to check if the value has been set.
func (o *Stakings) GetValidatorInfoOk() (*BabylonValidator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorInfo, true
}

// SetValidatorInfo sets field value
func (o *Stakings) SetValidatorInfo(v BabylonValidator) {
	o.ValidatorInfo = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Stakings) GetExtra() StakingsExtra {
	if o == nil || IsNil(o.Extra) {
		var ret StakingsExtra
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stakings) GetExtraOk() (*StakingsExtra, bool) {
	if o == nil || IsNil(o.Extra) {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Stakings) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given StakingsExtra and assigns it to the Extra field.
func (o *Stakings) SetExtra(v StakingsExtra) {
	o.Extra = &v
}

func (o Stakings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stakings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["address"] = o.Address
	toSerialize["amounts"] = o.Amounts
	toSerialize["pool_id"] = o.PoolId
	toSerialize["token_id"] = o.TokenId
	if !IsNil(o.RewardsInfo) {
		toSerialize["rewards_info"] = o.RewardsInfo
	}
	toSerialize["created_timestamp"] = o.CreatedTimestamp
	toSerialize["updated_timestamp"] = o.UpdatedTimestamp
	toSerialize["validator_info"] = o.ValidatorInfo
	if !IsNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	return toSerialize, nil
}

func (o *Stakings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"wallet_id",
		"address",
		"amounts",
		"pool_id",
		"token_id",
		"created_timestamp",
		"updated_timestamp",
		"validator_info",
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

	varStakings := _Stakings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStakings)

	if err != nil {
		return err
	}

	*o = Stakings(varStakings)

	return err
}

type NullableStakings struct {
	value *Stakings
	isSet bool
}

func (v NullableStakings) Get() *Stakings {
	return v.value
}

func (v *NullableStakings) Set(val *Stakings) {
	v.value = val
	v.isSet = true
}

func (v NullableStakings) IsSet() bool {
	return v.isSet
}

func (v *NullableStakings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStakings(val *Stakings) *NullableStakings {
	return &NullableStakings{value: val, isSet: true}
}

func (v NullableStakings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStakings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


