# ResumeObjectsPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var40** | Pointer to **NullableString** | URL изображения размером 40x40 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Var100** | Pointer to **NullableString** | URL изображения размером 100x100 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Var500** | Pointer to **NullableString** | URL изображения размером 500x500 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Medium** | **string** | URL среднего по размеру изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | 
**Small** | **string** | URL уменьшенного изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | 
**Id** | **string** | Идентификатор | 

## Methods

### NewResumeObjectsPhoto

`func NewResumeObjectsPhoto(medium string, small string, id string, ) *ResumeObjectsPhoto`

NewResumeObjectsPhoto instantiates a new ResumeObjectsPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsPhotoWithDefaults

`func NewResumeObjectsPhotoWithDefaults() *ResumeObjectsPhoto`

NewResumeObjectsPhotoWithDefaults instantiates a new ResumeObjectsPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar40

`func (o *ResumeObjectsPhoto) GetVar40() string`

GetVar40 returns the Var40 field if non-nil, zero value otherwise.

### GetVar40Ok

`func (o *ResumeObjectsPhoto) GetVar40Ok() (*string, bool)`

GetVar40Ok returns a tuple with the Var40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar40

`func (o *ResumeObjectsPhoto) SetVar40(v string)`

SetVar40 sets Var40 field to given value.

### HasVar40

`func (o *ResumeObjectsPhoto) HasVar40() bool`

HasVar40 returns a boolean if a field has been set.

### SetVar40Nil

`func (o *ResumeObjectsPhoto) SetVar40Nil(b bool)`

 SetVar40Nil sets the value for Var40 to be an explicit nil

### UnsetVar40
`func (o *ResumeObjectsPhoto) UnsetVar40()`

UnsetVar40 ensures that no value is present for Var40, not even an explicit nil
### GetVar100

`func (o *ResumeObjectsPhoto) GetVar100() string`

GetVar100 returns the Var100 field if non-nil, zero value otherwise.

### GetVar100Ok

`func (o *ResumeObjectsPhoto) GetVar100Ok() (*string, bool)`

GetVar100Ok returns a tuple with the Var100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar100

`func (o *ResumeObjectsPhoto) SetVar100(v string)`

SetVar100 sets Var100 field to given value.

### HasVar100

`func (o *ResumeObjectsPhoto) HasVar100() bool`

HasVar100 returns a boolean if a field has been set.

### SetVar100Nil

`func (o *ResumeObjectsPhoto) SetVar100Nil(b bool)`

 SetVar100Nil sets the value for Var100 to be an explicit nil

### UnsetVar100
`func (o *ResumeObjectsPhoto) UnsetVar100()`

UnsetVar100 ensures that no value is present for Var100, not even an explicit nil
### GetVar500

`func (o *ResumeObjectsPhoto) GetVar500() string`

GetVar500 returns the Var500 field if non-nil, zero value otherwise.

### GetVar500Ok

`func (o *ResumeObjectsPhoto) GetVar500Ok() (*string, bool)`

GetVar500Ok returns a tuple with the Var500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar500

`func (o *ResumeObjectsPhoto) SetVar500(v string)`

SetVar500 sets Var500 field to given value.

### HasVar500

`func (o *ResumeObjectsPhoto) HasVar500() bool`

HasVar500 returns a boolean if a field has been set.

### SetVar500Nil

`func (o *ResumeObjectsPhoto) SetVar500Nil(b bool)`

 SetVar500Nil sets the value for Var500 to be an explicit nil

### UnsetVar500
`func (o *ResumeObjectsPhoto) UnsetVar500()`

UnsetVar500 ensures that no value is present for Var500, not even an explicit nil
### GetMedium

`func (o *ResumeObjectsPhoto) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ResumeObjectsPhoto) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ResumeObjectsPhoto) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetSmall

`func (o *ResumeObjectsPhoto) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ResumeObjectsPhoto) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ResumeObjectsPhoto) SetSmall(v string)`

SetSmall sets Small field to given value.


### GetId

`func (o *ResumeObjectsPhoto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsPhoto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsPhoto) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


