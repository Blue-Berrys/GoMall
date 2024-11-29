export ROOT_MOD=github.com/Blue-Berrys/GoMall
export CWGO=/Users/mac/go/bin/cwgo

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && ${CWGO} server -I ../../idl --type RPC --module ${ROOT_MOD}/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && ${CWGO} server --type RPC --module ${ROOT_MOD}/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY:gen-frontend
gen-frontend:
	@cd app/frontend &&	${CWGO} server --type HTTP --idl  ../../idl/frontend/auth_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/

.PHONY:gen-user-client
gen-user-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service user --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto
.PHONY:gen-user-server
gen-user-server:
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --I ../../idl --idl ../../idl/user.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY:gen-product-client
gen-product-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service product --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/product.proto
.PHONY:gen-product-server
gen-product-server:
	@cd app/product && ${CWGO} server --type RPC --service product --module ${ROOT_MOD}/app/product --I ../../idl --idl ../../idl/product.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
.PHONY:gen-frontend-product
gen-frontend-product:
	@cd app/frontend &&	${CWGO} server --type HTTP --idl  ../../idl/frontend/product_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/
.PHONY:gen-frontend-category
gen-frontend-category:
	@cd app/frontend &&	${CWGO} server --type HTTP --idl  ../../idl/frontend/category_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/

.PHONY:gen-cart-client
gen-cart-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/cart.proto
.PHONY:gen-cart-server
gen-cart-server:
	@cd app/cart &&  ${CWGO} server --type RPC --service cart --module ${ROOT_MOD}/app/cart --I ../../idl --idl ../../idl/cart.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
.PHONY:gen-frontend-cart
gen-frontend-cart:
	@cd app/frontend &&	${CWGO} server --type HTTP --idl  ../../idl/frontend/cart_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/



