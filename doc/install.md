# Install RethinkDB
for Ubuntu 16.04.2 LTS
view version ubuntu
``` bash
lsb_release -a
```


```bash
source /etc/lsb-release && echo "deb http://download.rethinkdb.com/apt $DISTRIB_CODENAME main" | sudo tee /etc/apt/sources.list.d/rethinkdb.list
wget -qO- https://download.rethinkdb.com/apt/pubkey.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install rethinkdb
```

## Install Python driver
For backup and manage database need install Python Driver.

```bash

sudo apt install python-pip  ##  Install the python-pip package
sudo pip install rethinkdb   ##  Install the rethinkdb python client:
sudo apt upgrade             ##  Update the repository:
sudo apt update              ##  Update the repository: 
```

## Start Server

```bash
sudo cp /etc/rethinkdb/default.conf.sample /etc/rethinkdb/instances.d/instance1.conf
sudo vim /etc/rethinkdb/instances.d/instance1.conf
```
Then, restart the service:

```bash
sudo /etc/init.d/rethinkdb restart
```


# BACKUP DATABASE
## Backup All Databases in one file

```bash
rethinkdb dump -c localhost:28015 -f dball.dmp
```

## Backup DB test

```bash
rethinkdb dump -c localhost:28015 -f db.dmp -e test
```

### Path to backup file
home/user/db.dmp 
