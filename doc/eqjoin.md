## Работа с объединеными таблицами


### Regions
```json
[
{"id":  "14" ,"name":  "Николаевская область"} {
"id":  "10" ,"name":  "Киевская область"} {
"id":  "1" ,"name":  "АР Крым"} {
"id":  "13" ,"name":  "Львовская область"} {
"id":  "15" ,"name":  "Одесская область"}
]
```

### Cities
```json
[ {"id":"1001","name":"Константиновка","phone_code":"5256","region_id":"11"},
  {"id":"10","name":"Акимовка","phone_code":"6131","region_id":"8"},
  {"id":"1","name":"Абазовка","phone_code":"5322","region_id":"16"},
  {"id":"1006","name":"Копычинцы","phone_code":"3557","region_id":"19"},
  {"id":"1000","name":"Константиновка","phone_code":"6272","region_id":"5"},
  {"id":"1004","name":"Копылы","phone_code":"5322","region_id":"16"},
  {"id":"1003","name":"Копайгород","phone_code":"4341","region_id":"2"},
  {"id":"1007","name":"Копычинцы","phone_code":"3557","region_id":"19"},
  {"id":"1009","name":"Кореиз","phone_code":"654","region_id":"1"}
]
```

### Объединение двух таблиц с обязательным индексом "region_id" во второй таблице 

```javascript
r.db("Bar").table("regions")
 .eqJoin("id", r.db("Bar").table("cities"), {index:"region_id"})
 .map({ 
         "idcity": r.row("right")("id"),
         "nmcity": r.row("right")("name"),
         "idobl":  r.row("left")("id"),
         "nmobl":  r.row("left")("name"),
      
})
```

   
### Удалить поля в правом подключении  
```javascript
  .without({"right":{"name":true, "phone_code":true}})
  .zip()
```
  
   
### Выборка из областей и входящих в них городов 
```javascript
  r.db("Bar").table("regions")
   .filter({"id":"1"})
   .merge({"Cities":r.db("Bar").table("cities").filter({"region_id":"1"}).coerceTo('array')})
```

### Json Output

```json

{
"Cities": [ ... ] ,
"id":  "1" ,
"name":  "АР Крым"
}
``` 
  
   
   
    
    
    
