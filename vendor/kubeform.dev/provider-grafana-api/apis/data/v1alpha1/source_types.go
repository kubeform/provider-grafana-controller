/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Source struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourceSpec   `json:"spec,omitempty"`
	Status            SourceStatus `json:"status,omitempty"`
}

type SourceSpecJsonDataDerivedField struct {
	// +optional
	DatasourceUid *string `json:"datasourceUid,omitempty" tf:"datasource_uid"`
	// +optional
	MatcherRegex *string `json:"matcherRegex,omitempty" tf:"matcher_regex"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type SourceSpecJsonData struct {
	// (CloudWatch, Athena) The ARN of the role to be assumed by Grafana when using the CloudWatch or Athena data source.
	// +optional
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty" tf:"assume_role_arn"`
	// (CloudWatch, Athena) The authentication type used to access the data source.
	// +optional
	AuthType *string `json:"authType,omitempty" tf:"auth_type"`
	// (Stackdriver) The authentication type: `jwt` or `gce`.
	// +optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type"`
	// (Athena) Athena catalog.
	// +optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog"`
	// (Stackdriver) Service account email address.
	// +optional
	ClientEmail *string `json:"clientEmail,omitempty" tf:"client_email"`
	// (MySQL, PostgreSQL, and MSSQL) Maximum amount of time in seconds a connection may be reused (Grafana v5.4+).
	// +optional
	ConnMaxLifetime *int64 `json:"connMaxLifetime,omitempty" tf:"conn_max_lifetime"`
	// (CloudWatch) A comma-separated list of custom namespaces to be queried by the CloudWatch data source.
	// +optional
	CustomMetricsNamespaces *string `json:"customMetricsNamespaces,omitempty" tf:"custom_metrics_namespaces"`
	// (Athena) Name of the database within the catalog.
	// +optional
	Database *string `json:"database,omitempty" tf:"database"`
	// (InfluxDB) The default bucket for the data source.
	// +optional
	DefaultBucket *string `json:"defaultBucket,omitempty" tf:"default_bucket"`
	// (Stackdriver) The default project for the data source.
	// +optional
	DefaultProject *string `json:"defaultProject,omitempty" tf:"default_project"`
	// (CloudWatch, Athena) The default region for the data source.
	// +optional
	DefaultRegion *string `json:"defaultRegion,omitempty" tf:"default_region"`
	// (Loki) See https://grafana.com/docs/grafana/latest/datasources/loki/#derived-fields
	// +optional
	DerivedField []SourceSpecJsonDataDerivedField `json:"derivedField,omitempty" tf:"derived_field"`
	// (MSSQL) Connection SSL encryption handling: 'disable', 'false' or 'true'.
	// +optional
	Encrypt *string `json:"encrypt,omitempty" tf:"encrypt"`
	// (Elasticsearch) Elasticsearch semantic version (Grafana v8.0+).
	// +optional
	EsVersion *string `json:"esVersion,omitempty" tf:"es_version"`
	// (CloudWatch, Athena) If you are assuming a role in another account, that has been created with an external ID, specify the external ID here.
	// +optional
	ExternalID *string `json:"externalID,omitempty" tf:"external_id"`
	// (Github) Github URL
	// +optional
	GithubURL *string `json:"githubURL,omitempty" tf:"github_url"`
	// (Graphite) Graphite version.
	// +optional
	GraphiteVersion *string `json:"graphiteVersion,omitempty" tf:"graphite_version"`
	// (Prometheus) HTTP method to use for making requests.
	// +optional
	HttpMethod *string `json:"httpMethod,omitempty" tf:"http_method"`
	// (Elasticsearch) Index date time format. nil(No Pattern), 'Hourly', 'Daily', 'Weekly', 'Monthly' or 'Yearly'.
	// +optional
	Interval *string `json:"interval,omitempty" tf:"interval"`
	// (Elasticsearch) Which field should be used to indicate the priority of the log message.
	// +optional
	LogLevelField *string `json:"logLevelField,omitempty" tf:"log_level_field"`
	// (Elasticsearch) Which field should be used as the log message.
	// +optional
	LogMessageField *string `json:"logMessageField,omitempty" tf:"log_message_field"`
	// (Elasticsearch) Maximum number of concurrent shard requests.
	// +optional
	MaxConcurrentShardRequests *int64 `json:"maxConcurrentShardRequests,omitempty" tf:"max_concurrent_shard_requests"`
	// (MySQL, PostgreSQL and MSSQL) Maximum number of connections in the idle connection pool (Grafana v5.4+).
	// +optional
	MaxIdleConns *int64 `json:"maxIdleConns,omitempty" tf:"max_idle_conns"`
	// (Loki) Upper limit for the number of log lines returned by Loki
	// +optional
	MaxLines *int64 `json:"maxLines,omitempty" tf:"max_lines"`
	// (MySQL, PostgreSQL and MSSQL) Maximum number of open connections to the database (Grafana v5.4+).
	// +optional
	MaxOpenConns *int64 `json:"maxOpenConns,omitempty" tf:"max_open_conns"`
	// (Sentry) Organization slug.
	// +optional
	OrgSlug *string `json:"orgSlug,omitempty" tf:"org_slug"`
	// (InfluxDB) An organization is a workspace for a group of users. All dashboards, tasks, buckets, members, etc., belong to an organization.
	// +optional
	Organization *string `json:"organization,omitempty" tf:"organization"`
	// (Athena) AWS S3 bucket to store execution outputs. If not specified, the default query result location from the Workgroup configuration will be used.
	// +optional
	OutputLocation *string `json:"outputLocation,omitempty" tf:"output_location"`
	// (PostgreSQL) Postgres version as a number (903/904/905/906/1000) meaning v9.3, v9.4, etc.
	// +optional
	PostgresVersion *int64 `json:"postgresVersion,omitempty" tf:"postgres_version"`
	// (CloudWatch, Athena) The credentials profile name to use when authentication type is set as 'Credentials file'.
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// (Prometheus) Timeout for queries made to the Prometheus data source in seconds.
	// +optional
	QueryTimeout *string `json:"queryTimeout,omitempty" tf:"query_timeout"`
	// (Elasticsearch and Prometheus) Specifies the ARN of an IAM role to assume.
	// +optional
	Sigv4AssumeRoleArn *string `json:"sigv4AssumeRoleArn,omitempty" tf:"sigv4_assume_role_arn"`
	// (Elasticsearch and Prometheus) Enable usage of SigV4.
	// +optional
	Sigv4Auth *bool `json:"sigv4Auth,omitempty" tf:"sigv4_auth"`
	// (Elasticsearch and Prometheus) The Sigv4 authentication provider to use: 'default', 'credentials' or 'keys' (AMG: 'workspace-iam-role').
	// +optional
	Sigv4AuthType *string `json:"sigv4AuthType,omitempty" tf:"sigv4_auth_type"`
	// (Elasticsearch and Prometheus) When assuming a role in another account use this external ID.
	// +optional
	Sigv4ExternalID *string `json:"sigv4ExternalID,omitempty" tf:"sigv4_external_id"`
	// (Elasticsearch and Prometheus) Credentials profile name, leave blank for default.
	// +optional
	Sigv4Profile *string `json:"sigv4Profile,omitempty" tf:"sigv4_profile"`
	// (Elasticsearch and Prometheus) AWS region to use for Sigv4.
	// +optional
	Sigv4Region *string `json:"sigv4Region,omitempty" tf:"sigv4_region"`
	// (PostgreSQL) SSLmode. 'disable', 'require', 'verify-ca' or 'verify-full'.
	// +optional
	SslMode *string `json:"sslMode,omitempty" tf:"ssl_mode"`
	// (Elasticsearch) Which field that should be used as timestamp.
	// +optional
	TimeField *string `json:"timeField,omitempty" tf:"time_field"`
	// (Prometheus, Elasticsearch, InfluxDB, MySQL, PostgreSQL, and MSSQL) Lowest interval/step value that should be used for this data source.
	// +optional
	TimeInterval *string `json:"timeInterval,omitempty" tf:"time_interval"`
	// (PostgreSQL) Enable usage of TimescaleDB extension.
	// +optional
	Timescaledb *bool `json:"timescaledb,omitempty" tf:"timescaledb"`
	// (All) Enable TLS authentication using client cert configured in secure json data.
	// +optional
	TlsAuth *bool `json:"tlsAuth,omitempty" tf:"tls_auth"`
	// (All) Enable TLS authentication using CA cert.
	// +optional
	TlsAuthWithCaCert *bool `json:"tlsAuthWithCaCert,omitempty" tf:"tls_auth_with_ca_cert"`
	// (All) Controls whether a client verifies the server???s certificate chain and host name.
	// +optional
	TlsSkipVerify *bool `json:"tlsSkipVerify,omitempty" tf:"tls_skip_verify"`
	// (Stackdriver) The token URI used, provided in the service account key.
	// +optional
	TokenURI *string `json:"tokenURI,omitempty" tf:"token_uri"`
	// (OpenTSDB) Resolution.
	// +optional
	TsdbResolution *int64 `json:"tsdbResolution,omitempty" tf:"tsdb_resolution"`
	// (OpenTSDB) Version.
	// +optional
	TsdbVersion *int64 `json:"tsdbVersion,omitempty" tf:"tsdb_version"`
	// (InfluxDB) InfluxQL or Flux.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// (Athena) Workgroup to use.
	// +optional
	Workgroup *string `json:"workgroup,omitempty" tf:"workgroup"`
}

type SourceSpecSecureJSONData struct {
	// (CloudWatch, Athena) The access key to use to access the data source.
	// +optional
	AccessKey *string `json:"-" sensitive:"true" tf:"access_key"`
	// (Github) The access token to use to access the data source
	// +optional
	AccessToken *string `json:"-" sensitive:"true" tf:"access_token"`
	// (Sentry) Authorization token.
	// +optional
	AuthToken *string `json:"-" sensitive:"true" tf:"auth_token"`
	// (All) Password to use for basic authentication.
	// +optional
	BasicAuthPassword *string `json:"-" sensitive:"true" tf:"basic_auth_password"`
	// (All) Password to use for authentication.
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// (Stackdriver) The service account key `private_key` to use to access the data source.
	// +optional
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// (CloudWatch, Athena) The secret key to use to access the data source.
	// +optional
	SecretKey *string `json:"-" sensitive:"true" tf:"secret_key"`
	// (Elasticsearch and Prometheus) SigV4 access key. Required when using 'keys' auth provider.
	// +optional
	Sigv4AccessKey *string `json:"-" sensitive:"true" tf:"sigv4_access_key"`
	// (Elasticsearch and Prometheus) SigV4 secret key. Required when using 'keys' auth provider.
	// +optional
	Sigv4SecretKey *string `json:"-" sensitive:"true" tf:"sigv4_secret_key"`
	// (All) CA cert for out going requests.
	// +optional
	TlsCaCert *string `json:"-" sensitive:"true" tf:"tls_ca_cert"`
	// (All) TLS Client cert for outgoing requests.
	// +optional
	TlsClientCert *string `json:"-" sensitive:"true" tf:"tls_client_cert"`
	// (All) TLS Client key for outgoing requests.
	// +optional
	TlsClientKey *string `json:"-" sensitive:"true" tf:"tls_client_key"`
}

type SourceSpec struct {
	State *SourceSpecResource `json:"state,omitempty" tf:"-"`

	Resource SourceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SourceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The method by which Grafana will access the data source: `proxy` or `direct`.
	// +optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode"`
	// Whether to enable basic auth for the data source.
	// +optional
	BasicAuthEnabled *bool `json:"basicAuthEnabled,omitempty" tf:"basic_auth_enabled"`
	// Basic auth password.
	// +optional
	BasicAuthPassword *string `json:"-" sensitive:"true" tf:"basic_auth_password"`
	// Basic auth username.
	// +optional
	BasicAuthUsername *string `json:"basicAuthUsername,omitempty" tf:"basic_auth_username"`
	// (Required by some data source types) The name of the database to use on the selected data source server.
	// +optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`
	// Custom HTTP headers
	// +optional
	HttpHeaders *map[string]string `json:"-" sensitive:"true" tf:"http_headers"`
	// Whether to set the data source as default. This should only be `true` to a single data source.
	// +optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default"`
	// (Required by some data source types)
	// +optional
	JsonData []SourceSpecJsonData `json:"jsonData,omitempty" tf:"json_data"`
	// A unique name for the data source.
	Name *string `json:"name" tf:"name"`
	// (Required by some data source types) The password to use to authenticate to the data source.
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	SecureJSONData []SourceSpecSecureJSONData `json:"-" sensitive:"true" tf:"secure_json_data"`
	// The data source type. Must be one of the supported data source keywords.
	Type *string `json:"type" tf:"type"`
	// Unique identifier. If unset, this will be automatically generated.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// The URL for the data source. The type of URL required varies depending on the chosen data source type.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
	// (Required by some data source types) The username to use to authenticate to the data source.
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type SourceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SourceList is a list of Sources
type SourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Source CRD objects
	Items []Source `json:"items,omitempty"`
}
