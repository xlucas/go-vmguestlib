FILE = vmguestlib/session.go
TESTGEN = templates/test.template
TESTFILE = vmguestlib/session_accessor_test.go

gen-test:
	echo 'package vmguestlib\n' > $(TESTFILE)
	echo 'import "testing"' >> $(TESTFILE)
	for accessor in `egrep -i "Get[^ ]+\(\)" $(FILE) | egrep -o "Get[^(]+"`; do \
		FUNCTION=$$accessor envsubst <$(TESTGEN) >> $(TESTFILE); \
	done

test:
	go test -v ./...

.PHONY: gen-test test
