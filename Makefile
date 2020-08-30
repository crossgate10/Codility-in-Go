.PHONY: all

new:
	@read -p "Enter Project Name:" pn; mkdir $$pn; \
	make --no-print-directory create ProjectNameVar=$$pn

create:
	@cd $(ProjectNameVar) && echo > $(ProjectNameVar)"_test.go"
	@cd $(ProjectNameVar) && echo > $(ProjectNameVar)".go"
	@cd $(ProjectNameVar) && echo > README.md
	@echo "done"

