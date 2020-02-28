# SetLoadBalancerCommand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslEnabled** | **bool** | Is ssl enabled (only for \&quot;HTTP\&quot; load balancer service type) | [optional] [default to null]
**ServiceType** | **int32** | Load balancer service type | [default to null]
**PortNumber** | **int32** | Port number for \&quot;Port\&quot; load balancer service type | [optional] [default to null]
**TargetPortNumber** | **int32** | Port number for \&quot;TargetPort\&quot; load balancer service type | [optional] [default to null]
**SslTargetPortNumber** | **int32** | Ssl port number for \&quot;TargetPort\&quot; load balancer service type | [optional] [default to null]
**SessionPersistenceType** | **int32** | Session persistence type | [default to null]
**LoadBalancerAlgorithm** | **int32** | Load balancing algorithm | [default to null]
**IpVersion** | **int32** | Ip version for load balancing | [default to null]
**HealthCheckEnabled** | **bool** | Is health check enabled | [optional] [default to null]
**CommonPersistenceForHttpAndHttpsEnabled** | **bool** | Is common persistence for HTTP and HTTPS enabled (only for \&quot;HTTP\&quot; load balancer service type) | [optional] [default to null]
**LoadBalancerIpId** | **int32** | Public ip id for load balancer | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


