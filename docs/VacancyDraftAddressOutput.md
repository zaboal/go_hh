# VacancyDraftAddressOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Building** | Pointer to **NullableString** | Дом | [optional] 
**City** | Pointer to **NullableString** | Город | [optional] 
**Lat** | Pointer to **NullableFloat32** | Широта | [optional] 
**Lng** | Pointer to **NullableFloat32** | Долгота | [optional] 
**Street** | Pointer to **NullableString** | Улица | [optional] 
**Description** | Pointer to **NullableString** | Описание | [optional] 
**MetroStations** | Pointer to [**[]IncludesMetroStation**](IncludesMetroStation.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Адрес из [списка доступных адресов работодателя](#tag/Adresa-rabotodatelya/operation/get-employer-addresses) | [optional] 
**ShowMetroOnly** | Pointer to **NullableBool** | Показывать только метро для указанного адреса | [optional] 

## Methods

### NewVacancyDraftAddressOutput

`func NewVacancyDraftAddressOutput() *VacancyDraftAddressOutput`

NewVacancyDraftAddressOutput instantiates a new VacancyDraftAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftAddressOutputWithDefaults

`func NewVacancyDraftAddressOutputWithDefaults() *VacancyDraftAddressOutput`

NewVacancyDraftAddressOutputWithDefaults instantiates a new VacancyDraftAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *VacancyDraftAddressOutput) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *VacancyDraftAddressOutput) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *VacancyDraftAddressOutput) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *VacancyDraftAddressOutput) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *VacancyDraftAddressOutput) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *VacancyDraftAddressOutput) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *VacancyDraftAddressOutput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyDraftAddressOutput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyDraftAddressOutput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VacancyDraftAddressOutput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *VacancyDraftAddressOutput) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *VacancyDraftAddressOutput) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLat

`func (o *VacancyDraftAddressOutput) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VacancyDraftAddressOutput) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VacancyDraftAddressOutput) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VacancyDraftAddressOutput) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *VacancyDraftAddressOutput) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *VacancyDraftAddressOutput) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *VacancyDraftAddressOutput) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *VacancyDraftAddressOutput) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *VacancyDraftAddressOutput) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *VacancyDraftAddressOutput) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *VacancyDraftAddressOutput) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *VacancyDraftAddressOutput) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetStreet

`func (o *VacancyDraftAddressOutput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *VacancyDraftAddressOutput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *VacancyDraftAddressOutput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *VacancyDraftAddressOutput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *VacancyDraftAddressOutput) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *VacancyDraftAddressOutput) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetDescription

`func (o *VacancyDraftAddressOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyDraftAddressOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyDraftAddressOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyDraftAddressOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VacancyDraftAddressOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VacancyDraftAddressOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetroStations

`func (o *VacancyDraftAddressOutput) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *VacancyDraftAddressOutput) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *VacancyDraftAddressOutput) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *VacancyDraftAddressOutput) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetId

`func (o *VacancyDraftAddressOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftAddressOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftAddressOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDraftAddressOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyDraftAddressOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyDraftAddressOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetShowMetroOnly

`func (o *VacancyDraftAddressOutput) GetShowMetroOnly() bool`

GetShowMetroOnly returns the ShowMetroOnly field if non-nil, zero value otherwise.

### GetShowMetroOnlyOk

`func (o *VacancyDraftAddressOutput) GetShowMetroOnlyOk() (*bool, bool)`

GetShowMetroOnlyOk returns a tuple with the ShowMetroOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMetroOnly

`func (o *VacancyDraftAddressOutput) SetShowMetroOnly(v bool)`

SetShowMetroOnly sets ShowMetroOnly field to given value.

### HasShowMetroOnly

`func (o *VacancyDraftAddressOutput) HasShowMetroOnly() bool`

HasShowMetroOnly returns a boolean if a field has been set.

### SetShowMetroOnlyNil

`func (o *VacancyDraftAddressOutput) SetShowMetroOnlyNil(b bool)`

 SetShowMetroOnlyNil sets the value for ShowMetroOnly to be an explicit nil

### UnsetShowMetroOnly
`func (o *VacancyDraftAddressOutput) UnsetShowMetroOnly()`

UnsetShowMetroOnly ensures that no value is present for ShowMetroOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


