ucenter-rpc:
	goctl rpc protoc ucenter/api/register.proto --go_out=ucenter/api/types --go-grpc_out=ucenter/api/types --zrpc_out=ucenter/api/register --style go_zero
ucenter-rpc-protoc:
	protoc ucenter/api/register.proto --go_out=ucenter/api/types --go-grpc_out=ucenter/api/types
ucenter-api:
	goctl api new ucenterapi --style go_zero
code:
	curl -X POST -H "Content-Type: application/json"  -d '{"phone":"11","country":"cn" }' http://localhost:28888/uc/mobile/code