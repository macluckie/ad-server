# Ad-server
1.clone:
https://github.com/macluckie/ad-server.git  
2. start the stack: 

```console

docker-compose up -d

```
- the protofile are already compile by this command:   
protoc --go_out=./proto/  
 --go_opt=paths=source_relative  
 --go-grpc_out=./proto/   
--go-grpc_opt=paths=source_relative proto/ad.proto   


The server GRPC (ad-server) is up on  localhost:50051.  
To request the server we are going to use postman.  
- In postman import the protofile (ad.proto in ad-server). 
you will find 3 methods:   
CreateAd.  
GetAd.  
ServeAd.   

scenario:  
you can test the application by create an Ad and serve the Ad.   
