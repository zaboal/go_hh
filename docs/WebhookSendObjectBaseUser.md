# WebhookSendObjectBaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | **string** | Тип экшена | 
**Id** | **string** | Идентификатор сообщения | 
**Payload** | [**WebhookSendObjectBaseUserPayload**](WebhookSendObjectBaseUserPayload.md) |  | 
**SubscriptionId** | **string** | Идентификатор подписки | 
**UserId** | **string** | Идентификатор пользователя | 

## Methods

### NewWebhookSendObjectBaseUser

`func NewWebhookSendObjectBaseUser(actionType string, id string, payload WebhookSendObjectBaseUserPayload, subscriptionId string, userId string, ) *WebhookSendObjectBaseUser`

NewWebhookSendObjectBaseUser instantiates a new WebhookSendObjectBaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSendObjectBaseUserWithDefaults

`func NewWebhookSendObjectBaseUserWithDefaults() *WebhookSendObjectBaseUser`

NewWebhookSendObjectBaseUserWithDefaults instantiates a new WebhookSendObjectBaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *WebhookSendObjectBaseUser) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *WebhookSendObjectBaseUser) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *WebhookSendObjectBaseUser) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetId

`func (o *WebhookSendObjectBaseUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSendObjectBaseUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSendObjectBaseUser) SetId(v string)`

SetId sets Id field to given value.


### GetPayload

`func (o *WebhookSendObjectBaseUser) GetPayload() WebhookSendObjectBaseUserPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookSendObjectBaseUser) GetPayloadOk() (*WebhookSendObjectBaseUserPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookSendObjectBaseUser) SetPayload(v WebhookSendObjectBaseUserPayload)`

SetPayload sets Payload field to given value.


### GetSubscriptionId

`func (o *WebhookSendObjectBaseUser) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *WebhookSendObjectBaseUser) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *WebhookSendObjectBaseUser) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetUserId

`func (o *WebhookSendObjectBaseUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookSendObjectBaseUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookSendObjectBaseUser) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


