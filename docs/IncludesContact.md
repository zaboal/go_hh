# IncludesContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **NullableString** | Комментарий к контакту | [optional] 
**NeedVerification** | Pointer to **NullableBool** | Требуется ли подтверждение телефона | [optional] 
**Preferred** | **bool** | Является ли предпочтительным способом связи | 
**Type** | [**IncludesIdName**](IncludesIdName.md) | Тип контакта. Элемент справочника [preferred_contact_type](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Value** | [**IncludesContactPropertiesValue**](IncludesContactPropertiesValue.md) |  | 
**Verified** | Pointer to **NullableBool** | Является ли телефон подтвержденным | [optional] 

## Methods

### NewIncludesContact

`func NewIncludesContact(preferred bool, type_ IncludesIdName, value IncludesContactPropertiesValue, ) *IncludesContact`

NewIncludesContact instantiates a new IncludesContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncludesContactWithDefaults

`func NewIncludesContactWithDefaults() *IncludesContact`

NewIncludesContactWithDefaults instantiates a new IncludesContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IncludesContact) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IncludesContact) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IncludesContact) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IncludesContact) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *IncludesContact) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *IncludesContact) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetNeedVerification

`func (o *IncludesContact) GetNeedVerification() bool`

GetNeedVerification returns the NeedVerification field if non-nil, zero value otherwise.

### GetNeedVerificationOk

`func (o *IncludesContact) GetNeedVerificationOk() (*bool, bool)`

GetNeedVerificationOk returns a tuple with the NeedVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedVerification

`func (o *IncludesContact) SetNeedVerification(v bool)`

SetNeedVerification sets NeedVerification field to given value.

### HasNeedVerification

`func (o *IncludesContact) HasNeedVerification() bool`

HasNeedVerification returns a boolean if a field has been set.

### SetNeedVerificationNil

`func (o *IncludesContact) SetNeedVerificationNil(b bool)`

 SetNeedVerificationNil sets the value for NeedVerification to be an explicit nil

### UnsetNeedVerification
`func (o *IncludesContact) UnsetNeedVerification()`

UnsetNeedVerification ensures that no value is present for NeedVerification, not even an explicit nil
### GetPreferred

`func (o *IncludesContact) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *IncludesContact) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *IncludesContact) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.


### GetType

`func (o *IncludesContact) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncludesContact) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncludesContact) SetType(v IncludesIdName)`

SetType sets Type field to given value.


### GetValue

`func (o *IncludesContact) GetValue() IncludesContactPropertiesValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncludesContact) GetValueOk() (*IncludesContactPropertiesValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncludesContact) SetValue(v IncludesContactPropertiesValue)`

SetValue sets Value field to given value.


### GetVerified

`func (o *IncludesContact) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *IncludesContact) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *IncludesContact) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *IncludesContact) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *IncludesContact) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *IncludesContact) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


