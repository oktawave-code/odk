# \OCISchedulersApi

All URIs are relative to *https://api.oktawave.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceSchedulersDelete**](OCISchedulersApi.md#InstanceSchedulersDelete) | **Delete** /instances/schedulers/{id} | Deletes instance scheduler
[**InstanceSchedulersGet**](OCISchedulersApi.md#InstanceSchedulersGet) | **Get** /instances/schedulers/{id} | Gets scheduler by identifier
[**InstanceSchedulersGetInstanceSchedulers**](OCISchedulersApi.md#InstanceSchedulersGetInstanceSchedulers) | **Get** /instances/{id}/schedulers | Gets instance schedulers
[**InstanceSchedulersPost**](OCISchedulersApi.md#InstanceSchedulersPost) | **Post** /instances/{id}/schedulers | Creates instance scheduler
[**InstanceSchedulersPut**](OCISchedulersApi.md#InstanceSchedulersPut) | **Put** /instances/schedulers/{id} | Updates instance scheduler
[**InstancesGetSchedulers**](OCISchedulersApi.md#InstancesGetSchedulers) | **Get** /instances/schedulers | Gets schedulers by search params


# **InstanceSchedulersDelete**
> InstanceSchedulersDelete(ctx, id)
Deletes instance scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Scheduler id | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceSchedulersGet**
> InstanceScheduler InstanceSchedulersGet(ctx, id, optional)
Gets scheduler by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Scheduler id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Scheduler id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**InstanceScheduler**](InstanceScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceSchedulersGetInstanceSchedulers**
> ApiCollectionInstanceScheduler InstanceSchedulersGetInstanceSchedulers(ctx, id, optional)
Gets instance schedulers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Instance id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceScheduler**](ApiCollection[InstanceScheduler].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceSchedulersPost**
> InstanceScheduler InstanceSchedulersPost(ctx, id, command)
Creates instance scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
  **command** | [**CreateUpdateInstanceSchedulerCommand**](CreateUpdateInstanceSchedulerCommand.md)| Create instance scheduler command | 

### Return type

[**InstanceScheduler**](InstanceScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceSchedulersPut**
> InstanceScheduler InstanceSchedulersPut(ctx, id, command)
Updates instance scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Scheduler id | 
  **command** | [**CreateUpdateInstanceSchedulerCommand**](CreateUpdateInstanceSchedulerCommand.md)| Create instance scheduler command | 

### Return type

[**InstanceScheduler**](InstanceScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetSchedulers**
> ApiCollectionInstanceScheduler InstancesGetSchedulers(ctx, optional)
Gets schedulers by search params

Acceptable order values are: Name, CreationDate, StartDate, Instance, ActionType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **int32**| Instance id filter | 
 **actionTypeId** | **int32**| Action type id filter | 
 **query** | **string**| Query | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceScheduler**](ApiCollection[InstanceScheduler].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

