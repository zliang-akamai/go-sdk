# GetPayments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetPayments200ResponseDataInner**](GetPayments200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewGetPayments200Response

`func NewGetPayments200Response() *GetPayments200Response`

NewGetPayments200Response instantiates a new GetPayments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayments200ResponseWithDefaults

`func NewGetPayments200ResponseWithDefaults() *GetPayments200Response`

NewGetPayments200ResponseWithDefaults instantiates a new GetPayments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPayments200Response) GetData() []GetPayments200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPayments200Response) GetDataOk() (*[]GetPayments200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPayments200Response) SetData(v []GetPayments200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetPayments200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *GetPayments200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPayments200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPayments200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPayments200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GetPayments200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetPayments200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetPayments200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetPayments200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *GetPayments200Response) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetPayments200Response) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetPayments200Response) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetPayments200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


