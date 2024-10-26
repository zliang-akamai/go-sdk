/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetNodeBalancerConfigs200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeBalancerConfigs200ResponseDataInner{}

// GetNodeBalancerConfigs200ResponseDataInner A NodeBalancer config represents the configuration of this NodeBalancer on a single port.  For example, a NodeBalancer Config on port 80 would typically represent how this NodeBalancer response to HTTP requests. NodeBalancer configs have a list of backends, called \"nodes,\" that they forward requests between based on their configuration.
type GetNodeBalancerConfigs200ResponseDataInner struct {
	// What algorithm this NodeBalancer should use for routing traffic to backends.
	Algorithm *string `json:"algorithm,omitempty"`
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down.  - If `none` no check is performed. - `connection` requires only a connection to the backend to succeed. - `http` and `http_body` rely on the backend serving HTTP, and that the response returned matches what is expected.
	Check *string `json:"check,omitempty"`
	// How many times to attempt a check before considering a backend to be down.
	CheckAttempts *int32 `json:"check_attempts,omitempty"`
	// This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down.
	CheckBody *string `json:"check_body,omitempty"`
	// How often, in seconds, to check that backends are up and serving requests.  Must be greater than `check_timeout`.
	CheckInterval *int32 `json:"check_interval,omitempty"`
	// If true, any response from this backend with a `5xx` status code will be enough for it to be considered unhealthy and taken out of rotation.
	CheckPassive *bool `json:"check_passive,omitempty"`
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	CheckPath *string `json:"check_path,omitempty" validate:"regexp=^[a-zA-Z0-9\\/\\\\-%?&=.]*$"`
	// How long, in seconds, to wait for a check attempt before considering it failed.  Must be less than `check_interval`.
	CheckTimeout *int32 `json:"check_timeout,omitempty"`
	// What ciphers to use for SSL connections served by this NodeBalancer.  - `legacy` is considered insecure and should only be used if necessary.
	CipherSuite *string `json:"cipher_suite,omitempty"`
	// This config's unique ID.
	Id *int32 `json:"id,omitempty"`
	// The ID for the NodeBalancer this config belongs to.
	NodebalancerId *int32 `json:"nodebalancer_id,omitempty"`
	NodesStatus *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus `json:"nodes_status,omitempty"`
	// The port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example).  While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.
	Port *int32 `json:"port,omitempty"`
	// The protocol this port is configured to serve.  - The `http` and `tcp` protocols do not support `ssl_cert` and `ssl_key`.  - The `https` protocol is mutually required with `ssl_cert` and `ssl_key`.  Review our guide on [Available Protocols](https://www.linode.com/docs/products/networking/nodebalancers/guides/protocols/) for information on protocol features.
	Protocol *string `json:"protocol,omitempty"`
	// ProxyProtocol is a TCP extension that sends initial TCP connection information such as source/destination IPs and ports to backend devices. This information would be lost otherwise. Backend devices must be configured to work with ProxyProtocol if enabled.  - If omitted, or set to `none`, the NodeBalancer doesn't send any auxiliary data over TCP connections. This is the default. - If set to `v1`, the human-readable header format (Version 1) is used. Requires `tcp` protocol. - If set to `v2`, the binary header format (Version 2) is used. Requires `tcp` protocol.
	ProxyProtocol *string `json:"proxy_protocol,omitempty"`
	//  The PEM-formatted public SSL certificate (or the combined PEM-formatted SSL certificate and Certificate Authority chain) that should be served on this NodeBalancerConfig's port.  Line breaks must be represented as `\\n` in the string for requests (but not when using the Linode CLI).  [Diffie-Hellman Parameters](https://www.linode.com/docs/products/networking/nodebalancers/guides/ssl-termination/#diffie-hellman-parameters) can be included in this value to enable forward secrecy.  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, `<REDACTED>` will be printed where the field appears.  The read-only `ssl_commonname` and `ssl_fingerprint` fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig.
	SslCert NullableString `json:"ssl_cert,omitempty"`
	// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SslCommonname *string `json:"ssl_commonname,omitempty"`
	// The read-only SHA1-encoded fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SslFingerprint *string `json:"ssl_fingerprint,omitempty"`
	// The PEM-formatted private key for the SSL certificate set in the `ssl_cert` field.  Line breaks must be represented as `\\n` in the string for requests (but not when using the Linode CLI).  The contents of this field will not be shown in any responses that display the NodeBalancerConfig. Instead, `<REDACTED>` will be printed where the field appears.  The read-only `ssl_commonname` and `ssl_fingerprint` fields in a NodeBalancerConfig response are automatically derived from your certificate. Please refer to these fields to verify that the appropriate certificate was assigned to your NodeBalancerConfig.
	SslKey NullableString `json:"ssl_key,omitempty"`
	// Controls how session stickiness is handled on this port.  - If set to `none` connections will always be assigned a backend based on the algorithm configured. - If set to `table` sessions from the same remote address will be routed to the same backend. - For HTTP or HTTPS clients, `http_cookie` allows sessions to be routed to the same backend based on a cookie set by the NodeBalancer.
	Stickiness *string `json:"stickiness,omitempty"`
}

// NewGetNodeBalancerConfigs200ResponseDataInner instantiates a new GetNodeBalancerConfigs200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeBalancerConfigs200ResponseDataInner() *GetNodeBalancerConfigs200ResponseDataInner {
	this := GetNodeBalancerConfigs200ResponseDataInner{}
	var algorithm string = "roundrobin"
	this.Algorithm = &algorithm
	var check string = "none"
	this.Check = &check
	var checkAttempts int32 = 3
	this.CheckAttempts = &checkAttempts
	var checkInterval int32 = 31
	this.CheckInterval = &checkInterval
	var checkPassive bool = true
	this.CheckPassive = &checkPassive
	var checkTimeout int32 = 30
	this.CheckTimeout = &checkTimeout
	var cipherSuite string = "recommended"
	this.CipherSuite = &cipherSuite
	var port int32 = 80
	this.Port = &port
	var protocol string = "http"
	this.Protocol = &protocol
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol
	var stickiness string = "none"
	this.Stickiness = &stickiness
	return &this
}

// NewGetNodeBalancerConfigs200ResponseDataInnerWithDefaults instantiates a new GetNodeBalancerConfigs200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeBalancerConfigs200ResponseDataInnerWithDefaults() *GetNodeBalancerConfigs200ResponseDataInner {
	this := GetNodeBalancerConfigs200ResponseDataInner{}
	var algorithm string = "roundrobin"
	this.Algorithm = &algorithm
	var check string = "none"
	this.Check = &check
	var checkAttempts int32 = 3
	this.CheckAttempts = &checkAttempts
	var checkInterval int32 = 31
	this.CheckInterval = &checkInterval
	var checkPassive bool = true
	this.CheckPassive = &checkPassive
	var checkTimeout int32 = 30
	this.CheckTimeout = &checkTimeout
	var cipherSuite string = "recommended"
	this.CipherSuite = &cipherSuite
	var port int32 = 80
	this.Port = &port
	var protocol string = "http"
	this.Protocol = &protocol
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol
	var stickiness string = "none"
	this.Stickiness = &stickiness
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheck() string {
	if o == nil || IsNil(o.Check) {
		var ret string
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckOk() (*string, bool) {
	if o == nil || IsNil(o.Check) {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheck() bool {
	if o != nil && !IsNil(o.Check) {
		return true
	}

	return false
}

// SetCheck gets a reference to the given string and assigns it to the Check field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheck(v string) {
	o.Check = &v
}

// GetCheckAttempts returns the CheckAttempts field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckAttempts() int32 {
	if o == nil || IsNil(o.CheckAttempts) {
		var ret int32
		return ret
	}
	return *o.CheckAttempts
}

// GetCheckAttemptsOk returns a tuple with the CheckAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckAttempts) {
		return nil, false
	}
	return o.CheckAttempts, true
}

// HasCheckAttempts returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckAttempts() bool {
	if o != nil && !IsNil(o.CheckAttempts) {
		return true
	}

	return false
}

// SetCheckAttempts gets a reference to the given int32 and assigns it to the CheckAttempts field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckAttempts(v int32) {
	o.CheckAttempts = &v
}

// GetCheckBody returns the CheckBody field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckBody() string {
	if o == nil || IsNil(o.CheckBody) {
		var ret string
		return ret
	}
	return *o.CheckBody
}

// GetCheckBodyOk returns a tuple with the CheckBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckBodyOk() (*string, bool) {
	if o == nil || IsNil(o.CheckBody) {
		return nil, false
	}
	return o.CheckBody, true
}

// HasCheckBody returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckBody() bool {
	if o != nil && !IsNil(o.CheckBody) {
		return true
	}

	return false
}

// SetCheckBody gets a reference to the given string and assigns it to the CheckBody field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckBody(v string) {
	o.CheckBody = &v
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckInterval() int32 {
	if o == nil || IsNil(o.CheckInterval) {
		var ret int32
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckInterval) {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckInterval() bool {
	if o != nil && !IsNil(o.CheckInterval) {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given int32 and assigns it to the CheckInterval field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckInterval(v int32) {
	o.CheckInterval = &v
}

// GetCheckPassive returns the CheckPassive field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckPassive() bool {
	if o == nil || IsNil(o.CheckPassive) {
		var ret bool
		return ret
	}
	return *o.CheckPassive
}

// GetCheckPassiveOk returns a tuple with the CheckPassive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckPassiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckPassive) {
		return nil, false
	}
	return o.CheckPassive, true
}

// HasCheckPassive returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckPassive() bool {
	if o != nil && !IsNil(o.CheckPassive) {
		return true
	}

	return false
}

// SetCheckPassive gets a reference to the given bool and assigns it to the CheckPassive field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckPassive(v bool) {
	o.CheckPassive = &v
}

// GetCheckPath returns the CheckPath field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckPath() string {
	if o == nil || IsNil(o.CheckPath) {
		var ret string
		return ret
	}
	return *o.CheckPath
}

// GetCheckPathOk returns a tuple with the CheckPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckPathOk() (*string, bool) {
	if o == nil || IsNil(o.CheckPath) {
		return nil, false
	}
	return o.CheckPath, true
}

// HasCheckPath returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckPath() bool {
	if o != nil && !IsNil(o.CheckPath) {
		return true
	}

	return false
}

// SetCheckPath gets a reference to the given string and assigns it to the CheckPath field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckPath(v string) {
	o.CheckPath = &v
}

// GetCheckTimeout returns the CheckTimeout field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckTimeout() int32 {
	if o == nil || IsNil(o.CheckTimeout) {
		var ret int32
		return ret
	}
	return *o.CheckTimeout
}

// GetCheckTimeoutOk returns a tuple with the CheckTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCheckTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckTimeout) {
		return nil, false
	}
	return o.CheckTimeout, true
}

// HasCheckTimeout returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCheckTimeout() bool {
	if o != nil && !IsNil(o.CheckTimeout) {
		return true
	}

	return false
}

// SetCheckTimeout gets a reference to the given int32 and assigns it to the CheckTimeout field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCheckTimeout(v int32) {
	o.CheckTimeout = &v
}

// GetCipherSuite returns the CipherSuite field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCipherSuite() string {
	if o == nil || IsNil(o.CipherSuite) {
		var ret string
		return ret
	}
	return *o.CipherSuite
}

// GetCipherSuiteOk returns a tuple with the CipherSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetCipherSuiteOk() (*string, bool) {
	if o == nil || IsNil(o.CipherSuite) {
		return nil, false
	}
	return o.CipherSuite, true
}

// HasCipherSuite returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasCipherSuite() bool {
	if o != nil && !IsNil(o.CipherSuite) {
		return true
	}

	return false
}

// SetCipherSuite gets a reference to the given string and assigns it to the CipherSuite field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetCipherSuite(v string) {
	o.CipherSuite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetNodebalancerId returns the NodebalancerId field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetNodebalancerId() int32 {
	if o == nil || IsNil(o.NodebalancerId) {
		var ret int32
		return ret
	}
	return *o.NodebalancerId
}

// GetNodebalancerIdOk returns a tuple with the NodebalancerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetNodebalancerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NodebalancerId) {
		return nil, false
	}
	return o.NodebalancerId, true
}

// HasNodebalancerId returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasNodebalancerId() bool {
	if o != nil && !IsNil(o.NodebalancerId) {
		return true
	}

	return false
}

// SetNodebalancerId gets a reference to the given int32 and assigns it to the NodebalancerId field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetNodebalancerId(v int32) {
	o.NodebalancerId = &v
}

// GetNodesStatus returns the NodesStatus field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetNodesStatus() GetNodeBalancerConfigs200ResponseDataInnerNodesStatus {
	if o == nil || IsNil(o.NodesStatus) {
		var ret GetNodeBalancerConfigs200ResponseDataInnerNodesStatus
		return ret
	}
	return *o.NodesStatus
}

// GetNodesStatusOk returns a tuple with the NodesStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetNodesStatusOk() (*GetNodeBalancerConfigs200ResponseDataInnerNodesStatus, bool) {
	if o == nil || IsNil(o.NodesStatus) {
		return nil, false
	}
	return o.NodesStatus, true
}

// HasNodesStatus returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasNodesStatus() bool {
	if o != nil && !IsNil(o.NodesStatus) {
		return true
	}

	return false
}

// SetNodesStatus gets a reference to the given GetNodeBalancerConfigs200ResponseDataInnerNodesStatus and assigns it to the NodesStatus field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetNodesStatus(v GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) {
	o.NodesStatus = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetProxyProtocol() string {
	if o == nil || IsNil(o.ProxyProtocol) {
		var ret string
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetProxyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyProtocol) {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// HasProxyProtocol returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasProxyProtocol() bool {
	if o != nil && !IsNil(o.ProxyProtocol) {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given string and assigns it to the ProxyProtocol field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetProxyProtocol(v string) {
	o.ProxyProtocol = &v
}

// GetSslCert returns the SslCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslCert() string {
	if o == nil || IsNil(o.SslCert.Get()) {
		var ret string
		return ret
	}
	return *o.SslCert.Get()
}

// GetSslCertOk returns a tuple with the SslCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SslCert.Get(), o.SslCert.IsSet()
}

// HasSslCert returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasSslCert() bool {
	if o != nil && o.SslCert.IsSet() {
		return true
	}

	return false
}

// SetSslCert gets a reference to the given NullableString and assigns it to the SslCert field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslCert(v string) {
	o.SslCert.Set(&v)
}
// SetSslCertNil sets the value for SslCert to be an explicit nil
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslCertNil() {
	o.SslCert.Set(nil)
}

// UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
func (o *GetNodeBalancerConfigs200ResponseDataInner) UnsetSslCert() {
	o.SslCert.Unset()
}

// GetSslCommonname returns the SslCommonname field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslCommonname() string {
	if o == nil || IsNil(o.SslCommonname) {
		var ret string
		return ret
	}
	return *o.SslCommonname
}

// GetSslCommonnameOk returns a tuple with the SslCommonname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslCommonnameOk() (*string, bool) {
	if o == nil || IsNil(o.SslCommonname) {
		return nil, false
	}
	return o.SslCommonname, true
}

// HasSslCommonname returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasSslCommonname() bool {
	if o != nil && !IsNil(o.SslCommonname) {
		return true
	}

	return false
}

// SetSslCommonname gets a reference to the given string and assigns it to the SslCommonname field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslCommonname(v string) {
	o.SslCommonname = &v
}

// GetSslFingerprint returns the SslFingerprint field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslFingerprint() string {
	if o == nil || IsNil(o.SslFingerprint) {
		var ret string
		return ret
	}
	return *o.SslFingerprint
}

// GetSslFingerprintOk returns a tuple with the SslFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.SslFingerprint) {
		return nil, false
	}
	return o.SslFingerprint, true
}

// HasSslFingerprint returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasSslFingerprint() bool {
	if o != nil && !IsNil(o.SslFingerprint) {
		return true
	}

	return false
}

// SetSslFingerprint gets a reference to the given string and assigns it to the SslFingerprint field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslFingerprint(v string) {
	o.SslFingerprint = &v
}

// GetSslKey returns the SslKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslKey() string {
	if o == nil || IsNil(o.SslKey.Get()) {
		var ret string
		return ret
	}
	return *o.SslKey.Get()
}

// GetSslKeyOk returns a tuple with the SslKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetSslKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SslKey.Get(), o.SslKey.IsSet()
}

// HasSslKey returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasSslKey() bool {
	if o != nil && o.SslKey.IsSet() {
		return true
	}

	return false
}

// SetSslKey gets a reference to the given NullableString and assigns it to the SslKey field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslKey(v string) {
	o.SslKey.Set(&v)
}
// SetSslKeyNil sets the value for SslKey to be an explicit nil
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetSslKeyNil() {
	o.SslKey.Set(nil)
}

// UnsetSslKey ensures that no value is present for SslKey, not even an explicit nil
func (o *GetNodeBalancerConfigs200ResponseDataInner) UnsetSslKey() {
	o.SslKey.Unset()
}

// GetStickiness returns the Stickiness field value if set, zero value otherwise.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetStickiness() string {
	if o == nil || IsNil(o.Stickiness) {
		var ret string
		return ret
	}
	return *o.Stickiness
}

// GetStickinessOk returns a tuple with the Stickiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) GetStickinessOk() (*string, bool) {
	if o == nil || IsNil(o.Stickiness) {
		return nil, false
	}
	return o.Stickiness, true
}

// HasStickiness returns a boolean if a field has been set.
func (o *GetNodeBalancerConfigs200ResponseDataInner) HasStickiness() bool {
	if o != nil && !IsNil(o.Stickiness) {
		return true
	}

	return false
}

// SetStickiness gets a reference to the given string and assigns it to the Stickiness field.
func (o *GetNodeBalancerConfigs200ResponseDataInner) SetStickiness(v string) {
	o.Stickiness = &v
}

func (o GetNodeBalancerConfigs200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeBalancerConfigs200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Check) {
		toSerialize["check"] = o.Check
	}
	if !IsNil(o.CheckAttempts) {
		toSerialize["check_attempts"] = o.CheckAttempts
	}
	if !IsNil(o.CheckBody) {
		toSerialize["check_body"] = o.CheckBody
	}
	if !IsNil(o.CheckInterval) {
		toSerialize["check_interval"] = o.CheckInterval
	}
	if !IsNil(o.CheckPassive) {
		toSerialize["check_passive"] = o.CheckPassive
	}
	if !IsNil(o.CheckPath) {
		toSerialize["check_path"] = o.CheckPath
	}
	if !IsNil(o.CheckTimeout) {
		toSerialize["check_timeout"] = o.CheckTimeout
	}
	if !IsNil(o.CipherSuite) {
		toSerialize["cipher_suite"] = o.CipherSuite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NodebalancerId) {
		toSerialize["nodebalancer_id"] = o.NodebalancerId
	}
	if !IsNil(o.NodesStatus) {
		toSerialize["nodes_status"] = o.NodesStatus
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ProxyProtocol) {
		toSerialize["proxy_protocol"] = o.ProxyProtocol
	}
	if o.SslCert.IsSet() {
		toSerialize["ssl_cert"] = o.SslCert.Get()
	}
	if !IsNil(o.SslCommonname) {
		toSerialize["ssl_commonname"] = o.SslCommonname
	}
	if !IsNil(o.SslFingerprint) {
		toSerialize["ssl_fingerprint"] = o.SslFingerprint
	}
	if o.SslKey.IsSet() {
		toSerialize["ssl_key"] = o.SslKey.Get()
	}
	if !IsNil(o.Stickiness) {
		toSerialize["stickiness"] = o.Stickiness
	}
	return toSerialize, nil
}

type NullableGetNodeBalancerConfigs200ResponseDataInner struct {
	value *GetNodeBalancerConfigs200ResponseDataInner
	isSet bool
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInner) Get() *GetNodeBalancerConfigs200ResponseDataInner {
	return v.value
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInner) Set(val *GetNodeBalancerConfigs200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeBalancerConfigs200ResponseDataInner(val *GetNodeBalancerConfigs200ResponseDataInner) *NullableGetNodeBalancerConfigs200ResponseDataInner {
	return &NullableGetNodeBalancerConfigs200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetNodeBalancerConfigs200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeBalancerConfigs200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


