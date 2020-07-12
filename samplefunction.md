
## Работа с базой через использование массива
Сначала формируется массив данных и потм производится всавка


```golang
package main

// Пакеты
import (
	r "github.com/dancannon/gorethink"
	"log"
	"time"
	"fmt"
	"sync"
)

var (
	sessionArray         []*r.Session

    wg sync.WaitGroup
	CurTime              = time.Now().Format("2006-01-02 15:04:05")   // Формат
	CurTimeShort         = time.Now().Format("2006-01-02")            // Формат
	CurTimeUnix          = time.Now().Format(time.RFC3339)            // Дата UNIX  
	CurTimeNano          = time.Now().Format(time.RFC3339Nano)        // Дата UNIX nano
	ActiveIp      string = "localhost"                                // Активный адрес 
	TempStr       string = ":5555"                                    // Активный порт
	Remarks              = "Version testing - 5.027 17.06.2016 10:00"
	Term                 = "TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION"
	Servicestatus        = "Ok"
	Notify               = "Programm alredy started or port 5555 is busy other service."
)

// Declaration inetrfaces & structure & other type
type Mii interface{}                                          // Interface
type Mif []interface{}                                        // Cрез Interface
type Msr []string                                             // Срез String
type Mst map[string]interface{}                               // Map - string - interface
type Mss map[string]string                                    // Map - string - string
type Msi map[string]int64                                     // Map - string - int64
type Mis map[int64]string                                     // Map - int64 - string
type Msl map[int]string                                       // Map - int   - string
type Mil []int64
                     
/***************************************************************************************************************************************
 *
 *   Title        : Connection to DB
 *   Initialisation Service
 * 	 Date         : 2015-03-11
 *	 Description  : Initialization DB Connect
 *   Author       : SAVCHENKO ARTHUR 
 *   ☎           : 8-097-5547468
 *
 ****************************************************************************************************************************************/
func Dbini() {

    Conndb       := r.ConnectOpts{ Address: "localhost:28015", 
                                   Database: "test", 
                                   Username:"admin", Password: ""}
                                   
                                   
	session, err := r.Connect(Conndb)

	// Обработка ошибок
	if err != nil {
	   fmt.Println("NO CONNECTION.")
	   return
	}

	// Максимальное количество подключений
	// По умолчанию 200
	session.SetMaxOpenConns(200)
	session.SetMaxIdleConns(200)
	sessionArray = append(sessionArray, session)
}

// Init connect to database
func init(){
     Dbini()	
}
     

// Main procedure
func main(){
	Test_nsert_records()
}


// Вставка с макимальной производительностью
// Для максимальной производительности таблица должна бытьс создана с опцией Durability=Soft
// Формируется массив записей
// Поисходит вставка маасиива в базу
func Test_nsert_records(){
    tb:="W1"
    
    if DeleteTableData(tb) {
       fmt.Println("Error delete...")
       return
    }

    log.Println("Delete table...")

    var dts []Mst
    
    dt:=Mst{"Title": "Insert to table", 
            "Name":  "Todeush",
            "Date":  time.Now(),
            }


    // Внимание !
    // Максимально можно встаивть 100 000 записей !
    // Миллион не получилось !!!!!
    // Сначала формируем массив
    // Потом загоняем его в таблицу
    
    log.Println("Start create array...")    
	for i:=0; i<10000; i++{
	    dts = append(dts,dt)
	}
	
	log.Println("Finish create array...")    


    Insert_Table_Array(dts,tb)
    log.Println("Вставка произошла....")
}



func Inserttotable(){
	tb:="W2"
    
    if DeleteTableData(tb) {
       fmt.Println("Error delete...")
       return
    }

    var dts []Mst
    dt  :=Mst{"test":"Insert to table", "ssss":"sdddd"}

	for i:=0; i<100000; i++ {
	    dts = append(dts,dt)
	}

    Insert_Table_Array(dts,tb )
    log.Println("Вставка произошла....")

    It(dts,tb )
    log.Println("Вставка произошла....")

}    


// Insert Array
func It(Data []Mst, TableName string){
	
	err:=r.DB("Work").Table(TableName).Insert(Data).Exec(sessionArray[0])
     if err!=nil{
     	fmt.Println(err.Error())
     }
    //log.Println("Insert connect to DB")  
}  


func Insert_Table_Array(Data []Mst, TableName string){
   R  := r.InsertOpts{ReturnChanges: false, Durability:"soft"}	

	wg.Add(1)
	go func() {
		defer wg.Done()
	    err:=r.DB("Work").Table(TableName).Insert(Data,R).Exec(sessionArray[0])
        if err!=nil{
     	   fmt.Println(err.Error())
        }
	} ()
	wg.Wait()
}  


// Delete record int table
func DeleteTableData(Name string) bool {
	D := r.DeleteOpts{ReturnChanges: false, Durability:"soft"}	
	e := r.DB("Work").Table(Name).Delete(D).Exec(sessionArray[0])
	if e !=nil {
		return true
	}
    return false
}
```
