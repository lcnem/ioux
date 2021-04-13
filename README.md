# IOUX

The Docker image will be automatically built by Docker Hub when releases are created.

## Install

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt install docker.io -y
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.6/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo gpasswd -a $(whoami) docker
sudo chgrp docker /var/run/docker.sock
sudo systemctl enable docker
sudo systemctl restart docker
```

### Join network

```bash
docker run -v ~/.ioux:/root/.ioux lcnem/ioux iouxd init [moniker] --chain-id [chain-id]
mkdir ioux
cd ioux
curl -L https://raw.githubusercontent.com/lcnem/ioux/main/launch/[chain-id]/genesis.json > ~/.ioux/config/genesis.json
curl -O https://raw.githubusercontent.com/lcnem/ioux/main/docker-compose.yml
docker-compose up -d
```

## Deprecated way of Installation

### Environment setup

This is an example for Ubuntu.

```bash
sudo apt update
sudo apt install build-essential
cd ~
wget https://dl.google.com/go/go1.16.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
echo export PATH='$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Clone

```bash
git clone https://github.com/lcnem/ioux.git
cd ioux
make install
```

### Config daemon

```bash
iouxd init [moniker] --chain-id [chain-id]
cp launch/[chain-id]/genesis.json ~/.ioux/config/genesis.json
```

### Register daemon service

```bash
vi /etc/systemd/system/iouxd.service
```

```txt
[Unit]
Description=IOUX Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/iouxd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

```bash
systemctl enable iouxd
```

## License

Copyright © LCNEM, Inc. All rights reserved.
