# WebhookSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]WebhookSubscriptionUpdateActionsInner**](WebhookSubscriptionUpdateActionsInner.md) | Cписок событий | 
**Url** | **string** | URL, на который будет отправляться POST запрос при наступлении события | 
**Id** | **string** | Идентификатор подписки | 

## Methods

### NewWebhookSubscriptionItem

`func NewWebhookSubscriptionItem(actions []WebhookSubscriptionUpdateActionsInner, url string, id string, ) *WebhookSubscriptionItem`

NewWebhookSubscriptionItem instantiates a new WebhookSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionItemWithDefaults

`func NewWebhookSubscriptionItemWithDefaults() *WebhookSubscriptionItem`

NewWebhookSubscriptionItemWithDefaults instantiates a new WebhookSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *WebhookSubscriptionItem) GetActions() []WebhookSubscriptionUpdateActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WebhookSubscriptionItem) GetActionsOk() (*[]WebhookSubscriptionUpdateActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WebhookSubscriptionItem) SetActions(v []WebhookSubscriptionUpdateActionsInner)`

SetActions sets Actions field to given value.


### GetUrl

`func (o *WebhookSubscriptionItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscriptionItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscriptionItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *WebhookSubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSubscriptionItem) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


