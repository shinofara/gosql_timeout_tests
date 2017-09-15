VERSIONS := 21d7e97c9f760ca685a01ecea202e1c84276daa1 26471af196a17ee75a22e6481b5a5897fb16b081
WAIT_TIMES := 10 3

default:
	@for v in $(VERSIONS); \
	do \
		cd $${v}; \
		glide install; \
		cd ..; \
	done

test:
	@for t in $(WAIT_TIMES); \
	do \
		mysql -u root -ptest -h 127.0.0.1 -P 33306 -e "set global wait_timeout=$${t}; SHOW global VARIABLES LIKE \"wait_timeout\""; \
		for v in $(VERSIONS); \
		do \
			cd $${v}; \
			go test -v; \
			cd ..; \
		done \
	done

clean:
	@for v in $(VERSIONS); \
	do \
		cd $${v}; \
		rm -rf vendor; \
		cd ..; \
	done
