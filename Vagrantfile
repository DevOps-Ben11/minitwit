Vagrant.configure("2") do |config|
  config.env.enable # Enable vagrant-env(.env)
  config.vm.box = 'digital_ocean'
  config.vm.box_url = "https://github.com/devopsgroup-io/vagrant-digitalocean/raw/master/box/digital_ocean.box"
  config.ssh.private_key_path = '~/.ssh/id_rsa'

  config.vm.synced_folder "./scripts", "/minitwit/scripts", type: "rsync"

  # By default, vagrant will sync the current directory to /vagrant. We do not want that
  config.vm.synced_folder ".", "/vagrant", disabled: true

  config.vm.define "minitwit-prod" do |server|
    # Define the DigitalOcean provider
    server.vm.provider :digital_ocean do |provider, override|
      # https://github.com/devopsgroup-io/vagrant-digitalocean/issues/277
      override.nfs.functional = false
      override.vm.allowed_synced_folder_types = :rsync

      provider.ssh_key_name = ENV["SSH_KEY_NAME"]

      provider.token = ENV["DIGITAL_OCEAN_TOKEN"]
      provider.image = "ubuntu-22-04-x64"
      provider.region = "fra1"
      provider.size = "s-1vcpu-1gb"
    end

    server.vm.hostname = "minitwit-prod"

    # Setup environment variables on the server
    server.vm.provision "shell", inline: 'echo "export DOCKER_USERNAME=' + "'" + ENV["DOCKER_USERNAME"] + "'" + '" >> ~/.bash_profile'
    server.vm.provision "shell", inline: 'echo "export PSQL_CON_STR=' + "'" + ENV["PSQL_CON_STR"] + "'" + '" >> ~/.bash_profile'

    # # Give permissions to scripts
    server.vm.provision "shell", inline: 'chmod +x /minitwit/scripts/deploy.sh'
    server.vm.provision "shell", inline: 'chmod +x /minitwit/scripts/init.sh'

    # Run init server script
    server.vm.provision "shell", inline: 'bash /minitwit/scripts/init.sh'
  end
end