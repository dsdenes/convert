# SMSConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** |  | 
**LastInteractionAt** | **string** |  | 
**LastMessageText** | **string** |  | 
**LastMessageFailed** | **bool** |  | 

## Methods

### NewSMSConversation

`func NewSMSConversation(phoneNumber string, lastInteractionAt string, lastMessageText string, lastMessageFailed bool, ) *SMSConversation`

NewSMSConversation instantiates a new SMSConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConversationWithDefaults

`func NewSMSConversationWithDefaults() *SMSConversation`

NewSMSConversationWithDefaults instantiates a new SMSConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *SMSConversation) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SMSConversation) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SMSConversation) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetLastInteractionAt

`func (o *SMSConversation) GetLastInteractionAt() string`

GetLastInteractionAt returns the LastInteractionAt field if non-nil, zero value otherwise.

### GetLastInteractionAtOk

`func (o *SMSConversation) GetLastInteractionAtOk() (*string, bool)`

GetLastInteractionAtOk returns a tuple with the LastInteractionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInteractionAt

`func (o *SMSConversation) SetLastInteractionAt(v string)`

SetLastInteractionAt sets LastInteractionAt field to given value.


### GetLastMessageText

`func (o *SMSConversation) GetLastMessageText() string`

GetLastMessageText returns the LastMessageText field if non-nil, zero value otherwise.

### GetLastMessageTextOk

`func (o *SMSConversation) GetLastMessageTextOk() (*string, bool)`

GetLastMessageTextOk returns a tuple with the LastMessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageText

`func (o *SMSConversation) SetLastMessageText(v string)`

SetLastMessageText sets LastMessageText field to given value.


### GetLastMessageFailed

`func (o *SMSConversation) GetLastMessageFailed() bool`

GetLastMessageFailed returns the LastMessageFailed field if non-nil, zero value otherwise.

### GetLastMessageFailedOk

`func (o *SMSConversation) GetLastMessageFailedOk() (*bool, bool)`

GetLastMessageFailedOk returns a tuple with the LastMessageFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageFailed

`func (o *SMSConversation) SetLastMessageFailed(v bool)`

SetLastMessageFailed sets LastMessageFailed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


