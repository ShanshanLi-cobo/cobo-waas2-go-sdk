# TokenizationAllowlistActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**TokenizationTokenOperationSource**](TokenizationTokenOperationSource.md) |  | 
**Activation** | **bool** | Whether to activate the allowlist feature for the token. | 
**AppInitiator** | Pointer to **string** | The initiator of the tokenization activity. If you do not specify this property, the WaaS service will automatically designate the API key as the initiator. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewTokenizationAllowlistActivationRequest

`func NewTokenizationAllowlistActivationRequest(source TokenizationTokenOperationSource, activation bool, fee TransactionRequestFee, ) *TokenizationAllowlistActivationRequest`

NewTokenizationAllowlistActivationRequest instantiates a new TokenizationAllowlistActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationAllowlistActivationRequestWithDefaults

`func NewTokenizationAllowlistActivationRequestWithDefaults() *TokenizationAllowlistActivationRequest`

NewTokenizationAllowlistActivationRequestWithDefaults instantiates a new TokenizationAllowlistActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TokenizationAllowlistActivationRequest) GetSource() TokenizationTokenOperationSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenizationAllowlistActivationRequest) GetSourceOk() (*TokenizationTokenOperationSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenizationAllowlistActivationRequest) SetSource(v TokenizationTokenOperationSource)`

SetSource sets Source field to given value.


### GetActivation

`func (o *TokenizationAllowlistActivationRequest) GetActivation() bool`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *TokenizationAllowlistActivationRequest) GetActivationOk() (*bool, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *TokenizationAllowlistActivationRequest) SetActivation(v bool)`

SetActivation sets Activation field to given value.


### GetAppInitiator

`func (o *TokenizationAllowlistActivationRequest) GetAppInitiator() string`

GetAppInitiator returns the AppInitiator field if non-nil, zero value otherwise.

### GetAppInitiatorOk

`func (o *TokenizationAllowlistActivationRequest) GetAppInitiatorOk() (*string, bool)`

GetAppInitiatorOk returns a tuple with the AppInitiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInitiator

`func (o *TokenizationAllowlistActivationRequest) SetAppInitiator(v string)`

SetAppInitiator sets AppInitiator field to given value.

### HasAppInitiator

`func (o *TokenizationAllowlistActivationRequest) HasAppInitiator() bool`

HasAppInitiator returns a boolean if a field has been set.

### GetFee

`func (o *TokenizationAllowlistActivationRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TokenizationAllowlistActivationRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TokenizationAllowlistActivationRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


