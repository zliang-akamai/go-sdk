# PostRebuildNodeBalancerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | What algorithm this NodeBalancer should use for routing traffic to backends. | [optional] [default to "roundrobin"]
**Check** | Pointer to **string** | The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down.  - If &#x60;none&#x60; no check is performed. - &#x60;connection&#x60; requires only a connection to the backend to succeed. - &#x60;http&#x60; and &#x60;http_body&#x60; rely on the backend serving HTTP, and that the response returned matches what is expected. | [optional] [default to "none"]
**CheckAttempts** | Pointer to **int32** | How many times to attempt a check before considering a backend to be down. | [optional] [default to 3]
**CheckBody** | Pointer to **string** | This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down. | [optional] 
**CheckInterval** | Pointer to **int32** | How often, in seconds, to check that backends are up and serving requests.  Must be greater than &#x60;check_timeout&#x60;. | [optional] [default to 31]
**CheckPassive** | Pointer to **bool** | If true, any response from this backend with a &#x60;5xx&#x60; status code will be enough for it to be considered unhealthy and taken out of rotation. | [optional] [default to true]
**CheckPath** | Pointer to **string** | The URL path to check on each backend. If the backend does not respond to this request it is considered to be down. | [optional] 
**CheckTimeout** | Pointer to **int32** | How long, in seconds, to wait for a check attempt before considering it failed.  Must be less than &#x60;check_interval&#x60;. | [optional] [default to 30]
**CipherSuite** | Pointer to **string** | What ciphers to use for SSL connections served by this NodeBalancer.  - &#x60;legacy&#x60; is considered insecure and should only be used if necessary. | [optional] [default to "recommended"]
**Id** | Pointer to **int32** | This config&#39;s unique ID. | [optional] [readonly] 
**NodebalancerId** | Pointer to **int32** | The ID for the NodeBalancer this config belongs to. | [optional] [readonly] 
**NodesStatus** | Pointer to [**GetNodeBalancerConfigs200ResponseDataInnerNodesStatus**](GetNodeBalancerConfigs200ResponseDataInnerNodesStatus.md) |  | [optional] 
**Port** | Pointer to **int32** | The port this Config is for. These values must be unique across configs on a single NodeBalancer (you can&#39;t have two configs for port 80, for example).  While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. | [optional] [default to 80]
**Protocol** | Pointer to **string** | The protocol this port is configured to serve.  - The &#x60;http&#x60; and &#x60;tcp&#x60; protocols do not support &#x60;ssl_cert&#x60; and &#x60;ssl_key&#x60;.  - The &#x60;https&#x60; protocol is mutually required with &#x60;ssl_cert&#x60; and &#x60;ssl_key&#x60;.  Review our guide on [Available Protocols](https://www.linode.com/docs/products/networking/nodebalancers/guides/protocols/) for information on protocol features. | [optional] [default to "http"]
**ProxyProtocol** | Pointer to **string** | ProxyProtocol is a TCP extension that sends initial TCP connection information such as source/destination IPs and ports to backend devices. This information would be lost otherwise. Backend devices must be configured to work with ProxyProtocol if enabled.  - If omitted, or set to &#x60;none&#x60;, the NodeBalancer doesn&#39;t send any auxiliary data over TCP connections. This is the default. - If set to &#x60;v1&#x60;, the human-readable header format (Version 1) is used. Requires &#x60;tcp&#x60; protocol. - If set to &#x60;v2&#x60;, the binary header format (Version 2) is used. Requires &#x60;tcp&#x60; protocol. | [optional] [default to "none"]
**SslCert** | Pointer to **NullableString** |  The PEM-formatted public SSL certificate (or the combined PEM-formatted SSL certificate and Certificate Authority chain) that should be served on this NodeBalancerConfig&#39;s port.  Line breaks must be represented as &#x60;\\n&#x60; in the string for requests (but not when using the Linode CLI).  [Diffie-Hellman Parameters](https://www.linode.com/docs/products/networking/nodebalancers/guides/ssl-termination/#diffie-hellman-parameters) can be included in this value to enable forward secrecy.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, &#x60;&lt;REDACTED&gt;&#x60; will be printed where the field appears.  The read-only &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig. | [optional] 
**SslCommonname** | Pointer to **string** | The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig. | [optional] [readonly] 
**SslFingerprint** | Pointer to **string** | The read-only SHA1-encoded fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig. | [optional] [readonly] 
**SslKey** | Pointer to **NullableString** | The PEM-formatted private key for the SSL certificate set in the &#x60;ssl_cert&#x60; field.  Line breaks must be represented as &#x60;\\n&#x60; in the string for requests (but not when using the Linode CLI).  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, &#x60;&lt;REDACTED&gt;&#x60; will be printed where the field appears.  The read-only &#x60;ssl_commonname&#x60; and &#x60;ssl_fingerprint&#x60; fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig. | [optional] 
**Stickiness** | Pointer to **string** | Controls how session stickiness is handled on this port.  - If set to &#x60;none&#x60; connections will always be assigned a backend based on the algorithm configured. - If set to &#x60;table&#x60; sessions from the same remote address will be routed to the same backend. - For HTTP or HTTPS clients, &#x60;http_cookie&#x60; allows sessions to be routed to the same backend based on a cookie set by the NodeBalancer. | [optional] [default to "none"]
**Nodes** | [**[]PostRebuildNodeBalancerConfigRequestAllOfNodesInner**](PostRebuildNodeBalancerConfigRequestAllOfNodesInner.md) | The NodeBalancer Nodes that serve this Config.  Some considerations for Nodes when rebuilding a config:    - Current Nodes excluded from the request body will be deleted from the Config.   - Current Nodes (identified by their Node ID) will be updated.   - New Nodes (included without a Node ID) will be created. | 

## Methods

### NewPostRebuildNodeBalancerConfigRequest

`func NewPostRebuildNodeBalancerConfigRequest(nodes []PostRebuildNodeBalancerConfigRequestAllOfNodesInner, ) *PostRebuildNodeBalancerConfigRequest`

NewPostRebuildNodeBalancerConfigRequest instantiates a new PostRebuildNodeBalancerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRebuildNodeBalancerConfigRequestWithDefaults

`func NewPostRebuildNodeBalancerConfigRequestWithDefaults() *PostRebuildNodeBalancerConfigRequest`

NewPostRebuildNodeBalancerConfigRequestWithDefaults instantiates a new PostRebuildNodeBalancerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *PostRebuildNodeBalancerConfigRequest) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PostRebuildNodeBalancerConfigRequest) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PostRebuildNodeBalancerConfigRequest) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCheck

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckAttempts

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckAttempts() int32`

GetCheckAttempts returns the CheckAttempts field if non-nil, zero value otherwise.

### GetCheckAttemptsOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckAttemptsOk() (*int32, bool)`

GetCheckAttemptsOk returns a tuple with the CheckAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAttempts

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckAttempts(v int32)`

SetCheckAttempts sets CheckAttempts field to given value.

### HasCheckAttempts

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckAttempts() bool`

HasCheckAttempts returns a boolean if a field has been set.

### GetCheckBody

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckBody() string`

GetCheckBody returns the CheckBody field if non-nil, zero value otherwise.

### GetCheckBodyOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckBodyOk() (*string, bool)`

GetCheckBodyOk returns a tuple with the CheckBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckBody

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckBody(v string)`

SetCheckBody sets CheckBody field to given value.

### HasCheckBody

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckBody() bool`

HasCheckBody returns a boolean if a field has been set.

### GetCheckInterval

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetCheckPassive

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckPassive() bool`

GetCheckPassive returns the CheckPassive field if non-nil, zero value otherwise.

### GetCheckPassiveOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckPassiveOk() (*bool, bool)`

GetCheckPassiveOk returns a tuple with the CheckPassive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassive

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckPassive(v bool)`

SetCheckPassive sets CheckPassive field to given value.

### HasCheckPassive

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckPassive() bool`

HasCheckPassive returns a boolean if a field has been set.

### GetCheckPath

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckPath() string`

GetCheckPath returns the CheckPath field if non-nil, zero value otherwise.

### GetCheckPathOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckPathOk() (*string, bool)`

GetCheckPathOk returns a tuple with the CheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPath

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckPath(v string)`

SetCheckPath sets CheckPath field to given value.

### HasCheckPath

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckPath() bool`

HasCheckPath returns a boolean if a field has been set.

### GetCheckTimeout

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckTimeout() int32`

GetCheckTimeout returns the CheckTimeout field if non-nil, zero value otherwise.

### GetCheckTimeoutOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCheckTimeoutOk() (*int32, bool)`

GetCheckTimeoutOk returns a tuple with the CheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTimeout

`func (o *PostRebuildNodeBalancerConfigRequest) SetCheckTimeout(v int32)`

SetCheckTimeout sets CheckTimeout field to given value.

### HasCheckTimeout

`func (o *PostRebuildNodeBalancerConfigRequest) HasCheckTimeout() bool`

HasCheckTimeout returns a boolean if a field has been set.

### GetCipherSuite

`func (o *PostRebuildNodeBalancerConfigRequest) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *PostRebuildNodeBalancerConfigRequest) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *PostRebuildNodeBalancerConfigRequest) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetId

`func (o *PostRebuildNodeBalancerConfigRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostRebuildNodeBalancerConfigRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostRebuildNodeBalancerConfigRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *PostRebuildNodeBalancerConfigRequest) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *PostRebuildNodeBalancerConfigRequest) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.

### GetNodesStatus

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodesStatus() GetNodeBalancerConfigs200ResponseDataInnerNodesStatus`

GetNodesStatus returns the NodesStatus field if non-nil, zero value otherwise.

### GetNodesStatusOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodesStatusOk() (*GetNodeBalancerConfigs200ResponseDataInnerNodesStatus, bool)`

GetNodesStatusOk returns a tuple with the NodesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesStatus

`func (o *PostRebuildNodeBalancerConfigRequest) SetNodesStatus(v GetNodeBalancerConfigs200ResponseDataInnerNodesStatus)`

SetNodesStatus sets NodesStatus field to given value.

### HasNodesStatus

`func (o *PostRebuildNodeBalancerConfigRequest) HasNodesStatus() bool`

HasNodesStatus returns a boolean if a field has been set.

### GetPort

`func (o *PostRebuildNodeBalancerConfigRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostRebuildNodeBalancerConfigRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostRebuildNodeBalancerConfigRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *PostRebuildNodeBalancerConfigRequest) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetSslCert

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *PostRebuildNodeBalancerConfigRequest) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *PostRebuildNodeBalancerConfigRequest) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetSslCommonname

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslCommonname() string`

GetSslCommonname returns the SslCommonname field if non-nil, zero value otherwise.

### GetSslCommonnameOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslCommonnameOk() (*string, bool)`

GetSslCommonnameOk returns a tuple with the SslCommonname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCommonname

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslCommonname(v string)`

SetSslCommonname sets SslCommonname field to given value.

### HasSslCommonname

`func (o *PostRebuildNodeBalancerConfigRequest) HasSslCommonname() bool`

HasSslCommonname returns a boolean if a field has been set.

### GetSslFingerprint

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslFingerprint() string`

GetSslFingerprint returns the SslFingerprint field if non-nil, zero value otherwise.

### GetSslFingerprintOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslFingerprintOk() (*string, bool)`

GetSslFingerprintOk returns a tuple with the SslFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslFingerprint

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslFingerprint(v string)`

SetSslFingerprint sets SslFingerprint field to given value.

### HasSslFingerprint

`func (o *PostRebuildNodeBalancerConfigRequest) HasSslFingerprint() bool`

HasSslFingerprint returns a boolean if a field has been set.

### GetSslKey

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslKey() string`

GetSslKey returns the SslKey field if non-nil, zero value otherwise.

### GetSslKeyOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetSslKeyOk() (*string, bool)`

GetSslKeyOk returns a tuple with the SslKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslKey

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslKey(v string)`

SetSslKey sets SslKey field to given value.

### HasSslKey

`func (o *PostRebuildNodeBalancerConfigRequest) HasSslKey() bool`

HasSslKey returns a boolean if a field has been set.

### SetSslKeyNil

`func (o *PostRebuildNodeBalancerConfigRequest) SetSslKeyNil(b bool)`

 SetSslKeyNil sets the value for SslKey to be an explicit nil

### UnsetSslKey
`func (o *PostRebuildNodeBalancerConfigRequest) UnsetSslKey()`

UnsetSslKey ensures that no value is present for SslKey, not even an explicit nil
### GetStickiness

`func (o *PostRebuildNodeBalancerConfigRequest) GetStickiness() string`

GetStickiness returns the Stickiness field if non-nil, zero value otherwise.

### GetStickinessOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetStickinessOk() (*string, bool)`

GetStickinessOk returns a tuple with the Stickiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickiness

`func (o *PostRebuildNodeBalancerConfigRequest) SetStickiness(v string)`

SetStickiness sets Stickiness field to given value.

### HasStickiness

`func (o *PostRebuildNodeBalancerConfigRequest) HasStickiness() bool`

HasStickiness returns a boolean if a field has been set.

### GetNodes

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodes() []PostRebuildNodeBalancerConfigRequestAllOfNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *PostRebuildNodeBalancerConfigRequest) GetNodesOk() (*[]PostRebuildNodeBalancerConfigRequestAllOfNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *PostRebuildNodeBalancerConfigRequest) SetNodes(v []PostRebuildNodeBalancerConfigRequestAllOfNodesInner)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

