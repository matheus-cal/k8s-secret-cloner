provider "google" {
  credentials = file("~/.config/gcloud/application_default_credentials.json")
  project     = var.project
  region      = var.region
}

module "gke_cluster" {
  source       = "./modules/gke"
  project      = var.project
  region       = var.region
  cluster_name = var.cluster_name
  node_pool    = var.node_pool
}