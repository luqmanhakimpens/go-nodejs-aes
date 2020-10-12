// Node v6.9.0

//
// TEST FILE (cut down for simplicity)
// To ensure Golang encrypted string can be decrypted in NodeJS.
//

let crypto;
try {
  crypto = require('crypto');
} catch (err) {
  console.log('crypto support is disabled!');
}

var iv = Buffer.alloc(16, 0);
var key = "abcdefghijklmnop";

var encrypted = encrypt("hello my name is aloy", key, iv);
console.log("encrypted_msg: ", encrypted);
var decrypted = decrypt(encrypted, key, iv);
console.log("decrypted_msg: ", decrypted);

// Decrypts cipher text into plain text
function decrypt(cipherText, key, iv) {
  const decipher = crypto.createDecipheriv(ALGORITHM, CIPHER_KEY, iv);
  let decrypted = decipher.update(cipherText, 'hex', 'utf8');
  decrypted += decipher.final('utf8');
  return decrypted;
}

// Encrypts plain text into cipher text
function encrypt(plainText, key, iv) {
  const cipher = crypto.createCipheriv(ALGORITHM, CIPHER_KEY, iv);
  let cipherText;
  try {
    cipherText = cipher.update(plainText, 'utf8', 'hex');
    cipherText += cipher.final('hex');
  } catch (e) {
    cipherText = null;
  }
  return cipherText;
}