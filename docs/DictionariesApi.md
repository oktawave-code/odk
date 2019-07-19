# \DictionariesApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DictionariesGetDictionaries**](DictionariesApi.md#DictionariesGetDictionaries) | **Get** /dictionaries | Returns dictionaries
[**DictionariesGetDictionariesItems**](DictionariesApi.md#DictionariesGetDictionariesItems) | **Get** /dictionaries/{ids} | Returns dictionaries items
[**DictionariesGetLanguages**](DictionariesApi.md#DictionariesGetLanguages) | **Get** /dictionaries/languages | Returns languages


# **DictionariesGetDictionaries**
> ApiCollectionDictionary DictionariesGetDictionaries(ctx, optional)
Returns dictionaries

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

[**ApiCollectionDictionary**](ApiCollection[Dictionary].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DictionariesGetDictionariesItems**
> ApiCollectionDictionaryItem DictionariesGetDictionariesItems(ctx, ids, optional)
Returns dictionaries items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ids** | **string**| Dictionaries ids | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string**| Dictionaries ids | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionDictionaryItem**](ApiCollection[DictionaryItem].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DictionariesGetLanguages**
> ApiCollectionDictionaryItem DictionariesGetLanguages(ctx, optional)
Returns languages

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

[**ApiCollectionDictionaryItem**](ApiCollection[DictionaryItem].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

