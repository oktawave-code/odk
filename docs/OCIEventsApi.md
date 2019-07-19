# \OCIEventsApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsDelete**](OCIEventsApi.md#EventsDelete) | **Delete** /events/{id} | Deletes event
[**EventsGetEvent**](OCIEventsApi.md#EventsGetEvent) | **Get** /events/{id} | Returns instance event
[**EventsGetEvents**](OCIEventsApi.md#EventsGetEvents) | **Get** /events | Returns all instances events
[**InstancesDeleteEvents**](OCIEventsApi.md#InstancesDeleteEvents) | **Delete** /instances/{id}/events | Deletes instance events
[**InstancesGetEvents**](OCIEventsApi.md#InstancesGetEvents) | **Get** /instances/{id}/events | Returns instance events


# **EventsDelete**
> EventsDelete(ctx, id)
Deletes event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Event id | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventsGetEvent**
> ApiCollectionInstanceEvent EventsGetEvent(ctx, id, optional)
Returns instance event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| Event id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| Event id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceEvent**](ApiCollection[InstanceEvent].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventsGetEvents**
> ApiCollectionInstanceEvent EventsGetEvents(ctx, optional)
Returns all instances events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **typeIds** | **string**| Type IDs | 
 **statusIds** | **string**| Status IDs | 
 **dateFrom** | **time.Time**| Date from | 
 **dateTo** | **time.Time**| Date to | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceEvent**](ApiCollection[InstanceEvent].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesDeleteEvents**
> InstancesDeleteEvents(ctx, id, command)
Deletes instance events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
  **command** | [**DeleteEventsCommand**](DeleteEventsCommand.md)| Delete events command | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetEvents**
> ApiCollectionInstanceEvent InstancesGetEvents(ctx, id, optional)
Returns instance events

Acceptable order values are: OperationType, User, Date, Instance, OperationStatus

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Instance identifier | 
 **userId** | **int32**| User id | 
 **dateFrom** | **time.Time**| Date from | 
 **dateTo** | **time.Time**| Date to | 
 **operationTypeId** | **int32**| Operation type id | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceEvent**](ApiCollection[InstanceEvent].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

