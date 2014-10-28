package main
 
import (
	"log"
	"time"
 
	r "github.com/dancannon/gorethink"
)
 
type Aggregate struct {
	Id         string    `gorethink:"id,omitempty"`
	Language   string    `gorethink:"language"`
	Type       string    `gorethink:"type"`
	Start      time.Time `gorethink:"start"`
	End        time.Time `gorethink:"end"`
	Total      int       `gorethink:"total"`
	TotalDirty bool      `gorethink:"total_dirty"`
	Ratio      float64   `gorethink:"ratio"`
	RatioDirty bool      `gorethink:"ratio_dirty"`
	Rank       int       `gorethink:"rank"`
	RankDirty  bool      `gorethink:"rank_dirty"`
}
 
var (
	sess *r.Session
)
 
const (
	DName   = "platest"
	TName   = "created_aggregate"
	Address = "localhost:28015"
)
 
func must(args ...interface{}) {
	if len(args) == 0 {
		log.Fatal("no args passed")
	}
	last := args[len(args)-1]
	if last == nil {
		return
	}
	log.Fatal(last.(error))
}
 
func main() {
	var err error
 
	// Connect
	sess, err = r.Connect(r.ConnectOpts{
		Address: Address,
	})
	if err != nil {
		log.Fatal(err)
	}
 
	// Create
	must(CreateDatabaseIfNotExists(sess, DName))
	must(CreateTableIfNotExists(sess, DName, TName))
	must(CreateIndexIfNotExists(sess, DName, TName, "language"))
	must(CreateIndexIfNotExists(sess, DName, TName, "type"))
	must(CreateIndexIfNotExists(sess, DName, TName, "start"))
	must(CreateIndexIfNotExists(sess, DName, TName, "end"))
	must(CreateIndexIfNotExists(sess, DName, TName, "total"))
	must(CreateIndexIfNotExists(sess, DName, TName, "total_dirty"))
	must(CreateIndexIfNotExists(sess, DName, TName, "language", "type"))
	must(CreateIndexIfNotExists(sess, DName, TName, "type", "start"))
	must(CreateIndexIfNotExists(sess, DName, TName, "language", "type", "start"))
 
	// Stage
	must(r.Db(DName).Table(TName).Delete().RunWrite(sess))
	must(r.Db(DName).Table(TName).Insert(Aggregate{
		End:        time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
		Id:         "09ffb19d-3fcf-4f05-82ab-7afff125f47f",
		Language:   "Glyph",
		Rank:       112,
		RankDirty:  true,
		Ratio:      0,
		RatioDirty: true,
		Start:      time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
		Total:      0,
		TotalDirty: true,
		Type:       "year",
	}).RunWrite(sess))
 
	// Query
	cur, err := r.Db(DName).Table(TName).GetAllByIndex("total_dirty", true).Run(sess)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()
	aggregates := []Aggregate{}
	must(cur.All(&aggregates))
 
	log.Print(aggregates[0].Start)
}
