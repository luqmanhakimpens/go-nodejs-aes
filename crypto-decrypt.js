var crypto = require('crypto');

var mykey = crypto.createDecipher('aes-128-cbc', 'mypassword');
var mystr = mykey.update('NP65FMCZ3yV5S/nMuFvqcg==', 'base64', 'utf8')
mystr += mykey.final('utf8');

console.log(mystr); //abc