# Books

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Authors** | [**[]Author**](Author.md) |  | 
**Type** | [**AnyOfBookTypeAnyType**](anyOf&lt;BookType,AnyType&gt;.md) |  | 
**Editions** | [**map[string]Edition**](Edition.md) |  | 

## Methods

### NewBooks

`func NewBooks(title string, authors []Author, type_ AnyOfBookTypeAnyType, editions map[string]Edition, ) *Books`

NewBooks instantiates a new Books object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksWithDefaults

`func NewBooksWithDefaults() *Books`

NewBooksWithDefaults instantiates a new Books object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Books) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Books) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Books) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Books) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Books) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Books) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Books) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthors

`func (o *Books) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Books) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Books) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetType

`func (o *Books) GetType() AnyOfBookTypeAnyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Books) GetTypeOk() (*AnyOfBookTypeAnyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Books) SetType(v AnyOfBookTypeAnyType)`

SetType sets Type field to given value.


### GetEditions

`func (o *Books) GetEditions() map[string]Edition`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *Books) GetEditionsOk() (*map[string]Edition, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *Books) SetEditions(v map[string]Edition)`

SetEditions sets Editions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


