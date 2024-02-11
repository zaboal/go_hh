# WebhookSubscriptionCommonItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]WebhookSubscriptionUpdateActionsInner**](WebhookSubscriptionUpdateActionsInner.md) | Cписок событий | 
**Url** | **string** | URL, на который будет отправляться POST запрос при наступлении события | 

## Methods

### NewWebhookSubscriptionCommonItem

`func NewWebhookSubscriptionCommonItem(actions []WebhookSubscriptionUpdateActionsInner, url string, ) *WebhookSubscriptionCommonItem`

NewWebhookSubscriptionCommonItem instantiates a new WebhookSubscriptionCommonItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionCommonItemWithDefaults

`func NewWebhookSubscriptionCommonItemWithDefaults() *WebhookSubscriptionCommonItem`

NewWebhookSubscriptionCommonItemWithDefaults instantiates a new WebhookSubscriptionCommonItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *WebhookSubscriptionCommonItem) GetActions() []WebhookSubscriptionUpdateActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WebhookSubscriptionCommonItem) GetActionsOk() (*[]WebhookSubscriptionUpdateActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WebhookSubscriptionCommonItem) SetActions(v []WebhookSubscriptionUpdateActionsInner)`

SetActions sets Actions field to given value.


### GetUrl

`func (o *WebhookSubscriptionCommonItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscriptionCommonItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscriptionCommonItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


