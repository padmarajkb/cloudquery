#!/bin/bash
set -e

# Those are packages that were deprecated but still show up in the proxy

packages_re="github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/[a-z-]+/[a-z]+"

bad_packages_re="github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/web/armweb|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans"

# Get the content of the specified URL using curl
content=$(curl -s https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk)

# Use grep to search for a string that matches a regular expression
# and exclude a list of predefined strings
grep -oE "$packages_re" <<< "$content" | grep -v -E "$bad_packages_re" | xargs go get -u

# curl -s https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk | grep -o -E "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/[a-z-]+/[a-z]+" | xargs -t go get
