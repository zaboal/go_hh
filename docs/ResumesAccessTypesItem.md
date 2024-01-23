# ResumesAccessTypesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **NullableBool** | Выбран ли тип видимости | [optional] 
**Id** | **string** | Идентификатор типа видимости | 
**Limit** | Pointer to **NullableFloat32** | Максимальное количество компаний в списке видимости. Возвращается только для типов &#x60;blacklist&#x60; и &#x60;whitelist&#x60; | [optional] 
**ListUrl** | Pointer to **NullableString** | Ссылка на список видимости. Возвращается только для типов &#x60;blacklist&#x60; и &#x60;whitelist&#x60; | [optional] 
**Name** | **string** | Имя типа видимости | 
**Total** | Pointer to **NullableFloat32** | Количество компаний, добавленных в соответствующий список видимости. Возвращается только для типов &#x60;blacklist&#x60; и &#x60;whitelist&#x60; | [optional] 

## Methods

### NewResumesAccessTypesItem

`func NewResumesAccessTypesItem(id string, name string, ) *ResumesAccessTypesItem`

NewResumesAccessTypesItem instantiates a new ResumesAccessTypesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesAccessTypesItemWithDefaults

`func NewResumesAccessTypesItemWithDefaults() *ResumesAccessTypesItem`

NewResumesAccessTypesItemWithDefaults instantiates a new ResumesAccessTypesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ResumesAccessTypesItem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResumesAccessTypesItem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResumesAccessTypesItem) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResumesAccessTypesItem) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ResumesAccessTypesItem) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ResumesAccessTypesItem) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetId

`func (o *ResumesAccessTypesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumesAccessTypesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumesAccessTypesItem) SetId(v string)`

SetId sets Id field to given value.


### GetLimit

`func (o *ResumesAccessTypesItem) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResumesAccessTypesItem) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResumesAccessTypesItem) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResumesAccessTypesItem) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ResumesAccessTypesItem) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ResumesAccessTypesItem) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetListUrl

`func (o *ResumesAccessTypesItem) GetListUrl() string`

GetListUrl returns the ListUrl field if non-nil, zero value otherwise.

### GetListUrlOk

`func (o *ResumesAccessTypesItem) GetListUrlOk() (*string, bool)`

GetListUrlOk returns a tuple with the ListUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListUrl

`func (o *ResumesAccessTypesItem) SetListUrl(v string)`

SetListUrl sets ListUrl field to given value.

### HasListUrl

`func (o *ResumesAccessTypesItem) HasListUrl() bool`

HasListUrl returns a boolean if a field has been set.

### SetListUrlNil

`func (o *ResumesAccessTypesItem) SetListUrlNil(b bool)`

 SetListUrlNil sets the value for ListUrl to be an explicit nil

### UnsetListUrl
`func (o *ResumesAccessTypesItem) UnsetListUrl()`

UnsetListUrl ensures that no value is present for ListUrl, not even an explicit nil
### GetName

`func (o *ResumesAccessTypesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumesAccessTypesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumesAccessTypesItem) SetName(v string)`

SetName sets Name field to given value.


### GetTotal

`func (o *ResumesAccessTypesItem) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResumesAccessTypesItem) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResumesAccessTypesItem) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResumesAccessTypesItem) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ResumesAccessTypesItem) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ResumesAccessTypesItem) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


