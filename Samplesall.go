//https://gowalker.org/github.com/dancannon/gorethink#_exams


type Person struct {
    Id        string `gorethink:"id, omitempty"`
    FirstName string `gorethink:"first_name"`
    LastName  string `gorethink:"last_name"`
    Gender    string `gorethink:"gender"`
}

sess, err := r.Connect(r.ConnectOpts{
    Address: url,
    AuthKey: authKey,
})
if err != nil {
    log.Fatalf("Error connecting to DB: %s", err)
}

// Setup table
r.Db("test").TableDrop("table").Run(sess)
r.Db("test").TableCreate("table").Run(sess)
r.Db("test").Table("table").Insert(Person{"1", "John", "Smith", "M"}).Run(sess)

// Fetch the row from the database
res, err := r.Db("test").Table("table").Get("1").Run(sess)
if err != nil {
    log.Fatalf("Error finding person: %s", err)
}

if res.IsNil() {
    log.Fatalf("Person not found")
}

// Scan query result into the person variable
var person Person
err = res.One(&person)
if err != nil {
    log.Fatalf("Error scanning database result: %s", err)
}
fmt.Printf("%s %s (%s)", person.FirstName, person.LastName, person.Gender)
Output:
John Smith (M)
 
Example(GetAll_Compound)
Code:
type Person struct {
    Id        string `gorethink:"id, omitempty"`
    FirstName string `gorethink:"first_name"`
    LastName  string `gorethink:"last_name"`
    Gender    string `gorethink:"gender"`
}

sess, err := r.Connect(r.ConnectOpts{
    Address: url,
    AuthKey: authKey,
})
if err != nil {
    log.Fatalf("Error connecting to DB: %s", err)
}

// Setup table
r.Db("test").TableDrop("table").Run(sess)
r.Db("test").TableCreate("table").Run(sess)
r.Db("test").Table("table").Insert(Person{"1", "John", "Smith", "M"}).Run(sess)
r.Db("test").Table("table").IndexCreateFunc("full_name", func(row r.Term) interface{} {
    return []interface{}{row.Field("first_name"), row.Field("last_name")}
}).Run(sess)
r.Db("test").Table("table").IndexWait().Run(sess)

// Fetch the row from the database
res, err := r.Db("test").Table("table").GetAllByIndex("full_name", []interface{}{"John", "Smith"}).Run(sess)
if err != nil {
    log.Fatalf("Error finding person: %s", err)
}

if res.IsNil() {
    log.Fatalf("Person not found")
}

// Scan query result into the person variable
var person Person
err = res.One(&person)
if err == r.ErrEmptyResult {
    log.Fatalf("Person not found")
} else if err != nil {
    log.Fatalf("Error scanning database result: %s", err)
}

fmt.Printf("%s %s (%s)", person.FirstName, person.LastName, person.Gender)
Output:
John Smith (M)
 
Example(IndexCreate)
Code:
sess, err := r.Connect(r.ConnectOpts{
    Address: url,
    AuthKey: authKey,
})
if err != nil {
    log.Fatalf("Error connecting to DB: %s", err)
}

// Setup database
r.Db("test").TableDrop("table").Run(sess)
r.Db("test").TableCreate("table").Run(sess)

response, err := r.Db("test").Table("table").IndexCreate("name").RunWrite(sess)
if err != nil {
    log.Fatalf("Error creating index: %s", err)
}

fmt.Printf("%d index created", response.Created)
Output:
1 index created

Example(IndexCreate_compound)
Code:
sess, err := r.Connect(r.ConnectOpts{
    Address: url,
    AuthKey: authKey,
})
if err != nil {
    log.Fatalf("Error connecting to DB: %s", err)
}

// Setup database
r.Db("test").TableDrop("table").Run(sess)
r.Db("test").TableCreate("table").Run(sess)

response, err := r.Db("test").Table("table").IndexCreateFunc("full_name", func(row r.Term) interface{} {
    return []interface{}{row.Field("first_name"), row.Field("last_name")}
}).RunWrite(sess)
if err != nil {
    log.Fatalf("Error creating index: %s", err)
}

fmt.Printf("%d index created", response.Created)
Output:
1 index created
 
Example(TableCreate)
Code:
sess, err := r.Connect(r.ConnectOpts{
    Address: url,
    AuthKey: authKey,
})
if err != nil {
    log.Fatalf("Error connecting to DB: %s", err)
}

// Setup database
r.Db("test").TableDrop("table").Run(sess)

response, err := r.Db("test").TableCreate("table").RunWrite(sess)
if err != nil {
    log.Fatalf("Error creating table: %s", err)
}

fmt.Printf("%d table created", response.Created)
Output:
1 table created
Example()
Code:
package gorethink_test

import (
    "fmt"
    "log"
    "os"

    r "github.com/dancannon/gorethink"
)

var session *r.Session
var url, authKey string

func init() {
    // Needed for wercker. By default url is "localhost:28015"
    url = os.Getenv("RETHINKDB_URL")
    if url == "" {
        url = "localhost:28015"
    }

    // Needed for running tests for RethinkDB with a non-empty authkey
    authKey = os.Getenv("RETHINKDB_AUTHKEY")
}

func Example() {
    session, err := r.Connect(r.ConnectOpts{
        Address: url,
        AuthKey: authKey,
    })
    if err != nil {
        log.Fatalf("Error connecting to DB: %s", err)
    }

    res, err := r.Expr("Hello World").Run(session)
    if err != nil {
        log.Fatalln(err.Error())
    }

    var response string
    err = res.One(&response)
    if err != nil {
        log.Fatalln(err.Error())
    }

    fmt.Println(response)
}
 

