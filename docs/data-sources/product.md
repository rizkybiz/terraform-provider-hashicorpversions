---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hashicorpversions_product Data Source - terraform-provider-hashicorpversions"
subcategory: ""
description: |-
  
---

# hashicorpversions_product (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **builds** (List of Object) (see [below for nested schema](#nestedatt--builds))
- **shasums** (String)
- **shasums_signature** (String)
- **shasums_signatures** (List of String)
- **version** (String)

<a id="nestedatt--builds"></a>
### Nested Schema for `builds`

Read-Only:

- **arch** (String)
- **filename** (String)
- **name** (String)
- **os** (String)
- **url** (String)
- **version** (String)

