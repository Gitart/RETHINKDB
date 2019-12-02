const r = require('rethinkdb');

async function Test(){
    const connObj    = {host: 'localhost', port: 28015, db: 'test'}
    const connection = await r.connect(connObj)


    r.db('test').tableDrop('posts1').run(connection);
    r.db('test').tableDrop('posts2').run(connection);

    r.db('test').tableCreate('posts1').run(connection);
    r.db('test').tableCreate('posts2').run(connection);
    console.log('connection');

    r.db('Work').tableCreate('posts1').run(connection);
    r.db('Work').tableCreate('posts2').run(connection);
    console.log('connection');

    r.db('Work').table('posts1').insert({id:12, title:"tetst product"}).run(connection);
};


async function Test2(){
    const connObj    = {host: 'localhost', port: 28015, db: 'test'}
    const connection = await r.connect(connObj)
    r.db('Work').table('posts1').insert({ids:12, title:"tetst product"}).run(connection);
    console.log('Inserting....');
};

Test2();

                         
