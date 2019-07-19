# HealthCheckFullPage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethod** | [***DictionaryItem**](DictionaryItem.md) | Health check http method type | [optional] [default to null]
**ContentRegularExpression** | **string** | The content has to match the expression (GET and POST methods only) | [optional] [default to null]
**ContentNegativeRegularExpression** | **string** | The content cannot match the expression (GET and POST methods only) | [optional] [default to null]
**Port** | **int32** | Port | [optional] [default to null]
**PageTimeout** | **int32** | Time limit for the main page body [ms] | [optional] [default to null]
**ElementTimeout** | **int32** | Time limit for each page elements [ms] | [optional] [default to null]
**ElementsTotalTimeout** | **int32** | Time limit for all page elements [ms] | [optional] [default to null]
**FetchPageElements** | **bool** | Fetch page elements | [optional] [default to null]
**MaxRedirects** | **int32** | Maximum length of HTTP redirects sequence | [optional] [default to null]
**MaxParallelRequests** | **int32** | Maximum number of HTTP requests run in parallel | [optional] [default to null]
**GenerateHar** | **bool** | Generate a HAR file for each check | [optional] [default to null]
**AllowedElementErrorCount** | **int32** | Number of elements that may not be fetched correctly | [optional] [default to null]
**ContentSizeLimit** | **int32** | Content size limit (bytes) | [optional] [default to null]
**IgnoreHtmlParsingTime** | **bool** | Ignore HTML code processing time in results | [optional] [default to null]
**SaveCookies** | **bool** | Save cookies between checks | [optional] [default to null]
**DisableContentEncoding** | **bool** | Disable HTTP transfer compression | [optional] [default to null]
**Content** | **string** | Content | [optional] [default to null]
**ContentType** | **string** | Content type | [optional] [default to null]
**IgnoredElementsFilter** | **string** | Ignore errors for elements with URLs matching the expression (only if page elements are fetched) | [optional] [default to null]
**ElementsFilter** | **string** | Do not fetch elements with URLs that match the expression | [optional] [default to null]
**NotificationTypes** | [**[]DictionaryItem**](DictionaryItem.md) | Notification types enabled for a health check | [optional] [default to null]
**NotificationEventTypes** | [**[]DictionaryItem**](DictionaryItem.md) | Event types with enabled notification | [optional] [default to null]
**NotificationTime** | [***DictionaryItem**](DictionaryItem.md) | Time when notification is sent | [optional] [default to null]
**LocationsFailoverEnabled** | **bool** | Use random substitute locations in case of location breakdown | [optional] [default to null]
**ErrorTolerance** | **int32** | How many (%) locations have to report an error to consider it a breakdown | [optional] [default to null]
**Id** | **int32** | Id | [optional] [default to null]
**Interval** | **int32** | Interval | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Address** | **string** | Address | [optional] [default to null]
**ServiceType** | [***DictionaryItem**](DictionaryItem.md) | Type | [optional] [default to null]
**State** | [***DictionaryItem**](DictionaryItem.md) | State | [optional] [default to null]
**DetailsLocation** | **string** | Details url | [optional] [default to null]
**Paused** | **bool** | Is paused | [optional] [default to null]
**Suspended** | **bool** | Is suspended | [optional] [default to null]
**LastInvalidCheck** | [**time.Time**](time.Time.md) | Last invalid check | [optional] [default to null]
**LastValidCheck** | [**time.Time**](time.Time.md) | Last valid check | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


