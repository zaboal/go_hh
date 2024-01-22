# ResumeObjectsOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | [**ResumeObjectsOwnerComments**](ResumeObjectsOwnerComments.md) | Раздел [Комментарии к владельцу резюме](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list) | 
**Id** | **string** | Идентификатор владельца резюме | 

## Methods

### NewResumeObjectsOwner

`func NewResumeObjectsOwner(comments ResumeObjectsOwnerComments, id string, ) *ResumeObjectsOwner`

NewResumeObjectsOwner instantiates a new ResumeObjectsOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsOwnerWithDefaults

`func NewResumeObjectsOwnerWithDefaults() *ResumeObjectsOwner`

NewResumeObjectsOwnerWithDefaults instantiates a new ResumeObjectsOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *ResumeObjectsOwner) GetComments() ResumeObjectsOwnerComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ResumeObjectsOwner) GetCommentsOk() (*ResumeObjectsOwnerComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ResumeObjectsOwner) SetComments(v ResumeObjectsOwnerComments)`

SetComments sets Comments field to given value.


### GetId

`func (o *ResumeObjectsOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeObjectsOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeObjectsOwner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


