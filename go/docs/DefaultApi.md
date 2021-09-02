# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesSiteIDSmsConversationsGet**](DefaultApi.md#SitesSiteIDSmsConversationsGet) | **Get** /sites/{siteID}/sms-conversations | List of SMS conversations
[**SitesSiteIDSmsConversationsPhoneNumberGet**](DefaultApi.md#SitesSiteIDSmsConversationsPhoneNumberGet) | **Get** /sites/{siteID}/sms-conversations/{phoneNumber} | List of SMS conversation messages to a given phone number
[**SitesSiteIDSmsConversationsPhoneNumberMessagePost**](DefaultApi.md#SitesSiteIDSmsConversationsPhoneNumberMessagePost) | **Post** /sites/{siteID}/sms-conversations/{phoneNumber}/message | Send test message to a given conversation
[**SmsConversationPhoneNumberGet**](DefaultApi.md#SmsConversationPhoneNumberGet) | **Get** /sms-conversation/{phoneNumber} | List of SMS conversation messages to a given phone number
[**SmsConversationsGet**](DefaultApi.md#SmsConversationsGet) | **Get** /sms-conversations | List of SMS conversations
[**SmsConversationsPhoneNumberMessagePost**](DefaultApi.md#SmsConversationsPhoneNumberMessagePost) | **Post** /sms-conversations/{phoneNumber}/message | Send test message to a given conversation



## SitesSiteIDSmsConversationsGet

> SitesSiteIDSmsConversationsGet(ctx, siteID).Execute()

List of SMS conversations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteID := "siteID_example" // string | Site id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SitesSiteIDSmsConversationsGet(context.Background(), siteID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SitesSiteIDSmsConversationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteID** | **string** | Site id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesSiteIDSmsConversationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesSiteIDSmsConversationsPhoneNumberGet

> SitesSiteIDSmsConversationsPhoneNumberGet(ctx, siteID, phoneNumber).Execute()

List of SMS conversation messages to a given phone number

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteID := "siteID_example" // string | Site id
    phoneNumber := "phoneNumber_example" // string | Phone number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SitesSiteIDSmsConversationsPhoneNumberGet(context.Background(), siteID, phoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SitesSiteIDSmsConversationsPhoneNumberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteID** | **string** | Site id | 
**phoneNumber** | **string** | Phone number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesSiteIDSmsConversationsPhoneNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesSiteIDSmsConversationsPhoneNumberMessagePost

> SitesSiteIDSmsConversationsPhoneNumberMessagePost(ctx, siteID, phoneNumber).SendMessageToConversationRequest(sendMessageToConversationRequest).Execute()

Send test message to a given conversation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteID := "siteID_example" // string | Site id
    phoneNumber := "phoneNumber_example" // string | Phone number
    sendMessageToConversationRequest := *openapiclient.NewSendMessageToConversationRequest("Text_example") // SendMessageToConversationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SitesSiteIDSmsConversationsPhoneNumberMessagePost(context.Background(), siteID, phoneNumber).SendMessageToConversationRequest(sendMessageToConversationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SitesSiteIDSmsConversationsPhoneNumberMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteID** | **string** | Site id | 
**phoneNumber** | **string** | Phone number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesSiteIDSmsConversationsPhoneNumberMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sendMessageToConversationRequest** | [**SendMessageToConversationRequest**](SendMessageToConversationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmsConversationPhoneNumberGet

> SmsConversationPhoneNumberGet(ctx, phoneNumber).Execute()

List of SMS conversation messages to a given phone number

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    phoneNumber := "phoneNumber_example" // string | Phone number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SmsConversationPhoneNumberGet(context.Background(), phoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SmsConversationPhoneNumberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmsConversationPhoneNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmsConversationsGet

> SmsConversationsGet(ctx).Execute()

List of SMS conversations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SmsConversationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SmsConversationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSmsConversationsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmsConversationsPhoneNumberMessagePost

> SmsConversationsPhoneNumberMessagePost(ctx, phoneNumber).SendMessageToConversationRequest(sendMessageToConversationRequest).Execute()

Send test message to a given conversation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    phoneNumber := "phoneNumber_example" // string | Phone number
    sendMessageToConversationRequest := *openapiclient.NewSendMessageToConversationRequest("Text_example") // SendMessageToConversationRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SmsConversationsPhoneNumberMessagePost(context.Background(), phoneNumber).SendMessageToConversationRequest(sendMessageToConversationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SmsConversationsPhoneNumberMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmsConversationsPhoneNumberMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendMessageToConversationRequest** | [**SendMessageToConversationRequest**](SendMessageToConversationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

