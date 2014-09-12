#!/usr/bin/env bash

#
# Change ROOTPATH variable to represent your root src path
# for your go project. For example: github.com/adampresley.
# Project name should be specified. This will become a symbolic
# link to your vagrant folder.
#
ROOTPATH="github.com/adampresley"
PROJECTNAME="logamus-prime"

#
# Setup dependencies
#
apt-get update
apt-get -y install git mercurial

#
# Download and install Go
#
wget https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz -O /home/vagrant/go.tar.gz
tar -C /usr/local -xzf /home/vagrant/go.tar.gz

#
# Add Go to the profile path
#
echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/vagrant/.profile

#
# Setup the Go project directory expectations.
#
mkdir -p /home/vagrant/go/src/$ROOTPATH
mkdir /home/vagrant/go/bin
mkdir /home/vagrant/go/pkg

ln -sf /vagrant /home/vagrant/go/src/$ROOTPATH/$PROJECTNAME

chown -R vagrant:vagrant /home/vagrant/go

echo 'export GOPATH=/home/vagrant/go' >> /home/vagrant/.profile