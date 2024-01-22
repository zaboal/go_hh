# MetroLineStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор станции | 
**Lat** | **float32** | Широта расположения станции | 
**Line** | [**MetroMetroLine**](MetroMetroLine.md) | Линия метро | 
**Lng** | **float32** | Долгота расположения станции | 
**Name** | **string** | Название станции | 
**Order** | **int32** | Порядковый номер станции на линии, начиная с 0 | 

## Methods

### NewMetroLineStation

`func NewMetroLineStation(id string, lat float32, line MetroMetroLine, lng float32, name string, order int32, ) *MetroLineStation`

NewMetroLineStation instantiates a new MetroLineStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroLineStationWithDefaults

`func NewMetroLineStationWithDefaults() *MetroLineStation`

NewMetroLineStationWithDefaults instantiates a new MetroLineStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetroLineStation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetroLineStation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetroLineStation) SetId(v string)`

SetId sets Id field to given value.


### GetLat

`func (o *MetroLineStation) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *MetroLineStation) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *MetroLineStation) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLine

`func (o *MetroLineStation) GetLine() MetroMetroLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *MetroLineStation) GetLineOk() (*MetroMetroLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *MetroLineStation) SetLine(v MetroMetroLine)`

SetLine sets Line field to given value.


### GetLng

`func (o *MetroLineStation) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *MetroLineStation) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *MetroLineStation) SetLng(v float32)`

SetLng sets Lng field to given value.


### GetName

`func (o *MetroLineStation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetroLineStation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetroLineStation) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *MetroLineStation) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MetroLineStation) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MetroLineStation) SetOrder(v int32)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


