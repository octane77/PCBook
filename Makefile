gen:
	protoc --go_out=./pb .\proto\processor_message.proto

clean:
		rm pb/*.go

run:
	go run main.go

