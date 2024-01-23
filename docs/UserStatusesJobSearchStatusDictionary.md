# UserStatusesJobSearchStatusDictionary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Детальное описание статуса | [optional] 
**Id** | **string** | Идентификатор статуса | 
**Name** | **string** | Название статуса | 
**AllowedWhenCanStartWorkStatuses** | **[]string** | Массив с доступными для выбора &#x60;when_can_start_work_status&#x60; | 

## Methods

### NewUserStatusesJobSearchStatusDictionary

`func NewUserStatusesJobSearchStatusDictionary(id string, name string, allowedWhenCanStartWorkStatuses []string, ) *UserStatusesJobSearchStatusDictionary`

NewUserStatusesJobSearchStatusDictionary instantiates a new UserStatusesJobSearchStatusDictionary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusesJobSearchStatusDictionaryWithDefaults

`func NewUserStatusesJobSearchStatusDictionaryWithDefaults() *UserStatusesJobSearchStatusDictionary`

NewUserStatusesJobSearchStatusDictionaryWithDefaults instantiates a new UserStatusesJobSearchStatusDictionary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UserStatusesJobSearchStatusDictionary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserStatusesJobSearchStatusDictionary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserStatusesJobSearchStatusDictionary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserStatusesJobSearchStatusDictionary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UserStatusesJobSearchStatusDictionary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserStatusesJobSearchStatusDictionary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserStatusesJobSearchStatusDictionary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserStatusesJobSearchStatusDictionary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserStatusesJobSearchStatusDictionary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserStatusesJobSearchStatusDictionary) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedWhenCanStartWorkStatuses

`func (o *UserStatusesJobSearchStatusDictionary) GetAllowedWhenCanStartWorkStatuses() []string`

GetAllowedWhenCanStartWorkStatuses returns the AllowedWhenCanStartWorkStatuses field if non-nil, zero value otherwise.

### GetAllowedWhenCanStartWorkStatusesOk

`func (o *UserStatusesJobSearchStatusDictionary) GetAllowedWhenCanStartWorkStatusesOk() (*[]string, bool)`

GetAllowedWhenCanStartWorkStatusesOk returns a tuple with the AllowedWhenCanStartWorkStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedWhenCanStartWorkStatuses

`func (o *UserStatusesJobSearchStatusDictionary) SetAllowedWhenCanStartWorkStatuses(v []string)`

SetAllowedWhenCanStartWorkStatuses sets AllowedWhenCanStartWorkStatuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


