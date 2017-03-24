
Бинарный файл

```javascript
r.http('gravatar.com/avatar/0b1129eaca8152c556c200cd8d179187', {resultFormat: 'binary'})
```

Вставка в таблицу из ссылки
```javascript
r.table("ts").insert(r.http("http://beta.json-generator.com/api/json/get/BhzRccE"));
```

Вставка таблицы в таблицу
```javascript
r.table("tr").insert(r.table("ts"))
```

Поиск в первом уровне + во вторм вхождении
```javascript
r.table("tr").filter({
           index: 1,                                             -- индекс на первом уровне
           name:{                                                -- на первом уровне
                             first:"Britt",                      -- на втором уровне
                             last:"Donaldson"                    -- на втором уровне
                }
             });
```

Вывод второго уровня вложения
```javascript
r.table("tr")("name")("first")
```

Возвращает True если в списке второго уровня есть такое имя хотябы один раз в любой строчке
```javascript
r.table("tr")("name")("first").contains("Jan")
```

Вывод перечисленных полей - и обратите внимание !!!     
Поле name составное с уровнями и оно выводит свои под уровнями

```javascript
r.table("tr").pluck('index','id','isActive','name')

[
{
"id":  "314a57e9-58f9-4102-a081-b4c262d13c7a" ,
"index": 0 ,
"isActive": true ,
"name": {
     "first":  "Roach" ,
     "last":  "Brewer" 
        }
 } ,
 ....
 ]
 
r.table('posts').filter( r.row('category').eq('article').or(r.row('genre').eq('mystery'))).run(conn, callback);
r.db("foo").table("bar").insert(document).run(durability="soft")
```

Не работает!!
```javascript
r.table("Aliance").map(r.row.merge({Title: r.row("id")}).without("id"))
```

Уникальные значения в поле
```javascript
r.table("tr").pluck('age').distinct()
```

Фильтрация с последующим выводом опредленных полей
```javascript
r.table("tr").filter({age:38}).pluck('age','index','id','isActive','name')
```

Выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").pluck('age','index','id','isActive','name').orderBy('index','first')
```

Фильтрация с выводом определенных поле с сортировкой во втором уровне (first)
```javascript
r.table("tr").filter({eyeColor:"sd",age:32 }) .pluck('age','index','id','isActive','name') .orderBy('index','first','last')
```

Отобрать все не пустые значения в поле ege и вывести первіе 3

```javascript
r.table("tr").hasFields('age').limit(3)
```

Группировка с выводом на экран
```javascript
r.table("tr").group("index","age").pluck('age','index')
```

Обновление с поля ид поле Value
```javascript
r.table('Aliance').get(5).update({Value: r.row('count').default(0).add(1) });
```

Поиск в первом уровне по двум условиям
```javascript
r.table("Aliance").filter({"Title": "Aliance 4"})
r.table("tr").filter({eyeColor:"sd",age:38 });
r.table("tr").filter({age:38});
```

Фильтр по значению в поле Title='Aliance 5' или views=1250000
```javascript
r.table('Aliance').filter(r.row('Title').default('foo').eq('Aliance 5').or(r.row('views').default('foo').eq(1250000)))
```

Выбирает построчно начиная (с 3 строки и по 5 ) из набора записей
```javascript
r.table("tr").slice(3,5)
```

Связь по полю ID где поля есть в обеих таблицах
```javascript
r.table("tr").eqJoin('id',r.table("ts"));
```

Обновление по ключу
```javascript
r.table("tr").get("f261abe6-e44e-4b0b-bf07-e60abcb01e0b").update({eyeColor:"sd"})
```

Добавление нового значения в кoнец
```javascript
r.tableCreate('stargazers'); r.table('stargazers').insert( r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers'));
```

Обновление информации
```javascript
r.table('stargazers').update(r.http(r.row('url')))
```

По десять страниц
```javascript
r.http('https://api.github.com/repos/rethinkdb/rethinkdb/stargazers',       { page: 'link-next', pageLimit: 10 })
```

Вывод первого поля из таблицы
```javascript
r.table("tr")("eyeColor");
```

Пропустить две строки и с третьей вібрать 3 записи
```javascript
r.table("Aliance").orderBy("id").skip(2).limit(3)
```

Среднее значение по полю ИД
```javascript
r.table("Aliance").avg("id")
```

Группировка
```javascript
r.table("Aliance").group([r.row("date").year(), r.row("date").month()])
```

Создание индекса
```javascript
r.table('invoices').indexCreate('byDay', [r.row('date').year(), r.row('date').month(), r.row('date').day()])
```

Максимальная по группировке
```javascript
r.table("invoices").group({index: 'byDay'}).max('price')
```

Пример вставки документа
```javascript
r.table('Aliance').insert({
               _id: "5099803df3f4948bd2f98391",
               name: { first: "Alan", last: "Turing" },
               birth: 'Jun 23, 1912',
               death: 'Jun 07, 1954',
               contribs: [ "Turing machine", "Turing test", "Turingery" ],
               views : 1250000
            })
```

Исключение столбца после обновления
```javascript
r.table("Aliance").get("1").replace(r.row.without('Value'))
```

Неработает !!!
```javascript
r.table('posts').filter(r.row('Value').eq('Kyev').or(r.row('Title').eq('Aliance 1'))
```


Замена документа по ИД которій должен быть указан
```javascript
r.table("Aliance").get(1).replace({"id":1,"Title":"New Aliance","Value":"Kyev"})
```

Минуты
```javascript
r.now().minutes()
```

Дата
```javascript
r.now()    --- Thu Nov 06 2014 19:34:55 GMT+00:00
```

Вхождение
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('hijklmnopqrst')
```

Пересечение
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('^[abcdef]{3}')
{
"end": 3 ,
"groups": [ ],
"start": 0 ,
"str":  "abc"
}
```

Поиск вхождения
```javascript
r.expr('abcdefghijklmnopqrstuvwxyz').match('^abcdefghijkl')
```


Выражение
```javascript
r.expr('Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer vel est id leo pellentesque tempus. Vivamus condimentum porttitor lobortis. Donec scelerisque vel nisi vitae hendrerit. Curabitur rhoncus orci enim, eu faucibus nulla varius eu. Duis vel massa mi. Etiam non nibh non ligula dapibus congue. Donec sed lorem eu quam bibendum posuere nec lacinia nunc.')

r.expr(1).do(lambda value: value+2)
r.db('test').table(table['name']).between(None, table['ids'][i], right_bound='closed').replace(r.row.without('update_field'))
r.db('test').table(table['name']).get(table['ids'][i]).delete()

r.db('test').table(table['name']).get(table['ids'][i])
r.db('test').table(table['name']).get_all(str(i), index='field0')
r.db('test').table(table['name']).get_all(str(i), index='field1')
r.db('test').table(table['name']).between(table['ids'][i], table['ids'][i+100])
r.db('test').table(table['name']).between(table['ids'][i], table['ids'][i+100], index='field0')
r.db('test').table(table['name']).filter(r.row['boolean'])
r.db('test').table(table['name']).filter(True)
r.db('test').table(table['name']).filter(False)
r.db('test').table(table['name']).filter(r.row['missing_field'])
r.db('test').table(table['name']).filter(r.row['field0'].gt('5'))
r.db('test').table(table['name']).limit(10).inner_join(r.db('test').table(table['name']), lambda left, right: left['id'] == right['id'])
r.db('test').table(table['name']).limit(10).outer_join(r.db('test').table(table['name']), lambda left, right: left['id'] == right['id'])
r.db('test').table(table['name']).eq_join('id', r.db('test').table(table['name']))
r.db('test').table(table['name']).eq_join('id', r.db('test').table(table['name'])).zip()
r.db('test').table(table['name']).map(r.row['id'])
r.db('test').table(table['name']).with_fields('id')
r.db('test').table(table['name']).with_fields('non_existing_field')
r.db('test').table(table['name']).concat_map(r.row['array_num'].default([]))
r.db('test').table(table['name']).concat_map(r.row['array_str'].default([]))
r.db('test').table(table['name']).limit(900).order_by(r.row['id'])
r.db('test').table(table['name']).order_by(index='id')
r.db('test').table(table['name']).skip(0)
r.db('test').table(table['name']).skip(100)
r.db('test').table(table['name']).limit(0)
r.db('test').table(table['name']).limit(100)
r.db('test').table(table['name'])[0]
r.db('test').table(table['name']).has_fields('id')
r.db('test').table(table['name']).has_fields('missing_field')
r.db('test').table(table['name']).info()
r.db('test').table(table['name']).count()
r.db('test').table(table['name']).between(table['ids'][i], table['ids'][i+100]).count()
r.db('test').table(table['name']).between(table['ids'][i], table['ids'][i+100], index='field0').count()
r.db('test').table(table['name']).filter(r.expr(True)).count()
r.db('test').table(table['name']).get(table['ids'][i]).update({'update_field': 'value'})
# We just clean what we update using between and the fact that the array table['ids'] is sorted
r.db('test').table(table['name']).between(None, table['ids'][i], right_bound='closed').replace(r.row.without('update_field'))
r.db('test').table(table['name']).get(table['ids'][i]).replace(r.row.merge({'replace_field': 'value'}))
r.db('test').table(table['name']).between(None, table['ids'][i], right_bound='closed').replace(r.row.without('replace_field'))
r.db('test').table(table['name']).get(table['ids'][i]).delete()

r.table('users').filter(lambda user: user['age'] >= 18).map(lambda user: user['team_id']).map(lambda team_id: r.table('teams').get(team_id)).map(lambda team: team['team_name']).distinct())
r.table('users').filter(lambda user: user['age'] >= 18) .map(lambda user: r.tables('teams').get(user['team_id'])['name']).distinct())
r.expr(12).do(r.js('(function(x) { return x * x; })')).run()
```

Фибоначи
```javascript
r.table('fib').insert({'id': x,  'value': (r.table('fib').order_by('id').nth(x - 1)['value'] + r.table('fib').order_by('id').nth(x - 2)['value']) })).run(conn)
r.table('fib').order_by('id')['value'].run(conn)
r.expr([2, 3, 4, 5, 6, 7, 8, 9, 10, 11])


constant_queries = [
    # Terms
    "r.expr(123456789)",
    "r.expr('Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer vel est id leo pellentesque tempus. Vivamus condimentum porttitor lobortis. Donec scelerisque vel nisi vitae hendrerit. Curabitur rhoncus orci enim, eu faucibus nulla varius eu. Duis vel massa mi. Etiam non nibh non ligula dapibus congue. Donec sed lorem eu quam bibendum posuere nec lacinia nunc.')",
    "r.expr(None)",
    "r.expr(True)",
    "r.expr(False)",
    "r.expr([0, 1, 2, 3, 4, 5])",
    "r.expr({'a': 'aaaaa', 'b': 'bbbbb', 'c': 'ccccc'})",
    "r.expr({'a': 'aaaaa', 'b': 'bbbbb', 'c': 'ccccc'}).keys()",
    {"query": "r.expr({'a': 'aaaaa', 'c': 'ccccc', 'b': 'bbbbb', 'e': 'eeeee', 'd': 'ddddd', 'g': 'ggggg', 'f': 'fffff', 'i': 'iiiii', 'h': 'hhhhh', 'k': 'kkkkk', 'j': 'jjjjj', 'm': 'mmmmm', 'l': 'lllll', 'o': 'ooooo', 'n': 'nnnnn', 'q': 'qqqqq', 'p': 'ppppp', 's': 'sssss', 'r': 'rrrrr', 'u': 'uuuuu', 't': 'ttttt', 'w': 'wwwww', 'v': 'vvvvv', 'y': 'yyyyy', 'x': 'xxxxx', 'z': 'zzzzz'}).keys()", "tag": "keys for big object"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).insert_at(0, 100)", "tag": "insert_at start"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).insert_at(10, 100)", "tag": "insert_at middle"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).insert_at(20, 100)", "tag": "insert_at end"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).insert_at(20, 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa')", "tag": "insert_at long string"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).delete_at(0)", "tag": "delet_at start"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).delete_at(10)", "tag": "delete_at middle"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).delete_at(19)", "tag": "delete_at end"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).delete_at(5, 15)", "tag": "delete_at range"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).change_at(0, 999)", "tag": "change_at start"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).change_at(10, 999)", "tag": "change_at middle"},
    {"query": "r.expr([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]).change_at(19, 999)", "tag": "change_at end"},
    {
        "query": "r.expr([0, 1, 2, 3, 4, 5])[1:3]",
        "tag": "slice"
    },
    "r.expr('abcdefghijklmnopqrstuvwxyz').match('^abcdefghijkl')",
    "r.expr('abcdefghijklmnopqrstuvwxyz').match('qrstuvwxyz$')",
    "r.expr('abcdefghijklmnopqrstuvwxyz').match('hijklmnopqrst')",
    "r.expr('abcdefghijklmnopqrstuvwxyz').match('^[abcdef]{3}')",
    "(r.expr(1)+6)",
    "(r.expr('abcdef')+'abcdef')",
    "(r.expr([1,2,3,4,5,6,7,8,9])+[1,2,3,4,5,6,7,8,9])",
    "(r.now()+213414)",
    "(r.expr(12331321)-5884631)",
    "(r.now()-5884631)",
    "(r.now()-(r.now()-12312))",
    "(r.expr(12331321)*5884631)",
    "(r.expr([1,2,3])*5000)",
    "(r.expr(12345353213)/56631)",
    "(r.expr(12345353213)%56631)",
    "(r.expr(True) | r.expr(False))",
    "(r.expr(True) | (r.expr([1])*99000))",
    "(r.expr('abc') == r.expr('abc'))",
    "(r.expr('abc') == 2)",
    "(r.expr('abc') == r.expr(['a', 'b', 'c']))",
    "(r.expr('abc') != 'abc')",
    "(r.expr('aaaaaaaaaaaaaaaaaaaaaaaab') != 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa')",
    "(r.expr(1000) > 1)",
    "(r.expr(1000) > 1000)",
    "(r.expr(1000) > 2000)",
    "(r.expr(1000) >= 1)",
    "(r.expr(1000) >= 1000)",
    "(r.expr(1000) >= 2000)",
    "(r.expr(1000) < 1)",
    "(r.expr(1000) < 1000)",
    "(r.expr(1000) < 2000)",
    "(r.expr(1000) <= 1)",
    "(r.expr(1000) <= 1000)",
    "(r.expr(1000) <= 2000)",
    "(~r.expr(True))",
    "(r.expr(True).not_())",
    "(~r.expr(False))",
    "(r.expr(False).not_())",
    # Dates
    "r.now()",
    "r.time(1986, 11, 3, 'Z')",
    "r.epoch_time(531360000)",
    "r.iso8601('1986-11-03T08:30:00-07:00')",
    "r.now().in_timezone('-08:00')",
    "r.now().in_timezone('-08:00').timezone()",
    {"query": "r.now().during(r.time(2013, 12, 1, 'Z'), r.time(2015, 12, 10, 'Z'))", "tag": "during false"},
    {"query": "r.now().during(r.time(2013, 12, 1, 'Z'), r.time(2013, 12, 10, 'Z'))", "tag": "during true end"},
    "r.now().date()",
    "r.now().time_of_day()",
    "r.now().year()",
    "r.now().month()",
    "r.now().day()",
    "r.now().day_of_week()",
    "r.now().day_of_year()",
    "r.now().hours()",
    "r.now().minutes()",
    "r.now().seconds()",
    "r.now().to_iso8601()",
    "r.now().to_epoch_time()",
    "r.expr(1).do(lambda value: value+2)",
    "r.branch(r.expr(False), r.expr(1), r.expr(2))",
    "r.branch(r.expr(True), r.expr(1), r.expr(2))",
    "r.js('Math.random()')",
    "r.expr(1).coerce_to('STRING')",
    "r.expr('1').coerce_to('NUMBER')",
    "r.expr([['a', 'aaaaaa']]).coerce_to('OBJECT')",
    "r.expr({'a': 'aaaaaa'}).coerce_to('ARRAY')",
    "r.expr(1).type_of()",
    "r.expr('str').type_of()",
    "r.expr({'a': 'aaaa'}).type_of()",
    "r.expr([1,2,3]).type_of()",
    "r.expr(1).type_of()",
    "r.expr('str').type_of()",
    "r.expr({'a': 'aaaa'}).type_of()",
    "r.expr([1,2,3]).type_of()",
    "r.json('[1,2,3]').type_of()",
    {
        "query": "r.expr([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
        "tag": "big array"
    },
    "r.expr({'a': 'aaaaa'})['a'].default('default')",
    "r.expr({'a': 'aaaaa'})['b'].default('default')"
```


### Sample
```golang

************************************************************************************************************************
package rethinkdb_session_store

import (
     "errors"
     "net/http"

     "github.com/dancannon/gorethink"
     "github.com/gorilla/securecookie"
     "github.com/gorilla/sessions"
)

type RethinkDBStore struct {
     Codecs  []securecookie.Codec
     Options *sessions.Options // default configuration

     term             gorethink.Term
     rethinkdbSession *gorethink.Session
}

func NewRethinkDBStore(rethinkdbSession *gorethink.Session, db, table string, keyPairs ...[]byte) *RethinkDBStore {
     return &RethinkDBStore{
          Codecs: securecookie.CodecsFromPairs(keyPairs...),
          Options: &sessions.Options{
               Path:   "/",
               MaxAge: 86400 * 30,
          },
          term:             gorethink.Db(db).Table(table),
          rethinkdbSession: rethinkdbSession,
     }
}

func (s *RethinkDBStore) Get(r *http.Request, name string) (*sessions.Session, error) {
     return sessions.GetRegistry(r).Get(s, name)
}

func (s *RethinkDBStore) New(r *http.Request, name string) (*sessions.Session, error) {
     session := sessions.NewSession(s, name)
     opts := *s.Options
     session.Options = &opts
     session.IsNew = true
     var err error
     if c, errCookie := r.Cookie(name); errCookie == nil {
          err = securecookie.DecodeMulti(name, c.Value, &session.ID, s.Codecs...)
          if err == nil {
               err = s.load(session)
               if err == nil {
                    session.IsNew = false
               }
          }
     }
     return session, err
}

func (s *RethinkDBStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
     if err := s.save(session); err != nil {
          return err
     }
     encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, s.Codecs...)
     if err != nil {
          return err
     }
     http.SetCookie(w, sessions.NewCookie(session.Name(), encoded, session.Options))
     return nil
}

func (s *RethinkDBStore) save(session *sessions.Session) error {
     values := map[string]interface{}{}
     for k, v := range session.Values {
          kstr, ok := k.(string)
          if !ok {
               return errors.New("cannot serialize non-string value key")
          }

          values[kstr] = v
     }

     json := map[string]interface{}{
          "name":   session.Name(),
          "values": values,
     }

     var write gorethink.WriteResponse
     var err error
     if session.ID != "" {
          write, err = s.term.Get(session.ID).Update(json).RunWrite(s.rethinkdbSession)
          if err == nil && write.Updated == 0 {
               json["id"] = session.ID
          }
     }

     if write.Updated == 0 {
          write, err = s.term.Insert(json).RunWrite(s.rethinkdbSession)
          if err == nil && len(write.GeneratedKeys) > 0 {
               session.ID = write.GeneratedKeys[0]
          }
     }

     return err
}

func (s *RethinkDBStore) load(session *sessions.Session) error {
     if session.ID == "" {
          return errors.New("invalid session id")
     }

     json := map[string]interface{}{}
     cursor, err := s.term.Get(session.ID).Run(s.rethinkdbSession)
     if err != nil {
          return err
     }
     err = cursor.One(&json)
     if err != nil {
          return err
     }

     values, ok := json["values"].(map[string]interface{})
     if !ok {
          return errors.New("failed to decode session")
     }

     for k, v := range values {
          session.Values[k] = v
     }

     return nil
}
```

## Sample 2
```golang
----------------------------------------------------------------------------------
https://github.com/beshrkayali/rethinkdb-golang-api/blob/master/main_test.go
--------------------------------------------------------------------------
    package main

import (
    "log"
    "os"
    "net/http"
    "encoding/json"
    r "github.com/christopherhesse/rethinkgo"
)

var sessionArray []*r.Session

type Bookmark struct {
    Title string
    Url   string
}

func initDb() {
    session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
    if err != nil {
        log.Fatal(err)
        return
    }

    err = r.DbCreate("gettingstarted").Run(session).Exec()
    if err != nil {
      log.Println(err)
    }

    err = r.TableCreate("bookmarks").Run(session).Exec()
    if err != nil {
      log.Println(err)
    }

    sessionArray = append(sessionArray, session)
}

func main() {

    initDb()

    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/new", insertBookmark)

    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        log.Fatal("Error: %v", err)
    }
}

func insertBookmark(res http.ResponseWriter, req *http.Request) {
    session := sessionArray[0]

    b := new(Bookmark)
    json.NewDecoder(req.Body).Decode(b)

    var response r.WriteResponse

    err := r.Table("bookmarks").Insert(b).Run(session).One(&response)
    if err != nil {
        log.Fatal(err)
        return
    }
    data, _ := json.Marshal("{'bookmark':'saved'}")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
    session := sessionArray[0]
    var response []Bookmark

    err := r.Table("bookmarks").Run(session).All(&response)
    if err != nil {
        log.Fatal(err)
    }

    data, _ := json.Marshal(response)

    res.Header().Set("Content-Type", "application/json")
    res.Write(data)
}
```


## Simple 3
```golang
package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "encoding/json"
    "bytes"
)

func init() {
    initDb()
}

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
    request, _ := http.NewRequest("GET", "/", nil)
    response := httptest.NewRecorder()

    handleIndex(response, request)

    if response.Code != http.StatusOK {
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }
}

func TestHandInsertBookmarkWithStatusOK(t *testing.T) {
    bookmark := Bookmark{"wercker", "http://wercker.com"}

    b, err := json.Marshal(bookmark)
    if err != nil {
        t.Fatalf("Unable to marshal Bookmark")
    }

    request, _ := http.NewRequest("POST", "/new", bytes.NewReader(b))
    response := httptest.NewRecorder()

    insertBookmark(response, request)

    body := response.Body.String()
    if !strings.Contains(body, "{'bookmark':'saved'}") {
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "San Francisco", body)
    }
}

func TestHandleIndexReturnsJSON(t *testing.T) {
    request, _ := http.NewRequest("GET", "/", nil)
    response := httptest.NewRecorder()

    handleIndex(response, request)

    ct := response.HeaderMap["Content-Type"][0]
    if !strings.EqualFold(ct, "application/json") {
        t.Fatalf("Content-Type does not equal 'application/json'")
    }
}
```


## Simple 4

```golang
package main

import (
    "log"
    "os"
    "net/http"
    "encoding/json"
    r "github.com/christopherhesse/rethinkgo"
)

var sessionArray []*r.Session

type Bookmark struct {
    Title string
    Url   string
}

func initDb() {
    session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
    if err != nil {
        log.Fatal(err)
        return
    }

    err = r.DbCreate("gettingstarted").Run(session).Exec()
    if err != nil {
      log.Println(err)
    }

    err = r.TableCreate("bookmarks").Run(session).Exec()
    if err != nil {
      log.Println(err)
    }

    sessionArray = append(sessionArray, session)
}

func main() {

    initDb()

    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/new", insertBookmark)

    err := http.ListenAndServe(":5000", nil)
    if err != nil {
        log.Fatal("Error: %v", err)
    }
}

func insertBookmark(res http.ResponseWriter, req *http.Request) {
    session := sessionArray[0]

    b := new(Bookmark)
    json.NewDecoder(req.Body).Decode(b)

    var response r.WriteResponse

    err := r.Table("bookmarks").Insert(b).Run(session).One(&response)
    if err != nil {
        log.Fatal(err)
        return
    }
    data, _ := json.Marshal("{'bookmark':'saved'}")
    res.Header().Set("Content-Type", "application/json; charset=utf-8")
    res.Write(data)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
    session := sessionArray[0]
    var response []Bookmark

    err := r.Table("bookmarks").Run(session).All(&response)
    if err != nil {
        log.Fatal(err)
    }

    data, _ := json.Marshal(response)

    res.Header().Set("Content-Type", "application/json")
    res.Write(data)
}
```
