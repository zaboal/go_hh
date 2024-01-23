# VacancyAddressRawOutput

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

## Methods

### NewVacancyAddressRawOutput

`func NewVacancyAddressRawOutput() *VacancyAddressRawOutput`

NewVacancyAddressRawOutput instantiates a new VacancyAddressRawOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyAddressRawOutputWithDefaults

`func NewVacancyAddressRawOutputWithDefaults() *VacancyAddressRawOutput`

NewVacancyAddressRawOutputWithDefaults instantiates a new VacancyAddressRawOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *VacancyAddressRawOutput) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *VacancyAddressRawOutput) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *VacancyAddressRawOutput) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *VacancyAddressRawOutput) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *VacancyAddressRawOutput) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *VacancyAddressRawOutput) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *VacancyAddressRawOutput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyAddressRawOutput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyAddressRawOutput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VacancyAddressRawOutput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *VacancyAddressRawOutput) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *VacancyAddressRawOutput) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDescription

`func (o *VacancyAddressRawOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyAddressRawOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyAddressRawOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyAddressRawOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VacancyAddressRawOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VacancyAddressRawOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VacancyAddressRawOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyAddressRawOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyAddressRawOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyAddressRawOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyAddressRawOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyAddressRawOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLat

`func (o *VacancyAddressRawOutput) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VacancyAddressRawOutput) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VacancyAddressRawOutput) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VacancyAddressRawOutput) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *VacancyAddressRawOutput) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *VacancyAddressRawOutput) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *VacancyAddressRawOutput) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *VacancyAddressRawOutput) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *VacancyAddressRawOutput) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *VacancyAddressRawOutput) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *VacancyAddressRawOutput) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *VacancyAddressRawOutput) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetMetro

`func (o *VacancyAddressRawOutput) GetMetro() IncludesMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VacancyAddressRawOutput) GetMetroOk() (*IncludesMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VacancyAddressRawOutput) SetMetro(v IncludesMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VacancyAddressRawOutput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *VacancyAddressRawOutput) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *VacancyAddressRawOutput) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetMetroStations

`func (o *VacancyAddressRawOutput) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *VacancyAddressRawOutput) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *VacancyAddressRawOutput) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *VacancyAddressRawOutput) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetRaw

`func (o *VacancyAddressRawOutput) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *VacancyAddressRawOutput) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *VacancyAddressRawOutput) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *VacancyAddressRawOutput) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *VacancyAddressRawOutput) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *VacancyAddressRawOutput) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetStreet

`func (o *VacancyAddressRawOutput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *VacancyAddressRawOutput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *VacancyAddressRawOutput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *VacancyAddressRawOutput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *VacancyAddressRawOutput) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *VacancyAddressRawOutput) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


