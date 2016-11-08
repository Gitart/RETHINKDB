You can import data as follows:

# Import a JSON document into table `users` in database `my_db`
$ rethinkdb import -c HOST:PORT -f user_data.json --table my_db.users

# Import a CSV document
$ rethinkdb import -c HOST:PORT -f user_data.csv --format csv --table my_db.users
The export command works as follows:

# Export a table into a JSON file (placed in the directory
# rethinkdb_export_DATE_TIME by default)
$ rethinkdb export -c HOST:PORT -e my_db.users

# Export a table into a CSV file
$ rethinkdb export -c HOST:PORT -e my_db.users --format csv \
                   --fields first_name,last_name,address      # `--fields` is mandatory when exporting into CSV
