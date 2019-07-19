# LoadBalancerServer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | Ip address | [optional] [default to null]
**Port** | **int32** | Port | [optional] [default to null]
**Protocol** | **string** | Protocol | [optional] [default to null]
**ActiveServices** | **int64** | Active services. Null if HealthCheckEnabled in Load balancer disabled. | [optional] [default to null]
**InactiveServices** | **int64** | Inactive services. Null if HealthCheckEnabled in Load balancer disabled. | [optional] [default to null]
**RequestRate** | **int64** | Request rate | [optional] [default to null]
**ResponseRate** | **int64** | Response rate | [optional] [default to null]
**RequestRateBytes** | **int64** | Request rate in bytes | [optional] [default to null]
**ResponseRateBytes** | **int64** | Response rate in bytes | [optional] [default to null]
**CurrentClientConnections** | **int64** | Current client connections | [optional] [default to null]
**Status** | [***DictionaryItem**](DictionaryItem.md) | Status | [optional] [default to null]
**Services** | [**[]LoadBalancerService**](LoadBalancerService.md) | Services statistics. Filled only in load balancer details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


