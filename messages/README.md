#Challenge 1

##Dependencies
Golang 1.7 or later.<br />
**Note:** This program uses `github.com/gorilla/mux`  and `github.com/asaskevich/govalidator` as a third party dependency

##Setup
1. Clone the repo.
2. Setup go workspace appropriately (if not already setup)
3. `go get github.com/gorilla/mux`
4. `go get github.com/asaskevich/govalidator`
5. `cd $GOPATH/github.com/paxos/combos
6. `go build`
7. `./messages`

##Example success curl commands
```curl -X POST -H "Content-Type: application/json" -d '{"message":"Eva Fineberg"}' http://localhost:12345/message
{"digest":"9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4"}```<br />
```echo -n "Eva Fineberg" | shasum -a 256
   9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4``` <br />
```curl http://localhost:12345/message/9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4
   {"message":"Eva Fineberg"}```

##Example failed curl commands
```curl -i -X POST -H "Content-Type: application/json" -d '{"blah"}' http://localhost:12345/message
   HTTP/1.1 400 Bad Request
   Date: Sat, 28 Jan 2017 20:53:04 GMT
   Content-Length: 57
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Cannot process your request","err_code":400}```<br/>

```curl -i http://localhost:12345/message/blahblahblah
   HTTP/1.1 404 Not Found
   Date: Sat, 28 Jan 2017 20:54:09 GMT
   Content-Length: 47
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Message not found","err_code":404}```<br/>

```curl -i -X POST -H "Content-Type: application/json" -d '{"foo":"Eva Fineberg"}' http://localhost:12345/message
   HTTP/1.1 400 Bad Request
   Date: Sat, 28 Jan 2017 21:30:08 GMT
   Content-Length: 57
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Cannot process your request","err_code":400}```<br/>

```curl -i http://localhost:12345/message/
   HTTP/1.1 404 Not Found
   Content-Type: text/plain; charset=utf-8
   X-Content-Type-Options: nosniff
   Date: Sat, 28 Jan 2017 21:32:21 GMT
   Content-Length: 19

   404 page not found``` <br/>

##Notation
Usually for more extensive restful apis I would deploy with swagger documentation.