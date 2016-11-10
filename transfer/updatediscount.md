## Создание новых таблиц для  карточек

```Javascript

//var T="Discbalance";  
//var T="Discdrug";  
//var T="Disclink";  
var T="Disccard";  

//var D="C1476948733638";
var D="C1478265363906";

r.db(D).tableCreate(T, {primaryKey:"ID"});
r.db(D).table(T).indexCreate("DOC_DATE_TIME");
r.db(D).table(T).indexCreate("DOC_DATE_TIME_STR");
r.db(D).table(T).indexCreate("HDF_EDITOR");
r.db(D).table(T).indexCreate("HDF_SEQ");
r.db(D).table(T).indexCreate("ID_DRUG");

r.db(D).table("Seq").insert({"Created":  "2016-11-16 14:10:25" ,"Description":  "Счетчик" ,"SEQ": 1 ,"Status":  "New" ,"Update":  "2016-10-04 18:21:01" ,"id":  T});
r.db("C1478265363906").table("Seq").update({"UPD":"2016-11-10T10:11","Update":"2016-11-10T18:59", "Status":"Old"})

```
