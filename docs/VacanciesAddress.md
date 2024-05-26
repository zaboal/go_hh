# VacanciesAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Building** | Pointer to **NullableString** | Дом | [optional] 
**City** | Pointer to **NullableString** | Город | [optional] 
**Description** | Pointer to **NullableString** | Описание | [optional] 
**Id** | Pointer to **NullableString** | Адрес из [списка доступных адресов работодателя](https://api.hh.ru/openapi/redoc#tag/Adresa-rabotodatelya/operation/get-employer-addresses) | [optional] 
**Lat** | Pointer to **NullableFloat32** | Широта | [optional] 
**Lng** | Pointer to **NullableFloat32** | Долгота | [optional] 
**Metro** | Pointer to [**NullableIncludesMetroStation**](IncludesMetroStation.md) |  | [optional] 
**MetroStations** | Pointer to [**[]IncludesMetroStation**](IncludesMetroStation.md) |  | [optional] 
**Raw** | Pointer to **NullableString** | Полный адрес | [optional] 
**Street** | Pointer to **NullableString** | Улица | [optional] 
**ShowMetroOnly** | Pointer to **NullableBool** | Показывать только метро для указанного адреса | [optional] 

## Methods

### NewVacanciesAddress

`func NewVacanciesAddress() *VacanciesAddress`

NewVacanciesAddress instantiates a new VacanciesAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesAddressWithDefaults

`func NewVacanciesAddressWithDefaults() *VacanciesAddress`

NewVacanciesAddressWithDefaults instantiates a new VacanciesAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *VacanciesAddress) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *VacanciesAddress) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *VacanciesAddress) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *VacanciesAddress) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *VacanciesAddress) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *VacanciesAddress) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *VacanciesAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacanciesAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacanciesAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VacanciesAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *VacanciesAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *VacanciesAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDescription

`func (o *VacanciesAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacanciesAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VacanciesAddress) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VacanciesAddress) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VacanciesAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacanciesAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacanciesAddress) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacanciesAddress) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLat

`func (o *VacanciesAddress) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VacanciesAddress) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VacanciesAddress) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VacanciesAddress) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *VacanciesAddress) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *VacanciesAddress) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *VacanciesAddress) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *VacanciesAddress) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *VacanciesAddress) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *VacanciesAddress) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *VacanciesAddress) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *VacanciesAddress) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetMetro

`func (o *VacanciesAddress) GetMetro() IncludesMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VacanciesAddress) GetMetroOk() (*IncludesMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VacanciesAddress) SetMetro(v IncludesMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VacanciesAddress) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *VacanciesAddress) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *VacanciesAddress) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetMetroStations

`func (o *VacanciesAddress) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *VacanciesAddress) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *VacanciesAddress) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *VacanciesAddress) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetRaw

`func (o *VacanciesAddress) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *VacanciesAddress) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *VacanciesAddress) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *VacanciesAddress) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *VacanciesAddress) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *VacanciesAddress) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetStreet

`func (o *VacanciesAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *VacanciesAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *VacanciesAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *VacanciesAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *VacanciesAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *VacanciesAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetShowMetroOnly

`func (o *VacanciesAddress) GetShowMetroOnly() bool`

GetShowMetroOnly returns the ShowMetroOnly field if non-nil, zero value otherwise.

### GetShowMetroOnlyOk

`func (o *VacanciesAddress) GetShowMetroOnlyOk() (*bool, bool)`

GetShowMetroOnlyOk returns a tuple with the ShowMetroOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMetroOnly

`func (o *VacanciesAddress) SetShowMetroOnly(v bool)`

SetShowMetroOnly sets ShowMetroOnly field to given value.

### HasShowMetroOnly

`func (o *VacanciesAddress) HasShowMetroOnly() bool`

HasShowMetroOnly returns a boolean if a field has been set.

### SetShowMetroOnlyNil

`func (o *VacanciesAddress) SetShowMetroOnlyNil(b bool)`

 SetShowMetroOnlyNil sets the value for ShowMetroOnly to be an explicit nil

### UnsetShowMetroOnly
`func (o *VacanciesAddress) UnsetShowMetroOnly()`

UnsetShowMetroOnly ensures that no value is present for ShowMetroOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


