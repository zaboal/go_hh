# MetroMetroItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор города | 
**Lines** | [**[]MetroMetroLineWithStations**](MetroMetroLineWithStations.md) | Список линий метро в городе | 
**Name** | **string** | Название города | 
**Url** | **string** | URL запроса | 

## Methods

### NewMetroMetroItem

`func NewMetroMetroItem(id string, lines []MetroMetroLineWithStations, name string, url string, ) *MetroMetroItem`

NewMetroMetroItem instantiates a new MetroMetroItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroMetroItemWithDefaults

`func NewMetroMetroItemWithDefaults() *MetroMetroItem`

NewMetroMetroItemWithDefaults instantiates a new MetroMetroItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetroMetroItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetroMetroItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetroMetroItem) SetId(v string)`

SetId sets Id field to given value.


### GetLines

`func (o *MetroMetroItem) GetLines() []MetroMetroLineWithStations`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *MetroMetroItem) GetLinesOk() (*[]MetroMetroLineWithStations, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *MetroMetroItem) SetLines(v []MetroMetroLineWithStations)`

SetLines sets Lines field to given value.


### GetName

`func (o *MetroMetroItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetroMetroItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetroMetroItem) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *MetroMetroItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MetroMetroItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MetroMetroItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


