# ResumeObjectsMetroStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор станции метро | 
**Lat** | **float32** | Широта | 
**Line** | [**ResumeObjectsMetroLine**](ResumeObjectsMetroLine.md) | Линия метро | 
**Lng** | **float32** | Долгота | 
**Name** | Pointer to **string** | Название станции метро | [optional] 
**Order** | **float32** | Порядковый номер станции в линии метро | 

## Methods

### NewResumeObjectsMetroStation

`func NewResumeObjectsMetroStation(id string, lat float32, line ResumeObjectsMetroLine, lng float32, order float32, ) *ResumeObjectsMetroStation`

NewResumeObjectsMetroStation instantiates a new ResumeObjectsMetroStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsMetroStationWithDefaults

`func NewResumeObjectsMetroStationWithDefaults() *ResumeObjectsMetroStation`

NewResumeObjectsMetroStationWithDefaults instantiates a new ResumeObjectsMetroStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResumeObjectsMetroStation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsMetroStation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsMetroStation) SetId(v string)`

SetId sets Id field to given value.


### GetLat

`func (o *ResumeObjectsMetroStation) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *ResumeObjectsMetroStation) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *ResumeObjectsMetroStation) SetLat(v float32)`

SetLat sets Lat field to given value.


### GetLine

`func (o *ResumeObjectsMetroStation) GetLine() ResumeObjectsMetroLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ResumeObjectsMetroStation) GetLineOk() (*ResumeObjectsMetroLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ResumeObjectsMetroStation) SetLine(v ResumeObjectsMetroLine)`

SetLine sets Line field to given value.


### GetLng

`func (o *ResumeObjectsMetroStation) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *ResumeObjectsMetroStation) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *ResumeObjectsMetroStation) SetLng(v float32)`

SetLng sets Lng field to given value.


### GetName

`func (o *ResumeObjectsMetroStation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumeObjectsMetroStation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumeObjectsMetroStation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResumeObjectsMetroStation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *ResumeObjectsMetroStation) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ResumeObjectsMetroStation) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ResumeObjectsMetroStation) SetOrder(v float32)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


