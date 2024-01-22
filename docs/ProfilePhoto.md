# ProfilePhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var40** | Pointer to **string** |  | [optional] 
**Var100** | Pointer to **string** |  | [optional] 
**Var500** | Pointer to **string** |  | [optional] 
**Id** | **string** | Уникальный идентификатор изображения | 
**Medium** | **string** | URL среднего по размеру изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 
**Small** | **string** | URL уменьшенного изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 

## Methods

### NewProfilePhoto

`func NewProfilePhoto(id string, medium string, small string, ) *ProfilePhoto`

NewProfilePhoto instantiates a new ProfilePhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilePhotoWithDefaults

`func NewProfilePhotoWithDefaults() *ProfilePhoto`

NewProfilePhotoWithDefaults instantiates a new ProfilePhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar40

`func (o *ProfilePhoto) GetVar40() string`

GetVar40 returns the Var40 field if non-nil, zero value otherwise.

### GetVar40Ok

`func (o *ProfilePhoto) GetVar40Ok() (*string, bool)`

GetVar40Ok returns a tuple with the Var40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar40

`func (o *ProfilePhoto) SetVar40(v string)`

SetVar40 sets Var40 field to given value.

### HasVar40

`func (o *ProfilePhoto) HasVar40() bool`

HasVar40 returns a boolean if a field has been set.

### GetVar100

`func (o *ProfilePhoto) GetVar100() string`

GetVar100 returns the Var100 field if non-nil, zero value otherwise.

### GetVar100Ok

`func (o *ProfilePhoto) GetVar100Ok() (*string, bool)`

GetVar100Ok returns a tuple with the Var100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar100

`func (o *ProfilePhoto) SetVar100(v string)`

SetVar100 sets Var100 field to given value.

### HasVar100

`func (o *ProfilePhoto) HasVar100() bool`

HasVar100 returns a boolean if a field has been set.

### GetVar500

`func (o *ProfilePhoto) GetVar500() string`

GetVar500 returns the Var500 field if non-nil, zero value otherwise.

### GetVar500Ok

`func (o *ProfilePhoto) GetVar500Ok() (*string, bool)`

GetVar500Ok returns a tuple with the Var500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar500

`func (o *ProfilePhoto) SetVar500(v string)`

SetVar500 sets Var500 field to given value.

### HasVar500

`func (o *ProfilePhoto) HasVar500() bool`

HasVar500 returns a boolean if a field has been set.

### GetId

`func (o *ProfilePhoto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfilePhoto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfilePhoto) SetId(v string)`

SetId sets Id field to given value.


### GetMedium

`func (o *ProfilePhoto) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ProfilePhoto) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ProfilePhoto) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetSmall

`func (o *ProfilePhoto) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ProfilePhoto) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ProfilePhoto) SetSmall(v string)`

SetSmall sets Small field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


