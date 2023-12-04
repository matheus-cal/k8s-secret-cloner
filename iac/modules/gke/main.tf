resource "google_container_cluster" "gke_cluster" {
  name     = var.cluster_name
  project  = var.project
  location = var.region

  node_pool {
    name       = "default-pool"
  }
}