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

.PHONY:gen-payment-client
gen-payment-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service payment --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/payment.proto
.PHONY:gen-payment-server
gen-payment-server:
	@cd app/payment && ${CWGO} server --type RPC --service payment --module ${ROOT_MOD}/app/payment --I ../../idl --idl ../../idl/payment.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"


.PHONY:gen-checkout-client
gen-checkout-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/checkout.proto
.PHONY:gen-checkout-server
gen-checkout-server:
	@cd app/checkout && ${CWGO} server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --I ../../idl --idl ../../idl/checkout.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
.PHONY:gen-frontend-checkout
gen-frontend-checkout:
	@cd app/frontend &&	${CWGO} server --type HTTP --idl  ../../idl/frontend/checkout_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/

.PHONY:gen-order-client
gen-order-client:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/order.proto
.PHONY:gen-order-server
gen-order-server:
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --I ../../idl --idl ../../idl/order.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
gen-frontend-order:
	@cd app/frontend &&	cwgo server --type HTTP --idl  ../../idl/frontend/order_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl/

.PHONY:gen-email-client
gen-email-client:
	@cd rpc_gen && ${CWGO} client --type RPC --service email --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/email.proto
.PHONY:gen-email-server
gen-email-server:
	@cd app/email && ${CWGO} server --type RPC --service email --module ${ROOT_MOD}/app/email --I ../../idl --idl ../../idl/email.proto --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"

.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} . # v=v1.1.1打标签

.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .