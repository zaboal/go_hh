# EmployerAddressesEmployerAddressItem

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

### NewEmployerAddressesEmployerAddressItem

`func NewEmployerAddressesEmployerAddressItem(city NullableString, lat NullableFloat32, lng NullableFloat32, ) *EmployerAddressesEmployerAddressItem`

NewEmployerAddressesEmployerAddressItem instantiates a new EmployerAddressesEmployerAddressItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerAddressesEmployerAddressItemWithDefaults

`func NewEmployerAddressesEmployerAddressItemWithDefaults() *EmployerAddressesEmployerAddressItem`

NewEmployerAddressesEmployerAddressItemWithDefaults instantiates a new EmployerAddressesEmployerAddressItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *EmployerAddressesEmployerAddressItem) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *EmployerAddressesEmployerAddressItem) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *EmployerAddressesEmployerAddressItem) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *EmployerAddressesEmployerAddressItem) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *EmployerAddressesEmployerAddressItem) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *EmployerAddressesEmployerAddressItem) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCanEdit

`func (o *EmployerAddressesEmployerAddressItem) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *EmployerAddressesEmployerAddressItem) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *EmployerAddressesEmployerAddressItem) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *EmployerAddressesEmployerAddressItem) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetCity

`func (o *EmployerAddressesEmployerAddressItem) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EmployerAddressesEmployerAddressItem) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EmployerAddressesEmployerAddressItem) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *EmployerAddressesEmployerAddressItem) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *EmployerAddressesEmployerAddressItem) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDeleted

`func (o *EmployerAddressesEmployerAddressItem) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *EmployerAddressesEmployerAddressItem) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *EmployerAddressesEmployerAddressItem) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *EmployerAddressesEmployerAddressItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *EmployerAddressesEmployerAddressItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmployerAddressesEmployerAddressItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmployerAddressesEmployerAddressItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmployerAddressesEmployerAddressItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EmployerAddressesEmployerAddressItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EmployerAddressesEmployerAddressItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *EmployerAddressesEmployerAddressItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployerAddressesEmployerAddressItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployerAddressesEmployerAddressItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployerAddressesEmployerAddressItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLat

`func (o *EmployerAddressesEmployerAddressItem) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *EmployerAddressesEmployerAddressItem) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *EmployerAddressesEmployerAddressItem) SetLat(v float32)`

SetLat sets Lat field to given value.


### SetLatNil

`func (o *EmployerAddressesEmployerAddressItem) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *EmployerAddressesEmployerAddressItem) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *EmployerAddressesEmployerAddressItem) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *EmployerAddressesEmployerAddressItem) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *EmployerAddressesEmployerAddressItem) SetLng(v float32)`

SetLng sets Lng field to given value.


### SetLngNil

`func (o *EmployerAddressesEmployerAddressItem) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *EmployerAddressesEmployerAddressItem) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetManager

`func (o *EmployerAddressesEmployerAddressItem) GetManager() EmployerAddressesEmployerAddressItemManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *EmployerAddressesEmployerAddressItem) GetManagerOk() (*EmployerAddressesEmployerAddressItemManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *EmployerAddressesEmployerAddressItem) SetManager(v EmployerAddressesEmployerAddressItemManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *EmployerAddressesEmployerAddressItem) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *EmployerAddressesEmployerAddressItem) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *EmployerAddressesEmployerAddressItem) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetMetroStations

`func (o *EmployerAddressesEmployerAddressItem) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *EmployerAddressesEmployerAddressItem) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *EmployerAddressesEmployerAddressItem) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *EmployerAddressesEmployerAddressItem) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetRaw

`func (o *EmployerAddressesEmployerAddressItem) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EmployerAddressesEmployerAddressItem) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EmployerAddressesEmployerAddressItem) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EmployerAddressesEmployerAddressItem) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *EmployerAddressesEmployerAddressItem) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *EmployerAddressesEmployerAddressItem) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetStreet

`func (o *EmployerAddressesEmployerAddressItem) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *EmployerAddressesEmployerAddressItem) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *EmployerAddressesEmployerAddressItem) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *EmployerAddressesEmployerAddressItem) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *EmployerAddressesEmployerAddressItem) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *EmployerAddressesEmployerAddressItem) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


