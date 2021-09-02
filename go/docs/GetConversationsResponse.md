# GetConversationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SMSConversation**](SMSConversation.md) |  | 

## Methods

### NewGetConversationsResponse

`func NewGetConversationsResponse(data []SMSConversation, ) *GetConversationsResponse`

NewGetConversationsResponse instantiates a new GetConversationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversationsResponseWithDefaults

`func NewGetConversationsResponseWithDefaults() *GetConversationsResponse`

NewGetConversationsResponseWithDefaults instantiates a new GetConversationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetConversationsResponse) GetData() []SMSConversation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetConversationsResponse) GetDataOk() (*[]SMSConversation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetConversationsResponse) SetData(v []SMSConversation)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


