variable "project" {
  description = "GCP project ID"
}

variable "region" {
  description = "GCP Region"
}

variable "cluster_name" {
  description = "GKE cluster name"
}

variable "node_pool" {
  description =  "Node pool config"
}

variable "instance_name" {
  description = "Name of the GCE instance"
}

variable "startup_script" {
  description = "Startup script for Kubernetes installation"
}

variable "machine_type" {
  description = "Type of computer configuration"
}

variable "zone" {
  description = "Instance zone"
}
