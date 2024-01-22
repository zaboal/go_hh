# ResumeObjectsAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IncludesIdName**](IncludesIdName.md) | Определяет, кому будет доступно резюме в поиске и по прямой ссылке.

Установить значение параметра можно при [создании](#tag/Rezyume.-Sozdanie-i-obnovlenie/operation/create-resume) или [редактировании](#tag/Rezyume.-Sozdanie-i-obnovlenie/operation/edit-resume) резюме. Возможные значения приведены в поле &#x60;resume_access_type&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries).

С 1 сентября 2021 года тип видимости &#x60;everyone&#x60; стал недоступен для сохранения из-за законодательных ограничений.

Чтобы управлять списком работодателей, которые могут просматривать резюме, воспользуйтесь группой методов [Резюме. Списки видимости](#tag/Rezyume.-Spiski-vidimosti)
 | 

## Methods

### NewResumeObjectsAccess

`func NewResumeObjectsAccess(type_ IncludesIdName, ) *ResumeObjectsAccess`

NewResumeObjectsAccess instantiates a new ResumeObjectsAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsAccessWithDefaults

`func NewResumeObjectsAccessWithDefaults() *ResumeObjectsAccess`

NewResumeObjectsAccessWithDefaults instantiates a new ResumeObjectsAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResumeObjectsAccess) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResumeObjectsAccess) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResumeObjectsAccess) SetType(v IncludesIdName)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


