

# Многоуровневая система

```javascript
r.db("slbc").table("reports").count()
  .delete()
  
r.db("slbc").table("tasks").delete()
  
// Загрузка клиентов  
r.db("slbc").table("clients").insert(r.http("http://172.125.164.116:7777/report/data/clients.json"))
  
  
  .update({"BudgetAA":r.row("BudgetAA").coerceTo("number")})
  //.insert(r.http("http://172.125.164.116:5555/report/data/clients.json"))
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")("images")})
  r.db("test").table("all")("data")("children").nth(0)("data").filter({author:"cujax"})
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")})
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")("images")("resolution")})
  .map({"d": r.row("data")("children")("data")("approved_by")})
  .insert(r.http("https://www.reddit.com/r/golang/new.json"))
  
  r.db("test").table("all").insert(r.http("https://www.reddit.com/r/pics/new.json"))
    
 ```
    

## Доступ ко второму уровню

```javascript

r.db("slbc")
 .table("tasks")
 .get("fb09f477-d013-4682-8227-d53cc2ed9b14")("annotations")
 .map({"Description":r.row("description")})

// Заполнение таблицы
r.db("slbc")
 .table("tasks")
 .insert({
   "annotations" : [
      {
        "id"          : 1,
        "job_id"      : 42,
        "description" : "foo bar",
        "url"         : "http://example.com/job-42",
        "status"      : "passed"
      }
   ]
})
```

  
  

