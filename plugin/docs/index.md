---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "plugin Provider"
subcategory: ""
description: |-
  
---

# plugin Provider



## Example Usage

```terraform
terraform {
  required_providers {
    nuodbaas = {
      source = "hashicorp.com/edu/nuodbaas"
    }
  }
}

provider "nuodbaas" {
  organization = var.dbaas_credentials.organization
  username     = var.dbaas_credentials.username
  password     = sensitive(var.dbaas_credentials.password)
  url_base     = "http://localhost/nuodb-cp"
}

resource "nuodbaas_project" "dev" {
  organization = var.dbaas_credentials.organization
  name         = "dev"
  sla          = "dev"
  tier         = "n0.nano"
  maintenance = {
    is_disabled = false
  }

}

resource "nuodbaas_project" "nuodb" {
  organization = var.dbaas_credentials.organization
  name         = "nuodb"
  sla          = "dev"
  tier         = "n0.nano"
}

resource "nuodbaas_database" "nuodb" {
  organization = var.dbaas_credentials.organization
  project      = nuodbaas_project.nuodb.name
  name         = "nuodb"
  tier         = "n0.nano"
  dba_password = "helloworld"
  properties = {
    archive_disk_size = "1Gi"
    tier_parameters = {
      smReplicas = 1
      teReplicas = 2
    }
  }
  maintenance = {
    is_disabled = false
  }

  timeouts {
    create = "10m"
  }

}

# resource "nuodbaas_database" "dbaas" {
#   organization = var.dbaas_credentials.organization
#   project      = nuodbaas_project.nuodb.name
#   name         = "dbaas"
#   # tier         = "n1.small"
#   dba_password = "helloworld"

#   properties = {
#     archive_disk_size = "1Gi"
#   }
# }

# data "nuodbaas_projects" "projectsList" {
#   filter {

#   }
# }

# output "projectsList" {
#   value = data.nuodbaas_projects.projectsList
# }

# locals {
#   project_names = {
#     for project in data.nuodbaas_projects.projectsList.projects :
#     project.name => project
#   }
# }

# data "nuodbaas_project" "projectDetail" {
#   for_each = local.project_names
#   organization = var.dbaas_credentials.organization
#   name = "${each.key}"
# }

# output "projectDetail" {
#   value = data.nuodbaas_project.projectDetail
# }

# data "nuodbaas_databases" "databaseList" {
#   filter {
#     organization = var.dbaas_credentials.organization
#     project      = "nuodb"
#   }
# }

# output "databaseList" {
#   value = data.nuodbaas_databases.databaseList
# }
# locals {
#   database_names = {
#     for database in data.nuodbaas_databases.databaseList.databases :
#     database.name => database
#   }
# }

# data "nuodbaas_database" "databaseDetail" {
#   for_each = local.database_names
#   organization = var.dbaas_credentials.organization
#   project      = nuodbaas_project.nuodb.name
#   name = "${each.key}"
# }

# output "nuodbaas_databaseDetails" {
#   value = [for database in data.nuodbaas_database.databaseDetail: database]
# }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `organization` (String) Name of the organization
- `password` (String, Sensitive) Password for Dbaas Client
- `url_base` (String) The base URL for the server, including the protocol
- `username` (String) Username for Dbaas Client.