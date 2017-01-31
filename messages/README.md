#Challenge 1

##Dependencies
Golang 1.7 or later.<br />
**Note:** This program uses `github.com/gorilla/mux`  and `github.com/asaskevich/govalidator` as a third party dependency

##The Service
This service runs on an AWS instance.<br/>
Included in this repo is the init script used to run the service `messages.sh` as well as the very simple nginx configuration which works as a reverse proxy and lives on the AWS instance.

##Example success curl commands
```bash
curl -X POST -H "Content-Type: application/json" -d '{"message":"Eva Fineberg"}' http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message
{"digest":"9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4"}
echo -n "Eva Fineberg" | shasum -a 256
9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4
curl http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message/9856266689cd9b96c7c54fe80aed98e7acfe2ce6f9260b227a25bd57c5a9c9d4
{"message":"Eva Fineberg"}```
<br/>

##Example failed curl commands
```curl -i -X POST -H "Content-Type: application/json" -d '{"blah"}' http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message
   HTTP/1.1 400 Bad Request
   Date: Sat, 28 Jan 2017 20:53:04 GMT
   Content-Length: 57
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Cannot process your request","err_code":400}

curl -i http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message/blahblahblah
   HTTP/1.1 404 Not Found
   Date: Sat, 28 Jan 2017 20:54:09 GMT
   Content-Length: 47
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Message not found","err_code":404}

curl -i -X POST -H "Content-Type: application/json" -d '{"foo":"Eva Fineberg"}' http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message
   HTTP/1.1 400 Bad Request
   Date: Sat, 28 Jan 2017 21:30:08 GMT
   Content-Length: 57
   Content-Type: text/plain; charset=utf-8

   {"err_msg":"Cannot process your request","err_code":400}

curl -i http://ec2-54-214-173-4.us-west-2.compute.amazonaws.com/message/
   HTTP/1.1 404 Not Found
   Content-Type: text/plain; charset=utf-8
   X-Content-Type-Options: nosniff
   Date: Sat, 28 Jan 2017 21:32:21 GMT
   Content-Length: 19

   404 page not found``` <br/><br/>

##Notation
Usually for more extensive restful apis I would deploy with swagger documentation.