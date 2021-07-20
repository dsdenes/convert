/*
 * Converted from api.yaml with typeconv
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Books struct for Books
type Books struct {
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
	Authors []Author `json:"authors"`
	Type AnyOfBookTypeAnyType `json:"type"`
	Editions map[string]Edition `json:"editions"`
}

// NewBooks instantiates a new Books object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooks(title string, authors []Author, type_ AnyOfBookTypeAnyType, editions map[string]Edition) *Books {
	this := Books{}
	this.Title = title
	this.Authors = authors
	this.Type = type_
	this.Editions = editions
	return &this
}

// NewBooksWithDefaults instantiates a new Books object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooksWithDefaults() *Books {
	this := Books{}
	return &this
}

// GetTitle returns the Title field value
func (o *Books) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Books) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Books) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Books) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Books) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Books) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Books) SetDescription(v string) {
	o.Description = &v
}

// GetAuthors returns the Authors field value
func (o *Books) GetAuthors() []Author {
	if o == nil {
		var ret []Author
		return ret
	}

	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value
// and a boolean to check if the value has been set.
func (o *Books) GetAuthorsOk() (*[]Author, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Authors, true
}

// SetAuthors sets field value
func (o *Books) SetAuthors(v []Author) {
	o.Authors = v
}

// GetType returns the Type field value
func (o *Books) GetType() AnyOfBookTypeAnyType {
	if o == nil {
		var ret AnyOfBookTypeAnyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Books) GetTypeOk() (*AnyOfBookTypeAnyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Books) SetType(v AnyOfBookTypeAnyType) {
	o.Type = v
}

// GetEditions returns the Editions field value
func (o *Books) GetEditions() map[string]Edition {
	if o == nil {
		var ret map[string]Edition
		return ret
	}

	return o.Editions
}

// GetEditionsOk returns a tuple with the Editions field value
// and a boolean to check if the value has been set.
func (o *Books) GetEditionsOk() (*map[string]Edition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Editions, true
}

// SetEditions sets field value
func (o *Books) SetEditions(v map[string]Edition) {
	o.Editions = v
}

func (o Books) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["authors"] = o.Authors
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["editions"] = o.Editions
	}
	return json.Marshal(toSerialize)
}

type NullableBooks struct {
	value *Books
	isSet bool
}

func (v NullableBooks) Get() *Books {
	return v.value
}

func (v *NullableBooks) Set(val *Books) {
	v.value = val
	v.isSet = true
}

func (v NullableBooks) IsSet() bool {
	return v.isSet
}

func (v *NullableBooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooks(val *Books) *NullableBooks {
	return &NullableBooks{value: val, isSet: true}
}

func (v NullableBooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


