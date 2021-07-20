# Book

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Authors** | [**[]Author**](Author.md) |  | 
**Type** | [**AnyOfBookTypeAnyType**](anyOf&lt;BookType,AnyType&gt;.md) |  | 
**Editions** | [**map[string]Edition**](Edition.md) |  | 

## Methods

### NewBook

`func NewBook(title string, authors []Author, type_ AnyOfBookTypeAnyType, editions map[string]Edition, ) *Book`

NewBook instantiates a new Book object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookWithDefaults

`func NewBookWithDefaults() *Book`

NewBookWithDefaults instantiates a new Book object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Book) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Book) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Book) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Book) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Book) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Book) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Book) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthors

`func (o *Book) GetAuthors() []Author`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Book) GetAuthorsOk() (*[]Author, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Book) SetAuthors(v []Author)`

SetAuthors sets Authors field to given value.


### GetType

`func (o *Book) GetType() AnyOfBookTypeAnyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Book) GetTypeOk() (*AnyOfBookTypeAnyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Book) SetType(v AnyOfBookTypeAnyType)`

SetType sets Type field to given value.


### GetEditions

`func (o *Book) GetEditions() map[string]Edition`

GetEditions returns the Editions field if non-nil, zero value otherwise.

### GetEditionsOk

`func (o *Book) GetEditionsOk() (*map[string]Edition, bool)`

GetEditionsOk returns a tuple with the Editions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditions

`func (o *Book) SetEditions(v map[string]Edition)`

SetEditions sets Editions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


