# EmployersFunnelStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор этапа воронки | 
**State** | [**EmployersEmployersState**](EmployersEmployersState.md) |  | 
**Substate** | Pointer to [**NullableEmployersFunnelSubstate**](EmployersFunnelSubstate.md) |  | [optional] 

## Methods

### NewEmployersFunnelStage

`func NewEmployersFunnelStage(id string, state EmployersEmployersState, ) *EmployersFunnelStage`

NewEmployersFunnelStage instantiates a new EmployersFunnelStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersFunnelStageWithDefaults

`func NewEmployersFunnelStageWithDefaults() *EmployersFunnelStage`

NewEmployersFunnelStageWithDefaults instantiates a new EmployersFunnelStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployersFunnelStage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployersFunnelStage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployersFunnelStage) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *EmployersFunnelStage) GetState() EmployersEmployersState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EmployersFunnelStage) GetStateOk() (*EmployersEmployersState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EmployersFunnelStage) SetState(v EmployersEmployersState)`

SetState sets State field to given value.


### GetSubstate

`func (o *EmployersFunnelStage) GetSubstate() EmployersFunnelSubstate`

GetSubstate returns the Substate field if non-nil, zero value otherwise.

### GetSubstateOk

`func (o *EmployersFunnelStage) GetSubstateOk() (*EmployersFunnelSubstate, bool)`

GetSubstateOk returns a tuple with the Substate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstate

`func (o *EmployersFunnelStage) SetSubstate(v EmployersFunnelSubstate)`

SetSubstate sets Substate field to given value.

### HasSubstate

`func (o *EmployersFunnelStage) HasSubstate() bool`

HasSubstate returns a boolean if a field has been set.

### SetSubstateNil

`func (o *EmployersFunnelStage) SetSubstateNil(b bool)`

 SetSubstateNil sets the value for Substate to be an explicit nil

### UnsetSubstate
`func (o *EmployersFunnelStage) UnsetSubstate()`

UnsetSubstate ensures that no value is present for Substate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


