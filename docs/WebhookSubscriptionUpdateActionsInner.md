# WebhookSubscriptionUpdateActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**WebhookActionVacancyOnlyMineSettings**](WebhookActionVacancyOnlyMineSettings.md) |  | [optional] 
**Type** | **string** | Новый отклик на вакансию. Данное событие будет вызываться только на отклик со стороны соискателя | 

## Methods

### NewWebhookSubscriptionUpdateActionsInner

`func NewWebhookSubscriptionUpdateActionsInner(type_ string, ) *WebhookSubscriptionUpdateActionsInner`

NewWebhookSubscriptionUpdateActionsInner instantiates a new WebhookSubscriptionUpdateActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionUpdateActionsInnerWithDefaults

`func NewWebhookSubscriptionUpdateActionsInnerWithDefaults() *WebhookSubscriptionUpdateActionsInner`

NewWebhookSubscriptionUpdateActionsInnerWithDefaults instantiates a new WebhookSubscriptionUpdateActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *WebhookSubscriptionUpdateActionsInner) GetSettings() WebhookActionVacancyOnlyMineSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebhookSubscriptionUpdateActionsInner) GetSettingsOk() (*WebhookActionVacancyOnlyMineSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebhookSubscriptionUpdateActionsInner) SetSettings(v WebhookActionVacancyOnlyMineSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WebhookSubscriptionUpdateActionsInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *WebhookSubscriptionUpdateActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookSubscriptionUpdateActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookSubscriptionUpdateActionsInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


