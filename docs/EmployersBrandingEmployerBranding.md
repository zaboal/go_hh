# EmployersBrandingEmployerBranding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип брендирования   * constructor - Брендирование типом конструктор          * makeup      - Брендирование типом makeup (хамелеон)  | 
**Constructor** | [**EmployersBrandingConstructorConstructor**](EmployersBrandingConstructorConstructor.md) |  | 
**Makeup** | [**EmployersBrandingMakeupMakeup**](EmployersBrandingMakeupMakeup.md) |  | 
**TemplateCode** | **string** | Уникальный код активного бренд шаблона страницы в формате \&quot;makeup:\\d+\&quot; | 
**TemplateVersionId** | **string** | Идентификатор активной версии бренд шаблона страницы | 

## Methods

### NewEmployersBrandingEmployerBranding

`func NewEmployersBrandingEmployerBranding(type_ string, constructor EmployersBrandingConstructorConstructor, makeup EmployersBrandingMakeupMakeup, templateCode string, templateVersionId string, ) *EmployersBrandingEmployerBranding`

NewEmployersBrandingEmployerBranding instantiates a new EmployersBrandingEmployerBranding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersBrandingEmployerBrandingWithDefaults

`func NewEmployersBrandingEmployerBrandingWithDefaults() *EmployersBrandingEmployerBranding`

NewEmployersBrandingEmployerBrandingWithDefaults instantiates a new EmployersBrandingEmployerBranding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmployersBrandingEmployerBranding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmployersBrandingEmployerBranding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmployersBrandingEmployerBranding) SetType(v string)`

SetType sets Type field to given value.


### GetConstructor

`func (o *EmployersBrandingEmployerBranding) GetConstructor() EmployersBrandingConstructorConstructor`

GetConstructor returns the Constructor field if non-nil, zero value otherwise.

### GetConstructorOk

`func (o *EmployersBrandingEmployerBranding) GetConstructorOk() (*EmployersBrandingConstructorConstructor, bool)`

GetConstructorOk returns a tuple with the Constructor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructor

`func (o *EmployersBrandingEmployerBranding) SetConstructor(v EmployersBrandingConstructorConstructor)`

SetConstructor sets Constructor field to given value.


### GetMakeup

`func (o *EmployersBrandingEmployerBranding) GetMakeup() EmployersBrandingMakeupMakeup`

GetMakeup returns the Makeup field if non-nil, zero value otherwise.

### GetMakeupOk

`func (o *EmployersBrandingEmployerBranding) GetMakeupOk() (*EmployersBrandingMakeupMakeup, bool)`

GetMakeupOk returns a tuple with the Makeup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeup

`func (o *EmployersBrandingEmployerBranding) SetMakeup(v EmployersBrandingMakeupMakeup)`

SetMakeup sets Makeup field to given value.


### GetTemplateCode

`func (o *EmployersBrandingEmployerBranding) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *EmployersBrandingEmployerBranding) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *EmployersBrandingEmployerBranding) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetTemplateVersionId

`func (o *EmployersBrandingEmployerBranding) GetTemplateVersionId() string`

GetTemplateVersionId returns the TemplateVersionId field if non-nil, zero value otherwise.

### GetTemplateVersionIdOk

`func (o *EmployersBrandingEmployerBranding) GetTemplateVersionIdOk() (*string, bool)`

GetTemplateVersionIdOk returns a tuple with the TemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersionId

`func (o *EmployersBrandingEmployerBranding) SetTemplateVersionId(v string)`

SetTemplateVersionId sets TemplateVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


