# UserStatusesApplicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobSearchStatus** | Pointer to [**UserStatusesJobSearchStatus**](UserStatusesJobSearchStatus.md) |  | [optional] 
**WhenCanStartWorkStatus** | Pointer to [**UserStatusesWhenCanStartWorkStatus**](UserStatusesWhenCanStartWorkStatus.md) |  | [optional] 

## Methods

### NewUserStatusesApplicant

`func NewUserStatusesApplicant() *UserStatusesApplicant`

NewUserStatusesApplicant instantiates a new UserStatusesApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusesApplicantWithDefaults

`func NewUserStatusesApplicantWithDefaults() *UserStatusesApplicant`

NewUserStatusesApplicantWithDefaults instantiates a new UserStatusesApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobSearchStatus

`func (o *UserStatusesApplicant) GetJobSearchStatus() UserStatusesJobSearchStatus`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *UserStatusesApplicant) GetJobSearchStatusOk() (*UserStatusesJobSearchStatus, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *UserStatusesApplicant) SetJobSearchStatus(v UserStatusesJobSearchStatus)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *UserStatusesApplicant) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetWhenCanStartWorkStatus

`func (o *UserStatusesApplicant) GetWhenCanStartWorkStatus() UserStatusesWhenCanStartWorkStatus`

GetWhenCanStartWorkStatus returns the WhenCanStartWorkStatus field if non-nil, zero value otherwise.

### GetWhenCanStartWorkStatusOk

`func (o *UserStatusesApplicant) GetWhenCanStartWorkStatusOk() (*UserStatusesWhenCanStartWorkStatus, bool)`

GetWhenCanStartWorkStatusOk returns a tuple with the WhenCanStartWorkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenCanStartWorkStatus

`func (o *UserStatusesApplicant) SetWhenCanStartWorkStatus(v UserStatusesWhenCanStartWorkStatus)`

SetWhenCanStartWorkStatus sets WhenCanStartWorkStatus field to given value.

### HasWhenCanStartWorkStatus

`func (o *UserStatusesApplicant) HasWhenCanStartWorkStatus() bool`

HasWhenCanStartWorkStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


