---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_device_bfd Data Source - terraform-provider-fmc"
subcategory: "Device"
description: |-
  This data source can read the Device BFD.
---

# fmc_device_bfd (Data Source)

This data source can read the Device BFD.

## Example Usage

```terraform
data "fmc_device_bfd" "example" {
  id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) UUID of the parent device (fmc_device.example.id).

### Optional

- `domain` (String) The name of the FMC domain
- `id` (String) The id of the object
- `interface_logical_name` (String) Logical Name of the interface of BFD assignment if SINGLE_HOP selected.

### Read-Only

- `bfd_template_id` (String) ID of the BFD Template
- `destination_host_object_id` (String) The ID of the destination host object if MULTI_HOP selected.
- `hop_type` (String) BFD Hop type.
- `interface_id` (String) ID of the interface of BFD assignment if SINGLE_HOP selected.
- `slow_timer` (Number) BFD Slow Timer value in range: 1000-30000, default: 1000
- `source_host_object_id` (String) The ID of the source host object if MULTI_HOP selected.
- `type` (String) Type of the object; this value is always 'BFDPolicy'
