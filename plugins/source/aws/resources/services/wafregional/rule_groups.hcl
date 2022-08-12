service          = "aws"
output_directory = "."
add_generate     = true

resource "aws" "wafregional" "rule_groups" {
  path        = "github.com/aws/aws-sdk-go-v2/service/wafregional/types.RuleGroup"
  description = "A collection of predefined rules that you can add to a web ACL."
  multiplex "ServiceAccountRegionMultiplexer" {
<<<<<<< HEAD:plugins/source/aws/resources/services/wafregional/rule_groups.hcl
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ServiceAccountRegionMultiplexer"
  }
  ignoreError "IgnoreAccessDenied" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.IgnoreAccessDeniedServiceDisabled"
=======
    path   = "github.com/cloudquery/cq-provider-aws/client.ServiceAccountRegionMultiplexer"
    params = ["waf-regional"]
  }
  ignoreError "IgnoreCommonErrors" {
    path = "github.com/cloudquery/cq-provider-aws/client.IgnoreCommonErrors"
>>>>>>> d1d8e7efe (Add ‘unused resource’ example policy (#1378)):resources/services/wafregional/rule_groups.hcl
  }
  deleteFilter "AccountRegionFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.DeleteAccountRegionFilter"
  }

  options {
    primary_keys = ["account_id", "region", "id"]
  }

  userDefinedColumn "account_id" {
    description = "The AWS Account ID of the resource."
    type        = "string"
    resolver "resolveAWSAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSAccount"
    }
  }
  userDefinedColumn "region" {
    description = "The AWS Region of the resource."
    type        = "string"
    resolver "resolveAWSRegion" {
      path = "github.com/cloudquery/cloudquery/plugins/source/aws/client.ResolveAWSRegion"
    }
  }

  column "rule_group_id" {
    rename = "id"
  }

  userDefinedColumn "arn" {
    type              = "string"
    description       = "ARN of the rule group."
    generate_resolver = true
  }

  userDefinedColumn "tags" {
    type              = "json"
    generate_resolver = true
    description       = "Rule group tags."
  }
}
