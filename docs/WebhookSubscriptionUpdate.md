# WebhookSubscriptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]WebhookSubscriptionUpdateActionsInner**](WebhookSubscriptionUpdateActionsInner.md) | Cписок событий, на которые нужно подписаться | [optional] 
**Url** | Pointer to **string** | URL, на который будет отправляться POST запрос при наступлении события | [optional] 

## Methods

### NewWebhookSubscriptionUpdate

`func NewWebhookSubscriptionUpdate() *WebhookSubscriptionUpdate`

NewWebhookSubscriptionUpdate instantiates a new WebhookSubscriptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionUpdateWithDefaults

`func NewWebhookSubscriptionUpdateWithDefaults() *WebhookSubscriptionUpdate`

NewWebhookSubscriptionUpdateWithDefaults instantiates a new WebhookSubscriptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *WebhookSubscriptionUpdate) GetActions() []WebhookSubscriptionUpdateActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WebhookSubscriptionUpdate) GetActionsOk() (*[]WebhookSubscriptionUpdateActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WebhookSubscriptionUpdate) SetActions(v []WebhookSubscriptionUpdateActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *WebhookSubscriptionUpdate) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookSubscriptionUpdate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscriptionUpdate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscriptionUpdate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookSubscriptionUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


