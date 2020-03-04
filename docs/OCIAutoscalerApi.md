# odk\OCIAutoscalerApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesDisableInstanceAutoscaler**](OCIAutoscalerApi.md#InstancesDisableInstanceAutoscaler) | **Post** /instances/{id}/autoscaler/disable_ticket | Disables autoscaling for instance
[**InstancesEnableInstanceAutoscaler**](OCIAutoscalerApi.md#InstancesEnableInstanceAutoscaler) | **Post** /instances/{id}/autoscaler/enable_ticket | Enables autoscaling for instance
[**InstancesGetAutoscaler**](OCIAutoscalerApi.md#InstancesGetAutoscaler) | **Get** /instances/{id}/autoscaler | Returns instace autoscaler configuration
[**InstancesUpdateAutoscaler**](OCIAutoscalerApi.md#InstancesUpdateAutoscaler) | **Put** /instances/{id}/autoscaler | Updates instance autoscaler configuration


# **InstancesDisableInstanceAutoscaler**
> Ticket InstancesDisableInstanceAutoscaler(ctx, id)
Disables autoscaling for instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesEnableInstanceAutoscaler**
> Ticket InstancesEnableInstanceAutoscaler(ctx, id)
Enables autoscaling for instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetAutoscaler**
> Autoscaler InstancesGetAutoscaler(ctx, id, optional)
Returns instace autoscaler configuration

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

[**Autoscaler**](Autoscaler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesUpdateAutoscaler**
> Object InstancesUpdateAutoscaler(ctx, id, command)
Updates instance autoscaler configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 
  **command** | [**AutoscalerUpdateCommand**](AutoscalerUpdateCommand.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

