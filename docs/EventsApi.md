# \EventsApi

All URIs are relative to *https://api-kourou-0-0-1.endpoints.labs-console-universe-sh.cloud.goog*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvents**](EventsApi.md#AddEvents) | **Post** /v1/events | 



## AddEvents

> Generic AddEvents(ctx, xTOKEN)



Add properties of events

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xTOKEN** | **string**| string X-TOKEN (name or id) of the events | 

### Return type

[**Generic**](generic.md)

### Authorization

[okta_jwt](../README.md#okta_jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

