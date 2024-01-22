# ResumeObjectsRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **string** | Контакт | [optional] 
**Name** | **string** | Имя выдавшего рекомендацию | 
**Organization** | **string** | Организация | 
**Position** | **string** | Должность | 

## Methods

### NewResumeObjectsRecommendation

`func NewResumeObjectsRecommendation(name string, organization string, position string, ) *ResumeObjectsRecommendation`

NewResumeObjectsRecommendation instantiates a new ResumeObjectsRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsRecommendationWithDefaults

`func NewResumeObjectsRecommendationWithDefaults() *ResumeObjectsRecommendation`

NewResumeObjectsRecommendationWithDefaults instantiates a new ResumeObjectsRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *ResumeObjectsRecommendation) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeObjectsRecommendation) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeObjectsRecommendation) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ResumeObjectsRecommendation) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetName

`func (o *ResumeObjectsRecommendation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumeObjectsRecommendation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumeObjectsRecommendation) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *ResumeObjectsRecommendation) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ResumeObjectsRecommendation) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ResumeObjectsRecommendation) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetPosition

`func (o *ResumeObjectsRecommendation) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResumeObjectsRecommendation) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResumeObjectsRecommendation) SetPosition(v string)`

SetPosition sets Position field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


