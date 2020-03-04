# odk\StatisticsApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatisticsGetClientStatistics**](StatisticsApi.md#StatisticsGetClientStatistics) | **Get** /statistics/account | Gets client statistics
[**StatisticsGetInstanceStatistics**](StatisticsApi.md#StatisticsGetInstanceStatistics) | **Get** /statistics/instances/{id} | Gets instance statistics
[**StatisticsGetStatisticIntervals**](StatisticsApi.md#StatisticsGetStatisticIntervals) | **Get** /statistics/dictionaries/intervals | Gets statistic interval types


# **StatisticsGetClientStatistics**
> ApiCollectionClientStatistics StatisticsGetClientStatistics(ctx, dateFrom, dateTo, statisticTypes, optional)
Gets client statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dateFrom** | **time.Time**| Date from (utc timezone) | 
  **dateTo** | **time.Time**| Date to (utc timezone) | 
  **statisticTypes** | [**[]int32**](int32.md)| Statistic types | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **time.Time**| Date from (utc timezone) | 
 **dateTo** | **time.Time**| Date to (utc timezone) | 
 **statisticTypes** | [**[]int32**](int32.md)| Statistic types | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionClientStatistics**](ApiCollection[ClientStatistics].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGetInstanceStatistics**
> ApiCollectionInstanceStatistics StatisticsGetInstanceStatistics(ctx, id, statisticInterval, dateFrom, dateTo, statisticTypes, optional)
Gets instance statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id | 
  **statisticInterval** | **int32**| Statistic interval | [default to 1519]
  **dateFrom** | **time.Time**| Date from (utc timezone) | 
  **dateTo** | **time.Time**| Date to (utc timezone) | 
  **statisticTypes** | [**[]int32**](int32.md)| Statistic types | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id | 
 **statisticInterval** | **int32**| Statistic interval | [default to 1519]
 **dateFrom** | **time.Time**| Date from (utc timezone) | 
 **dateTo** | **time.Time**| Date to (utc timezone) | 
 **statisticTypes** | [**[]int32**](int32.md)| Statistic types | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceStatistics**](ApiCollection[InstanceStatistics].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGetStatisticIntervals**
> ApiCollectionDictionaryItem StatisticsGetStatisticIntervals(ctx, optional)
Gets statistic interval types

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

