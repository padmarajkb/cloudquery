# Table: azure_containerservice_managed_clusters

https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/list?tabs=HTTP#managedcluster

The primary key for this table is **id**.

## Relations

The following tables depend on azure_containerservice_managed_clusters:
  - [azure_containerservice_managed_cluster_upgrade_profiles](azure_containerservice_managed_cluster_upgrade_profiles)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|properties|JSON|
|sku|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|