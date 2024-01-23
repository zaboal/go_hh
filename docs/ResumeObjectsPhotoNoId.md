# ResumeObjectsPhotoNoId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var40** | Pointer to **NullableString** | URL изображения размером 40x40 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Var100** | Pointer to **NullableString** | URL изображения размером 100x100 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Var500** | Pointer to **NullableString** | URL изображения размером 500x500 пикселей. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | [optional] 
**Medium** | **string** | URL среднего по размеру изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | 
**Small** | **string** | URL уменьшенного изображения. Изображение по данному URL доступно ограниченное время после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернется ошибка &#x60;404 Not Found&#x60; | 

## Methods

### NewResumeObjectsPhotoNoId

`func NewResumeObjectsPhotoNoId(medium string, small string, ) *ResumeObjectsPhotoNoId`

NewResumeObjectsPhotoNoId instantiates a new ResumeObjectsPhotoNoId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsPhotoNoIdWithDefaults

`func NewResumeObjectsPhotoNoIdWithDefaults() *ResumeObjectsPhotoNoId`

NewResumeObjectsPhotoNoIdWithDefaults instantiates a new ResumeObjectsPhotoNoId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar40

`func (o *ResumeObjectsPhotoNoId) GetVar40() string`

GetVar40 returns the Var40 field if non-nil, zero value otherwise.

### GetVar40Ok

`func (o *ResumeObjectsPhotoNoId) GetVar40Ok() (*string, bool)`

GetVar40Ok returns a tuple with the Var40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar40

`func (o *ResumeObjectsPhotoNoId) SetVar40(v string)`

SetVar40 sets Var40 field to given value.

### HasVar40

`func (o *ResumeObjectsPhotoNoId) HasVar40() bool`

HasVar40 returns a boolean if a field has been set.

### SetVar40Nil

`func (o *ResumeObjectsPhotoNoId) SetVar40Nil(b bool)`

 SetVar40Nil sets the value for Var40 to be an explicit nil

### UnsetVar40
`func (o *ResumeObjectsPhotoNoId) UnsetVar40()`

UnsetVar40 ensures that no value is present for Var40, not even an explicit nil
### GetVar100

`func (o *ResumeObjectsPhotoNoId) GetVar100() string`

GetVar100 returns the Var100 field if non-nil, zero value otherwise.

### GetVar100Ok

`func (o *ResumeObjectsPhotoNoId) GetVar100Ok() (*string, bool)`

GetVar100Ok returns a tuple with the Var100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar100

`func (o *ResumeObjectsPhotoNoId) SetVar100(v string)`

SetVar100 sets Var100 field to given value.

### HasVar100

`func (o *ResumeObjectsPhotoNoId) HasVar100() bool`

HasVar100 returns a boolean if a field has been set.

### SetVar100Nil

`func (o *ResumeObjectsPhotoNoId) SetVar100Nil(b bool)`

 SetVar100Nil sets the value for Var100 to be an explicit nil

### UnsetVar100
`func (o *ResumeObjectsPhotoNoId) UnsetVar100()`

UnsetVar100 ensures that no value is present for Var100, not even an explicit nil
### GetVar500

`func (o *ResumeObjectsPhotoNoId) GetVar500() string`

GetVar500 returns the Var500 field if non-nil, zero value otherwise.

### GetVar500Ok

`func (o *ResumeObjectsPhotoNoId) GetVar500Ok() (*string, bool)`

GetVar500Ok returns a tuple with the Var500 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar500

`func (o *ResumeObjectsPhotoNoId) SetVar500(v string)`

SetVar500 sets Var500 field to given value.

### HasVar500

`func (o *ResumeObjectsPhotoNoId) HasVar500() bool`

HasVar500 returns a boolean if a field has been set.

### SetVar500Nil

`func (o *ResumeObjectsPhotoNoId) SetVar500Nil(b bool)`

 SetVar500Nil sets the value for Var500 to be an explicit nil

### UnsetVar500
`func (o *ResumeObjectsPhotoNoId) UnsetVar500()`

UnsetVar500 ensures that no value is present for Var500, not even an explicit nil
### GetMedium

`func (o *ResumeObjectsPhotoNoId) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ResumeObjectsPhotoNoId) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ResumeObjectsPhotoNoId) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetSmall

`func (o *ResumeObjectsPhotoNoId) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ResumeObjectsPhotoNoId) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ResumeObjectsPhotoNoId) SetSmall(v string)`

SetSmall sets Small field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


