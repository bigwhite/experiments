broker_addr=localhost:30083

all:
	go test ./... -addr $(broker_addr) -json|go-test-report 

clean:
	rm -fr test_report.html


