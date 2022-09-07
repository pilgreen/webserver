releases:
	mkdir -p dist
	env GOOS=darwin go build && tar -czf dist/webserver-osx.tar.gz webserver
	env GOOS=linux go build && tar -czf dist/webserver-linux.tar.gz webserver
	rm webserver
