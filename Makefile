all:
	go build -o solve

clean:
	rm -f solve
	go clean