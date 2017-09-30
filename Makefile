build:
	go fmt && docker build -t loggenerator .
run:
	docker run --rm -it -e DEBUG=true loggenerator
run-file:
	docker run --rm -it -e OUTPUT=/tmp/abc/def.log -e DEBUG=true loggenerator
