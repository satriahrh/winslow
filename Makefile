start_mysql:
	brew services start mysql@5.7
entc_generate:
	entc generate ./ent/schema
install:
	rm /Users/satriahrh/go/bin/winslow
	go install