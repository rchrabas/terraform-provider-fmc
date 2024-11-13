resource "fmc_ftd_nat_policy" "example" {
  name        = "nat_policy_1"
  description = "My nat policy"
  manual_nat_rules = [
    {
      name                           = "my_rule_1"
      description                    = "My manual nat rule 1"
      enabled                        = true
      section                        = "before_auto"
      fall_through                   = false
      nat_type                       = "STATIC"
      ipv6                           = false
      unidirectional                 = true
      source_interface_id            = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      original_source_id             = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      original_source_port_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      original_destination_id        = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      original_destination_port_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      destination_interface_id       = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_source_id           = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_source_port_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_destination_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_destination_port_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  auto_nat_rules = [
    {
      nat_type                 = "STATIC"
      description              = "My auto nat rule 1"
      destination_interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      fall_through             = false
      ipv6                     = false
      original_network_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      original_port            = 8022
      protocol                 = "TCP"
      source_interface_id      = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_network_id    = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      translated_port          = 22
    }
  ]
}
