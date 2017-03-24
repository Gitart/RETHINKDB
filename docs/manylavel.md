

# Многоуровневая система

```javascript
r.db("slbc").table("reports")
 .count()
  .delete()
  
r.db("slbc").table("tasks")
  .delete()
  
r.db("slbc")
  .table("clients")
  .insert(r.http("http://172.25.64.16:5555/report/data/clients.json"))
  //.delete()
  
  .update({"BudgetAA":r.row("BudgetAA").coerceTo("number")})
  //.insert(r.http("http://172.25.64.16:5555/report/data/clients.json"))
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")("images")})
  r.db("test").table("all")("data")("children").nth(0)("data").filter({author:"cujax"})
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")})
  .map({"Tre": r.row("data")("children").nth(0)("data")("preview")("images")("resolution")})
  .map({"d": r.row("data")("children")("data")("approved_by")})
  .insert(r.http("https://www.reddit.com/r/golang/new.json"))
  
  r.db("test").table("all").insert(r.http("https://www.reddit.com/r/pics/new.json"))
    
 ```
    
    
