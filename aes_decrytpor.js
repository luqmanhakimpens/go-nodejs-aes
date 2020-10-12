//npm install --save-dev crypto-js
var CryptoJS = require("crypto-js");
var encrypted_message_b64 = 'BcFCQ4goMgfOcXamt+ODm9AHhMuPt1E+lx+q5XN6Wua5v9jv8CvVszloqy1W9iwRlKekBq7b36TTH+OC64ZMHlbU8VUsKWPYBUST4NwaGBM=';

var iv = CryptoJS.enc.Hex.parse( '00000000000000000000000000000000' ); //16 bytes length hex
var key= CryptoJS.enc.Hex.parse( '6162636465666768696A6B6C6D6E6F70' );

// Decrypt
var bytes  = CryptoJS.AES.decrypt( encrypted_message_b64, key , { iv: iv} );
var plaintext = bytes.toString(CryptoJS.enc.Base64);
var decoded_b64msg =  new Buffer(plaintext , 'base64').toString('ascii');
var decoded_msg =     new Buffer( decoded_b64msg , 'base64').toString('ascii');

console.log("decoded_msg: ", decoded_msg)
