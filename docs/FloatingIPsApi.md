# \FloatingIPsApi

All URIs are relative to *https://api.oktawave.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FloatingIpsBookNewIp**](FloatingIPsApi.md#FloatingIpsBookNewIp) | **Post** /floating_ips | Book new IP address
[**FloatingIpsChangeIpSubregionTicket**](FloatingIPsApi.md#FloatingIpsChangeIpSubregionTicket) | **Post** /floating_ips/change_ip_subregion_ticket | Change IP subregion
[**FloatingIpsDeleteIp**](FloatingIPsApi.md#FloatingIpsDeleteIp) | **Delete** /floating_ips/{ip} | Deletes IP address
[**FloatingIpsGetIp**](FloatingIPsApi.md#FloatingIpsGetIp) | **Get** /floating_ips/{ip} | Returns IP by Ipv4 address
[**FloatingIpsGetIps**](FloatingIPsApi.md#FloatingIpsGetIps) | **Get** /floating_ips | Returns public ip list
[**FloatingIpsPostAttachIpTicket**](FloatingIPsApi.md#FloatingIpsPostAttachIpTicket) | **Post** /floating_ips/attach_ip_ticket | Attach public IP to instance
[**FloatingIpsPostDetachIpTicket**](FloatingIPsApi.md#FloatingIpsPostDetachIpTicket) | **Post** /floating_ips/detach_ip_ticket | Detach public IP from instance
[**FloatingIpsUpdateIp**](FloatingIPsApi.md#FloatingIpsUpdateIp) | **Put** /floating_ips/{ip} | Updates IP address


# **FloatingIpsBookNewIp**
> Ip FloatingIpsBookNewIp(ctx, command)
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

# **FloatingIpsChangeIpSubregionTicket**
> Ticket FloatingIpsChangeIpSubregionTicket(ctx, ipV4, optional)
Change IP subregion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ipV4** | **string**| IPv4 address identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipV4** | **string**| IPv4 address identifier. | 
 **subregionId** | **int32**| Subregion Id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FloatingIpsDeleteIp**
> Object FloatingIpsDeleteIp(ctx, ip)
Deletes IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ip** | **string**| IPv4 address | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FloatingIpsGetIp**
> Ip FloatingIpsGetIp(ctx, ip, optional)
Returns IP by Ipv4 address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ip** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string**|  | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Ip**](Ip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FloatingIpsGetIps**
> ApiCollectionIp FloatingIpsGetIps(ctx, optional)
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

# **FloatingIpsPostAttachIpTicket**
> Ticket FloatingIpsPostAttachIpTicket(ctx, instanceId, optional)
Attach public IP to instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **instanceId** | **int32**| Instance identifier. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **int32**| Instance identifier. | 
 **ipV4** | **string**| IPv4 address identifier. Optional value, if null random ip will be attached. | 
 **ipV6Only** | **bool**| If attach IPv6 only. Optional value, if null IPv4 and IPv6 will be attached. | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FloatingIpsPostDetachIpTicket**
> Ticket FloatingIpsPostDetachIpTicket(ctx, ipV4, instanceId)
Detach public IP from instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ipV4** | **string**| IPv4 address identifier | 
  **instanceId** | **int32**| Instance identifier. | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FloatingIpsUpdateIp**
> Object FloatingIpsUpdateIp(ctx, ip, command)
Updates IP address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ip** | **string**|  | 
  **command** | [**UpdateIpCommand**](UpdateIpCommand.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

