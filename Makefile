.PHONY: all

pn=""
new:
	@read -p "Enter Project Name:" pn; \
    mkdir $$pn && cd $$pn && \
    echo > $$pn"_test.go" && \
    echo > $$pn".go" && \
    echo > README.md