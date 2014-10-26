// Создание базы
  r.dbCreate("News") 


// Cоздание таблицы
r.db("News").table_create("post")

// Loads
r.db("News").table("post").delete();

r.db("News").table("post").insert({id:"a-125", name:"Stepanov"});
r.db("News").table("post").insert({id:"a-125-20012014", name:"Stepanov"});
r.db("News").table("post").insert( [{"id":"a-125","name":"Stepanov"},
                                    {"id":"a-123","name":"Stepan"},
                                    {"id":"a-125-20012014","name":"Stepanov"},
                                    {"id":"a-124","name":"Stepanov"}] );



r.db("News").table("post");



// Clear
r.db("News").table("post").delete();


// Insert
r.db("News").table("post").insert({id:"a-121", name:"Stepanov", nomer:"123-ss"});

r.db("News").table("post").insert({id:"a-125-20012014", name:"Stepanov"});

r.db("News").table("post").insert( [{"id":"a-125","name":"Stepanov","nomer":"Primer"},
                                    {"id":"a-123","name":"Stepan"},
                                    {"id":"a-125-20012014","name":"Stepanov"},
                                    {"id":"a-124","name":"Stepanov"}] );


// insert
r.db("News").table("post").insert( [{"id":"a-126","name":"Stepanov","nomer":"Primer"},
                                    {"id":"a-127","name":"Stepan"},
                                    {"id":"a-128-20012014","name":"Stepanov"},
                                    {"id":"a-129","name":"Stepanov"}] );

// Sort
r.db("News").table("post").orderBy({index:"id"});

// Create Index
r.db("News").table("post").indexCreate("name");

//
r.db("News").table("post").orderBy({index:"name"});

/r.db("Headoffice").table("S_SUPPLIERS").delete();  
r.db("Headoffice").table("S_SUPPLIERS");    
r.db("Headoffice").table("S_SUPPLIERS").insert({id_sup:1, Name:'Suppliers1', Is_seles:2 });
r.db("Headoffice").table("S_SUPPLIERS").insert({id_sup:1, Name:'Suppliers Name 2', Is_seles:2 });
r.db("Headoffice").table("S_SUPPLIERS").insert({id_sup:1, 
                                                Name:'Suppliers Name 3',                                                                                  Is_seles:1,                          //dd
                                                Price:123.78,
  Name_ext: "Name ext"
                                                                       });

// Create Index
r.db("Headoffice").table("S_SUPPLIERS").orderBy("Name");    
// Create Index
//r.db("Headoffice").table("S_SUPPLIERS").indexCreate("Is_seles");
//   
r.db("Headoffice").table("S_SUPPLIERS").orderBy({index: r.desc("Is_seles")});  

r.expr({a:"b"}).merge({b:[5,2,7]})
