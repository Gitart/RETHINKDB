package main
 
import (
	"log"
	"strings"
 
	r "github.com/dancannon/gorethink"
)
 
func CreateDatabaseIfNotExists(sess *r.Session, dName string) error {
	cur, err := r.DbList().Run(sess)
	if err != nil {
		return err
	}
	dbName := ""
	found := false
	for cur.Next(&dbName) {
		if dbName == dName {
			found = true
			break
		}
	}
	if !found {
		log.Printf("creating database %s\n", dName)
		_, err := r.DbCreate(dName).RunWrite(sess)
		if err != nil {
			return err
		}
	}
	return nil
}
 
func CreateTableIfNotExists(sess *r.Session, dName, tName string) error {
	cur, err := r.Db(dName).TableList().Run(sess)
	if err != nil {
		return err
	}
	defer cur.Close()
	tableName := ""
	found := false
	for cur.Next(&tableName) {
		if tableName == tName {
			found = true
			break
		}
	}
	if !found {
		log.Printf("creating table %s\n", tName)
		_, err := r.Db(dName).TableCreate(tName).RunWrite(sess)
		if err != nil {
			return err
		}
	}
	return nil
}
 
func IndexName(indexes ...string) string {
	return strings.Join(indexes, ":")
}
 
func CreateIndexIfNotExists(
	sess *r.Session,
	dName, tName string,
	indexes ...string,
) error {
	cur, err := r.Db(dName).Table(tName).IndexList().Run(sess)
	if err != nil {
		return err
	}
	defer cur.Close()
	indexName := IndexName(indexes...)
	n := ""
	found := false
	for cur.Next(&n) {
		if n == indexName {
			found = true
			break
		}
	}
	if found {
		return nil
	}
	log.Printf("creating index %s.%s\n", tName, indexName)
	if len(indexes) == 1 {
		_, err := r.Db(dName).Table(tName).IndexCreate(indexes[0]).RunWrite(
			sess)
		if err != nil {
			return nil
		}
	} else {
		_, err := r.Db(dName).Table(tName).IndexCreateFunc(indexName,
			func(row r.Term) interface{} {
				index := make([]interface{}, len(indexes))
				for i, in := range indexes {
					index[i] = row.Field(in)
				}
				return index
			}).RunWrite(sess)
		if err != nil {
			return err
		}
	}
	_, err = r.Db(dName).Table(tName).IndexWait(indexName).Run(sess)
	return err
}
