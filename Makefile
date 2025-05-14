ucenter-rpc:
	goctl rpc protoc ucenter/api/register.proto --go_out=ucenter/api/types --go-grpc_out=ucenter/api/types --zrpc_out=ucenter/api/register --style go_zero
ucenter-api:
	goctl api new ucenterapi --style go_zero