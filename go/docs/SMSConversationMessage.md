# SMSConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** |  | 
**IsSuccessful** | **bool** |  | 
**IsAutomatic** | **bool** |  | 
**Timestamp** | **string** |  | 
**Text** | **string** |  | 

## Methods

### NewSMSConversationMessage

`func NewSMSConversationMessage(direction string, isSuccessful bool, isAutomatic bool, timestamp string, text string, ) *SMSConversationMessage`

NewSMSConversationMessage instantiates a new SMSConversationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConversationMessageWithDefaults

`func NewSMSConversationMessageWithDefaults() *SMSConversationMessage`

NewSMSConversationMessageWithDefaults instantiates a new SMSConversationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SMSConversationMessage) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SMSConversationMessage) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SMSConversationMessage) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetIsSuccessful

`func (o *SMSConversationMessage) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *SMSConversationMessage) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *SMSConversationMessage) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.


### GetIsAutomatic

`func (o *SMSConversationMessage) GetIsAutomatic() bool`

GetIsAutomatic returns the IsAutomatic field if non-nil, zero value otherwise.

### GetIsAutomaticOk

`func (o *SMSConversationMessage) GetIsAutomaticOk() (*bool, bool)`

GetIsAutomaticOk returns a tuple with the IsAutomatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomatic

`func (o *SMSConversationMessage) SetIsAutomatic(v bool)`

SetIsAutomatic sets IsAutomatic field to given value.


### GetTimestamp

`func (o *SMSConversationMessage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SMSConversationMessage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SMSConversationMessage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetText

`func (o *SMSConversationMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SMSConversationMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SMSConversationMessage) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


