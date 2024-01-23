# ManagerAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employer** | [**ManagerAccountCompany**](ManagerAccountCompany.md) |  | 
**Id** | **string** | Идентификатор рабочего аккаунта менеджера | 

## Methods

### NewManagerAccount

`func NewManagerAccount(employer ManagerAccountCompany, id string, ) *ManagerAccount`

NewManagerAccount instantiates a new ManagerAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerAccountWithDefaults

`func NewManagerAccountWithDefaults() *ManagerAccount`

NewManagerAccountWithDefaults instantiates a new ManagerAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployer

`func (o *ManagerAccount) GetEmployer() ManagerAccountCompany`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ManagerAccount) GetEmployerOk() (*ManagerAccountCompany, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ManagerAccount) SetEmployer(v ManagerAccountCompany)`

SetEmployer sets Employer field to given value.


### GetId

`func (o *ManagerAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagerAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagerAccount) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


