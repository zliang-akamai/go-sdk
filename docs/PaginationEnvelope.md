# PaginationEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewPaginationEnvelope

`func NewPaginationEnvelope() *PaginationEnvelope`

NewPaginationEnvelope instantiates a new PaginationEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationEnvelopeWithDefaults

`func NewPaginationEnvelopeWithDefaults() *PaginationEnvelope`

NewPaginationEnvelopeWithDefaults instantiates a new PaginationEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaginationEnvelope) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginationEnvelope) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginationEnvelope) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginationEnvelope) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *PaginationEnvelope) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PaginationEnvelope) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PaginationEnvelope) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PaginationEnvelope) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *PaginationEnvelope) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginationEnvelope) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginationEnvelope) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginationEnvelope) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


