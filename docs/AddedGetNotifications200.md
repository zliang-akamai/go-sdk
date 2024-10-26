# AddedGetNotifications200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetNotifications200ResponseDataInner**](GetNotifications200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewAddedGetNotifications200

`func NewAddedGetNotifications200() *AddedGetNotifications200`

NewAddedGetNotifications200 instantiates a new AddedGetNotifications200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedGetNotifications200WithDefaults

`func NewAddedGetNotifications200WithDefaults() *AddedGetNotifications200`

NewAddedGetNotifications200WithDefaults instantiates a new AddedGetNotifications200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddedGetNotifications200) GetData() []GetNotifications200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddedGetNotifications200) GetDataOk() (*[]GetNotifications200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddedGetNotifications200) SetData(v []GetNotifications200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AddedGetNotifications200) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *AddedGetNotifications200) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AddedGetNotifications200) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AddedGetNotifications200) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AddedGetNotifications200) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *AddedGetNotifications200) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AddedGetNotifications200) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AddedGetNotifications200) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AddedGetNotifications200) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *AddedGetNotifications200) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AddedGetNotifications200) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AddedGetNotifications200) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *AddedGetNotifications200) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


