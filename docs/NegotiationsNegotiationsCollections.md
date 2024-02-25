# NegotiationsNegotiationsCollections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | [**[]NegotiationsNegotiationsCollection**](NegotiationsNegotiationsCollection.md) | Коллекции откликов/приглашений для данной вакансии | 
**EmployerStates** | [**[]EmployersEmployersState**](EmployersEmployersState.md) | Состояния [откликов/приглашений](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#term-employer-state) вакансии для работодателя  | 
**GeneratedCollections** | [**[]NegotiationsNegotiationsCollection**](NegotiationsNegotiationsCollection.md) | Сгенерированные коллекции откликов/приглашений для данной вакансии | 

## Methods

### NewNegotiationsNegotiationsCollections

`func NewNegotiationsNegotiationsCollections(collections []NegotiationsNegotiationsCollection, employerStates []EmployersEmployersState, generatedCollections []NegotiationsNegotiationsCollection, ) *NegotiationsNegotiationsCollections`

NewNegotiationsNegotiationsCollections instantiates a new NegotiationsNegotiationsCollections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationsCollectionsWithDefaults

`func NewNegotiationsNegotiationsCollectionsWithDefaults() *NegotiationsNegotiationsCollections`

NewNegotiationsNegotiationsCollectionsWithDefaults instantiates a new NegotiationsNegotiationsCollections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *NegotiationsNegotiationsCollections) GetCollections() []NegotiationsNegotiationsCollection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *NegotiationsNegotiationsCollections) GetCollectionsOk() (*[]NegotiationsNegotiationsCollection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *NegotiationsNegotiationsCollections) SetCollections(v []NegotiationsNegotiationsCollection)`

SetCollections sets Collections field to given value.


### GetEmployerStates

`func (o *NegotiationsNegotiationsCollections) GetEmployerStates() []EmployersEmployersState`

GetEmployerStates returns the EmployerStates field if non-nil, zero value otherwise.

### GetEmployerStatesOk

`func (o *NegotiationsNegotiationsCollections) GetEmployerStatesOk() (*[]EmployersEmployersState, bool)`

GetEmployerStatesOk returns a tuple with the EmployerStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerStates

`func (o *NegotiationsNegotiationsCollections) SetEmployerStates(v []EmployersEmployersState)`

SetEmployerStates sets EmployerStates field to given value.


### GetGeneratedCollections

`func (o *NegotiationsNegotiationsCollections) GetGeneratedCollections() []NegotiationsNegotiationsCollection`

GetGeneratedCollections returns the GeneratedCollections field if non-nil, zero value otherwise.

### GetGeneratedCollectionsOk

`func (o *NegotiationsNegotiationsCollections) GetGeneratedCollectionsOk() (*[]NegotiationsNegotiationsCollection, bool)`

GetGeneratedCollectionsOk returns a tuple with the GeneratedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedCollections

`func (o *NegotiationsNegotiationsCollections) SetGeneratedCollections(v []NegotiationsNegotiationsCollection)`

SetGeneratedCollections sets GeneratedCollections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


