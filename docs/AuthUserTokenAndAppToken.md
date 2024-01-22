# AuthUserTokenAndAppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**ExpiresIn** | **int64** | Время жизни токена в секундах | 
**RefreshToken** | **string** | Токен, который можно использовать для продления срока жизни соответствующего access токена | 
**TokenType** | **string** | Тип выданного токена. Всегда принимает значение &#x60;bearer&#x60; | 

## Methods

### NewAuthUserTokenAndAppToken

`func NewAuthUserTokenAndAppToken(accessToken string, expiresIn int64, refreshToken string, tokenType string, ) *AuthUserTokenAndAppToken`

NewAuthUserTokenAndAppToken instantiates a new AuthUserTokenAndAppToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserTokenAndAppTokenWithDefaults

`func NewAuthUserTokenAndAppTokenWithDefaults() *AuthUserTokenAndAppToken`

NewAuthUserTokenAndAppTokenWithDefaults instantiates a new AuthUserTokenAndAppToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AuthUserTokenAndAppToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthUserTokenAndAppToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthUserTokenAndAppToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *AuthUserTokenAndAppToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AuthUserTokenAndAppToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AuthUserTokenAndAppToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.


### GetRefreshToken

`func (o *AuthUserTokenAndAppToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthUserTokenAndAppToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthUserTokenAndAppToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetTokenType

`func (o *AuthUserTokenAndAppToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AuthUserTokenAndAppToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AuthUserTokenAndAppToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


