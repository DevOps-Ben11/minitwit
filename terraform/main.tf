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

resource "digitalocean_droplet" "manager" {
  image    = "docker-20-04"
  name     = "manager"
  region   = "fra1"
  size     = "s-1vcpu-1gb"
  ssh_keys = [data.digitalocean_ssh_key.Viktoria.id]

  provisioner "remote-exec" {
    inline = [
      "echo 'export DOCKER_USERNAME=${var.DOCKER_USERNAME}' >> ~/.bash_profile",
      "echo 'export PSQL_CON_STR=${var.PSQL_CON_STR}' >> ~/.bash_profile",
      "echo 'export SECRET_COOKIE_HMAC=${var.SECRET_COOKIE_HMAC}' >> ~/.bash_profile",
      "echo 'export SECRET_COOKIE_AES=${var.SECRET_COOKIE_AES}' >> ~/.bash_profile",
      "mkdir /minitwit",
      "mkdir /minitwit/scripts",
      "mkdir /minitwit/config", 
      "docker swarm init --advertise-addr ${self.ipv4_address_private}",
      "SWARM_TOKEN=$(docker swarm join-token -q worker)",
      "echo $SWARM_TOKEN > /tmp/swarm_token",
    ]

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file(var.pvt_key)
      host        = self.ipv4_address
    }
  }

  provisioner "file" {
    source      = "../config"
    destination = "/minitwit"

    connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
    }
  }

  provisioner "file" {
    source      = "../scripts"
    destination = "/minitwit"

    connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /minitwit/scripts/deploy.sh" 
    ]

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file(var.pvt_key)
      host        = self.ipv4_address
    }
  }
}

resource "digitalocean_droplet" "worker-1" {
  image    = "docker-20-04"
  name     = "worker-1"
  region   = "fra1"
  size     = "s-1vcpu-1gb"
  ssh_keys = [data.digitalocean_ssh_key.Viktoria.id]

  connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
  }

  provisioner "file" {
    source      = var.pvt_key
    destination = "/root/.ssh/id_rsa"
    
    connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod 600 /root/.ssh/id_rsa", 
      "mkdir -p /tmp",
      "scp -i /root/.ssh/id_rsa root@${digitalocean_droplet.manager.ipv4_address_private}:/tmp/swarm_token /tmp",
      "docker swarm join --token \"$(cat /tmp/swarm_token)\" ${digitalocean_droplet.manager.ipv4_address_private}:2377",
    ]

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file(var.pvt_key)
      host        = self.ipv4_address
    }
  }
}

resource "digitalocean_droplet" "worker-2" {
  image    = "docker-20-04"
  name     = "worker-2"
  region   = "fra1"
  size     = "s-1vcpu-1gb"
  ssh_keys = [data.digitalocean_ssh_key.Viktoria.id]

  connection {
      host        = self.ipv4_address
      user        = "root"
      type        = "ssh"
      private_key = file(var.pvt_key)
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir -p /tmp",
      "scp -i ${var.pvt_key} root@${digitalocean_droplet.manager.ipv4_address_private}:/tmp/swarm_token /tmp",
      "docker swarm join --token \"$(cat /tmp/swarm_token)\" ${digitalocean_droplet.manager.ipv4_address_private}:2377",
    ]

    connection {
      type        = "ssh"
      user        = "root"
      private_key = file(var.pvt_key)
      host        = self.ipv4_address
    }
  }
}


