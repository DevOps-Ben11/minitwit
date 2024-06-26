#https://registry.terraform.io/providers/digitalocean/digitalocean/latest/docs

terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

variable "do_token" {
}

variable "pvt_key" {
  type    = string
  default = "~/.ssh/id_rsa"
}

variable "SECRET_COOKIE_HMAC" {
}

variable "SECRET_COOKIE_AES" {
}

variable "PSQL_CON_STR" {
}

variable "DOCKER_USERNAME" {
}

data "digitalocean_ssh_key" "Viktoria" {
  name = "Viktoria"
}

data "digitalocean_droplet" "manager" {
  name = digitalocean_droplet.manager.name
}

resource "digitalocean_droplet" "manager" {
  image    = "docker-20-04"
  name     = "manager"
  region   = "fra1"
  size     = "s-1vcpu-1gb"
  ssh_keys = [data.digitalocean_ssh_key.Viktoria.id]

  connection {
    type        = "ssh"
    user        = "root"
    private_key = file(var.pvt_key)
    host        = self.ipv4_address
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir -p /minitwit",
      "mkdir -p /minitwit/grafana_data",
      "mkdir -p /minitwit/prom_data",
      "mkdir -p /minitwit/loki_data",
    ]
  }

  provisioner "file" {
    source      = "../config"
    destination = "/minitwit"
  }

  provisioner "file" {
    source      = "../scripts"
    destination = "/minitwit"
  }

  provisioner "remote-exec" {
    inline = [
      "docker plugin install --grant-all-permissions grafana/loki-docker-driver:latest --alias loki",
      "chmod +x /minitwit/scripts/deploy.sh",  
      "echo 'export DOCKER_USERNAME=${var.DOCKER_USERNAME}' >> ~/.bash_profile",
      "echo 'export PSQL_CON_STR=${var.PSQL_CON_STR}' >> ~/.bash_profile",
      "echo 'export SECRET_COOKIE_HMAC=${var.SECRET_COOKIE_HMAC}' >> ~/.bash_profile",
      "echo 'export SECRET_COOKIE_AES=${var.SECRET_COOKIE_AES}' >> ~/.bash_profile",
      "echo 'export MANAGER_PRIVATE_IP=${self.ipv4_address_private}' >> ~/.bash_profile",
      "sudo ufw allow 2377/tcp",
      "sudo ufw allow 7946/tcp",
      "sudo ufw allow 7946/udp",
      "sudo ufw allow 4789/udp",
      "sudo ufw reload",
      "docker swarm init --advertise-addr ${self.ipv4_address_private}",
      "SWARM_TOKEN=$(docker swarm join-token -q worker)",
      "echo $SWARM_TOKEN > /tmp/swarm_token",
      "docker run -it -d -p 7070:8080 -v /var/run/docker.sock:/var/run/docker.sock dockersamples/visualizer",
    ]
  }
}

locals {
  manager_ip = data.digitalocean_droplet.manager.ipv4_address_private 
}

resource "digitalocean_droplet" "worker" {
  image    = "docker-20-04"
  for_each = toset(["worker-1", "worker-2"])
  name     = each.key
  region   = "fra1"
  size     = "s-1vcpu-1gb"
  ssh_keys = [data.digitalocean_ssh_key.Viktoria.id]

  connection {
    type        = "ssh"
    user        = "root"
    private_key = file(var.pvt_key)
    host        = self.ipv4_address
  }

  provisioner "file" {
    source      = var.pvt_key
    destination = "/root/.ssh/id_rsa"
  }

  provisioner "remote-exec" {
    inline = [
      "docker plugin install --grant-all-permissions grafana/loki-docker-driver:latest --alias loki",
      "chmod 600 /root/.ssh/id_rsa",
      "sudo ufw allow 2377/tcp",
      "sudo ufw allow 7946/tcp",
      "sudo ufw allow 7946/udp",
      "sudo ufw allow 4789/udp",
      "sudo ufw reload",
      "mkdir -p /tmp",
      "scp -o StrictHostKeyChecking=no -o BatchMode=yes -i /root/.ssh/id_rsa root@${local.manager_ip}:/tmp/swarm_token /tmp",
      "SWARM_TOKEN=$(cat /tmp/swarm_token)",
      "docker swarm join --token $SWARM_TOKEN ${local.manager_ip}:2377",
    ]
  }
}

