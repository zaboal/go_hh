# ArtifactsFieldsFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | **float32** | Максимальный размер файла | 
**MimeType** | **[]string** | Список допустимых [MIME-типов](https://www.iana.org/assignments/media-types/media-types.xhtml) файлов | 
**Required** | **bool** | Является ли поле &#x60;file&#x60; обязательным | 

## Methods

### NewArtifactsFieldsFile

`func NewArtifactsFieldsFile(maxSize float32, mimeType []string, required bool, ) *ArtifactsFieldsFile`

NewArtifactsFieldsFile instantiates a new ArtifactsFieldsFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsFieldsFileWithDefaults

`func NewArtifactsFieldsFileWithDefaults() *ArtifactsFieldsFile`

NewArtifactsFieldsFileWithDefaults instantiates a new ArtifactsFieldsFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *ArtifactsFieldsFile) GetMaxSize() float32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *ArtifactsFieldsFile) GetMaxSizeOk() (*float32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *ArtifactsFieldsFile) SetMaxSize(v float32)`

SetMaxSize sets MaxSize field to given value.


### GetMimeType

`func (o *ArtifactsFieldsFile) GetMimeType() []string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ArtifactsFieldsFile) GetMimeTypeOk() (*[]string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ArtifactsFieldsFile) SetMimeType(v []string)`

SetMimeType sets MimeType field to given value.


### GetRequired

`func (o *ArtifactsFieldsFile) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ArtifactsFieldsFile) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ArtifactsFieldsFile) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


