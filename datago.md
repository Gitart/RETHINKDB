

## Read data with RetinkDB 
Чтение глубокой ветки в сложной структуре
https://www.reddit.com/r/golang/new.json

```Javascript
   r.db("test").table("all")("data")("children").nth(0)("data")("preview")("images").nth(0)("resolutions").nth(0)
```  

## Sample with filtered
```Javascript
r.db("test").table("all")("data")("children").nth(0)("data").filter({author:"cujax"})
```


## Rethinkdb Command

```Javascript
r.db('test').table('Loaddata').delete();
r.db('test').table('Loaddata').insert({id: 2,  name: 'Post2', category: 'banan',   points: 123.99  });
r.db('test').table('Loaddata').insert({id: 3,  name: 'Post3', category: 'banan',   points: 133.45  });
r.db('test').table('Loaddata').insert({id: 4,  name: 'Post5', category: 'apelsin', points: 1345.03 });
r.db('test').table('Loaddata').insert({id: 5,  name: 'Post6', category: 'gruz',    points: 23.22   });
r.db('test').table('Loaddata').insert({id: 6,  name: 'Post2', category: 'gruz',    points: 523.34  });
r.db('test').table('Loaddata').insert({id: 7,  name: 'Post7', category: 'gruz',    points: 223.11  });
r.db('test').table('Loaddata').insert({id: 8,  name: 'Post8', category: 'gruz',    points: 223.55  });
r.db('test').table('Loaddata').insert({id: 9,  name: 'Post9', category: 'gruz',    points: 7123.45 });
r.db('test').table('Loaddata').insert({id: 10, name: 'Post6', category: 'gruz',    points: 823.12  });
r.db('test').table('Loaddata').insert({id: 11, name: 'Post2', category: 'gruz',    points: 9123.55 });
r.db('test').table('Loaddata').insert({id: 12, name: 'Post2', category: 'gruz',    points: 8123.45 });
r.db('test').table('Loaddata').insert({id: 13, name: 'Post7', category: 'gruz',    points: 3123.45 });
r.db('test').table('Loaddata').insert({id: 14, name: 'Post6', category: 'gruz',    points: 2323.67 });
```

## Commands RETHINKDB:
```Javascript
GO:
 ☐ https://github.com/coreos/etcd
```


 ### In ReQL, you can chain commands at the end of other commands using the . operator:
 ```Javascript
 ☐ # Get an iterable cursor to the `users` table (we've seen this above)    r.table('users') 
  ☐ # Return only the `last_name` field of the documents                     r.table('users').pluck('last_name') 
  ☐ # Get all the distinct last names (remove duplicates)                    r.table('users').pluck('last_name').distinct() 
  ☐ # Count the number of distinct last names                                r.table('users').pluck('last_name').distinct().count() 
  ☐ # create a table `users`                                                 r.table_create('users')   
  ☐ # get an iterable cursor to the `users` table                            r.table('users')  
```

### Server-side execution:
 ```Javascript
 ☐ # Create the query to get distinct last names                             distinct_lastnames_query = r.table('users').pluck('last_name').distinct()
 ☐ # Send it to the server and execute                                       distinct_lastnames_query 
 ☐ # Get up to five user documents that have the `age` field defined         r.table('users').has_fields('age').limit(5) 
 ☐ # Get all users older than 30                                             r.table('users').filter(lambda user: user['age'] > 30) 
 ☐ # avoid writing lambdas, RethinkDB supports an alternative syntax:        r.table('users').filter(r.row['age'] > 30) 
 ☐ # WRONG: Get all users older than 30 using the `if` statement             r.table('users').filter(lambda user: True if user['age'] > 30 else False) 
 ☐ # RIGHT: Get all users older than 30 using the `r.branch` command         r.table('users').filter(lambda user: r.branch(user['age'] > 30, True, False)) 
```

### Composing simple commands:
 ```Javascript
 ☐ # Evaluate a JavaScript expression on the server and get the result       r.js('1 + 1') 
 ☐ # Get all users older than 30 (we've seen this above)                     r.table('users').filter(lambda user: user['age'] > 30) 
 ☐ # Get all users older than 30 using server-side JavaScript                r.table('users').filter(r.js('(function (user) { return user.age > 30; })')) 
 ☐ # Find all authors whose last names are also in the `users` table         r.table('authors').filter(lambda author: r.table('users').pluck('last_name').contains(author.('last_name'))). run(conn)
 ☐ # expression                         r.table('users').filter(lambda user: user['salary'] + user['bonus'] < 90000).update(lambda user: {'salary': user['salary'] + user['salary'] * 0.1})
 ☐ # insert                                                                  r.table('fib').insert([{'id': 0, 'value': 0}, {'id': 1, 'value': 1}]) 
 ```

 r.expr([2, 3, 4, 5, 6, 7, 8, 9, 10, 11]).for_each(lambda x:  r.table('fib').insert({'id': x,
                         'value': (r.table('fib').order_by('id').nth(x - 1)['value'] +
                                   r.table('fib').order_by('id').nth(x - 2)['value'])
                        })) 
r.table('fib').order_by('id')['value'] 


Changes:
```Javascript
 ☐ r.table('users').changes .each{|change| p(change)} 
```

### Filtering events:
```Javascript
 ☐ r.table('messages').changes.filter{ |row|  row['new_val']['room_id'].eq(ROOM_ID)} 
 ☐ r.table('scores').changes.filter{ |change|  change['new_val']['score'] > change['old_val']['score']}['new_val'] 
```

### Мои команды:
```Javascript
 ☐ r.dbList()                                                                                                 -- Перечень баз данных на сервере
 ☐ r.db("test").table("Test") ;                                                                               -- все записи из теста   
 ☐ r.db("test").table("Test").delete();                                                                       -- удаление данных в таблице
 ☐ r.db("test").table("Test").insert({ _id: 5, type: "food", item: "aaa", ratings: [ 5, 8, 9 ] })             -- вставка с массивом в поле
 ☐ r.db("test").table("Test").get("26");                                                                      -- показать запись с ключем 26 
 ☐ r.db("test").table("Test").get("26").update({title: "News 0102002"});                                      -- обновить запись с ключем 26 поле Title
 ☐ r.db("test").table("Test").update({title: "News 010200-ssss2"});                                           -- обновить все записи по полю Title
 ☐ r.table('tv_shows').count();                                                                               -- подсчет количества записей в таблице 
 ☐ r.table('tv_shows').filter(r.row('episodes').lt(100))                                                      -- Фильтрация Больше -gt меньше - lt
 ☐ r.use('ho')                                                                                                -- переход в базу НО  
 ☐ r.dbDrop("Planeta")                                                                                        -- удаление базы данных Planeta   
 ☐ r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers').pluck('login', 'id').orderBy('id')   -- Показ двух полей с сортировкой по Id
 ☐ r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers')                                      -- Показ всех полей
 ☐ r.db("test").table("Test").insert(r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers').pluck('login', 'id').orderBy('id'))   -- вставка с сети в таблицу тест с сортировкой
 ☐ r.table('Loaddata').group('category');                                                                      -- Группировка по поле категории
 ☐ r.table('Loaddata').filter({name: 'Post2', category:'gruz'})                                                -- Фильтр по двум полям
 ☐ r.table('Loaddata').group('category').sum('points');                                                        -- Cуммирование по категориям  
 ☐ r.table('Loaddata').hasFields('name').limit(3);                                                             -- Вывод три записи из записей которые имеют не пустое значение в поле name 
 ☐ r.table('Loaddata').filter({points: 1000}).limit(1);                                                        -- у кого поинтов = 1000 и вывести только первую запись  
 ☐ r.table('Loaddata').filter(r.row('points').gt(100))  ;                                                      -- все у кого поиты больше 100   
 ☐ r.table('Loaddata').filter(r.row('points').lt(100))  ;                                                      -- все у кого поиты меньше 100   
```






```javascript
  r.db("test").table("Test").insert({id:      "7644aaf2-9928-4231-aa68-4e65e31bf219",
                                   title:    "The line must be drawn here",
                                   content:  "This far, no further! ...",
                                   category: "Fiction" });
```


### Получить букву 'a'
```javascript
r.db('test').table('Loaddata').insert({id: 2, list: ['a','b','c']});
r.table("Loaddata")  ;
r.db('test').table('Loaddata').get(2).do(r.row('list').nth(0)) ;    // Получить букву 'a' 1 вариант
r.db('test').table('Loaddata').get(2)('list').nth(0);               // Получить букву 'a' 2 вариант
```




### Insert
```javascript
r.db('test').table('Loaddata').delete();
r.db('test').table('Loaddata').insert({id: 2,  name: 'Post2', category: 'banan',   points: 123.99  });
r.db('test').table('Loaddata').insert({id: 3,  name: 'Post3', category: 'banan',   points: 133.45  });
r.db('test').table('Loaddata').insert({id: 4,  name: 'Post5', category: 'apelsin', points: 1345.03 });
r.db('test').table('Loaddata').insert({id: 5,  name: 'Post6', category: 'gruz',    points: 23.22   });
r.db('test').table('Loaddata').insert({id: 6,  name: 'Post2', category: 'gruz',    points: 523.34  });
r.db('test').table('Loaddata').insert({id: 7,  name: 'Post7', category: 'gruz',    points: 223.11  });
r.db('test').table('Loaddata').insert({id: 8,  name: 'Post8', category: 'gruz',    points: 223.55  });
r.db('test').table('Loaddata').insert({id: 9,  name: 'Post9', category: 'gruz',    points: 7123.45 });
r.db('test').table('Loaddata').insert({id: 10, name: 'Post6', category: 'gruz',    points: 823.12  });
r.db('test').table('Loaddata').insert({id: 11, name: 'Post2', category: 'gruz',    points: 9123.55 });
r.db('test').table('Loaddata').insert({id: 12, name: 'Post2', category: 'gruz',    points: 8123.45 });
r.db('test').table('Loaddata').insert({id: 13, name: 'Post7', category: 'gruz',    points: 3123.45 });
r.db('test').table('Loaddata').insert({id: 14, name: 'Post6', category: 'gruz',    points: 2323.67 });
*/
```



```javascript
r.db('test').table('Loaddata');
r.table('Loaddata').group('category');
r.table('Loaddata').filter({name: 'Post2', category:'gruz'});
r.expr([3, 5, 7]).sum();
r.table('Loaddata').sum('points');
```  
