# ResumeObjectsRelocationPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to [**[]IncludesArea**](IncludesArea.md) | Список городов, в которые возможен переезд. Имеет смысл только с соответствующим полем &#x60;type&#x60; | [optional] 
**District** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список районов, в которые возможен переезд. Имеет смысл только с соответствующим полем &#x60;type&#x60; | [optional] 
**Type** | [**IncludesIdName**](IncludesIdName.md) | Готовность к переезду. Элемент справочника [relocation_type](#tag/Obshie-spravochniki/operation/get-dictionaries) | 

## Methods

### NewResumeObjectsRelocationPublic

`func NewResumeObjectsRelocationPublic(type_ IncludesIdName, ) *ResumeObjectsRelocationPublic`

NewResumeObjectsRelocationPublic instantiates a new ResumeObjectsRelocationPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsRelocationPublicWithDefaults

`func NewResumeObjectsRelocationPublicWithDefaults() *ResumeObjectsRelocationPublic`

NewResumeObjectsRelocationPublicWithDefaults instantiates a new ResumeObjectsRelocationPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ResumeObjectsRelocationPublic) GetArea() []IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeObjectsRelocationPublic) GetAreaOk() (*[]IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeObjectsRelocationPublic) SetArea(v []IncludesArea)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeObjectsRelocationPublic) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetDistrict

`func (o *ResumeObjectsRelocationPublic) GetDistrict() []IncludesIdName`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *ResumeObjectsRelocationPublic) GetDistrictOk() (*[]IncludesIdName, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *ResumeObjectsRelocationPublic) SetDistrict(v []IncludesIdName)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *ResumeObjectsRelocationPublic) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### GetType

`func (o *ResumeObjectsRelocationPublic) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResumeObjectsRelocationPublic) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResumeObjectsRelocationPublic) SetType(v IncludesIdName)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


