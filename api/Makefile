GCC=go
GCMD=run
GPATH=main.go


RA512=4096
RA256=2048

run:
	make build
	$(GCC) $(GCMD) $(GPATH)


build:
	make build_db


build_db:
	rm
	go run pkg/db/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/