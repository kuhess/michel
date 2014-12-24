# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

$script = <<SCRIPT
# Install Go
sudo apt-get update
sudo apt-get install -y build-essential mercurial
sudo hg clone -u release https://code.google.com/p/go /opt/go
cd /opt/go/src
sudo ./all.bash

# Setup the GOPATH
mkdir -p /opt/gopath
cat <<EOF | sudo tee /etc/profile.d/gopath.sh
export GOPATH="/opt/gopath"
export PATH="/opt/go/bin:/opt/gopath/bin:$PATH"
EOF
sudo chmod 0755 /etc/profile.d/gopath.sh

# Make sure the gopath is usable by vagrant
sudo chown -R vagrant:vagrant /opt/go
sudo chown -R vagrant:vagrant /opt/gopath

# Install some other stuff we need
sudo apt-get install -y curl git-core

# Install some go tools
source /etc/profile.d/gopath.sh
go get code.google.com/p/go.tools/cmd/vet
go get github.com/golang/lint/golint
go get code.google.com/p/go.tools/cmd/godoc
SCRIPT

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.synced_folder ".", "/opt/gopath/src/github.com/kuhess/michel"

  # Expose godoc (run a server: godoc -http ":6060")
  config.vm.network "forwarded_port", guest: 6060, host: 6060

  config.vm.provision "shell", inline: $script
end