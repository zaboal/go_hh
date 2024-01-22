# MetroCityMetroItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор города | 
**Lines** | [**[]MetroMetroLineWithStations**](MetroMetroLineWithStations.md) | Список линий метро в городе | 
**Name** | **string** | Название города | 

## Methods

### NewMetroCityMetroItem

`func NewMetroCityMetroItem(id string, lines []MetroMetroLineWithStations, name string, ) *MetroCityMetroItem`

NewMetroCityMetroItem instantiates a new MetroCityMetroItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroCityMetroItemWithDefaults

`func NewMetroCityMetroItemWithDefaults() *MetroCityMetroItem`

NewMetroCityMetroItemWithDefaults instantiates a new MetroCityMetroItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetroCityMetroItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetroCityMetroItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetroCityMetroItem) SetId(v string)`

SetId sets Id field to given value.


### GetLines

`func (o *MetroCityMetroItem) GetLines() []MetroMetroLineWithStations`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MetroCityMetroItem) GetLinesOk() (*[]MetroMetroLineWithStations, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MetroCityMetroItem) SetLines(v []MetroMetroLineWithStations)`

SetLines sets Lines field to given value.


### GetName

`func (o *MetroCityMetroItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetroCityMetroItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetroCityMetroItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


