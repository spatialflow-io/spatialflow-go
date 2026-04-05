# \BillingAPI

All URIs are relative to *https://api.spatialflow.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsSubscriptionsBillingApiAddPaymentMethod**](BillingAPI.md#AppsSubscriptionsBillingApiAddPaymentMethod) | **Post** /api/v1/billing/payment-methods | Add Payment Method
[**AppsSubscriptionsBillingApiChangePlan**](BillingAPI.md#AppsSubscriptionsBillingApiChangePlan) | **Post** /api/v1/billing/change-plan | Change Plan
[**AppsSubscriptionsBillingApiCreateSetupIntent**](BillingAPI.md#AppsSubscriptionsBillingApiCreateSetupIntent) | **Post** /api/v1/billing/create-setup-intent | Create Setup Intent
[**AppsSubscriptionsBillingApiDownloadInvoice**](BillingAPI.md#AppsSubscriptionsBillingApiDownloadInvoice) | **Get** /api/v1/billing/invoices/{invoice_id}/download | Download Invoice
[**AppsSubscriptionsBillingApiGetInvoice**](BillingAPI.md#AppsSubscriptionsBillingApiGetInvoice) | **Get** /api/v1/billing/invoices/{invoice_id} | Get Invoice
[**AppsSubscriptionsBillingApiListInvoices**](BillingAPI.md#AppsSubscriptionsBillingApiListInvoices) | **Get** /api/v1/billing/invoices | List Invoices
[**AppsSubscriptionsBillingApiListPaymentMethods**](BillingAPI.md#AppsSubscriptionsBillingApiListPaymentMethods) | **Get** /api/v1/billing/payment-methods | List Payment Methods
[**AppsSubscriptionsBillingApiPreviewPlanChange**](BillingAPI.md#AppsSubscriptionsBillingApiPreviewPlanChange) | **Post** /api/v1/billing/preview-plan-change | Preview Plan Change
[**AppsSubscriptionsBillingApiRemovePaymentMethod**](BillingAPI.md#AppsSubscriptionsBillingApiRemovePaymentMethod) | **Delete** /api/v1/billing/payment-methods/{pm_id} | Remove Payment Method
[**AppsSubscriptionsBillingApiSetDefaultPaymentMethod**](BillingAPI.md#AppsSubscriptionsBillingApiSetDefaultPaymentMethod) | **Put** /api/v1/billing/payment-methods/{pm_id}/default | Set Default Payment Method



## AppsSubscriptionsBillingApiAddPaymentMethod

> map[string]interface{} AppsSubscriptionsBillingApiAddPaymentMethod(ctx).PaymentMethodRequest(paymentMethodRequest).Execute()

Add Payment Method



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	paymentMethodRequest := *openapiclient.NewPaymentMethodRequest("PaymentMethodId_example") // PaymentMethodRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiAddPaymentMethod(context.Background()).PaymentMethodRequest(paymentMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiAddPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiAddPaymentMethod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiAddPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiAddPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodRequest** | [**PaymentMethodRequest**](PaymentMethodRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiChangePlan

> map[string]interface{} AppsSubscriptionsBillingApiChangePlan(ctx).PlanChangeRequest(planChangeRequest).Execute()

Change Plan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	planChangeRequest := *openapiclient.NewPlanChangeRequest("NewPlanId_example") // PlanChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiChangePlan(context.Background()).PlanChangeRequest(planChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiChangePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiChangePlan`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiChangePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiChangePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planChangeRequest** | [**PlanChangeRequest**](PlanChangeRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiCreateSetupIntent

> SetupIntentResponse AppsSubscriptionsBillingApiCreateSetupIntent(ctx).Execute()

Create Setup Intent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiCreateSetupIntent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiCreateSetupIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiCreateSetupIntent`: SetupIntentResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiCreateSetupIntent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiCreateSetupIntentRequest struct via the builder pattern


### Return type

[**SetupIntentResponse**](SetupIntentResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiDownloadInvoice

> map[string]interface{} AppsSubscriptionsBillingApiDownloadInvoice(ctx, invoiceId).Execute()

Download Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	invoiceId := "invoiceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiDownloadInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiDownloadInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiDownloadInvoice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiDownloadInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiDownloadInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiGetInvoice

> InvoiceResponse AppsSubscriptionsBillingApiGetInvoice(ctx, invoiceId).Execute()

Get Invoice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	invoiceId := "invoiceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiGetInvoice(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiGetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiGetInvoice`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiGetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiListInvoices

> InvoiceListResponse AppsSubscriptionsBillingApiListInvoices(ctx).Limit(limit).StartingAfter(startingAfter).Execute()

List Invoices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	limit := int32(56) // int32 |  (optional) (default to 10)
	startingAfter := "startingAfter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiListInvoices(context.Background()).Limit(limit).StartingAfter(startingAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiListInvoices`: InvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 10]
 **startingAfter** | **string** |  | 

### Return type

[**InvoiceListResponse**](InvoiceListResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiListPaymentMethods

> []PaymentMethodResponse AppsSubscriptionsBillingApiListPaymentMethods(ctx).Execute()

List Payment Methods



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiListPaymentMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiListPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiListPaymentMethods`: []PaymentMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiListPaymentMethodsRequest struct via the builder pattern


### Return type

[**[]PaymentMethodResponse**](PaymentMethodResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiPreviewPlanChange

> PlanChangePreviewResponse AppsSubscriptionsBillingApiPreviewPlanChange(ctx).PlanChangeRequest(planChangeRequest).Execute()

Preview Plan Change



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	planChangeRequest := *openapiclient.NewPlanChangeRequest("NewPlanId_example") // PlanChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiPreviewPlanChange(context.Background()).PlanChangeRequest(planChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiPreviewPlanChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiPreviewPlanChange`: PlanChangePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiPreviewPlanChange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiPreviewPlanChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planChangeRequest** | [**PlanChangeRequest**](PlanChangeRequest.md) |  | 

### Return type

[**PlanChangePreviewResponse**](PlanChangePreviewResponse.md)

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiRemovePaymentMethod

> map[string]interface{} AppsSubscriptionsBillingApiRemovePaymentMethod(ctx, pmId).Execute()

Remove Payment Method



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	pmId := "pmId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiRemovePaymentMethod(context.Background(), pmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiRemovePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiRemovePaymentMethod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiRemovePaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiRemovePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsSubscriptionsBillingApiSetDefaultPaymentMethod

> map[string]interface{} AppsSubscriptionsBillingApiSetDefaultPaymentMethod(ctx, pmId).Execute()

Set Default Payment Method



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/generated"
)

func main() {
	pmId := "pmId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AppsSubscriptionsBillingApiSetDefaultPaymentMethod(context.Background(), pmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AppsSubscriptionsBillingApiSetDefaultPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsSubscriptionsBillingApiSetDefaultPaymentMethod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AppsSubscriptionsBillingApiSetDefaultPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsSubscriptionsBillingApiSetDefaultPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[JWTBearer](../README.md#JWTBearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

