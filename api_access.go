/*
ProxMox VE API

ProxMox VE API

API version: 8.3
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiflat

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type AccessApi interface {

	/*
	CreateAccessTicket createAccessTicket

	Create or verify authentication ticket.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAccessTicketRequest
	*/
	CreateAccessTicket(ctx context.Context) ApiCreateAccessTicketRequest

	// CreateAccessTicketExecute executes the request
	//  @return CreateAccessTicket200Response
	CreateAccessTicketExecute(r ApiCreateAccessTicketRequest) (*CreateAccessTicket200Response, *http.Response, error)

	/*
	GetAccessTicket getAccessTicket

	Dummy. Useful for formatters which want to provide a login page.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccessTicketRequest
	*/
	GetAccessTicket(ctx context.Context) ApiGetAccessTicketRequest

	// GetAccessTicketExecute executes the request
	//  @return CreateVM200Response
	GetAccessTicketExecute(r ApiGetAccessTicketRequest) (*CreateVM200Response, *http.Response, error)
}

// AccessApiService AccessApi service
type AccessApiService service

type ApiCreateAccessTicketRequest struct {
	ctx context.Context
	ApiService AccessApi
	createAccessTicketRequest *CreateAccessTicketRequest
}

func (r ApiCreateAccessTicketRequest) CreateAccessTicketRequest(createAccessTicketRequest CreateAccessTicketRequest) ApiCreateAccessTicketRequest {
	r.createAccessTicketRequest = &createAccessTicketRequest
	return r
}

func (r ApiCreateAccessTicketRequest) Execute() (*CreateAccessTicket200Response, *http.Response, error) {
	return r.ApiService.CreateAccessTicketExecute(r)
}

/*
CreateAccessTicket createAccessTicket

Create or verify authentication ticket.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateAccessTicketRequest
*/
func (a *AccessApiService) CreateAccessTicket(ctx context.Context) ApiCreateAccessTicketRequest {
	return ApiCreateAccessTicketRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateAccessTicket200Response
func (a *AccessApiService) CreateAccessTicketExecute(r ApiCreateAccessTicketRequest) (*CreateAccessTicket200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateAccessTicket200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessApiService.CreateAccessTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/access/ticket"

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
	localVarPostBody = r.createAccessTicketRequest
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

type ApiGetAccessTicketRequest struct {
	ctx context.Context
	ApiService AccessApi
}

func (r ApiGetAccessTicketRequest) Execute() (*CreateVM200Response, *http.Response, error) {
	return r.ApiService.GetAccessTicketExecute(r)
}

/*
GetAccessTicket getAccessTicket

Dummy. Useful for formatters which want to provide a login page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAccessTicketRequest
*/
func (a *AccessApiService) GetAccessTicket(ctx context.Context) ApiGetAccessTicketRequest {
	return ApiGetAccessTicketRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateVM200Response
func (a *AccessApiService) GetAccessTicketExecute(r ApiGetAccessTicketRequest) (*CreateVM200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateVM200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessApiService.GetAccessTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/access/ticket"

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
