---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hightouch Provider"
subcategory: ""
description: |-
  Hightouch API: Hightouch Public Rest API to access syncs, models, sources and destinations
---

# hightouch Provider

Hightouch API: Hightouch Public Rest API to access syncs, models, sources and destinations

## Example Usage

```terraform
terraform {
  required_providers {
    hightouch = {
      source  = "de-tf-providers/hightouch"
      version = "4.3.0"
    }
  }
}

provider "hightouch" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bearer_auth` (String, Sensitive)
- `server_url` (String) Server URL (defaults to https://api.hightouch.com/api/v1)
