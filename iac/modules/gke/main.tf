resource "google_container_cluster" "gke_cluster" {
  name     = var.cluster_name
  project  = var.project
  location = var.region
  deletion_protection = false

  node_pool {
    name       = "default-pool"
  }
}