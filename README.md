# hangon
Sample Go server with Clean Architecture
## Usage
```$bash
# for first run
# run build when you change something at `code` folder
make build_prod
# only for migrate_db or first run
make init_prod
# run docker-compose with some setting files
make run_prod
# To stop run this command on the other shell
make stop_prod
# you can also stop docker server with ctrl-C
```

## Dir
### GO
- main.go
#### clean architecture
- domain/
- usecases/
- interface_adapters/
- framework_drivers/
#### net/http
- factories/ (handler)
#### go modules
- go.mod
### Database
- db/
### docker etc. (AutoDeploy)
- Dockerfile
- docker-compose.yml
- .dockerignore
- Makefile
### Nginx
- lib/
- static/
### others
- README.md (this file)
- .gitignore
- LICENCE

## References
Thanks for every people created programs written articles
### Clean architecture
#### official
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
#### Qiita etc.
##### for understand
- https://qiita.com/gki/items/91386b082c57123f1ba0 (structure!)
- https://qiita.com/gki/items/91386b082c57123f1ba0 (why we use clean arch.)
- https://qiita.com/gki/items/ee4a49a881e577c6dfd5 (importance of usecase)
- https://tech-blog.optim.co.jp/entry/2019/01/29/173000 (understand architecture, example with echo)
- https://syfm.hatenablog.com/entry/2017/12/20/021707 (DIP, DI)
- https://postd.cc/golang-clean-archithecture/ (mock etc.)
- https://qiita.com/hirotakan/items/698c1f5773a3cca6193e (example)
- https://qiita.com/mikesorae/items/ff8192fb9cf106262dbf (repository)
- https://qiita.com/koutalou/items/07a4f9cf51a2d13e4cdc (ios)
- https://qiita.com/nrslib/items/a5f902c4defc83bd46b8 (C#)
- https://qiita.com/muroon/items/8add8da911341312176d

### go
#### official etc
https://golang.org/doc/effective_go.html

### go net/http
##### create api server
https://thenewstack.io/make-a-restful-json-api-go/
##### grouping res model
https://deeeet.com/writing/2016/11/01/go-api-client/

### plugins
#### router
[chi - router](https://github.com/go-chi/chi)
https://github.com/go-chi/chi/blob/master/_examples/rest/main.go
#### db
https://wapa5pow.com/compare-golang-database-libraries
- [goose(forked)](https://github.com/pressly/goose)
  - https://dev.to/msh5/start-db-schema-migration-with-goose-bhg
- [gorm](https://github.com/jinzhu/gorm)
  - http://doc.gorm.io/
