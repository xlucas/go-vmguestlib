PACKAGE = vmguestlib
TEMPLATES = templates

gen-accessor-test:
	PACKAGE=$(PACKAGE) envsubst <$(TEMPLATES)/test_header.template > $(PACKAGE)/session_accessor_test.go
	for accessor in `egrep -i "Get[^ ]+\(\)" $(PACKAGE)/session_accessor.go | egrep -o "Get[^(]+"`; do \
		FUNCTION=$$accessor envsubst <$(TEMPLATES)/test.template >> $(PACKAGE)/session_accessor_test.go; \
	done

gen-accessor-code:

gen-accessor: gen-accessor-code gen-accessor-test

test:
	go test -v ./...

.PHONY: gen-accessor gen-accessor-code gen-accessor-test test
