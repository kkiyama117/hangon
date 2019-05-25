PROGRAM = pumpkin

#all: init_prod_db
# docker development
run_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml up
down_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml down -v
check_dev: docker-compose.yml lib/docker/development.yml
	docker-compose -f docker-compose.yml -f lib/docker/development.yml config
# docker production
build_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml build
run_prod: build_prod
	docker-compose -f docker-compose.yml -f lib/docker/production.yml up
init_prod_db: db/pq run_prod
	docker-compose -f docker-compose.yml -f lib/docker/production.yml exec app cp db/pq/* db/migrations
	docker-compose -f docker-compose.yml -f lib/docker/production.yml exec app goose -env production up
stop_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml stop
down_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml down -v
config_prod: docker-compose.yml lib/docker/production.yml
	docker-compose -f docker-compose.yml -f lib/docker/production.yml config

# goose commands (debug, deprecated)
# TODO: Rewrite
# Deprecateted
status_prod_db: set_prod_db_sql_file
	goose -env production status
migrate_prod_db: set_prod_db_sql_file
	goose -env production up
down_prod_db: set_prod_db_sql_file
	goose -env production down
set_prod_db_sql_file: $(PROGRAM) db/pq
	mkdir db/migrations
	cp db/pq/* db/migrations/
