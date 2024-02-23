Vagrant.configure("2") do |config|
  config.env.enable # Enable vagrant-env(.env)
  config.vm.box = 'digital_ocean'
  config.vm.box_url = "https://github.com/devopsgroup-io/vagrant-digitalocean/raw/master/box/digital_ocean.box"
  config.ssh.private_key_path = '~/.ssh/id_rsa'
  config.vm.synced_folder ".", "/minitwit", type: "rsync"
  
  config.vm.define "droplet" do |server|
    # Define the DigitalOcean provider
    server.vm.provider :digital_ocean do |provider, override|
      provider.ssh_key_name = ENV["SSH_KEY_NAME"]

      provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
      provider.image = "fedora-39-x64"              # Choose your preferred OS image
      provider.region = "fra1"                         # Choose your preferred region
      provider.size = "s-1vcpu-1gb"                    # Choose your preferred droplet size
    end

    server.vm.hostname = "minitwit"

    # Configure Docker provisioner
    server.vm.provision "docker" do |docker|
      docker.build_image "/minitwit",
        args: "-t minitwit"
        
      docker.run "minitwit",
        args: "-d -p 8080:8080"

    end
  end
end