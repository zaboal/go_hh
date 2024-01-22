# ArtifactsArtifactItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор изображения | 
**Medium** | Pointer to **NullableString** | URL для получения среднего по размеру изображения или &#x60;null&#x60;, если изображение ещё не готово | [optional] 
**Small** | Pointer to **NullableString** | URL для получения уменьшенного изображения или &#x60;null&#x60;, если изображение ещё не готово | [optional] 
**State** | [**ArtifactsState**](ArtifactsState.md) |  | 

## Methods

### NewArtifactsArtifactItem

`func NewArtifactsArtifactItem(id string, state ArtifactsState, ) *ArtifactsArtifactItem`

NewArtifactsArtifactItem instantiates a new ArtifactsArtifactItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsArtifactItemWithDefaults

`func NewArtifactsArtifactItemWithDefaults() *ArtifactsArtifactItem`

NewArtifactsArtifactItemWithDefaults instantiates a new ArtifactsArtifactItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtifactsArtifactItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsArtifactItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsArtifactItem) SetId(v string)`

SetId sets Id field to given value.


### GetMedium

`func (o *ArtifactsArtifactItem) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ArtifactsArtifactItem) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ArtifactsArtifactItem) SetMedium(v string)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *ArtifactsArtifactItem) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### SetMediumNil

`func (o *ArtifactsArtifactItem) SetMediumNil(b bool)`

 SetMediumNil sets the value for Medium to be an explicit nil

### UnsetMedium
`func (o *ArtifactsArtifactItem) UnsetMedium()`

UnsetMedium ensures that no value is present for Medium, not even an explicit nil
### GetSmall

`func (o *ArtifactsArtifactItem) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ArtifactsArtifactItem) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ArtifactsArtifactItem) SetSmall(v string)`

SetSmall sets Small field to given value.

### HasSmall

`func (o *ArtifactsArtifactItem) HasSmall() bool`

HasSmall returns a boolean if a field has been set.

### SetSmallNil

`func (o *ArtifactsArtifactItem) SetSmallNil(b bool)`

 SetSmallNil sets the value for Small to be an explicit nil

### UnsetSmall
`func (o *ArtifactsArtifactItem) UnsetSmall()`

UnsetSmall ensures that no value is present for Small, not even an explicit nil
### GetState

`func (o *ArtifactsArtifactItem) GetState() ArtifactsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ArtifactsArtifactItem) GetStateOk() (*ArtifactsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ArtifactsArtifactItem) SetState(v ArtifactsState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


