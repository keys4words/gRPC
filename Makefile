geng:
	protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

genc:
	protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.