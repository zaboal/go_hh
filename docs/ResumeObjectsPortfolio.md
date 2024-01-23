# ResumeObjectsPortfolio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Описание изображения в портфолио | [optional] 
**Id** | **string** | Уникальный идентификатор изображения | 
**Medium** | **string** | URL среднего по размеру изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 
**Small** | **string** | URL уменьшенного изображения. Изображение по данному url доступно ограниченное время, после получения ответа. Приложение должно быть готово к тому, что на запрос изображения вернётся &#x60;404 Not Found&#x60;  | 

## Methods

### NewResumeObjectsPortfolio

`func NewResumeObjectsPortfolio(id string, medium string, small string, ) *ResumeObjectsPortfolio`

NewResumeObjectsPortfolio instantiates a new ResumeObjectsPortfolio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsPortfolioWithDefaults

`func NewResumeObjectsPortfolioWithDefaults() *ResumeObjectsPortfolio`

NewResumeObjectsPortfolioWithDefaults instantiates a new ResumeObjectsPortfolio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResumeObjectsPortfolio) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResumeObjectsPortfolio) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResumeObjectsPortfolio) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResumeObjectsPortfolio) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResumeObjectsPortfolio) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResumeObjectsPortfolio) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ResumeObjectsPortfolio) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsPortfolio) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsPortfolio) SetId(v string)`

SetId sets Id field to given value.


### GetMedium

`func (o *ResumeObjectsPortfolio) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ResumeObjectsPortfolio) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ResumeObjectsPortfolio) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetSmall

`func (o *ResumeObjectsPortfolio) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *ResumeObjectsPortfolio) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *ResumeObjectsPortfolio) SetSmall(v string)`

SetSmall sets Small field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


