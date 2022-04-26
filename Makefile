geng:
	protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

genc:
	protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.

genb:
	protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.

runblog:
	docker-compose -f blog/docker-compose.yml up -d
	sleep 9
	go run blog/blog_server/server.go

blog_evans:
	evans -p 50051 -r