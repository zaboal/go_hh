# NegotiationsPhoneCallItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **string** | Дата и время создания звонка | 
**DurationSeconds** | Pointer to **NullableFloat32** | Длительность звонка в секундах | [optional] 
**Id** | **float32** | Идентификатор звонка | 
**LastChangeTime** | Pointer to **NullableString** | Дата и время обновления звонка | [optional] 
**Status** | **string** | Статус звонка.  Возможные значения перечислены в разделе &#x60;phone_call_status&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 

## Methods

### NewNegotiationsPhoneCallItem

`func NewNegotiationsPhoneCallItem(creationTime string, id float32, status string, ) *NegotiationsPhoneCallItem`

NewNegotiationsPhoneCallItem instantiates a new NegotiationsPhoneCallItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsPhoneCallItemWithDefaults

`func NewNegotiationsPhoneCallItemWithDefaults() *NegotiationsPhoneCallItem`

NewNegotiationsPhoneCallItemWithDefaults instantiates a new NegotiationsPhoneCallItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *NegotiationsPhoneCallItem) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NegotiationsPhoneCallItem) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NegotiationsPhoneCallItem) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.


### GetDurationSeconds

`func (o *NegotiationsPhoneCallItem) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *NegotiationsPhoneCallItem) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *NegotiationsPhoneCallItem) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *NegotiationsPhoneCallItem) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *NegotiationsPhoneCallItem) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *NegotiationsPhoneCallItem) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetId

`func (o *NegotiationsPhoneCallItem) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsPhoneCallItem) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsPhoneCallItem) SetId(v float32)`

SetId sets Id field to given value.


### GetLastChangeTime

`func (o *NegotiationsPhoneCallItem) GetLastChangeTime() string`

GetLastChangeTime returns the LastChangeTime field if non-nil, zero value otherwise.

### GetLastChangeTimeOk

`func (o *NegotiationsPhoneCallItem) GetLastChangeTimeOk() (*string, bool)`

GetLastChangeTimeOk returns a tuple with the LastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeTime

`func (o *NegotiationsPhoneCallItem) SetLastChangeTime(v string)`

SetLastChangeTime sets LastChangeTime field to given value.

### HasLastChangeTime

`func (o *NegotiationsPhoneCallItem) HasLastChangeTime() bool`

HasLastChangeTime returns a boolean if a field has been set.

### SetLastChangeTimeNil

`func (o *NegotiationsPhoneCallItem) SetLastChangeTimeNil(b bool)`

 SetLastChangeTimeNil sets the value for LastChangeTime to be an explicit nil

### UnsetLastChangeTime
`func (o *NegotiationsPhoneCallItem) UnsetLastChangeTime()`

UnsetLastChangeTime ensures that no value is present for LastChangeTime, not even an explicit nil
### GetStatus

`func (o *NegotiationsPhoneCallItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NegotiationsPhoneCallItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NegotiationsPhoneCallItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


