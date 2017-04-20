# Install RethinkDB
for Ubuntu 16.04.2 LTS
[Link](https://www.linkedin.com/pulse/rethinkdb-installation-ubuntu-1404-deepak-gupta)


View version ubuntu :
``` bash
lsb_release -a
```

## Install Database
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


# RESTORE

Восстановление из Бекапа
с --force
будут добавлены в конец таблиц
Будут восстановлены из архива те базы которые там были в архиве!

1. Если базы нет - она будет создана из архива
2. Если нет таблицы - она будет создана
3. Если не персикаются первичные ключи - записи будут добавлены в таблицы
4. Если пересекаются первичные ключи - будет пропуск этих записей


```bash
rethinkdb restore  db.dmp --force
```

[Генерация JSON](https://www.mockaroo.com/schemas/download)
