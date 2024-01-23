# IncludesContactProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** | Комментарий к контакту | [optional] 
**NeedVerification** | Pointer to **NullableBool** | Требуется ли подтверждение телефона | [optional] 
**Preferred** | Pointer to **bool** | Является ли предпочтительным способом связи | [optional] 
**Type** | Pointer to [**IncludesIdName**](IncludesIdName.md) | Тип контакта. Элемент справочника [preferred_contact_type](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Value** | Pointer to [**IncludesContactPropertiesValue**](IncludesContactPropertiesValue.md) |  | [optional] 
**Verified** | Pointer to **NullableBool** | Является ли телефон подтвержденным | [optional] 

## Methods

### NewIncludesContactProperties

`func NewIncludesContactProperties() *IncludesContactProperties`

NewIncludesContactProperties instantiates a new IncludesContactProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludesContactPropertiesWithDefaults

`func NewIncludesContactPropertiesWithDefaults() *IncludesContactProperties`

NewIncludesContactPropertiesWithDefaults instantiates a new IncludesContactProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IncludesContactProperties) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IncludesContactProperties) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IncludesContactProperties) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IncludesContactProperties) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *IncludesContactProperties) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *IncludesContactProperties) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetNeedVerification

`func (o *IncludesContactProperties) GetNeedVerification() bool`

GetNeedVerification returns the NeedVerification field if non-nil, zero value otherwise.

### GetNeedVerificationOk

`func (o *IncludesContactProperties) GetNeedVerificationOk() (*bool, bool)`

GetNeedVerificationOk returns a tuple with the NeedVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedVerification

`func (o *IncludesContactProperties) SetNeedVerification(v bool)`

SetNeedVerification sets NeedVerification field to given value.

### HasNeedVerification

`func (o *IncludesContactProperties) HasNeedVerification() bool`

HasNeedVerification returns a boolean if a field has been set.

### SetNeedVerificationNil

`func (o *IncludesContactProperties) SetNeedVerificationNil(b bool)`

 SetNeedVerificationNil sets the value for NeedVerification to be an explicit nil

### UnsetNeedVerification
`func (o *IncludesContactProperties) UnsetNeedVerification()`

UnsetNeedVerification ensures that no value is present for NeedVerification, not even an explicit nil
### GetPreferred

`func (o *IncludesContactProperties) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *IncludesContactProperties) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *IncludesContactProperties) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *IncludesContactProperties) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetType

`func (o *IncludesContactProperties) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludesContactProperties) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludesContactProperties) SetType(v IncludesIdName)`

SetType sets Type field to given value.

### HasType

`func (o *IncludesContactProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IncludesContactProperties) GetValue() IncludesContactPropertiesValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncludesContactProperties) GetValueOk() (*IncludesContactPropertiesValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncludesContactProperties) SetValue(v IncludesContactPropertiesValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *IncludesContactProperties) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVerified

`func (o *IncludesContactProperties) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *IncludesContactProperties) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *IncludesContactProperties) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *IncludesContactProperties) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *IncludesContactProperties) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *IncludesContactProperties) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


