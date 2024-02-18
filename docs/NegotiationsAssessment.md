# NegotiationsAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор | 
**Name** | **string** | Название | 
**Actions** | [**[]NegotiationsAction**](NegotiationsAction.md) | Инструменты оценки | 

## Methods

### NewNegotiationsAssessment

`func NewNegotiationsAssessment(id string, name string, actions []NegotiationsAction, ) *NegotiationsAssessment`

NewNegotiationsAssessment instantiates a new NegotiationsAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsAssessmentWithDefaults

`func NewNegotiationsAssessmentWithDefaults() *NegotiationsAssessment`

NewNegotiationsAssessmentWithDefaults instantiates a new NegotiationsAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NegotiationsAssessment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsAssessment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsAssessment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NegotiationsAssessment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NegotiationsAssessment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NegotiationsAssessment) SetName(v string)`

SetName sets Name field to given value.


### GetActions

`func (o *NegotiationsAssessment) GetActions() []NegotiationsAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NegotiationsAssessment) GetActionsOk() (*[]NegotiationsAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NegotiationsAssessment) SetActions(v []NegotiationsAction)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


