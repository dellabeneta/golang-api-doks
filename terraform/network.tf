
resource "digitalocean_vpc" "vpc-doks" {
  name     = "vpc-app"
  region   = var.doks_region
  ip_range = "10.20.10.0/24"
}

