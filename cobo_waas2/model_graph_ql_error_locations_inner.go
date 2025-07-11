/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the GraphQLErrorLocationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLErrorLocationsInner{}

// GraphQLErrorLocationsInner struct for GraphQLErrorLocationsInner
type GraphQLErrorLocationsInner struct {
	// The line number in the query where the error occurred.
	Line *int32 `json:"line,omitempty"`
	// The column number in the query where the error occurred.
	Column *int32 `json:"column,omitempty"`
}

// NewGraphQLErrorLocationsInner instantiates a new GraphQLErrorLocationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLErrorLocationsInner() *GraphQLErrorLocationsInner {
	this := GraphQLErrorLocationsInner{}
	return &this
}

// NewGraphQLErrorLocationsInnerWithDefaults instantiates a new GraphQLErrorLocationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLErrorLocationsInnerWithDefaults() *GraphQLErrorLocationsInner {
	this := GraphQLErrorLocationsInner{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *GraphQLErrorLocationsInner) GetLine() int32 {
	if o == nil || IsNil(o.Line) {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLErrorLocationsInner) GetLineOk() (*int32, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *GraphQLErrorLocationsInner) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *GraphQLErrorLocationsInner) SetLine(v int32) {
	o.Line = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *GraphQLErrorLocationsInner) GetColumn() int32 {
	if o == nil || IsNil(o.Column) {
		var ret int32
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLErrorLocationsInner) GetColumnOk() (*int32, bool) {
	if o == nil || IsNil(o.Column) {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *GraphQLErrorLocationsInner) HasColumn() bool {
	if o != nil && !IsNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given int32 and assigns it to the Column field.
func (o *GraphQLErrorLocationsInner) SetColumn(v int32) {
	o.Column = &v
}

func (o GraphQLErrorLocationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLErrorLocationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	return toSerialize, nil
}

type NullableGraphQLErrorLocationsInner struct {
	value *GraphQLErrorLocationsInner
	isSet bool
}

func (v NullableGraphQLErrorLocationsInner) Get() *GraphQLErrorLocationsInner {
	return v.value
}

func (v *NullableGraphQLErrorLocationsInner) Set(val *GraphQLErrorLocationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLErrorLocationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLErrorLocationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLErrorLocationsInner(val *GraphQLErrorLocationsInner) *NullableGraphQLErrorLocationsInner {
	return &NullableGraphQLErrorLocationsInner{value: val, isSet: true}
}

func (v NullableGraphQLErrorLocationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLErrorLocationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


