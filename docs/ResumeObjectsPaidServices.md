# ResumeObjectsPaidServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Активна ли в данный момент услуга | 
**Expires** | Pointer to **string** | Время окончания действия услуги, если услуга активна | [optional] 
**Id** | **string** | Идентификатор услуги | 
**Name** | **string** | Название услуги | 

## Methods

### NewResumeObjectsPaidServices

`func NewResumeObjectsPaidServices(active bool, id string, name string, ) *ResumeObjectsPaidServices`

NewResumeObjectsPaidServices instantiates a new ResumeObjectsPaidServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsPaidServicesWithDefaults

`func NewResumeObjectsPaidServicesWithDefaults() *ResumeObjectsPaidServices`

NewResumeObjectsPaidServicesWithDefaults instantiates a new ResumeObjectsPaidServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ResumeObjectsPaidServices) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResumeObjectsPaidServices) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResumeObjectsPaidServices) SetActive(v bool)`

SetActive sets Active field to given value.


### GetExpires

`func (o *ResumeObjectsPaidServices) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ResumeObjectsPaidServices) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ResumeObjectsPaidServices) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ResumeObjectsPaidServices) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetId

`func (o *ResumeObjectsPaidServices) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsPaidServices) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsPaidServices) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResumeObjectsPaidServices) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumeObjectsPaidServices) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumeObjectsPaidServices) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


