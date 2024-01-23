# VacancyPicture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlurredPath** | Pointer to **NullableString** | Путь до маленькой (порядка 4% от изначального размера) размытой картинки. При показе ее нужно преобразовать к нужному размеру | [optional] 
**Height** | Pointer to **float32** | Высота картинки | [optional] 
**Path** | Pointer to **string** | Адрес картинки | [optional] 
**Width** | Pointer to **float32** | Ширина картинки | [optional] 

## Methods

### NewVacancyPicture

`func NewVacancyPicture() *VacancyPicture`

NewVacancyPicture instantiates a new VacancyPicture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyPictureWithDefaults

`func NewVacancyPictureWithDefaults() *VacancyPicture`

NewVacancyPictureWithDefaults instantiates a new VacancyPicture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlurredPath

`func (o *VacancyPicture) GetBlurredPath() string`

GetBlurredPath returns the BlurredPath field if non-nil, zero value otherwise.

### GetBlurredPathOk

`func (o *VacancyPicture) GetBlurredPathOk() (*string, bool)`

GetBlurredPathOk returns a tuple with the BlurredPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlurredPath

`func (o *VacancyPicture) SetBlurredPath(v string)`

SetBlurredPath sets BlurredPath field to given value.

### HasBlurredPath

`func (o *VacancyPicture) HasBlurredPath() bool`

HasBlurredPath returns a boolean if a field has been set.

### SetBlurredPathNil

`func (o *VacancyPicture) SetBlurredPathNil(b bool)`

 SetBlurredPathNil sets the value for BlurredPath to be an explicit nil

### UnsetBlurredPath
`func (o *VacancyPicture) UnsetBlurredPath()`

UnsetBlurredPath ensures that no value is present for BlurredPath, not even an explicit nil
### GetHeight

`func (o *VacancyPicture) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VacancyPicture) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VacancyPicture) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VacancyPicture) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetPath

`func (o *VacancyPicture) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VacancyPicture) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VacancyPicture) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VacancyPicture) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetWidth

`func (o *VacancyPicture) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VacancyPicture) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VacancyPicture) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VacancyPicture) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


