# ErrorsCommonCaptchaError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptchaUrl** | Pointer to **string** | Адрес веб-страницы, на которой можно пройти капчу.  После прохождения капчи аналогичный запрос в API должен выполниться успешно.  Приложение должно добавить в captcha_url обязательный параметр backurl,на который произойдет редирект   после прохождения капчи.  Backurl должен обязательно содержать схему, например, https:// или схему приложения  | [optional] 
**FallbackUrl** | Pointer to **string** | Адрес веб-страницы, на котором можно капчу. Аналогично параметру captcha_url | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Необходимо пройти капчу - &#x60;captcha_required&#x60;  | 

## Methods

### NewErrorsCommonCaptchaError

`func NewErrorsCommonCaptchaError(type_ string, value string, ) *ErrorsCommonCaptchaError`

NewErrorsCommonCaptchaError instantiates a new ErrorsCommonCaptchaError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonCaptchaErrorWithDefaults

`func NewErrorsCommonCaptchaErrorWithDefaults() *ErrorsCommonCaptchaError`

NewErrorsCommonCaptchaErrorWithDefaults instantiates a new ErrorsCommonCaptchaError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptchaUrl

`func (o *ErrorsCommonCaptchaError) GetCaptchaUrl() string`

GetCaptchaUrl returns the CaptchaUrl field if non-nil, zero value otherwise.

### GetCaptchaUrlOk

`func (o *ErrorsCommonCaptchaError) GetCaptchaUrlOk() (*string, bool)`

GetCaptchaUrlOk returns a tuple with the CaptchaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaUrl

`func (o *ErrorsCommonCaptchaError) SetCaptchaUrl(v string)`

SetCaptchaUrl sets CaptchaUrl field to given value.

### HasCaptchaUrl

`func (o *ErrorsCommonCaptchaError) HasCaptchaUrl() bool`

HasCaptchaUrl returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *ErrorsCommonCaptchaError) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *ErrorsCommonCaptchaError) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *ErrorsCommonCaptchaError) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *ErrorsCommonCaptchaError) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### GetType

`func (o *ErrorsCommonCaptchaError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonCaptchaError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonCaptchaError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonCaptchaError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonCaptchaError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonCaptchaError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


