variable "project" {
  description = "GCP project ID"
}

variable "region" {
  description = "GCP Region"
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