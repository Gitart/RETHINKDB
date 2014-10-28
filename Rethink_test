package main
 
import (
	r "github.com/dancannon/gorethink"
	"log"
	"time"
)
 
func main() {
	log.Println("BEGIN")
	session, err := r.Connect(map[string]interface{}{
		"address": "rethinkdb:28015",
	})
	if err != nil {
		log.Fatalln(err)
	}
 
	db := r.Db("test")
	tableName := "books"
 
	log.Println("create table ")
	if _, err := db.TableCreate(tableName).RunWrite(session); err != nil {
		log.Fatalln(err)
	}
 
	table := db.Table(tableName)
 
	var keys []string
 
	log.Println("insert data")
	if res, err := table.Insert(newData("one")).RunWrite(session); err != nil {
		log.Fatalln(err)
	} else {
		keys = append(keys, res.GeneratedKeys...)
	}
 
	log.Println("bulk insert")
 
	var datas []interface{}
	for _, v := range []string{"two", "three", "four", "five"} {
		datas = append(datas, newData(v))
	}
	if res, err := table.Insert(datas).RunWrite(session); err != nil {
		log.Fatalln(err)
	} else {
		keys = append(keys, res.GeneratedKeys...)
	}
 
	log.Println("make simple index. range query needs index.")
	if _, err := table.IndexCreate("no").RunWrite(session); err != nil {
		log.Fatalln(err)
	}
 
	log.Println("update one record")
	if _, err := table.Get(keys[2]).Update(
		map[string]r.RqlTerm{
			"datas": r.Row.Field("datas").Append("$$$"),
		}).RunWrite(session); err != nil {
		log.Fatalln(err)
	}
 
	log.Println("range query and sort")
	if rows, err := table.Between(2, 4, r.BetweenOpts{Index: "no"}).OrderBy("timestamp").Run(session); err != nil {
		log.Fatalln(err)
	} else {
		for rows.Next() {
			var row interface{}
			rows.Scan(&row)
			log.Println(row)
		}
	}
 
	log.Println("delete all records")
	if _, err := table.GetAll(convert(keys)...).Delete().RunWrite(session); err != nil {
		log.Fatalln(err)
	}
 
	log.Println("drop table")
	if _, err := db.TableDrop(tableName).RunWrite(session); err != nil {
		log.Fatalln(err)
	}
 
	log.Println("END")
}
 
var no int
 
func newData(n string) map[string]interface{} {
	no++
	return map[string]interface{}{
		"name":      n,
		"datas":     []string{"eins", "zwei", "drei"},
		"timestamp": time.Now(),
		"no":        no,
		"nested": map[string]string{
			"kkk": "vvvv",
			"eee": "aaaa",
		},
	}
}
 
func convert(keys []string) []interface{} {
	ks := make([]interface{}, len(keys))
	for i, v := range keys {
		ks[i] = v
	}
	return ks
 
}
