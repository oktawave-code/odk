# \TicketsApi

All URIs are relative to *https://api.oktawave.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketsGet**](TicketsApi.md#TicketsGet) | **Get** /tickets | Returns ticket collection
[**TicketsGet_0**](TicketsApi.md#TicketsGet_0) | **Get** /tickets/{id} | Returns ticket by identifier


# **TicketsGet**
> ApiCollectionTicket TicketsGet(ctx, optional)
Returns ticket collection

Acceptable order values are: CreationDate, Status, OperationType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusId** | **int32**| Tickets status id | 
 **creationDateFrom** | **time.Time**| Tickets creation date from | 
 **creationDateTo** | **time.Time**| Tickets creation date to | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionTicket**](ApiCollection[Ticket].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TicketsGet_0**
> Ticket TicketsGet_0(ctx, id, optional)
Returns ticket by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int64**| Ticket identifier | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| Ticket identifier | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

