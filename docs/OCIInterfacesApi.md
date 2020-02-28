# odk\OCIInterfacesApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesAttachOpn**](OCIInterfacesApi.md#InstancesAttachOpn) | **Post** /instances/{id}/attach_opn_ticket | Attach instance to OPN
[**InstancesBookNewIp**](OCIInterfacesApi.md#InstancesBookNewIp) | **Post** /instances/ip_addresses | Book new IP address
[**InstancesChangeOpn**](OCIInterfacesApi.md#InstancesChangeOpn) | **Post** /instances/{id}/change_opn_ticket | Change OPN on network interface
[**InstancesDeleteIp**](OCIInterfacesApi.md#InstancesDeleteIp) | **Delete** /instances/ip_addresses/{id} | Deletes IP address
[**InstancesDetachFromOpn**](OCIInterfacesApi.md#InstancesDetachFromOpn) | **Post** /instances/{id}/detach_opn_ticket | Detach instance from OPN
[**InstancesGetAllNetworkInterfaces**](OCIInterfacesApi.md#InstancesGetAllNetworkInterfaces) | **Get** /instances/interfaces | Returns all network interfaces
[**InstancesGetInstanceIp**](OCIInterfacesApi.md#InstancesGetInstanceIp) | **Get** /instances/ip_addresses/{id} | Returns IP by id
[**InstancesGetInstanceIps**](OCIInterfacesApi.md#InstancesGetInstanceIps) | **Get** /instances/{id}/ip_addresses | Returns instance public ip list
[**InstancesGetInstanceNetworkInterfaces**](OCIInterfacesApi.md#InstancesGetInstanceNetworkInterfaces) | **Get** /instances/{id}/interfaces | Returns instance network interfaces
[**InstancesGetIps**](OCIInterfacesApi.md#InstancesGetIps) | **Get** /instances/ip_addresses | Returns public ip list
[**InstancesGetOpns**](OCIInterfacesApi.md#InstancesGetOpns) | **Get** /instances/{id}/opns | Returns instance OPN&#39;s
[**InstancesPostAttachIpTicket**](OCIInterfacesApi.md#InstancesPostAttachIpTicket) | **Post** /instances/{id}/attach_ip_ticket | Attach public IP to instance
[**InstancesPostDetachIpTicket**](OCIInterfacesApi.md#InstancesPostDetachIpTicket) | **Post** /instances/{id}/detach_ip_ticket | Detach public IP from instance
[**InstancesUpdateIp**](OCIInterfacesApi.md#InstancesUpdateIp) | **Put** /instances/ip_addresses/{id} | Updates IP address


# **InstancesAttachOpn**
> Ticket InstancesAttachOpn(ctx, id, command)
Attach instance to OPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **command** | [**AttachInstanceToOpnCommand**](AttachInstanceToOpnCommand.md)| Attach instance to OPN command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesBookNewIp**
> Ip InstancesBookNewIp(ctx, command)
Book new IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**BookIpCommand**](BookIpCommand.md)|  | 

### Return type

[**Ip**](Ip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesChangeOpn**
> Ticket InstancesChangeOpn(ctx, id, command)
Change OPN on network interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **command** | [**ChangeOpnCommand**](ChangeOpnCommand.md)| Change OPN command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesDeleteIp**
> Object InstancesDeleteIp(ctx, id)
Deletes IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| IP address identifier | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesDetachFromOpn**
> Ticket InstancesDetachFromOpn(ctx, id, command)
Detach instance from OPN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **command** | [**DetachInstanceFromOpnCommand**](DetachInstanceFromOpnCommand.md)| Detach instance from OPN command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetAllNetworkInterfaces**
> NetworkInterface InstancesGetAllNetworkInterfaces(ctx, optional)
Returns all network interfaces

Acceptable order values are: MacAddress, Instance, Opn, Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**NetworkInterface**](NetworkInterface.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceIp**
> Ip InstancesGetInstanceIp(ctx, id, optional)
Returns IP by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**|  | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Ip**](Ip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceIps**
> ApiCollectionIp InstancesGetInstanceIps(ctx, id, optional)
Returns instance public ip list

Acceptable order values are: Address, Subregion, Comment, Type

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
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionIp**](ApiCollection[Ip].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceNetworkInterfaces**
> NetworkInterface InstancesGetInstanceNetworkInterfaces(ctx, id, optional)
Returns instance network interfaces

Acceptable order values are: MacAddress, Instance, Opn, Address

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
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**NetworkInterface**](NetworkInterface.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetIps**
> ApiCollectionIp InstancesGetIps(ctx, optional)
Returns public ip list

Acceptable order values are: Address, Subregion, Comment, Type.

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
 **onlyFree** | **bool**| Only free | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionIp**](ApiCollection[Ip].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetOpns**
> ApiCollectionOpn InstancesGetOpns(ctx, id, optional)
Returns instance OPN's

Acceptable order values are: Name

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

# **InstancesPostAttachIpTicket**
> Ticket InstancesPostAttachIpTicket(ctx, id, optional)
Attach public IP to instance

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
 **ipId** | **int32**| IP address identifier. Optional value, if null random ip will be attached. | 
 **ipV6** | **bool**| If attach IPv6 only. Optional value, if null IPv4 and IPv6 will be attached. | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesPostDetachIpTicket**
> Ticket InstancesPostDetachIpTicket(ctx, id, ipId)
Detach public IP from instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **ipId** | **int32**| IP address identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesUpdateIp**
> Object InstancesUpdateIp(ctx, id, command)
Updates IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 
  **command** | [**UpdateIpCommand**](UpdateIpCommand.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

