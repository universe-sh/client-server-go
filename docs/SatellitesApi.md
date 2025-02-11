# \SatellitesApi

All URIs are relative to *https://api-kourou-0-0-1.endpoints.labs-console-universe-sh.cloud.goog*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetrics**](SatellitesApi.md#CreateMetrics) | **Post** /v1/satellites/{cloud}/{satellite}/{pool}/metrics | 
[**ListPools**](SatellitesApi.md#ListPools) | **Get** /v1/satellites/{cloud}/{satellite}/pools | 
[**ReadSatellite**](SatellitesApi.md#ReadSatellite) | **Get** /v1/satellites/{cloud}/{satellite} | 



## CreateMetrics

> Generic CreateMetrics(ctx, cloud, satellite, pool)



Create properties of metrics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloud** | **string**| string cloud (name or id) of the metrics | 
**satellite** | **string**| string satellite (name or id) of the metrics | 
**pool** | **string**| string pool (name or id) of the metrics | 

### Return type

[**Generic**](generic.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPools

> []Pool ListPools(ctx, cloud, satellite)



List properties of pools

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloud** | **string**| string cloud (name or id) of the pools | 
**satellite** | **string**| string satellite (name or id) of the pools | 

### Return type

[**[]Pool**](pool.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSatellite

> Satellite ReadSatellite(ctx, cloud, satellite)



Read properties of satellite

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloud** | **string**| string cloud (name or id) of the satellite | 
**satellite** | **string**| string satellite (name or id) of the satellite | 

### Return type

[**Satellite**](satellite.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

