# LoadBalancer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | Group id | [optional] [default to null]
**GroupName** | **string** | Group name | [optional] [default to null]
**IpAddress** | **string** | IPv4 address | [optional] [default to null]
**IpV6Address** | **string** | IPv6 address | [optional] [default to null]
**ServiceType** | [***DictionaryItem**](DictionaryItem.md) | Service type (HTTP, SMTP, Port...) | [optional] [default to null]
**PortNumber** | **int32** | Port number for \&quot;Port\&quot; service type | [optional] [default to null]
**TargetPortNumber** | **int32** | Port number for \&quot;TargetPort\&quot; service type | [optional] [default to null]
**SslTargetPortNumber** | **int32** | Ssl port number for \&quot;TargetPort\&quot; service type | [optional] [default to null]
**SessionPersistenceType** | [***DictionaryItem**](DictionaryItem.md) | Session persistence type | [optional] [default to null]
**Algorithm** | [***DictionaryItem**](DictionaryItem.md) | Algorithm | [optional] [default to null]
**IpVersion** | [***DictionaryItem**](DictionaryItem.md) | IP version | [optional] [default to null]
**HealthCheckEnabled** | **bool** | Is health check enabled | [optional] [default to null]
**SslEnabled** | **bool** | Is ssl enabled (only for \&quot;HTTP\&quot; service type) | [optional] [default to null]
**CommonPersistenceForHttpAndHttpsEnabled** | **bool** | Is common persistence for HTTP and HTTPS enabled (only for \&quot;HTTP\&quot; service type) | [optional] [default to null]
**Servers** | [**[]LoadBalancerServer**](LoadBalancerServer.md) | Services | [optional] [default to null]
**ProxyProtocolVersion** | [***DictionaryItem**](DictionaryItem.md) | Proxy-protocol version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


