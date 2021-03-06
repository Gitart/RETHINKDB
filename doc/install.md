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
rethinkdb dump -c localhost:28016 -f /Users/UnixRoot/Desktop/backup.zip
```

## Backup DB test

```bash
rethinkdb dump -c localhost:28015 -f db.dmp -e test
```

### Path to backup file
home/user/db.dmp 


# Performing automatic backups

#### backupdb.sh
```bash
#!/bin/bash

echo "($(date -u)) Starting RethinkDB daily backup"
/usr/local/bin/rethinkdb-dump -f /etc/rethinkdbbackup/backup_."$(date)".zip
echo "($(date -u)) Finished creating backup"
```

```
sudo chmod +x shellfile.sh
```

### Crontab
Просомтр текущих заданий
```
crontab -e
```

Now we will add our cron code; here it is:

```
00 00 * * * /Users/UnixRoot/Desktop/backup.sh >>/Users/UnixRoot/Desktop/backup.log
```


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
rethinkdb dump -c localhost:28016 -e company.employees -f  /Users/UnixRoot/Desktop/backup.zip
```


# Import data

```
rethinkdb import -f user.json --table test.users --pkey id
```

# Export

Экспорт в JSON тлоько таблицы ***People*** из базы данных ***test***
```
rethinkdb export -c localhost:28015  --e test.People
```
Будет сформирована директоррия   
rethinkdb_export_2017-04-20T09:46:23


Экспорт в JSON всех таблиц из базы данных ***test***
```
rethinkdb export -c localhost:28015  --e test
```
Будет сформирована директоррия   
rethinkdb_export_2017-04-20T09:46:23


Экспорт в JSON всех таблиц из базы данных ***test*** в папку back
папка будет сформирована автоматически

```
rethinkdb export -c localhost:28015  --e test -d back
```





# Cluster
```bash
rethinkdb --port-offset 1 --directory rethinkdb_data2 --join localhost:29015
```

```bash
rethinkdb --join 104.121.23.24:29015 --bind all
```


By executing the rethinkdb command as follows:

```bash
sudo /etc/init.d/rethinkdb restart
```

By using the service command as follwos:

```bash
sudo service rethinkdb restart
```


By using the systemctl command as follows:

```bash
sudo systemctl rethinkdb restart
```





[Генерация JSON](https://www.mockaroo.com/schemas/download)
