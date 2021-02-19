terraform {
  required_providers {
    curl = {
      version = "0.1"
      source  = "schoewe.me/terraform/curl"
    }
  }
}

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