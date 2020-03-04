# odk\NetworkingApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpnsDelete**](NetworkingApi.md#OpnsDelete) | **Delete** /opns/{id} | Deletes an OPN
[**OpnsGet**](NetworkingApi.md#OpnsGet) | **Get** /opns | Returns OPNs
[**OpnsPost**](NetworkingApi.md#OpnsPost) | **Post** /opns | Creates an OPN
[**OpnsPut**](NetworkingApi.md#OpnsPut) | **Put** /opns/{id} | Updates OPN


# **OpnsDelete**
> Ticket OpnsDelete(ctx, id)
Deletes an OPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| OPN id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpnsGet**
> ApiCollectionOpn OpnsGet(ctx, optional)
Returns OPNs

Acceptable order values are: Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **int32**| Instance id | 
 **query** | **string**| Query | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionOpn**](ApiCollection[Opn].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpnsPost**
> Ticket OpnsPost(ctx, command)
Creates an OPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateOpnCommand**](CreateOpnCommand.md)| Create OPN command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpnsPut**
> Object OpnsPut(ctx, id, command)
Updates OPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| OPN id | 
  **command** | [**UpdateOpnCommand**](UpdateOpnCommand.md)| Update OPN command | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

