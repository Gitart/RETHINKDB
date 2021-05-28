package main

// Пакеты
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "time"
    "flag"
 Db "github.com/dancannon/gorethink"
)

var (
	Cnn []*Db.Session
)



func DBini(){
	// OpcDb := Db.ConnectOpts{Address: "2815", Database: "wrk", AuthKey: "12FF00##", Username:"admin", Password: ""}
	OpcDb := Db.ConnectOpts{Address:":28015", Database: "wrk"}

	session, err := Db.Connect(OpcDb)

	// Обработка ошибок
	if err != nil {
	   fmt.Println("NO CONNECTION !")
	   return
	}

	// Максимальное количество подключений
	// По умолчанию 200
	session.SetMaxOpenConns(200)
	session.SetMaxIdleConns(200)
    
        fmt.Println("CONNECTION to DB...")
	Cnn = append(Cnn, session)
}

// Start service func
func main(){
     DBini()
     Port:= flag.String("p", "1968",   "Input Port")                      // Port by default
     flag.Parse()
  
     http.HandleFunc("/",                             StartPage)     // Start page
     http.HandleFunc("/news/",                        TemplateNews)  // test page
     http.HandleFunc("/add/",                         AddNews)       // Add data
     
     srv := &http.Server{Addr        : ":" + *Port, 
                IdleTimeout : 120 * time.Second, 
                ReadTimeout : 10  * time.Second, 
                WriteTimeout: 10  * time.Second}
                      
     fmt.Println("server was strted in port :", *Port)

     if err := srv.ListenAndServe(); err != nil {
        log.Println(err)
     }
}

//************************************************************
// Вывод сообщений
//************************************************************
func StartPage(w http.ResponseWriter, r *http.Request) {
	t := `<html>
        Сайт сформирован в директории 
        <b>out</b>. 
        Готов к заливке на сайт. 
        <b>Arttech.inf.ua</b>
        </html>`
  
  log.Println("Сайт сформирован.")
  w.Write([]byte(t))
}

//************************************************************
// Вывод сообщений
//************************************************************
func TemplateNews(w http.ResponseWriter, r *http.Request) {
	IOptc := Db.InsertOpts{Durability: "soft"}

	// Cтруктура
	type Inventory struct {
		Country string
		Index   string
	}

	// Данные для зарядки
	sweaters := []Inventory {
		{"Добавление нового документа",             "/docs/"},
		{"Работа с документом по ID документу",     "/docs/id/"},
		{"Получение набора документов по фильтру",  "/docs/filterbyseq/"},
		{"Получение максимального сиквенса",        "/maxseq/"},
		{"Информация по сервису",                   "/docs/info/"},
	}

     _, errst := Db.DB("wrk").Table("settings").Insert(sweaters, IOptc).Run(Cnn[0])
     

if errst != nil {
   w.WriteHeader(208)
   return
} 
 log.Println("Вставка данных произошла успешно.")
}


//************************************************************
// Вывод сообщений
//************************************************************
func AddNews (w http.ResponseWriter, r *http.Request) {
	IOptc := Db.InsertOpts{Durability: "soft"}

	// Чтение тела документа
	reads, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// Проверка ввода пустого значения в теле документа
	if len(string(reads)) == 0 {
		log.Println("Document is Empty and donot insert to Database....\n")
		return
	}

	// var CurTm = time.Now().Format(time.RFC3339Nano)
	// Контроль вывода
	// DocumentStr := string(reads)
	// Формирование для нескольких документов в Json файле документа
	// var m []*Mst // Автоматически подходит для всех форматов Json

	// Формирование для одного документа
	// m := make(Mst)

    var m []Mst

	// Cтрогая опредленная структура заранее
	// m := DocHeader{}
	errj := json.Unmarshal([]byte(reads), &m)

   	// Обработка ошибок
	if errj != nil {
		fmt.Println(errj.Error())
		return
	}

 	 fmt.Println("-->",m)	
	 Db.DB("wrk").Table("settings").Insert(m, IOptc).Run(Cnn[0])
	 fmt.Println("OK...")
	 w.WriteHeader(200)
	 t := `<html>
	    Сайт сформирован в директории 
        <b>out</b>. 
        Готов к заливке на сайт. 
        <b>Arttech.inf.ua</b>
        </html>`
  
        log.Println("Сайт сформирован.")
        w.Write([]byte(t))
}
