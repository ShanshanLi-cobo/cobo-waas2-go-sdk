/*
Cobo Wallet as a Service 2.0

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waas2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MPCWalletInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MPCWalletInfo{}

// MPCWalletInfo struct for MPCWalletInfo
type MPCWalletInfo struct {
	// The wallet ID.
	WalletId string `json:"wallet_id"`
	WalletType WalletType `json:"wallet_type"`
	WalletSubtype WalletSubtype `json:"wallet_subtype"`
	// The wallet name.
	Name string `json:"name"`
	// The ID of the owning organization.
	OrgId string `json:"org_id"`
	// The project ID.
	ProjectId *string `json:"project_id,omitempty"`
	// The ID of the owning vault.
	VaultId string `json:"vault_id"`
}

type _MPCWalletInfo MPCWalletInfo

// NewMPCWalletInfo instantiates a new MPCWalletInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMPCWalletInfo(walletId string, walletType WalletType, walletSubtype WalletSubtype, name string, orgId string, vaultId string) *MPCWalletInfo {
	this := MPCWalletInfo{}
	this.WalletId = walletId
	this.WalletType = walletType
	this.WalletSubtype = walletSubtype
	this.Name = name
	this.OrgId = orgId
	this.VaultId = vaultId
	return &this
}

// NewMPCWalletInfoWithDefaults instantiates a new MPCWalletInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMPCWalletInfoWithDefaults() *MPCWalletInfo {
	this := MPCWalletInfo{}
	return &this
}

// GetWalletId returns the WalletId field value
func (o *MPCWalletInfo) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetWalletIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *MPCWalletInfo) SetWalletId(v string) {
	o.WalletId = v
}

// GetWalletType returns the WalletType field value
func (o *MPCWalletInfo) GetWalletType() WalletType {
	if o == nil {
		var ret WalletType
		return ret
	}

	return o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetWalletTypeOk() (*WalletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletType, true
}

// SetWalletType sets field value
func (o *MPCWalletInfo) SetWalletType(v WalletType) {
	o.WalletType = v
}

// GetWalletSubtype returns the WalletSubtype field value
func (o *MPCWalletInfo) GetWalletSubtype() WalletSubtype {
	if o == nil {
		var ret WalletSubtype
		return ret
	}

	return o.WalletSubtype
}

// GetWalletSubtypeOk returns a tuple with the WalletSubtype field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetWalletSubtypeOk() (*WalletSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletSubtype, true
}

// SetWalletSubtype sets field value
func (o *MPCWalletInfo) SetWalletSubtype(v WalletSubtype) {
	o.WalletSubtype = v
}

// GetName returns the Name field value
func (o *MPCWalletInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MPCWalletInfo) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *MPCWalletInfo) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *MPCWalletInfo) SetOrgId(v string) {
	o.OrgId = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *MPCWalletInfo) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *MPCWalletInfo) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *MPCWalletInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetVaultId returns the VaultId field value
func (o *MPCWalletInfo) GetVaultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultId
}

// GetVaultIdOk returns a tuple with the VaultId field value
// and a boolean to check if the value has been set.
func (o *MPCWalletInfo) GetVaultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultId, true
}

// SetVaultId sets field value
func (o *MPCWalletInfo) SetVaultId(v string) {
	o.VaultId = v
}

func (o MPCWalletInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MPCWalletInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet_id"] = o.WalletId
	toSerialize["wallet_type"] = o.WalletType
	toSerialize["wallet_subtype"] = o.WalletSubtype
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["vault_id"] = o.VaultId
	return toSerialize, nil
}

func (o *MPCWalletInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet_id",
		"wallet_type",
		"wallet_subtype",
		"name",
		"org_id",
		"vault_id",
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

	varMPCWalletInfo := _MPCWalletInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMPCWalletInfo)

	if err != nil {
		return err
	}

	*o = MPCWalletInfo(varMPCWalletInfo)

	return err
}

type NullableMPCWalletInfo struct {
	value *MPCWalletInfo
	isSet bool
}

func (v NullableMPCWalletInfo) Get() *MPCWalletInfo {
	return v.value
}

func (v *NullableMPCWalletInfo) Set(val *MPCWalletInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMPCWalletInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMPCWalletInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMPCWalletInfo(val *MPCWalletInfo) *NullableMPCWalletInfo {
	return &NullableMPCWalletInfo{value: val, isSet: true}
}

func (v NullableMPCWalletInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMPCWalletInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

