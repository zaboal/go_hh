# ResumesCreationAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **float32** | Количество созданных резюме | 
**IsCreationAvailable** | **bool** | Доступно ли создание новых резюме для данного пользователя | 
**Max** | **float32** | Максимально возможное количество резюме | 
**Remaining** | **float32** | Количество доступных для создания резюме | 

## Methods

### NewResumesCreationAvailability

`func NewResumesCreationAvailability(created float32, isCreationAvailable bool, max float32, remaining float32, ) *ResumesCreationAvailability`

NewResumesCreationAvailability instantiates a new ResumesCreationAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesCreationAvailabilityWithDefaults

`func NewResumesCreationAvailabilityWithDefaults() *ResumesCreationAvailability`

NewResumesCreationAvailabilityWithDefaults instantiates a new ResumesCreationAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResumesCreationAvailability) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResumesCreationAvailability) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResumesCreationAvailability) SetCreated(v float32)`

SetCreated sets Created field to given value.


### GetIsCreationAvailable

`func (o *ResumesCreationAvailability) GetIsCreationAvailable() bool`

GetIsCreationAvailable returns the IsCreationAvailable field if non-nil, zero value otherwise.

### GetIsCreationAvailableOk

`func (o *ResumesCreationAvailability) GetIsCreationAvailableOk() (*bool, bool)`

GetIsCreationAvailableOk returns a tuple with the IsCreationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreationAvailable

`func (o *ResumesCreationAvailability) SetIsCreationAvailable(v bool)`

SetIsCreationAvailable sets IsCreationAvailable field to given value.


### GetMax

`func (o *ResumesCreationAvailability) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ResumesCreationAvailability) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ResumesCreationAvailability) SetMax(v float32)`

SetMax sets Max field to given value.


### GetRemaining

`func (o *ResumesCreationAvailability) GetRemaining() float32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *ResumesCreationAvailability) GetRemainingOk() (*float32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *ResumesCreationAvailability) SetRemaining(v float32)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


