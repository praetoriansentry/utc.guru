UTC:=main

$(UTC):
	GOOS=linux GOARCH=amd64 go build src/main.go

clean:
	$(RM) $(UTC)
	$(RM) main.zip

package:
	zip main.zip main
