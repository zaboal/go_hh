# MetroMetroLineWithStations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HexColor** | **string** | Цвет линии в HEX-формате &#x60;RRGGBB&#x60; (от &#x60;000000&#x60; до &#x60;FFFFFF&#x60;) | 
**Id** | **string** | Идентификатор линии | 
**Name** | **string** | Название линии | 
**Stations** | [**[]MetroLineStation**](MetroLineStation.md) | Список станций метро на линии | 

## Methods

### NewMetroMetroLineWithStations

`func NewMetroMetroLineWithStations(hexColor string, id string, name string, stations []MetroLineStation, ) *MetroMetroLineWithStations`

NewMetroMetroLineWithStations instantiates a new MetroMetroLineWithStations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroMetroLineWithStationsWithDefaults

`func NewMetroMetroLineWithStationsWithDefaults() *MetroMetroLineWithStations`

NewMetroMetroLineWithStationsWithDefaults instantiates a new MetroMetroLineWithStations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHexColor

`func (o *MetroMetroLineWithStations) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *MetroMetroLineWithStations) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *MetroMetroLineWithStations) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.


### GetId

`func (o *MetroMetroLineWithStations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetroMetroLineWithStations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetroMetroLineWithStations) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MetroMetroLineWithStations) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetroMetroLineWithStations) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetroMetroLineWithStations) SetName(v string)`

SetName sets Name field to given value.


### GetStations

`func (o *MetroMetroLineWithStations) GetStations() []MetroLineStation`

GetStations returns the Stations field if non-nil, zero value otherwise.

### GetStationsOk

`func (o *MetroMetroLineWithStations) GetStationsOk() (*[]MetroLineStation, bool)`

GetStationsOk returns a tuple with the Stations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStations

`func (o *MetroMetroLineWithStations) SetStations(v []MetroLineStation)`

SetStations sets Stations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


