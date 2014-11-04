package main

import (
	"bufio"
	"encoding/json"
	r "github.com/dancannon/gorethink"
	"labix.org/v2/mgo"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

type Zip struct {
	City  string    `json:"city"`
	Loc   []float64 `json:"loc"`
	Pop   int       `json:"pop"`
	State string    `json:"state"`
	Id    string    `json:"_id" gorethink:"zip,omitempty" mgo:"zip"`
}

func main() {
	log.Println(runtime.GOMAXPROCS(6))
	log.Println(runtime.NumCPU())
	DbSession, err := r.Connect(map[string]interface{}{
		"address":     "localhost:28015",
		"database":    "test",
		"maxIdle":     10,
		"idleTimeout": time.Second * 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Db("test").TableDrop("zips").RunWrite(DbSession)
	if err != nil || w.FirstError != "" {
		log.Fatal(err, w.FirstError)
	}
	w, err = r.Db("test").TableCreate("zips").RunWrite(DbSession)
	if err != nil {
		log.Fatal(err, w.FirstError)
	}

	MgoSession, err := mgo.Dial("localhost:27017")
	col := MgoSession.DB("test").C("zips")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("zips.json")
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan int, 2000)
	var wg sync.WaitGroup
	var wg1 sync.WaitGroup

	start := time.Now()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wg.Add(1)
		c <- 1
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		go func(line []byte) {
			defer wg.Done()
			var zip Zip
			err := json.Unmarshal(line, &zip)
			// log.Println(string(line))
			err = col.Insert(zip)
			if err != nil {
				log.Fatal(err)
			}
			<-c
		}(line)
	}

	wg.Wait()
	log.Printf("Inserts to Mongodb took %s", time.Since(start))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f.Seek(0, 0)

	start = time.Now()
	scanner1 := bufio.NewScanner(f)
	for scanner1.Scan() {
		wg1.Add(1)
		c <- 1
		line := make([]byte, len(scanner1.Bytes()))
		copy(line, scanner1.Bytes())
		go func(line []byte) {
			defer wg1.Done()
			var zip Zip
			err := json.Unmarshal(line, &zip)
			// log.Println(string(line))
			w, err := r.Table("zips").Insert(zip).RunWrite(DbSession)
			if err != nil {
				log.Fatal(err)
			}
			if w.FirstError != "" {
				log.Fatal(w.FirstError)
			}
			<-c
		}(line)
	}

	wg1.Wait()
	log.Printf("Inserts to Rethinkdb took %s seconds", time.Since(start))

	if err := scanner1.Err(); err != nil {
		log.Fatal(err)
	}

}
