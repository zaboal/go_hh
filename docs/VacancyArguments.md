# VacancyArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор аргумента. Возможные значения:  * &#x60;resume_id&#x60; — идентификатор резюме. * &#x60;vacancy_id&#x60; — идентификатор вакансии. * &#x60;message&#x60; — сообщение, которое будет отправлено соискателю на электронную почту. Используйте [шаблоны](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-mail-templates) для получения текстов. * &#x60;send_sms&#x60; — уведомлять ли соискателя о приглашении с помощью SMS. Значение по умолчанию — &#x60;false&#x60;. Обратите внимание: в SMS-сообщении используется стандартный текст, изменить его нельзя * &#x60;address_id&#x60; — идентификатор [адреса](#tag/Adresa-rabotodatelya), который будет указан в приглашении  | 
**Required** | **bool** | Обязательность аргумента | 
**RequiredArguments** | [**[]IncludesId**](IncludesId.md) | Идентификаторы аргументов, которые необходимо приложить, если указан данный аргумент. Например, адрес является необязательным, но при его указании необходимо указать также и сообщение | 

## Methods

### NewVacancyArguments

`func NewVacancyArguments(id string, required bool, requiredArguments []IncludesId, ) *VacancyArguments`

NewVacancyArguments instantiates a new VacancyArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyArgumentsWithDefaults

`func NewVacancyArgumentsWithDefaults() *VacancyArguments`

NewVacancyArgumentsWithDefaults instantiates a new VacancyArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyArguments) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyArguments) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyArguments) SetId(v string)`

SetId sets Id field to given value.


### GetRequired

`func (o *VacancyArguments) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *VacancyArguments) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *VacancyArguments) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetRequiredArguments

`func (o *VacancyArguments) GetRequiredArguments() []IncludesId`

GetRequiredArguments returns the RequiredArguments field if non-nil, zero value otherwise.

### GetRequiredArgumentsOk

`func (o *VacancyArguments) GetRequiredArgumentsOk() (*[]IncludesId, bool)`

GetRequiredArgumentsOk returns a tuple with the RequiredArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredArguments

`func (o *VacancyArguments) SetRequiredArguments(v []IncludesId)`

SetRequiredArguments sets RequiredArguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


