# greeting gRPC service
say "hello" to you
- unary rpc
- server streaming rpc
- client streaming rpc
- bi-directional streaming rpc
- unary with deadline rpc
- add ssl server certs support for unary rpc

# calculator gRPC service
1. sum of two numbers (unary example)
2. prime number decomposition (server streaming example)
3. compute average (client streaming example)
4. find maximum (bi-directioanl streaming example)
5. square root (unary with errors)
6. add reflection & evans cli grpc to calc server
- start calc server
- evans -p 50051 -r
- show package/service/message
- desc <any-message>
- call <any-rpc> (use ctrl-d for streaming calc)