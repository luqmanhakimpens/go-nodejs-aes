var crypto = require('crypto');

var mykey = crypto.createCipher('aes-128-cbc', 'mypassword');
var mystr = mykey.update('abc', 'utf8', 'base64')
mystr += mykey.final('base64');

console.log(mystr); //34feb914c099df25794bf9ccb85bea72