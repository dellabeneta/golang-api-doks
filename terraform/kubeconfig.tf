
locals {
  kubeconfig = digitalocean_kubernetes_cluster.doks.kube_config[0].raw_config
}

# Salvar o kubeconfig em um arquivo local
resource "local_file" "kubeconfig" {
  content  = local.kubeconfig
  filename = "${path.module}/kubeconfig.yaml"
}