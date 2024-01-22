# DictionariesBranchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор отрасли | 
**Industries** | [**[]IncludesIdName**](IncludesIdName.md) | Сферы деятельности | 
**Name** | **string** | Название отрасли | 

## Methods

### NewDictionariesBranchItem

`func NewDictionariesBranchItem(id string, industries []IncludesIdName, name string, ) *DictionariesBranchItem`

NewDictionariesBranchItem instantiates a new DictionariesBranchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionariesBranchItemWithDefaults

`func NewDictionariesBranchItemWithDefaults() *DictionariesBranchItem`

NewDictionariesBranchItemWithDefaults instantiates a new DictionariesBranchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DictionariesBranchItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DictionariesBranchItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DictionariesBranchItem) SetId(v string)`

SetId sets Id field to given value.


### GetIndustries

`func (o *DictionariesBranchItem) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *DictionariesBranchItem) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *DictionariesBranchItem) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.


### GetName

`func (o *DictionariesBranchItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DictionariesBranchItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DictionariesBranchItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


