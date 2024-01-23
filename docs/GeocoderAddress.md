# GeocoderAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Building** | Pointer to **NullableString** | Дом | [optional] 
**City** | Pointer to **NullableString** | Город | [optional] 
**Lat** | Pointer to **NullableFloat32** | Широта | [optional] 
**Lng** | Pointer to **NullableFloat32** | Долгота | [optional] 
**Street** | Pointer to **NullableString** | Улица | [optional] 

## Methods

### NewGeocoderAddress

`func NewGeocoderAddress() *GeocoderAddress`

NewGeocoderAddress instantiates a new GeocoderAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocoderAddressWithDefaults

`func NewGeocoderAddressWithDefaults() *GeocoderAddress`

NewGeocoderAddressWithDefaults instantiates a new GeocoderAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilding

`func (o *GeocoderAddress) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *GeocoderAddress) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *GeocoderAddress) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *GeocoderAddress) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *GeocoderAddress) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *GeocoderAddress) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *GeocoderAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GeocoderAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GeocoderAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GeocoderAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *GeocoderAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *GeocoderAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetLat

`func (o *GeocoderAddress) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GeocoderAddress) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GeocoderAddress) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *GeocoderAddress) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *GeocoderAddress) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *GeocoderAddress) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLng

`func (o *GeocoderAddress) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GeocoderAddress) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GeocoderAddress) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *GeocoderAddress) HasLng() bool`

HasLng returns a boolean if a field has been set.

### SetLngNil

`func (o *GeocoderAddress) SetLngNil(b bool)`

 SetLngNil sets the value for Lng to be an explicit nil

### UnsetLng
`func (o *GeocoderAddress) UnsetLng()`

UnsetLng ensures that no value is present for Lng, not even an explicit nil
### GetStreet

`func (o *GeocoderAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GeocoderAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GeocoderAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GeocoderAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *GeocoderAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *GeocoderAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


