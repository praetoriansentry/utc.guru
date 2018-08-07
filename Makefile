UTC:=main

$(UTC):
	go build src/main.go

clean:
	$(RM) $(UTC)
