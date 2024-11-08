/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LinodeStackScriptsAPIService LinodeStackScriptsAPI service
type LinodeStackScriptsAPIService service

type ApiDeleteStackScriptRequest struct {
	ctx context.Context
	ApiService *LinodeStackScriptsAPIService
	apiVersion string
	stackscriptId string
}

func (r ApiDeleteStackScriptRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteStackScriptExecute(r)
}

/*
DeleteStackScript Delete a StackScript

Deletes a private StackScript you have permission to `read_write`. You cannot delete a public StackScript.


<<LB>>

---


- __CLI__.

    ```
    linode-cli stackscripts delete 10079
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    stackscripts:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param stackscriptId The ID of the StackScript to look up.
 @return ApiDeleteStackScriptRequest
*/
func (a *LinodeStackScriptsAPIService) DeleteStackScript(ctx context.Context, apiVersion string, stackscriptId string) ApiDeleteStackScriptRequest {
	return ApiDeleteStackScriptRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		stackscriptId: stackscriptId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LinodeStackScriptsAPIService) DeleteStackScriptExecute(r ApiDeleteStackScriptRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeStackScriptsAPIService.DeleteStackScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/linode/stackscripts/{stackscriptId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stackscriptId"+"}", url.PathEscape(parameterValueToString(r.stackscriptId, "stackscriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStackScriptRequest struct {
	ctx context.Context
	ApiService *LinodeStackScriptsAPIService
	apiVersion string
	stackscriptId string
}

func (r ApiGetStackScriptRequest) Execute() (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	return r.ApiService.GetStackScriptExecute(r)
}

/*
GetStackScript Get a StackScript

Returns all of the information about a specified StackScript, including the contents of the script.


<<LB>>

---


- __CLI__.

    ```
    linode-cli stackscripts view 10079
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    stackscripts:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param stackscriptId The ID of the StackScript to look up.
 @return ApiGetStackScriptRequest
*/
func (a *LinodeStackScriptsAPIService) GetStackScript(ctx context.Context, apiVersion string, stackscriptId string) ApiGetStackScriptRequest {
	return ApiGetStackScriptRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		stackscriptId: stackscriptId,
	}
}

// Execute executes the request
//  @return GetStackScripts200ResponseDataInner
func (a *LinodeStackScriptsAPIService) GetStackScriptExecute(r ApiGetStackScriptRequest) (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStackScripts200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeStackScriptsAPIService.GetStackScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/linode/stackscripts/{stackscriptId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stackscriptId"+"}", url.PathEscape(parameterValueToString(r.stackscriptId, "stackscriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStackScriptsRequest struct {
	ctx context.Context
	ApiService *LinodeStackScriptsAPIService
	apiVersion string
	page *int32
	pageSize *int32
}

// The page of a collection to return.
func (r ApiGetStackScriptsRequest) Page(page int32) ApiGetStackScriptsRequest {
	r.page = &page
	return r
}

// The number of items to return per page.
func (r ApiGetStackScriptsRequest) PageSize(pageSize int32) ApiGetStackScriptsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetStackScriptsRequest) Execute() (*GetStackScripts200Response, *http.Response, error) {
	return r.ApiService.GetStackScriptsExecute(r)
}

/*
GetStackScripts List StackScripts

If the request is not authenticated, only public StackScripts are returned.

For more information on StackScripts, please read our [StackScripts documentation](https://www.linode.com/docs/products/tools/stackscripts/).


<<LB>>

---


- __CLI__.

    ```
    linode-cli stackscripts list
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    stackscripts:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiGetStackScriptsRequest
*/
func (a *LinodeStackScriptsAPIService) GetStackScripts(ctx context.Context, apiVersion string) ApiGetStackScriptsRequest {
	return ApiGetStackScriptsRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetStackScripts200Response
func (a *LinodeStackScriptsAPIService) GetStackScriptsExecute(r ApiGetStackScriptsRequest) (*GetStackScripts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStackScripts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeStackScriptsAPIService.GetStackScripts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/linode/stackscripts"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostAddStackScriptRequest struct {
	ctx context.Context
	ApiService *LinodeStackScriptsAPIService
	apiVersion string
	postAddStackScriptRequest *PostAddStackScriptRequest
}

// The properties to set for the new StackScript.
func (r ApiPostAddStackScriptRequest) PostAddStackScriptRequest(postAddStackScriptRequest PostAddStackScriptRequest) ApiPostAddStackScriptRequest {
	r.postAddStackScriptRequest = &postAddStackScriptRequest
	return r
}

func (r ApiPostAddStackScriptRequest) Execute() (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PostAddStackScriptExecute(r)
}

/*
PostAddStackScript Create a StackScript

Creates a StackScript in your Account.


<<LB>>

---


- __CLI__.

    ```
    linode-cli stackscripts create \
  --label a-stackscript \
  --description "This StackScript install and configures MySQL" \
  --images "linode/debian9" \
  --images "linode/debian8" \
  --is_public true \
  --rev_note "Set up MySQL" \
  --script '#!/bin/bash'
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    stackscripts:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiPostAddStackScriptRequest
*/
func (a *LinodeStackScriptsAPIService) PostAddStackScript(ctx context.Context, apiVersion string) ApiPostAddStackScriptRequest {
	return ApiPostAddStackScriptRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetStackScripts200ResponseDataInner
func (a *LinodeStackScriptsAPIService) PostAddStackScriptExecute(r ApiPostAddStackScriptRequest) (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStackScripts200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeStackScriptsAPIService.PostAddStackScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/linode/stackscripts"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postAddStackScriptRequest == nil {
		return localVarReturnValue, nil, reportError("postAddStackScriptRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postAddStackScriptRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutStackScriptRequest struct {
	ctx context.Context
	ApiService *LinodeStackScriptsAPIService
	apiVersion string
	stackscriptId string
	getStackScripts200ResponseDataInner *GetStackScripts200ResponseDataInner
}

// The fields to update.
func (r ApiPutStackScriptRequest) GetStackScripts200ResponseDataInner(getStackScripts200ResponseDataInner GetStackScripts200ResponseDataInner) ApiPutStackScriptRequest {
	r.getStackScripts200ResponseDataInner = &getStackScripts200ResponseDataInner
	return r
}

func (r ApiPutStackScriptRequest) Execute() (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PutStackScriptExecute(r)
}

/*
PutStackScript Update a StackScript

Updates a StackScript.

__Once a StackScript is made public, it cannot be made private.__


<<LB>>

---


- __CLI__.

    ```
    linode-cli stackscripts update 10079 \
  --label a-stackscript \
  --description "This StackScript installs \
    and configures MySQL" \
  --images "linode/debian9" \
  --images "linode/debian8" \
  --is_public true \
  --rev_note "Set up MySQL" \
  --script '#!/bin/bash'
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    stackscripts:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param stackscriptId The ID of the StackScript to look up.
 @return ApiPutStackScriptRequest
*/
func (a *LinodeStackScriptsAPIService) PutStackScript(ctx context.Context, apiVersion string, stackscriptId string) ApiPutStackScriptRequest {
	return ApiPutStackScriptRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		stackscriptId: stackscriptId,
	}
}

// Execute executes the request
//  @return GetStackScripts200ResponseDataInner
func (a *LinodeStackScriptsAPIService) PutStackScriptExecute(r ApiPutStackScriptRequest) (*GetStackScripts200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStackScripts200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinodeStackScriptsAPIService.PutStackScript")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/linode/stackscripts/{stackscriptId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stackscriptId"+"}", url.PathEscape(parameterValueToString(r.stackscriptId, "stackscriptId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.getStackScripts200ResponseDataInner
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
