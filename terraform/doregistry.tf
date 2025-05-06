
resource "digitalocean_container_registry" "doregistry" {
    name                   = "golang-api"
    subscription_tier_slug = "basic"
}

output "registry_endpoint" {
    value = digitalocean_container_registry.doregistry.endpoint
}