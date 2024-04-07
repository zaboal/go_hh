# NegotiationsNegotiationsStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageReplyTime** | Pointer to **NullableFloat32** | Среднее время (в днях) между получением отклика и отправкой сообщения | [optional] 
**Politeness** | Pointer to [**NullableNegotiationsObjectsPoliteness**](NegotiationsObjectsPoliteness.md) |  | [optional] 
**Received** | **float32** | Количество откликов на вакансии, полученных за период (последние 30 дней) | 
**RepliedPercent** | Pointer to **NullableFloat32** | Процент откликов на вакансии, перемещенных в любую другую [коллекцию](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#term-collection) с отправкой сообщения, за период  | [optional] 
**ViewedPercent** | Pointer to **NullableFloat32** | Процент прочитанных откликов за период | [optional] 

## Methods

### NewNegotiationsNegotiationsStatistics

`func NewNegotiationsNegotiationsStatistics(received float32, ) *NegotiationsNegotiationsStatistics`

NewNegotiationsNegotiationsStatistics instantiates a new NegotiationsNegotiationsStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationsStatisticsWithDefaults

`func NewNegotiationsNegotiationsStatisticsWithDefaults() *NegotiationsNegotiationsStatistics`

NewNegotiationsNegotiationsStatisticsWithDefaults instantiates a new NegotiationsNegotiationsStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageReplyTime

`func (o *NegotiationsNegotiationsStatistics) GetAverageReplyTime() float32`

GetAverageReplyTime returns the AverageReplyTime field if non-nil, zero value otherwise.

### GetAverageReplyTimeOk

`func (o *NegotiationsNegotiationsStatistics) GetAverageReplyTimeOk() (*float32, bool)`

GetAverageReplyTimeOk returns a tuple with the AverageReplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageReplyTime

`func (o *NegotiationsNegotiationsStatistics) SetAverageReplyTime(v float32)`

SetAverageReplyTime sets AverageReplyTime field to given value.

### HasAverageReplyTime

`func (o *NegotiationsNegotiationsStatistics) HasAverageReplyTime() bool`

HasAverageReplyTime returns a boolean if a field has been set.

### SetAverageReplyTimeNil

`func (o *NegotiationsNegotiationsStatistics) SetAverageReplyTimeNil(b bool)`

 SetAverageReplyTimeNil sets the value for AverageReplyTime to be an explicit nil

### UnsetAverageReplyTime
`func (o *NegotiationsNegotiationsStatistics) UnsetAverageReplyTime()`

UnsetAverageReplyTime ensures that no value is present for AverageReplyTime, not even an explicit nil
### GetPoliteness

`func (o *NegotiationsNegotiationsStatistics) GetPoliteness() NegotiationsObjectsPoliteness`

GetPoliteness returns the Politeness field if non-nil, zero value otherwise.

### GetPolitenessOk

`func (o *NegotiationsNegotiationsStatistics) GetPolitenessOk() (*NegotiationsObjectsPoliteness, bool)`

GetPolitenessOk returns a tuple with the Politeness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliteness

`func (o *NegotiationsNegotiationsStatistics) SetPoliteness(v NegotiationsObjectsPoliteness)`

SetPoliteness sets Politeness field to given value.

### HasPoliteness

`func (o *NegotiationsNegotiationsStatistics) HasPoliteness() bool`

HasPoliteness returns a boolean if a field has been set.

### SetPolitenessNil

`func (o *NegotiationsNegotiationsStatistics) SetPolitenessNil(b bool)`

 SetPolitenessNil sets the value for Politeness to be an explicit nil

### UnsetPoliteness
`func (o *NegotiationsNegotiationsStatistics) UnsetPoliteness()`

UnsetPoliteness ensures that no value is present for Politeness, not even an explicit nil
### GetReceived

`func (o *NegotiationsNegotiationsStatistics) GetReceived() float32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *NegotiationsNegotiationsStatistics) GetReceivedOk() (*float32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *NegotiationsNegotiationsStatistics) SetReceived(v float32)`

SetReceived sets Received field to given value.


### GetRepliedPercent

`func (o *NegotiationsNegotiationsStatistics) GetRepliedPercent() float32`

GetRepliedPercent returns the RepliedPercent field if non-nil, zero value otherwise.

### GetRepliedPercentOk

`func (o *NegotiationsNegotiationsStatistics) GetRepliedPercentOk() (*float32, bool)`

GetRepliedPercentOk returns a tuple with the RepliedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepliedPercent

`func (o *NegotiationsNegotiationsStatistics) SetRepliedPercent(v float32)`

SetRepliedPercent sets RepliedPercent field to given value.

### HasRepliedPercent

`func (o *NegotiationsNegotiationsStatistics) HasRepliedPercent() bool`

HasRepliedPercent returns a boolean if a field has been set.

### SetRepliedPercentNil

`func (o *NegotiationsNegotiationsStatistics) SetRepliedPercentNil(b bool)`

 SetRepliedPercentNil sets the value for RepliedPercent to be an explicit nil

### UnsetRepliedPercent
`func (o *NegotiationsNegotiationsStatistics) UnsetRepliedPercent()`

UnsetRepliedPercent ensures that no value is present for RepliedPercent, not even an explicit nil
### GetViewedPercent

`func (o *NegotiationsNegotiationsStatistics) GetViewedPercent() float32`

GetViewedPercent returns the ViewedPercent field if non-nil, zero value otherwise.

### GetViewedPercentOk

`func (o *NegotiationsNegotiationsStatistics) GetViewedPercentOk() (*float32, bool)`

GetViewedPercentOk returns a tuple with the ViewedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedPercent

`func (o *NegotiationsNegotiationsStatistics) SetViewedPercent(v float32)`

SetViewedPercent sets ViewedPercent field to given value.

### HasViewedPercent

`func (o *NegotiationsNegotiationsStatistics) HasViewedPercent() bool`

HasViewedPercent returns a boolean if a field has been set.

### SetViewedPercentNil

`func (o *NegotiationsNegotiationsStatistics) SetViewedPercentNil(b bool)`

 SetViewedPercentNil sets the value for ViewedPercent to be an explicit nil

### UnsetViewedPercent
`func (o *NegotiationsNegotiationsStatistics) UnsetViewedPercent()`

UnsetViewedPercent ensures that no value is present for ViewedPercent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


