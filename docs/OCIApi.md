# \OCIApi

All URIs are relative to *https://api.oktawave.com/services*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstancesChangeHotPlugSetting**](OCIApi.md#InstancesChangeHotPlugSetting) | **Post** /instances/{id}/change_hotplug_ticket | Change instance hotplug setting
[**InstancesChangeName**](OCIApi.md#InstancesChangeName) | **Post** /instances/{id}/change_name_ticket | Change instance name
[**InstancesChangeSubregion**](OCIApi.md#InstancesChangeSubregion) | **Post** /instances/{id}/change_subregion_ticket | Change instance subregion
[**InstancesChangeType**](OCIApi.md#InstancesChangeType) | **Post** /instances/{id}/change_scsi_controllers_type_ticket | Change SCSI controllers type
[**InstancesChangeType_0**](OCIApi.md#InstancesChangeType_0) | **Post** /instances/{id}/change_type_ticket | Change instance type
[**InstancesClone**](OCIApi.md#InstancesClone) | **Post** /instances/{id}/clone_ticket | Clone instance
[**InstancesConvertToTemplate**](OCIApi.md#InstancesConvertToTemplate) | **Post** /instances/{id}/convert_to_template_ticket | Converts instance to template
[**InstancesCreateVncConnection**](OCIApi.md#InstancesCreateVncConnection) | **Post** /instances/{id}/remote_console_connection | Creates remote console connection
[**InstancesDelete**](OCIApi.md#InstancesDelete) | **Delete** /instances/{id} | Delete instance
[**InstancesDeleteVncConnection**](OCIApi.md#InstancesDeleteVncConnection) | **Delete** /instances/{id}/remote_console_connection | Deletes remote console connection
[**InstancesGet**](OCIApi.md#InstancesGet) | **Get** /instances | Returns instance list
[**InstancesGetAccessData**](OCIApi.md#InstancesGetAccessData) | **Get** /instances/{id}/access_data | Returns instance access data
[**InstancesGetDisks**](OCIApi.md#InstancesGetDisks) | **Get** /instances/{id}/disks | Returns instance disk list
[**InstancesGetInstanceInitScript**](OCIApi.md#InstancesGetInstanceInitScript) | **Get** /instances/{id}/init_script | Returns instance init script
[**InstancesGetInstanceSoftware**](OCIApi.md#InstancesGetInstanceSoftware) | **Get** /instances/{id}/software | Returns instance software
[**InstancesGetInstanceType**](OCIApi.md#InstancesGetInstanceType) | **Get** /instances/types/{id} | Returns instance type
[**InstancesGetInstancesTypes**](OCIApi.md#InstancesGetInstancesTypes) | **Get** /instances/types | Returns all available instances types
[**InstancesGetScreenshot**](OCIApi.md#InstancesGetScreenshot) | **Get** /instances/{id}/screenshot | Returns instance screenshot
[**InstancesGetSoftware**](OCIApi.md#InstancesGetSoftware) | **Get** /instances/software | Returns software
[**InstancesGetSshKeys**](OCIApi.md#InstancesGetSshKeys) | **Get** /instances/{id}/ssh_keys | Returns SSH keys uploaded to instances
[**InstancesGetTemplateByBaseVirtualMachineId**](OCIApi.md#InstancesGetTemplateByBaseVirtualMachineId) | **Get** /instances/{id}/template | Returns the template by source virtual machine id
[**InstancesGetVncConnection**](OCIApi.md#InstancesGetVncConnection) | **Get** /instances/{id}/remote_console_connection | Returns remote console connection
[**InstancesGet_0**](OCIApi.md#InstancesGet_0) | **Get** /instances/{id} | Returns instance by identifier
[**InstancesPost**](OCIApi.md#InstancesPost) | **Post** /instances | Creates instance
[**InstancesPowerOff**](OCIApi.md#InstancesPowerOff) | **Post** /instances/{id}/power_off_ticket | Power off instance
[**InstancesPowerOn**](OCIApi.md#InstancesPowerOn) | **Post** /instances/{id}/power_on_ticket | Power on instance
[**InstancesReboot**](OCIApi.md#InstancesReboot) | **Post** /instances/{id}/reboot_ticket | Reboot instance
[**InstancesReset**](OCIApi.md#InstancesReset) | **Post** /instances/{id}/reset_ticket | Reset instance
[**InstancesShutdown**](OCIApi.md#InstancesShutdown) | **Post** /instances/{id}/shutdown_ticket | Shutdown instance
[**InstancesUpdateVncConnection**](OCIApi.md#InstancesUpdateVncConnection) | **Put** /instances/{id}/remote_console_connection | Updates remote console connection


# **InstancesChangeHotPlugSetting**
> Ticket InstancesChangeHotPlugSetting(ctx, id, hotPlugEnabled)
Change instance hotplug setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 
  **hotPlugEnabled** | **bool**| Is hot plug enabled | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesChangeName**
> Ticket InstancesChangeName(ctx, id, name)
Change instance name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **name** | **string**| Name of an instance | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesChangeSubregion**
> Ticket InstancesChangeSubregion(ctx, id, command)
Change instance subregion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **command** | [**ChangeInstanceSubregionCommand**](ChangeInstanceSubregionCommand.md)| Change instance subregion command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesChangeType**
> Ticket InstancesChangeType(ctx, id, scsiControllerTypeId)
Change SCSI controllers type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **scsiControllerTypeId** | **int32**| Type id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesChangeType_0**
> Ticket InstancesChangeType_0(ctx, id, typeId)
Change instance type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **typeId** | **int32**| Type id | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesClone**
> Ticket InstancesClone(ctx, id, command)
Clone instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance Id | 
  **command** | [**CloneInstanceCommand**](CloneInstanceCommand.md)| Clone instance command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesConvertToTemplate**
> Ticket InstancesConvertToTemplate(ctx, id, command)
Converts instance to template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 
  **command** | [**ConvertInstanceToTemplateCommand**](ConvertInstanceToTemplateCommand.md)| Convert instance to template command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesCreateVncConnection**
> VncConnection InstancesCreateVncConnection(ctx, id, command)
Creates remote console connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
  **command** | [**CreateUpdateVncConnectionCommand**](CreateUpdateVncConnectionCommand.md)| Create remote console connection command | 

### Return type

[**VncConnection**](VncConnection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesDelete**
> Ticket InstancesDelete(ctx, id, optional)
Delete instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifie | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Instance identifie | 
 **deep** | **bool**| Deletes also additional disks attached to instance | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesDeleteVncConnection**
> Object InstancesDeleteVncConnection(ctx, id)
Deletes remote console connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGet**
> ApiCollectionInstance InstancesGet(ctx, optional)
Returns instance list

Acceptable order values are: Type, Status, CreationDate, Name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateTypeId** | **int32**| Template type id eg marketplace, oci instance | 
 **isTurnedOn** | **bool**| Indicates wether an instance is turned on | 
 **subregionId** | **int32**| Subregion Id | 
 **typeId** | **int32**| Type Id | 
 **query** | **string**| Query | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstance**](ApiCollection[Instance].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetAccessData**
> AccessData InstancesGetAccessData(ctx, id, optional)
Returns instance access data

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
 **fields** | **string**| Response fields filter | 

### Return type

[**AccessData**](AccessData.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetDisks**
> ApiCollectionDisk InstancesGetDisks(ctx, id, optional)
Returns instance disk list

Acceptable order values are: SpaceCapacity, Name, Tier, IsShared, Subregion

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
 **diskType** | **string**| Disk type | 
 **showDeleted** | **bool**| Show deleted | 
 **query** | **string**| Query | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionDisk**](ApiCollection[Disk].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceInitScript**
> string InstancesGetInstanceInitScript(ctx, id, optional)
Returns instance init script

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

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceSoftware**
> ApiCollectionSoftware InstancesGetInstanceSoftware(ctx, id, optional)
Returns instance software

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

[**ApiCollectionSoftware**](ApiCollection[Software].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstanceType**
> ApiCollectionInstanceType InstancesGetInstanceType(ctx, id, optional)
Returns instance type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance type id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Instance type id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceType**](ApiCollection[InstanceType].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetInstancesTypes**
> ApiCollectionInstanceType InstancesGetInstancesTypes(ctx, optional)
Returns all available instances types

Acceptable order values are: Category, Cpu, Ram, Name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **int32**| Category id | 
 **availableForFreemium** | **bool**| Is available for freemium | 
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceType**](ApiCollection[InstanceType].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetScreenshot**
> string InstancesGetScreenshot(ctx, id, optional)
Returns instance screenshot

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
 **width** | **int32**| The pixel width of the scaled image | 
 **height** | **int32**| The pixel height of the scaled image | 
 **fields** | **string**| Response fields filter | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetSoftware**
> ApiCollectionSoftware InstancesGetSoftware(ctx, optional)
Returns software

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

[**ApiCollectionSoftware**](ApiCollection[Software].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetSshKeys**
> ApiCollectionInstanceSshKey InstancesGetSshKeys(ctx, id, optional)
Returns SSH keys uploaded to instances

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
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionInstanceSshKey**](ApiCollection[InstanceSshKey].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetTemplateByBaseVirtualMachineId**
> Template InstancesGetTemplateByBaseVirtualMachineId(ctx, id, optional)
Returns the template by source virtual machine id

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

[**Template**](Template.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetVncConnection**
> VncConnection InstancesGetVncConnection(ctx, id, optional)
Returns remote console connection

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
 **fields** | **string**| Response fields filter | 

### Return type

[**VncConnection**](VncConnection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGet_0**
> Instance InstancesGet_0(ctx, id, optional)
Returns instance by identifier

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
 **fields** | **string**| Response fields filter | 

### Return type

[**Instance**](Instance.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesPost**
> Ticket InstancesPost(ctx, command)
Creates instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateInstanceCommand**](CreateInstanceCommand.md)| Create instance command | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesPowerOff**
> Ticket InstancesPowerOff(ctx, id)
Power off instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesPowerOn**
> Ticket InstancesPowerOn(ctx, id)
Power on instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesReboot**
> Ticket InstancesReboot(ctx, id)
Reboot instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesReset**
> Ticket InstancesReset(ctx, id)
Reset instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesShutdown**
> Ticket InstancesShutdown(ctx, id)
Shutdown instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance identifier | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesUpdateVncConnection**
> VncConnection InstancesUpdateVncConnection(ctx, id, command)
Updates remote console connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
  **command** | [**CreateUpdateVncConnectionCommand**](CreateUpdateVncConnectionCommand.md)| Update remote console connection command | 

### Return type

[**VncConnection**](VncConnection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

