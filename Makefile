dev:
	go build cmd/main.go && mv main simulation
	./simulation

build:
	go build cmd/main.go && mv main simulation