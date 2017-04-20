# Install RethinkDB


```bash
source /etc/lsb-release && echo "deb http://download.rethinkdb.com/apt $DISTRIB_CODENAME main" | sudo tee /etc/apt/sources.list.d/rethinkdb.list
wget -qO- https://download.rethinkdb.com/apt/pubkey.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install rethinkdb
```

## Install Python driver
For backup and manage database need install Python Driver.

```bash
sudo apt install python-pip  
sudo pip install rethinkdb
sudo apt upgrade
sudo apt update
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
