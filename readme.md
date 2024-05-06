source - [GitHub repo](https://github.com/techschool/simplebank/tree/master)

- created sql script `simpleback.sql`
- installed go , vscode , tableplus , docker (postgres) , brew (golang-migrate) , make
- created docker compose
- create migrate up and down sql files
- create makeFile
- sqlc yaml , queries , gencode
- test files for db/ssqlc (main_test , table tests etc)
- create db/store to add transaction capabalities to db (currently we only support queries)
- create test file for store
- handle deadlocks for concurrent transactions
- implement restful http api using Gin for accounts , entry , transfer etc
- install viper to manage config files
- install mockgen to mock database (writing tests for api)

