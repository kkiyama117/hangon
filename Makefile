PROGRAM = pumpkin

#all: init_prod_db
# docker
run_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml up
down_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml down -v
check_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml config

# goose commands
init_prod_db: migrate_db
status_prod_db: init_prod_db
	goose -env production status
migrate_db: $(PROGRAM)
	goose -env production up
