# EstimateWithdrawFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | [**ActivityType**](ActivityType.md) |  | 
**RequestId** | Pointer to **string** | The request ID that is used to track a request. The request ID is provided by you and must be unique within your organization. | [optional] 
**StakingId** | **string** | The ID of the corresponding staking position. | 
**Amount** | Pointer to **string** | The amount to withdraw. | [optional] 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 

## Methods

### NewEstimateWithdrawFee

`func NewEstimateWithdrawFee(activityType ActivityType, stakingId string, fee TransactionRequestFee, ) *EstimateWithdrawFee`

NewEstimateWithdrawFee instantiates a new EstimateWithdrawFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateWithdrawFeeWithDefaults

`func NewEstimateWithdrawFeeWithDefaults() *EstimateWithdrawFee`

NewEstimateWithdrawFeeWithDefaults instantiates a new EstimateWithdrawFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *EstimateWithdrawFee) GetActivityType() ActivityType`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *EstimateWithdrawFee) GetActivityTypeOk() (*ActivityType, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *EstimateWithdrawFee) SetActivityType(v ActivityType)`

SetActivityType sets ActivityType field to given value.


### GetRequestId

`func (o *EstimateWithdrawFee) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EstimateWithdrawFee) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EstimateWithdrawFee) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EstimateWithdrawFee) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStakingId

`func (o *EstimateWithdrawFee) GetStakingId() string`

GetStakingId returns the StakingId field if non-nil, zero value otherwise.

### GetStakingIdOk

`func (o *EstimateWithdrawFee) GetStakingIdOk() (*string, bool)`

GetStakingIdOk returns a tuple with the StakingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingId

`func (o *EstimateWithdrawFee) SetStakingId(v string)`

SetStakingId sets StakingId field to given value.


### GetAmount

`func (o *EstimateWithdrawFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EstimateWithdrawFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EstimateWithdrawFee) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EstimateWithdrawFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *EstimateWithdrawFee) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EstimateWithdrawFee) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EstimateWithdrawFee) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


