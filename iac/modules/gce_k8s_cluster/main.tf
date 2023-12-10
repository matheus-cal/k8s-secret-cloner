resource "google_compute_instance" "gce_instance" {
  name         = var.instance_name
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12-bookworm-v20231115"
      size = 30
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }

  metadata_startup_script = var.startup_script

}