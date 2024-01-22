# WebhookActionNegotiationEmployerStateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**NullableWebhookActionVacancyOnlyMineSettings**](WebhookActionVacancyOnlyMineSettings.md) |  | [optional] 
**Type** | **string** | Перемещение резюме кандидата на другой этап вакансии.  Содержит следующую информацию:  * кто переместил резюме;  * идентификатор резюме; * исходный статус резюме;  * новый статус резюме  | 

## Methods

### NewWebhookActionNegotiationEmployerStateChange

`func NewWebhookActionNegotiationEmployerStateChange(type_ string, ) *WebhookActionNegotiationEmployerStateChange`

NewWebhookActionNegotiationEmployerStateChange instantiates a new WebhookActionNegotiationEmployerStateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookActionNegotiationEmployerStateChangeWithDefaults

`func NewWebhookActionNegotiationEmployerStateChangeWithDefaults() *WebhookActionNegotiationEmployerStateChange`

NewWebhookActionNegotiationEmployerStateChangeWithDefaults instantiates a new WebhookActionNegotiationEmployerStateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *WebhookActionNegotiationEmployerStateChange) GetSettings() WebhookActionVacancyOnlyMineSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebhookActionNegotiationEmployerStateChange) GetSettingsOk() (*WebhookActionVacancyOnlyMineSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebhookActionNegotiationEmployerStateChange) SetSettings(v WebhookActionVacancyOnlyMineSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WebhookActionNegotiationEmployerStateChange) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *WebhookActionNegotiationEmployerStateChange) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *WebhookActionNegotiationEmployerStateChange) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetType

`func (o *WebhookActionNegotiationEmployerStateChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookActionNegotiationEmployerStateChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookActionNegotiationEmployerStateChange) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


