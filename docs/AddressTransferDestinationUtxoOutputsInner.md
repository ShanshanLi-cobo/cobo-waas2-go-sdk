# AddressTransferDestinationUtxoOutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The destination address. | 
**Amount** | **string** | The transfer amount. For example, if you trade 1.5 BTC, then the value is &#x60;1.5&#x60;.  | 

## Methods

### NewAddressTransferDestinationUtxoOutputsInner

`func NewAddressTransferDestinationUtxoOutputsInner(address string, amount string, ) *AddressTransferDestinationUtxoOutputsInner`

NewAddressTransferDestinationUtxoOutputsInner instantiates a new AddressTransferDestinationUtxoOutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransferDestinationUtxoOutputsInnerWithDefaults

`func NewAddressTransferDestinationUtxoOutputsInnerWithDefaults() *AddressTransferDestinationUtxoOutputsInner`

NewAddressTransferDestinationUtxoOutputsInnerWithDefaults instantiates a new AddressTransferDestinationUtxoOutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressTransferDestinationUtxoOutputsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransferDestinationUtxoOutputsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransferDestinationUtxoOutputsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


