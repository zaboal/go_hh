# ResumeObjectsPortfolioNoId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Описание изображения в портфолио | [optional] 
**Medium** | **string** | URL среднего по размеру изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 
**Small** | **string** | URL уменьшенного изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 

## Methods

### NewResumeObjectsPortfolioNoId

`func NewResumeObjectsPortfolioNoId(medium string, small string, ) *ResumeObjectsPortfolioNoId`

NewResumeObjectsPortfolioNoId instantiates a new ResumeObjectsPortfolioNoId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsPortfolioNoIdWithDefaults

`func NewResumeObjectsPortfolioNoIdWithDefaults() *ResumeObjectsPortfolioNoId`

NewResumeObjectsPortfolioNoIdWithDefaults instantiates a new ResumeObjectsPortfolioNoId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResumeObjectsPortfolioNoId) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResumeObjectsPortfolioNoId) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResumeObjectsPortfolioNoId) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResumeObjectsPortfolioNoId) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMedium

`func (o *ResumeObjectsPortfolioNoId) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ResumeObjectsPortfolioNoId) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ResumeObjectsPortfolioNoId) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetSmall

`func (o *ResumeObjectsPortfolioNoId) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ResumeObjectsPortfolioNoId) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ResumeObjectsPortfolioNoId) SetSmall(v string)`

SetSmall sets Small field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


