PROGRAM = pumpkin

#all: init_prod_db
# docker
run_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml up
down_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml down -v
check_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml config

build_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml build
run_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml up
stop_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml stop
down_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml down -v
check_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml config

# goose commands
init_prod_db: migrate_prod_db
status_prod_db: set_prod_db_sql_file
	goose -env production status
migrate_prod_db: set_prod_db_sql_file
	goose -env production up
down_prod_db: set_prod_db_sql_file
	goose -env production down
set_prod_db_sql_file: $(PROGRAM)
	cp db/pq/* db/migrations/
