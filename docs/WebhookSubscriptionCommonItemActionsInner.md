# WebhookSubscriptionCommonItemActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**NullableWebhookActionVacancyOnlyMineSettings**](WebhookActionVacancyOnlyMineSettings.md) |  | [optional] 
**Type** | **string** | Новый отклик на вакансию. Данное событие будет вызываться только на отклик со стороны соискателя | 

## Methods

### NewWebhookSubscriptionCommonItemActionsInner

`func NewWebhookSubscriptionCommonItemActionsInner(type_ string, ) *WebhookSubscriptionCommonItemActionsInner`

NewWebhookSubscriptionCommonItemActionsInner instantiates a new WebhookSubscriptionCommonItemActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionCommonItemActionsInnerWithDefaults

`func NewWebhookSubscriptionCommonItemActionsInnerWithDefaults() *WebhookSubscriptionCommonItemActionsInner`

NewWebhookSubscriptionCommonItemActionsInnerWithDefaults instantiates a new WebhookSubscriptionCommonItemActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *WebhookSubscriptionCommonItemActionsInner) GetSettings() WebhookActionVacancyOnlyMineSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebhookSubscriptionCommonItemActionsInner) GetSettingsOk() (*WebhookActionVacancyOnlyMineSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebhookSubscriptionCommonItemActionsInner) SetSettings(v WebhookActionVacancyOnlyMineSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WebhookSubscriptionCommonItemActionsInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *WebhookSubscriptionCommonItemActionsInner) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *WebhookSubscriptionCommonItemActionsInner) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetType

`func (o *WebhookSubscriptionCommonItemActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookSubscriptionCommonItemActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookSubscriptionCommonItemActionsInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


