build:
	go build main.go -o bin/main

run:
	rm output.ppm 2> /dev/null; go run main.go >> output.ppm

clean:
	rm output.ppm 2> /dev/null