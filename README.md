# prox
## Installation
```
go install -v github.com/sharif1337/prox@latest
```
## Usage
## Example input file:
```
$ cat params.txt
https://example.com/path?one=1&two=2
https://example.com/path?two=2&one=1
```
## Replaced Values
```
$ cat params.txt | prox Test
https://example.com/path?one=Test&two=Test
https://example.com/path?two=Test&one=Test
```
## Appended Values
```
$ cat params.txt | prox -a Test
https://example.com/path?one=1Test&two=2Test
https://example.com/path?two=2Test&one=1Test
```
[![Facebook](https://img.shields.io/badge/Facebook-Profile-blue?style=flat-square&logo=facebook)](https://www.facebook.com/sharifansari00)
