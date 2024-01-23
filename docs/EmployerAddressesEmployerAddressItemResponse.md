# EmployerAddressesEmployerAddressItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Building** | Pointer to **NullableString** | Номер дома | [optional] 
**CanEdit** | Pointer to **bool** | Имеет ли текущий пользователь право редактировать этот адрес | [optional] 
**City** | **NullableString** | Город | 
**Deleted** | Pointer to **bool** | Удалён ли адрес | [optional] 
**Description** | Pointer to **NullableString** | Дополнительная информация об адресе | [optional] 
**Id** | Pointer to **string** | Идентификатор адреса | [optional] 
**Lat** | **NullableFloat32** | Географическая широта | 
**Lng** | **NullableFloat32** | Географическая долгота | 
**Manager** | Pointer to [**NullableEmployerAddressesEmployerAddressItemManager**](EmployerAddressesEmployerAddressItemManager.md) |  | [optional] 
**MetroStations** | Pointer to [**[]IncludesMetroStation**](IncludesMetroStation.md) |  | [optional] 
**Raw** | Pointer to **NullableString** | Полный адрес | [optional] 
**Street** | Pointer to **NullableString** | Улица | [optional] 

## Methods

### NewEmployerAddressesEmployerAddressItemResponse

`func NewEmployerAddressesEmployerAddressItemResponse(city NullableString, lat NullableFloat32, lng NullableFloat32, ) *EmployerAddressesEmployerAddressItemResponse`

NewEmployerAddressesEmployerAddressItemResponse instantiates a new EmployerAddressesEmployerAddressItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerAddressesEmployerAddressItemResponseWithDefaults

`func NewEmployerAddressesEmployerAddressItemResponseWithDefaults() *EmployerAddressesEmployerAddressItemResponse`

NewEmployerAddressesEmployerAddressItemResponseWithDefaults instantiates a new EmployerAddressesEmployerAddressItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *EmployerAddressesEmployerAddressItemResponse) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *EmployerAddressesEmployerAddressItemResponse) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *EmployerAddressesEmployerAddressItemResponse) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCanEdit

`func (o *EmployerAddressesEmployerAddressItemResponse) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *EmployerAddressesEmployerAddressItemResponse) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *EmployerAddressesEmployerAddressItemResponse) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCity

`func (o *EmployerAddressesEmployerAddressItemResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EmployerAddressesEmployerAddressItemResponse) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDeleted

`func (o *EmployerAddressesEmployerAddressItemResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *EmployerAddressesEmployerAddressItemResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *EmployerAddressesEmployerAddressItemResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *EmployerAddressesEmployerAddressItemResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmployerAddressesEmployerAddressItemResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmployerAddressesEmployerAddressItemResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *EmployerAddressesEmployerAddressItemResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployerAddressesEmployerAddressItemResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployerAddressesEmployerAddressItemResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLat

`func (o *EmployerAddressesEmployerAddressItemResponse) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *EmployerAddressesEmployerAddressItemResponse) SetLat(v float32)`

SetLat sets Lat field to given value.


### SetLatNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *EmployerAddressesEmployerAddressItemResponse) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *EmployerAddressesEmployerAddressItemResponse) SetLng(v float32)`

SetLng sets Lng field to given value.


### SetLngNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetManager

`func (o *EmployerAddressesEmployerAddressItemResponse) GetManager() EmployerAddressesEmployerAddressItemManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetManagerOk() (*EmployerAddressesEmployerAddressItemManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *EmployerAddressesEmployerAddressItemResponse) SetManager(v EmployerAddressesEmployerAddressItemManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *EmployerAddressesEmployerAddressItemResponse) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMetroStations

`func (o *EmployerAddressesEmployerAddressItemResponse) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *EmployerAddressesEmployerAddressItemResponse) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *EmployerAddressesEmployerAddressItemResponse) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetRaw

`func (o *EmployerAddressesEmployerAddressItemResponse) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EmployerAddressesEmployerAddressItemResponse) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EmployerAddressesEmployerAddressItemResponse) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetStreet

`func (o *EmployerAddressesEmployerAddressItemResponse) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *EmployerAddressesEmployerAddressItemResponse) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *EmployerAddressesEmployerAddressItemResponse) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *EmployerAddressesEmployerAddressItemResponse) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *EmployerAddressesEmployerAddressItemResponse) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *EmployerAddressesEmployerAddressItemResponse) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


