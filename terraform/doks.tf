
resource "digitalocean_kubernetes_cluster" "doks" {
  name     = "golang-api-cluster"
  region   = var.doks_region
  version  = var.doks_version
  vpc_uuid = digitalocean_vpc.vpc-doks.id
  registry_integration = true

  node_pool {
    name       = "golang-api-pool"
    size       = "s-2vcpu-4gb"
    auto_scale = true
    min_nodes  = 2
    max_nodes  = 5
  }
}