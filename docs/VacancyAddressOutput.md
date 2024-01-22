# VacancyAddressOutput

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

## Methods

### NewVacancyAddressOutput

`func NewVacancyAddressOutput() *VacancyAddressOutput`

NewVacancyAddressOutput instantiates a new VacancyAddressOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyAddressOutputWithDefaults

`func NewVacancyAddressOutputWithDefaults() *VacancyAddressOutput`

NewVacancyAddressOutputWithDefaults instantiates a new VacancyAddressOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *VacancyAddressOutput) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *VacancyAddressOutput) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *VacancyAddressOutput) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *VacancyAddressOutput) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *VacancyAddressOutput) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *VacancyAddressOutput) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *VacancyAddressOutput) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyAddressOutput) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyAddressOutput) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VacancyAddressOutput) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *VacancyAddressOutput) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *VacancyAddressOutput) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLat

`func (o *VacancyAddressOutput) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VacancyAddressOutput) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VacancyAddressOutput) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VacancyAddressOutput) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *VacancyAddressOutput) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *VacancyAddressOutput) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *VacancyAddressOutput) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *VacancyAddressOutput) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *VacancyAddressOutput) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *VacancyAddressOutput) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *VacancyAddressOutput) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *VacancyAddressOutput) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetStreet

`func (o *VacancyAddressOutput) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *VacancyAddressOutput) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *VacancyAddressOutput) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *VacancyAddressOutput) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *VacancyAddressOutput) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *VacancyAddressOutput) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetDescription

`func (o *VacancyAddressOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyAddressOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyAddressOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyAddressOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VacancyAddressOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VacancyAddressOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetroStations

`func (o *VacancyAddressOutput) GetMetroStations() []IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *VacancyAddressOutput) GetMetroStationsOk() (*[]IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *VacancyAddressOutput) SetMetroStations(v []IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *VacancyAddressOutput) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


