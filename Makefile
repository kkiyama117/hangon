PROGRAM = pumpkin

all: init_prod_db
init_prod_db: migrate_db
status_prod_db: init_prod_db
	goose -env production status
migrate_db: $(PROGRAM)
	goose -env production up
