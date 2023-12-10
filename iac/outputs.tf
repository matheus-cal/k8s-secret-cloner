output "gke_cluster_name" {
  value = module.gke_cluster.cluster_name
}

output "instance_cluster_name" {
  value = module.gce_instance.instance_cluster_name
}