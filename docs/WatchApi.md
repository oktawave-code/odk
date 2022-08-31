# \WatchApi

All URIs are relative to *https://api.oktawave.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchAddSelectedMonitoringStations**](WatchApi.md#WatchAddSelectedMonitoringStations) | **Post** /watch/sensors/assignments | Add monitoring sensor
[**WatchCreateDnsHealthCheck**](WatchApi.md#WatchCreateDnsHealthCheck) | **Post** /watch/healthchecks/dns | Creates dns health check
[**WatchCreateFullPageHealthCheck**](WatchApi.md#WatchCreateFullPageHealthCheck) | **Post** /watch/healthchecks/fullpage | Creates FullPage health check
[**WatchCreateFullPageHttpsHealthCheck**](WatchApi.md#WatchCreateFullPageHttpsHealthCheck) | **Post** /watch/healthchecks/fullpagehttps | Creates FullPage Https health check
[**WatchCreateHealthCheckNotification**](WatchApi.md#WatchCreateHealthCheckNotification) | **Post** /watch/healthchecks/notifications | Creates health check notification
[**WatchCreateHttpHealthCheck**](WatchApi.md#WatchCreateHttpHealthCheck) | **Post** /watch/healthchecks/http | Creates http health check
[**WatchCreateHttpsHealthCheck**](WatchApi.md#WatchCreateHttpsHealthCheck) | **Post** /watch/healthchecks/https | Creates https health check
[**WatchCreateImapHealthCheck**](WatchApi.md#WatchCreateImapHealthCheck) | **Post** /watch/healthchecks/imap | Creates imap health check
[**WatchCreateImapSslHealthCheck**](WatchApi.md#WatchCreateImapSslHealthCheck) | **Post** /watch/healthchecks/imapssl | Creates imap ssl health check
[**WatchCreatePingHealthCheck**](WatchApi.md#WatchCreatePingHealthCheck) | **Post** /watch/healthchecks/ping | Creates ping health check
[**WatchCreateSipHealthCheck**](WatchApi.md#WatchCreateSipHealthCheck) | **Post** /watch/healthchecks/sip | Creates sip health check
[**WatchCreateSmtpHealthCheck**](WatchApi.md#WatchCreateSmtpHealthCheck) | **Post** /watch/healthchecks/smtp | Creates smtp health check
[**WatchCreateTcpHealthCheck**](WatchApi.md#WatchCreateTcpHealthCheck) | **Post** /watch/healthchecks/tcp | Creates Tcp health check
[**WatchDeleteHealthCheck**](WatchApi.md#WatchDeleteHealthCheck) | **Delete** /watch/healthchecks/{id} | Deletes health check
[**WatchDeleteHealthCheckNotification**](WatchApi.md#WatchDeleteHealthCheckNotification) | **Delete** /watch/healthchecks/notifications/{id} | Deletes health check notification
[**WatchDeleteSelectedMonitoringStations**](WatchApi.md#WatchDeleteSelectedMonitoringStations) | **Delete** /watch/sensors/assignments/{id} | Remove monitoring sensor
[**WatchGetAvailableMonitoringStations**](WatchApi.md#WatchGetAvailableMonitoringStations) | **Get** /watch/sensors/all | Gets all available monitoring sensors
[**WatchGetDnsHealthCheck**](WatchApi.md#WatchGetDnsHealthCheck) | **Get** /watch/healthchecks/dns/{id} | Returns dns health check details
[**WatchGetFullPageHealthCheck**](WatchApi.md#WatchGetFullPageHealthCheck) | **Get** /watch/healthchecks/fullpage/{id} | Returns FullPage health check details
[**WatchGetFullPageHttpsHealthCheck**](WatchApi.md#WatchGetFullPageHttpsHealthCheck) | **Get** /watch/healthchecks/fullpagehttps/{id} | Returns FullPage Https health check details
[**WatchGetHealthCheck**](WatchApi.md#WatchGetHealthCheck) | **Get** /watch/healthchecks/{id} | Returns health check
[**WatchGetHealthCheckNotification**](WatchApi.md#WatchGetHealthCheckNotification) | **Get** /watch/healthchecks/notifications/{id} | Returns health check notification details
[**WatchGetHealthCheckNotifications**](WatchApi.md#WatchGetHealthCheckNotifications) | **Get** /watch/healthchecks/notifications | Returns a list of configured health check notifications
[**WatchGetHealthChecks**](WatchApi.md#WatchGetHealthChecks) | **Get** /watch/healthchecks | Returns a list of configured health checks
[**WatchGetHttpHealthCheck**](WatchApi.md#WatchGetHttpHealthCheck) | **Get** /watch/healthchecks/http/{id} | Returns http health check details
[**WatchGetHttpsHealthCheck**](WatchApi.md#WatchGetHttpsHealthCheck) | **Get** /watch/healthchecks/https/{id} | Returns https health check details
[**WatchGetImapHealthCheck**](WatchApi.md#WatchGetImapHealthCheck) | **Get** /watch/healthchecks/imap/{id} | Returns imap health check details
[**WatchGetImapSslHealthCheck**](WatchApi.md#WatchGetImapSslHealthCheck) | **Get** /watch/healthchecks/imapssl/{id} | Returns imap ssl health check details
[**WatchGetPingHealthCheck**](WatchApi.md#WatchGetPingHealthCheck) | **Get** /watch/healthchecks/ping/{id} | Returns ping health check details
[**WatchGetSelectedMonitoringStation**](WatchApi.md#WatchGetSelectedMonitoringStation) | **Get** /watch/sensors/assignments/{id} | Gets selected monitoring sensor
[**WatchGetSelectedMonitoringStations**](WatchApi.md#WatchGetSelectedMonitoringStations) | **Get** /watch/sensors/assignments | Gets selected monitoring sensors
[**WatchGetSipHealthCheck**](WatchApi.md#WatchGetSipHealthCheck) | **Get** /watch/healthchecks/sip/{id} | Returns sip health check details
[**WatchGetSmtpHealthCheck**](WatchApi.md#WatchGetSmtpHealthCheck) | **Get** /watch/healthchecks/smtp/{id} | Returns smtp health check details
[**WatchGetTcpHealthCheck**](WatchApi.md#WatchGetTcpHealthCheck) | **Get** /watch/healthchecks/tcp/{id} | Returns tcp health check details
[**WatchUpdateDnsHealthCheck**](WatchApi.md#WatchUpdateDnsHealthCheck) | **Put** /watch/healthchecks/dns/{id} | Updates dns health check
[**WatchUpdateFullPageHealthCheck**](WatchApi.md#WatchUpdateFullPageHealthCheck) | **Put** /watch/healthchecks/fullpage/{id} | Updates FullPage health check
[**WatchUpdateFullPageHttpsHealthCheck**](WatchApi.md#WatchUpdateFullPageHttpsHealthCheck) | **Put** /watch/healthchecks/fullpagehttps/{id} | Updates FullPage Https health check
[**WatchUpdateHealthCheckNotification**](WatchApi.md#WatchUpdateHealthCheckNotification) | **Put** /watch/healthchecks/notifications/{id} | Updates health check notification
[**WatchUpdateHttpHealthCheck**](WatchApi.md#WatchUpdateHttpHealthCheck) | **Put** /watch/healthchecks/http/{id} | Updates http health check
[**WatchUpdateHttpsHealthCheck**](WatchApi.md#WatchUpdateHttpsHealthCheck) | **Put** /watch/healthchecks/https/{id} | Updates https health check
[**WatchUpdateImapHealthCheck**](WatchApi.md#WatchUpdateImapHealthCheck) | **Put** /watch/healthchecks/imap/{id} | Updates imap health check
[**WatchUpdateImapHealthCheck_0**](WatchApi.md#WatchUpdateImapHealthCheck_0) | **Put** /watch/healthchecks/sip/{id} | Updates sip health check
[**WatchUpdateImapSslHealthCheck**](WatchApi.md#WatchUpdateImapSslHealthCheck) | **Put** /watch/healthchecks/imapssl/{id} | Updates imap ssl health check
[**WatchUpdatePingHealthCheck**](WatchApi.md#WatchUpdatePingHealthCheck) | **Put** /watch/healthchecks/ping/{id} | Updates ping health check
[**WatchUpdateSmtpHealthCheck**](WatchApi.md#WatchUpdateSmtpHealthCheck) | **Put** /watch/healthchecks/smtp/{id} | Updates smtp health check
[**WatchUpdateTcpHealthCheck**](WatchApi.md#WatchUpdateTcpHealthCheck) | **Put** /watch/healthchecks/tcp/{id} | Updates Tcp health check


# **WatchAddSelectedMonitoringStations**
> MonitoringSensor WatchAddSelectedMonitoringStations(ctx, station)
Add monitoring sensor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **station** | [**AssignNewMonitoringSensorCommand**](AssignNewMonitoringSensorCommand.md)| Station | 

### Return type

[**MonitoringSensor**](MonitoringSensor.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateDnsHealthCheck**
> HealthCheckDns WatchCreateDnsHealthCheck(ctx, command)
Creates dns health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckDnsCommand**](CreateUpdateHealthCheckDnsCommand.md)| Create dns health check command | 

### Return type

[**HealthCheckDns**](HealthCheckDns.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateFullPageHealthCheck**
> HealthCheckFullPage WatchCreateFullPageHealthCheck(ctx, command)
Creates FullPage health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckFullPageCommand**](CreateUpdateHealthCheckFullPageCommand.md)| Create FullPage health check command | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateFullPageHttpsHealthCheck**
> HealthCheckFullPage WatchCreateFullPageHttpsHealthCheck(ctx, command)
Creates FullPage Https health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckFullPageHttpsCommand**](CreateUpdateHealthCheckFullPageHttpsCommand.md)| Create FullPage health check command | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateHealthCheckNotification**
> HealthCheckNotification WatchCreateHealthCheckNotification(ctx, command)
Creates health check notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckNotificationCommand**](CreateUpdateHealthCheckNotificationCommand.md)| Create health check notification command | 

### Return type

[**HealthCheckNotification**](HealthCheckNotification.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateHttpHealthCheck**
> HealthCheckHttp WatchCreateHttpHealthCheck(ctx, command)
Creates http health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckHttpCommand**](CreateUpdateHealthCheckHttpCommand.md)| Create http health check command | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateHttpsHealthCheck**
> HealthCheckHttp WatchCreateHttpsHealthCheck(ctx, command)
Creates https health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckHttpsCommand**](CreateUpdateHealthCheckHttpsCommand.md)| Create https health check command | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateImapHealthCheck**
> HealthCheckImap WatchCreateImapHealthCheck(ctx, command)
Creates imap health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckImapCommand**](CreateUpdateHealthCheckImapCommand.md)| Create imap health check command | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateImapSslHealthCheck**
> HealthCheckImap WatchCreateImapSslHealthCheck(ctx, command)
Creates imap ssl health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckImapSslCommand**](CreateUpdateHealthCheckImapSslCommand.md)| Create imap ssl health check command | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreatePingHealthCheck**
> HealthCheckPing WatchCreatePingHealthCheck(ctx, command)
Creates ping health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckPingCommand**](CreateUpdateHealthCheckPingCommand.md)| Create ping health check command | 

### Return type

[**HealthCheckPing**](HealthCheckPing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateSipHealthCheck**
> HealthCheckSip WatchCreateSipHealthCheck(ctx, command)
Creates sip health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckSipCommand**](CreateUpdateHealthCheckSipCommand.md)| Create sip health check command | 

### Return type

[**HealthCheckSip**](HealthCheckSip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateSmtpHealthCheck**
> HealthCheckSmtp WatchCreateSmtpHealthCheck(ctx, command)
Creates smtp health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckSmtpCommand**](CreateUpdateHealthCheckSmtpCommand.md)| Create smtp health check command | 

### Return type

[**HealthCheckSmtp**](HealthCheckSmtp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchCreateTcpHealthCheck**
> HealthCheckTcp WatchCreateTcpHealthCheck(ctx, command)
Creates Tcp health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateUpdateHealthCheckTcpCommand**](CreateUpdateHealthCheckTcpCommand.md)| Create Tcp health check command | 

### Return type

[**HealthCheckTcp**](HealthCheckTcp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchDeleteHealthCheck**
> WatchDeleteHealthCheck(ctx, id)
Deletes health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchDeleteHealthCheckNotification**
> WatchDeleteHealthCheckNotification(ctx, id)
Deletes health check notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check notification | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchDeleteSelectedMonitoringStations**
> MonitoringSensor WatchDeleteSelectedMonitoringStations(ctx, id)
Remove monitoring sensor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Sensor id | 

### Return type

[**MonitoringSensor**](MonitoringSensor.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetAvailableMonitoringStations**
> MonitoringSensor WatchGetAvailableMonitoringStations(ctx, optional)
Gets all available monitoring sensors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Response fields filter | 

### Return type

[**MonitoringSensor**](MonitoringSensor.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetDnsHealthCheck**
> HealthCheckDns WatchGetDnsHealthCheck(ctx, id, optional)
Returns dns health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckDns**](HealthCheckDns.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetFullPageHealthCheck**
> HealthCheckFullPage WatchGetFullPageHealthCheck(ctx, id, optional)
Returns FullPage health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetFullPageHttpsHealthCheck**
> HealthCheckFullPage WatchGetFullPageHttpsHealthCheck(ctx, id, optional)
Returns FullPage Https health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHealthCheck**
> HealthCheck WatchGetHealthCheck(ctx, id, optional)
Returns health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheck**](HealthCheck.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHealthCheckNotification**
> HealthCheckNotification WatchGetHealthCheckNotification(ctx, id, optional)
Returns health check notification details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check notification | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check notification | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckNotification**](HealthCheckNotification.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHealthCheckNotifications**
> ApiCollectionHealthCheckNotification WatchGetHealthCheckNotifications(ctx, optional)
Returns a list of configured health check notifications

Acceptable order values are: Address, Id

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

[**ApiCollectionHealthCheckNotification**](ApiCollection[HealthCheckNotification].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHealthChecks**
> ApiCollectionHealthCheck WatchGetHealthChecks(ctx, optional)
Returns a list of configured health checks

Acceptable order values are: Name, Type

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

[**ApiCollectionHealthCheck**](ApiCollection[HealthCheck].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHttpHealthCheck**
> HealthCheckHttp WatchGetHttpHealthCheck(ctx, id, optional)
Returns http health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetHttpsHealthCheck**
> HealthCheckHttp WatchGetHttpsHealthCheck(ctx, id, optional)
Returns https health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetImapHealthCheck**
> HealthCheckImap WatchGetImapHealthCheck(ctx, id, optional)
Returns imap health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetImapSslHealthCheck**
> HealthCheckImap WatchGetImapSslHealthCheck(ctx, id, optional)
Returns imap ssl health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetPingHealthCheck**
> HealthCheckPing WatchGetPingHealthCheck(ctx, id, optional)
Returns ping health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckPing**](HealthCheckPing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetSelectedMonitoringStation**
> MonitoringSensor WatchGetSelectedMonitoringStation(ctx, id, optional)
Gets selected monitoring sensor

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

[**MonitoringSensor**](MonitoringSensor.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetSelectedMonitoringStations**
> MonitoringSensor WatchGetSelectedMonitoringStations(ctx, optional)
Gets selected monitoring sensors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string**| Response fields filter | 

### Return type

[**MonitoringSensor**](MonitoringSensor.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetSipHealthCheck**
> HealthCheckSip WatchGetSipHealthCheck(ctx, id, optional)
Returns sip health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckSip**](HealthCheckSip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetSmtpHealthCheck**
> HealthCheckSmtp WatchGetSmtpHealthCheck(ctx, id, optional)
Returns smtp health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckSmtp**](HealthCheckSmtp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchGetTcpHealthCheck**
> HealthCheckTcp WatchGetTcpHealthCheck(ctx, id, optional)
Returns tcp health check details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a health check | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a health check | 
 **fields** | **string**| Response fields filter | 

### Return type

[**HealthCheckTcp**](HealthCheckTcp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateDnsHealthCheck**
> HealthCheckDns WatchUpdateDnsHealthCheck(ctx, id, command)
Updates dns health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckDnsCommand**](CreateUpdateHealthCheckDnsCommand.md)| Create dns health check command | 

### Return type

[**HealthCheckDns**](HealthCheckDns.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateFullPageHealthCheck**
> HealthCheckFullPage WatchUpdateFullPageHealthCheck(ctx, id, command)
Updates FullPage health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckFullPageCommand**](CreateUpdateHealthCheckFullPageCommand.md)| Create FullPage health check command | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateFullPageHttpsHealthCheck**
> HealthCheckFullPage WatchUpdateFullPageHttpsHealthCheck(ctx, id, command)
Updates FullPage Https health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckFullPageHttpsCommand**](CreateUpdateHealthCheckFullPageHttpsCommand.md)| Create FullPage health check command | 

### Return type

[**HealthCheckFullPage**](HealthCheckFullPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateHealthCheckNotification**
> HealthCheckNotification WatchUpdateHealthCheckNotification(ctx, id, command)
Updates health check notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check notification id | 
  **command** | [**CreateUpdateHealthCheckNotificationCommand**](CreateUpdateHealthCheckNotificationCommand.md)| Create health check notification command | 

### Return type

[**HealthCheckNotification**](HealthCheckNotification.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateHttpHealthCheck**
> HealthCheckHttp WatchUpdateHttpHealthCheck(ctx, id, command)
Updates http health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckHttpCommand**](CreateUpdateHealthCheckHttpCommand.md)| Create http health check command | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateHttpsHealthCheck**
> HealthCheckHttp WatchUpdateHttpsHealthCheck(ctx, id, command)
Updates https health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckHttpsCommand**](CreateUpdateHealthCheckHttpsCommand.md)| Create https health check command | 

### Return type

[**HealthCheckHttp**](HealthCheckHttp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateImapHealthCheck**
> HealthCheckImap WatchUpdateImapHealthCheck(ctx, id, command)
Updates imap health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckImapCommand**](CreateUpdateHealthCheckImapCommand.md)| Create imap health check command | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateImapHealthCheck_0**
> HealthCheckSip WatchUpdateImapHealthCheck_0(ctx, id, command)
Updates sip health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckSipCommand**](CreateUpdateHealthCheckSipCommand.md)| Create sip health check command | 

### Return type

[**HealthCheckSip**](HealthCheckSip.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateImapSslHealthCheck**
> HealthCheckImap WatchUpdateImapSslHealthCheck(ctx, id, command)
Updates imap ssl health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckImapSslCommand**](CreateUpdateHealthCheckImapSslCommand.md)| Create imap ssl health check command | 

### Return type

[**HealthCheckImap**](HealthCheckImap.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdatePingHealthCheck**
> HealthCheckPing WatchUpdatePingHealthCheck(ctx, id, command)
Updates ping health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckPingCommand**](CreateUpdateHealthCheckPingCommand.md)| Create ping health check command | 

### Return type

[**HealthCheckPing**](HealthCheckPing.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateSmtpHealthCheck**
> HealthCheckSmtp WatchUpdateSmtpHealthCheck(ctx, id, command)
Updates smtp health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckSmtpCommand**](CreateUpdateHealthCheckSmtpCommand.md)| Create smtp health check command | 

### Return type

[**HealthCheckSmtp**](HealthCheckSmtp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchUpdateTcpHealthCheck**
> HealthCheckTcp WatchUpdateTcpHealthCheck(ctx, id, command)
Updates Tcp health check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Health check id | 
  **command** | [**CreateUpdateHealthCheckTcpCommand**](CreateUpdateHealthCheckTcpCommand.md)| Create Tcp health check command | 

### Return type

[**HealthCheckTcp**](HealthCheckTcp.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

