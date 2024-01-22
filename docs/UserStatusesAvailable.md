# UserStatusesAvailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobSearchStatuses** | Pointer to [**[]UserStatusesJobSearchStatusDictionary**](UserStatusesJobSearchStatusDictionary.md) |  | [optional] 
**JobSearchStatusesApplicant** | [**[]IncludesIdName**](IncludesIdName.md) |  | 
**JobSearchStatusesEmployer** | [**[]IncludesIdName**](IncludesIdName.md) |  | 
**WhenCanStartWorkStatuses** | Pointer to [**[]UserStatusesWhenCanStartWorkStatus**](UserStatusesWhenCanStartWorkStatus.md) |  | [optional] 

## Methods

### NewUserStatusesAvailable

`func NewUserStatusesAvailable(jobSearchStatusesApplicant []IncludesIdName, jobSearchStatusesEmployer []IncludesIdName, ) *UserStatusesAvailable`

NewUserStatusesAvailable instantiates a new UserStatusesAvailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusesAvailableWithDefaults

`func NewUserStatusesAvailableWithDefaults() *UserStatusesAvailable`

NewUserStatusesAvailableWithDefaults instantiates a new UserStatusesAvailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobSearchStatuses

`func (o *UserStatusesAvailable) GetJobSearchStatuses() []UserStatusesJobSearchStatusDictionary`

GetJobSearchStatuses returns the JobSearchStatuses field if non-nil, zero value otherwise.

### GetJobSearchStatusesOk

`func (o *UserStatusesAvailable) GetJobSearchStatusesOk() (*[]UserStatusesJobSearchStatusDictionary, bool)`

GetJobSearchStatusesOk returns a tuple with the JobSearchStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatuses

`func (o *UserStatusesAvailable) SetJobSearchStatuses(v []UserStatusesJobSearchStatusDictionary)`

SetJobSearchStatuses sets JobSearchStatuses field to given value.

### HasJobSearchStatuses

`func (o *UserStatusesAvailable) HasJobSearchStatuses() bool`

HasJobSearchStatuses returns a boolean if a field has been set.

### GetJobSearchStatusesApplicant

`func (o *UserStatusesAvailable) GetJobSearchStatusesApplicant() []IncludesIdName`

GetJobSearchStatusesApplicant returns the JobSearchStatusesApplicant field if non-nil, zero value otherwise.

### GetJobSearchStatusesApplicantOk

`func (o *UserStatusesAvailable) GetJobSearchStatusesApplicantOk() (*[]IncludesIdName, bool)`

GetJobSearchStatusesApplicantOk returns a tuple with the JobSearchStatusesApplicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatusesApplicant

`func (o *UserStatusesAvailable) SetJobSearchStatusesApplicant(v []IncludesIdName)`

SetJobSearchStatusesApplicant sets JobSearchStatusesApplicant field to given value.


### GetJobSearchStatusesEmployer

`func (o *UserStatusesAvailable) GetJobSearchStatusesEmployer() []IncludesIdName`

GetJobSearchStatusesEmployer returns the JobSearchStatusesEmployer field if non-nil, zero value otherwise.

### GetJobSearchStatusesEmployerOk

`func (o *UserStatusesAvailable) GetJobSearchStatusesEmployerOk() (*[]IncludesIdName, bool)`

GetJobSearchStatusesEmployerOk returns a tuple with the JobSearchStatusesEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatusesEmployer

`func (o *UserStatusesAvailable) SetJobSearchStatusesEmployer(v []IncludesIdName)`

SetJobSearchStatusesEmployer sets JobSearchStatusesEmployer field to given value.


### GetWhenCanStartWorkStatuses

`func (o *UserStatusesAvailable) GetWhenCanStartWorkStatuses() []UserStatusesWhenCanStartWorkStatus`

GetWhenCanStartWorkStatuses returns the WhenCanStartWorkStatuses field if non-nil, zero value otherwise.

### GetWhenCanStartWorkStatusesOk

`func (o *UserStatusesAvailable) GetWhenCanStartWorkStatusesOk() (*[]UserStatusesWhenCanStartWorkStatus, bool)`

GetWhenCanStartWorkStatusesOk returns a tuple with the WhenCanStartWorkStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenCanStartWorkStatuses

`func (o *UserStatusesAvailable) SetWhenCanStartWorkStatuses(v []UserStatusesWhenCanStartWorkStatus)`

SetWhenCanStartWorkStatuses sets WhenCanStartWorkStatuses field to given value.

### HasWhenCanStartWorkStatuses

`func (o *UserStatusesAvailable) HasWhenCanStartWorkStatuses() bool`

HasWhenCanStartWorkStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


