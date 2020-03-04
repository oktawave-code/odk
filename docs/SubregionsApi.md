# odk\SubregionsApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubregionsGet**](SubregionsApi.md#SubregionsGet) | **Get** /subregions | Gets subregions
[**SubregionsGet_0**](SubregionsApi.md#SubregionsGet_0) | **Get** /subregions/{id} | Gets subregion


# **SubregionsGet**
> ApiCollectionSubregion SubregionsGet(ctx, optional)
Gets subregions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionSubregion**](ApiCollection[Subregion].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubregionsGet_0**
> Subregion SubregionsGet_0(ctx, id, optional)
Gets subregion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Subregion id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Subregion id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Subregion**](Subregion.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

