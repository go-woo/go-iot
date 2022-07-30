#### copy files between host and guest OS
``` 
sudo apt install -y open-vm-tools
sudo apt install -y open-vm-tools-desktop

sudo gedit /etc/gdm3/custom.conf
WaylandEnable=false
reboot

vm-options-guest disable copy-past and disable p&p
close vm-workstation and restart
change file ext from zip to zip-1 etc...
```
#### hidpi
```
gsettings set org.gnome.desktop.interface scaling-factor 2
```
#### install git curl
``` 
sudo apt-get install -y curl git tree net-tools
```
#### go
```
sudo mv go /usr/local/
```
```
gedit .profile

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$HOME/.local/bin
export GOSUMDB=off
export DEBUG=1
export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```
#### ide
``` 
echo -e \
[Desktop Entry]"\n"\
Name=GoLand"\n"\
Exec=$HOME/GoLand/bin/goland.sh"\n"\
Icon=$HOME/GoLand/bin/goland.png"\n"\
Terminal=false"\n"\
X-MultipleArgs=false"\n"\
Type=Application"\n"\
Encoding=UTF-8"\n"\
Categories=Application"\n"\
StartupNotify=false"\n"\
StartupWMClass=jetbrains-goland"\n"\
 | sudo tee /usr/share/applications/goland.desktop
```
#### docker
``` 
sudo apt-get update
sudo apt-get -y install apt-transport-https ca-certificates curl software-properties-common git
sudo apt install -y docker.io
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo service docker restart
sudo gedit /etc/docker/daemon.json
{
  "registry-mirrors": ["https://hub-mirror.c.163.com"],
  "log-driver": "json-file",
  "log-opts": {
        "max-size" :"50m",
        "max-file":"3"
  }
}

sudo systemctl restart docker.service
```

#### docker-compose
``` 
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/v2.7.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```
