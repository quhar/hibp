# hibp
Tool which checks if your password has been compromised.

## Usage
Install the tool.
```
$ go get github.com/quhar/hibp
```

And run it
```
$ hibp --pass_file pwned-passwords-sha1-ordered-by-count-v4.txt 
Type your password

You have been pwned!
Found hash "01D0260B3765E465D6B880FEFE3C05CE1EDB16F8" with 1429 hits, it was 137313 password in the list
```
