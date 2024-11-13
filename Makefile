.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/Blue-Berrys/GoMall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/Blue-Berrys/GoMall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY:gen-frontend
gen-frontend:
	@cd app/frontend &&	cwgo server --type HTTP --idl  ../../idl/frontend/auth_page.proto --service frontend -module github.com/Blue-Berrys/GoMall/app/frontend -I ../../idl/

.PHONY:gen-user-client
gen-user-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/Blue-Berrys/GoMall/rpc_gen --I ../idl --idl ../idl/user.proto

.PHONY:gen-user-server
gen-user-server:
	@cd app/user && cwgo server --type RPC --service user --module github.com/Blue-Berrys/GoMall/app/user --I ../../idl --idl ../../idl/user.proto --pass "-use github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen"
