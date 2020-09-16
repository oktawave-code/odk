# \OCIBackupsApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportsGet**](OCIBackupsApi.md#ExportsGet) | **Get** /exports | Returns exports list
[**ExportsGet_0**](OCIBackupsApi.md#ExportsGet_0) | **Get** /exports/{id} | Gets export by identifier
[**ImportsCreateImport**](OCIBackupsApi.md#ImportsCreateImport) | **Post** /imports | Creates an import
[**ImportsDeleteImport**](OCIBackupsApi.md#ImportsDeleteImport) | **Delete** /imports/{id} | Deletes an import
[**ImportsGetImport**](OCIBackupsApi.md#ImportsGetImport) | **Get** /imports/{id} | Returns an import
[**ImportsGetImportDisks**](OCIBackupsApi.md#ImportsGetImportDisks) | **Get** /imports/{id}/disks | Returns an import disks
[**ImportsGetImports**](OCIBackupsApi.md#ImportsGetImports) | **Get** /imports | Returns a list of imports
[**ImportsRunImport**](OCIBackupsApi.md#ImportsRunImport) | **Post** /imports/{id}/run_ticket | Run an import
[**InstancesGetExports**](OCIBackupsApi.md#InstancesGetExports) | **Get** /instances/{id}/exports | Returns instace&#39;s exports
[**InstancesPostExport**](OCIBackupsApi.md#InstancesPostExport) | **Post** /instances/{id}/exports | Create export


# **ExportsGet**
> ApiCollectionExport ExportsGet(ctx, optional)
Returns exports list

Acceptable order values are: TotalSpaceCapacity, CreationDate, OcsLocation, Name, StartDate, Status.

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
 **statusId** | **int32**| Status id | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionExport**](ApiCollection[Export].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportsGet_0**
> Export ExportsGet_0(ctx, id, optional)
Gets export by identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Export id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Export id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Export**](Export.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsCreateImport**
> ModelImport ImportsCreateImport(ctx, command)
Creates an import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateImportCommand**](CreateImportCommand.md)| Create import command | 

### Return type

[**ModelImport**](Import.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsDeleteImport**
> ImportsDeleteImport(ctx, id)
Deletes an import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Import identifier | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsGetImport**
> ModelImport ImportsGetImport(ctx, id, optional)
Returns an import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Import identifier | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Import identifier | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ModelImport**](Import.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsGetImportDisks**
> ApiCollectionImportDisk ImportsGetImportDisks(ctx, id, optional)
Returns an import disks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Import identifier | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Import identifier | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionImportDisk**](ApiCollection[ImportDisk].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsGetImports**
> ApiCollectionImport ImportsGetImports(ctx, optional)
Returns a list of imports

Acceptable order values are: Name, Status, CreationDate, DisksCount, NetworkInterfacesCount, Size

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

[**ApiCollectionImport**](ApiCollection[Import].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportsRunImport**
> Ticket ImportsRunImport(ctx, id, command)
Run an import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Import identifier | 
  **command** | [**RunImportCommand**](RunImportCommand.md)| Run import command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetExports**
> ApiCollectionExport InstancesGetExports(ctx, id, optional)
Returns instace's exports

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
 **statusId** | **int32**| Status id | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionExport**](ApiCollection[Export].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesPostExport**
> Export InstancesPostExport(ctx, id, command)
Create export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
  **command** | [**CreateExportCommand**](CreateExportCommand.md)| Create export command | 

### Return type

[**Export**](Export.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

