# GetConversationMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SMSConversationMessage**](SMSConversationMessage.md) |  | 

## Methods

### NewGetConversationMessagesResponse

`func NewGetConversationMessagesResponse(data []SMSConversationMessage, ) *GetConversationMessagesResponse`

NewGetConversationMessagesResponse instantiates a new GetConversationMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversationMessagesResponseWithDefaults

`func NewGetConversationMessagesResponseWithDefaults() *GetConversationMessagesResponse`

NewGetConversationMessagesResponseWithDefaults instantiates a new GetConversationMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetConversationMessagesResponse) GetData() []SMSConversationMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetConversationMessagesResponse) GetDataOk() (*[]SMSConversationMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetConversationMessagesResponse) SetData(v []SMSConversationMessage)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


