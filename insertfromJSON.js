   
 r.db("test")
  .table("People")
  .insert(r.http('https://raw.githubusercontent.com/codeforgeek/sample-mockup-data/master/employeeRecordMockup.json',{timeout : 220,resultFormat : 'json'}))
  
