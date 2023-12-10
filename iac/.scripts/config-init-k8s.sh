#!/bin/bash

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

start_docker() {
  if ! command_exists docker; then
    echo "Instaling Docker..."
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker $USER
    rm get-docker.sh
    echo "Docker installed."
  fi

  if ! sudo systemctl is-active --quiet docker; then
    echo "Initiating Docker..."
    sudo systemctl start docker
    echo "Docker running."
  else
    echo "Docker already running."
  fi
}

check_docker_group() {
  if ! groups | grep -q docker; then
    echo "Adding user to Docker group..."
    sudo usermod -aG docker $USER
    echo "User added to Docker group. Logout and login again or reset the system to set changes."
  else
    echo "User already in Docker group."
  fi
}

fix_docker_socket_permissions() {
  if [ -S /var/run/docker.sock ]; then
    echo "Fixing Docker socket permissions..."
    sudo chmod 666 /var/run/docker.sock
    echo "Docker socket permissions fixed."
  fi
}

install_kind() {
  if ! command_exists kind; then
    echo "Instalando o Kind..."
    curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
    chmod +x ./kind
    sudo mv ./kind /usr/local/bin/kind
    echo "Kind installed."
  else
    echo "Kind already on machine."
  fi
}

install_kubectl() {
  if ! command_exists kubectl; then
    echo "Instalando o kubectl..."
    sudo apt-get update
    sudo apt-get install -y kubectl
    echo "kubectl fixed."
  else
    echo "kubectl already on machine."
  fi
}

create_kind_cluster() {
  if command_exists kind; then
    echo "Creating a Kind cluster..."
    kind create cluster
    kubectl cluster-info --context kind-kind
    echo "Cluster Kind sucesfully created."
  else
    echo "An error occurred. Kind cluster not installed"
  fi
}

start_docker
check_docker_group
fix_docker_socket_permissions
install_kind
install_kubectl
create_kind_cluster

echo "Script done."
