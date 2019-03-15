'use strict';

const express = require('express');
const request = require('request');


// Constants
const PORT = 8080;
const HOST = '0.0.0.0';
const BACKEND = '10.245.4.236'

// App
const app = express();
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
  res.render('pages/index');
});

app.get('/currency', (req, resp) => {
  console.log('http://'+BACKEND+"/?currency="+req.query.currency+" API request send");
  request('http://'+BACKEND+"/?currency="+req.query.currency, { json: true }, (err, res, body) => {
  if (err) { return console.log(err); }
  resp.send(body);
});
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
