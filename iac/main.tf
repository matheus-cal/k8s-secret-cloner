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

module "gce_instance" {
  source          = "./modules/gce_k8s_cluster"
  project         = var.project
  region          = var.region
  instance_name   = var.instance_name
  zone = var.zone
  machine_type = var.machine_type
  startup_script  = file(var.startup_script)
}