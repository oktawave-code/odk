# CreateUpdateHealthCheckTcpCommand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Port | [default to 80]
**Timeout** | **int32** | Time the server has to complete the request in [ms] | [default to 10000]
**ErrorTolerance** | **int32** | How many (%) locations have to report an error to consider it a breakdown | [default to 51]
**Name** | **string** | Health check name | [default to null]
**Address** | **string** | URL or IP address of the monitored service | [default to null]
**Interval** | **int32** | Time interval between health checks, in seconds | [default to 60]
**Paused** | **bool** | Is paused | [default to null]
**LocationsFailoverEnabled** | **bool** | Use random substitute locations in case of location breakdown | [default to null]
**NotificationTypeIds** | **[]int32** | Notification types enabled for a health check | [optional] [default to null]
**NotificationEventTypeIds** | **[]int32** | Event types with enabled notification | [optional] [default to null]
**NotificationTimeId** | **int32** | Time when notification is sent | [default to 1594]
**Description** | **string** | Description | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


