# Table: gcp_compute_ssl_certificates

https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates#SslCertificate

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|certificate|String|
|creation_timestamp|String|
|description|String|
|expire_time|String|
|id|Int|
|kind|String|
|managed|JSON|
|name|String|
|private_key|String|
|region|String|
|self_link (PK)|String|
|self_managed|JSON|
|subject_alternative_names|StringArray|
|type|String|