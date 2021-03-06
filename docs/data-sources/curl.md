---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "curl Data Source - terraform-provider-curl"
subcategory: ""
description: |-
  
---

# curl (Data Source)

Sends HTTP(s) requests, optionally with OAuth2 access_token in the 'Authorization' header.

## Example Usage

```hcl
provider "curl" {
//  client_id = "<client id of this app, registered in Azure AD>"
//  resource = "https://vault.azure.net" //example of the scope/resource to call Azure KeyVault APIs
//  tenant_id = "<azure tenant id>"
//  secret = "" //taken from environment variable 'CURL_CLIENT_SECRET'
}

data "curl" "getTodos" {
  http_method = "GET"
  uri = "https://jsonplaceholder.typicode.com/todos/1"
}

locals {
  json_data = jsondecode(data.curl.getTodos.response)
}

# Returns all Todos
output "all_todos" {
  value = local.json_data
}

//# Returns the title of todo
output "todo_title" {
  value = local.json_data.title
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **http_method** (String) HTTP method like GET, POST, PUT, DELETE, PATCH.
- **uri** (String) URI of resource you'd like to retrieve via HTTP(s).

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **response** (String) Valued returned by the HTTP request.


