# GetLinodeDisks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetLinodeDisks200ResponseDataInner**](GetLinodeDisks200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewGetLinodeDisks200Response

`func NewGetLinodeDisks200Response() *GetLinodeDisks200Response`

NewGetLinodeDisks200Response instantiates a new GetLinodeDisks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeDisks200ResponseWithDefaults

`func NewGetLinodeDisks200ResponseWithDefaults() *GetLinodeDisks200Response`

NewGetLinodeDisks200ResponseWithDefaults instantiates a new GetLinodeDisks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetLinodeDisks200Response) GetData() []GetLinodeDisks200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLinodeDisks200Response) GetDataOk() (*[]GetLinodeDisks200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLinodeDisks200Response) SetData(v []GetLinodeDisks200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetLinodeDisks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *GetLinodeDisks200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetLinodeDisks200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetLinodeDisks200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetLinodeDisks200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GetLinodeDisks200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetLinodeDisks200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetLinodeDisks200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetLinodeDisks200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *GetLinodeDisks200Response) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetLinodeDisks200Response) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetLinodeDisks200Response) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetLinodeDisks200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


