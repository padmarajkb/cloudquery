# Table: aws_apigateway_domain_name_base_path_mappings

https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_domain_names](aws_apigateway_domain_names).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|domain_name_arn|String|
|arn (PK)|String|
|base_path|String|
|rest_api_id|String|
|stage|String|