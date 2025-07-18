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

// checks if the ApprovalUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalUser{}

// ApprovalUser The information of the user.
type ApprovalUser struct {
	// The user email.
	Email string `json:"email"`
	// The user name.
	Name *string `json:"name,omitempty"`
	Status ApprovalStatus `json:"status"`
	// The time when the approval was created, in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp int64 `json:"created_timestamp"`
}

type _ApprovalUser ApprovalUser

// NewApprovalUser instantiates a new ApprovalUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalUser(email string, status ApprovalStatus, createdTimestamp int64) *ApprovalUser {
	this := ApprovalUser{}
	this.Email = email
	this.Status = status
	this.CreatedTimestamp = createdTimestamp
	return &this
}

// NewApprovalUserWithDefaults instantiates a new ApprovalUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalUserWithDefaults() *ApprovalUser {
	this := ApprovalUser{}
	return &this
}

// GetEmail returns the Email field value
func (o *ApprovalUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ApprovalUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ApprovalUser) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApprovalUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApprovalUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApprovalUser) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value
func (o *ApprovalUser) GetStatus() ApprovalStatus {
	if o == nil {
		var ret ApprovalStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApprovalUser) GetStatusOk() (*ApprovalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApprovalUser) SetStatus(v ApprovalStatus) {
	o.Status = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *ApprovalUser) GetCreatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ApprovalUser) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *ApprovalUser) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = v
}

func (o ApprovalUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["status"] = o.Status
	toSerialize["created_timestamp"] = o.CreatedTimestamp
	return toSerialize, nil
}

func (o *ApprovalUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"status",
		"created_timestamp",
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

	varApprovalUser := _ApprovalUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApprovalUser)

	if err != nil {
		return err
	}

	*o = ApprovalUser(varApprovalUser)

	return err
}

type NullableApprovalUser struct {
	value *ApprovalUser
	isSet bool
}

func (v NullableApprovalUser) Get() *ApprovalUser {
	return v.value
}

func (v *NullableApprovalUser) Set(val *ApprovalUser) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalUser) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalUser(val *ApprovalUser) *NullableApprovalUser {
	return &NullableApprovalUser{value: val, isSet: true}
}

func (v NullableApprovalUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


