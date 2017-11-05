"use strict"

let request = require('request');

module.exports = (context, callback) => {
    let url = context;
    request.get(url, (err, res, body) => {
      callback(err, {status: "done", "body": body});
    });
}
