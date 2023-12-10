project        = "inlaid-theater-406815"
region         = "us-central1"
cluster_name   = "kcc-test-7207"
instance_name  = "kcc-test-instance-7207"
machine_type   = "e2-micro"
zone           = "us-central1-c"
startup_script = "./.scripts/config-init-k8s.sh"

node_pool = {
  node_count = 3
}