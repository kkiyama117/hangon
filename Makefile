PROGRAM = pumpkin

# Macro User Define function
dev_dc = docker-compose -f docker-compose.yml -f lib/docker/development.yml $(1)
stage_dc = docker-compose -f lib/docker/staging.yml $(1)
prod_dc = docker-compose -f docker-compose.yml -f lib/docker/production.yml $(1)
stop_prod = @$(call prod_dc, stop)
stop_stage = @$(call prod_dc, stop)

#all: init_prod_db
# docker development
run_dev: docker-compose.yml lib/docker/development.yml
	$(call dev_dc, up)
down_dev: docker-compose.yml lib/docker/development.yml
	$(call dev_dc, down -v)
config_dev: docker-compose.yml lib/docker/development.yml
	$(call dev_dc, config)

# docker production
build_prod: docker-compose.yml lib/docker/production.yml
	$(call prod_dc, build)
# call after build
run_prod:
	$(call prod_dc, up)
run_prod_shadow:
	@$(call prod_dc, up -d)
init_prod: db/pq run_prod_shadow
	set_prod_db_sql_files
	migrate_prod_db
	$(call prod_dc, stop)
stop_prod: docker-compose.yml lib/docker/production.yml
	$(call stop_prod)
down_prod: docker-compose.yml lib/docker/production.yml
	$(call prod_dc, down -v)
config_prod: docker-compose.yml lib/docker/production.yml
	$(call prod_dc, config)

# docker production
build_stage: docker-compose.yml lib/docker/staging.yml
	$(call stage_dc, build)
# call after build
run_stage:
	$(call stage_dc, up)
run_stage_shadow:
	@$(call stage_dc, up -d)
init_stage: db/pq run_stage_shadow
	set_stage_db_sql_files
	migrate_stage_db
	$(call stage_dc, stop)
stop_stage: lib/docker/staging.yml
	$(call stop_stage)
down_stage: lib/docker/staging.yml
	$(call stage_dc, down -v)
config_stage: lib/docker/staging.yml
	$(call stage_dc, config)

# goose commands
# prod
# call after `run_prod` command
status_prod_db: set_prod_db_sql_files
	$(call prod_dc, exec app goose -env production status)
	$(call stop_prod)
migrate_prod_db: set_prod_db_sql_files
	$(call prod_dc, exec app goose -env production up)
	$(call stop_prod)
down_prod_db: set_prod_db_sql_files
	$(call prod_dc, exec app goose -env production down)
	$(call stop_prod)
set_prod_db_sql_files: $(PROGRAM) db/pq run_prod_shadow
	@$(call prod_dc, exec app cp db/pq/* db/migrations)

# stage
# call after `run_stage` command
status_stage_db: set_stage_db_sql_files
	$(call stage_dc, exec app goose -env production status)
	$(call stop_stage)
migrate_stage_db: set_stage_db_sql_files
	$(call stage_dc, exec app goose -env production up)
	$(call stop_stage)
down_stage_db: set_stage_db_sql_files
	$(call stage_dc, exec app goose -env production down)
	$(call stop_stage)
set_stage_db_sql_files: $(PROGRAM) db/pq run_stage_shadow
	@$(call stage_dc, exec app cp db/pq/* db/migrations)

# dev
status_dev_db: set_dev_db_sql_files↲
	goose -env development status
migrate_dev_db: set_dev_db_sql_files↲
	goose -env development up
down_dev_db: set_dev_db_sql_files↲
	goose -env development down
set_dev_db_sql_files: $(PROGRAM) db/sqlite
	cp db/sqlite/* db/migrations

clean:
	$(rm) pumpkin*

