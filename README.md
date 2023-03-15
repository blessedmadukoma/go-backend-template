## Golang Backend Template

This is a template guide on how my Go backend projects are structured.

Features Covered:
1. user authentication i.e. registration and login using email with access and refresh tokens
2. session management
3. Database documentation

### Links
1. [Github](https://github.com/blessedmadukoma/go-backend-template)
2. [Database Documentation](https://dbdocs.io/blessedmadukoma/backend-template)


How to run:
1. Clone the repository.
2. Change the name of the `go.mod` file to a mod name of your choice, and replace in all locations.
3. Install all the packages: `go mod tidy`.
4. Set the env configuration.
5. Run `make sqlc` to update db queries and models.
6. First set up your database, the run `make migrate` to perform a migration.
7. Run `make test` to make sure all tests pass.
8. Spin up the project: `go run main.go` or for `air` users: `air`.

To set up the DB Docs using DBML code: [dbdocs](https://dbdocs.io)
1.  install dbdocs using yarn: `yarn global add dbdocs`
2.  create a `docs/db.dbml` which is used to generate docs for the first migration file
3.  login to dbdocs in the terminal/CLI using `dbdocs login`
4.  generate the docs: `dbdocs build <path to your dbml file>/database.dbml` e.g. `dbdocs build docs/db.dbml`
5.  set password for the db docs (not necessary): `dbdocs password --set <password> --project <project name>`

Extras
1. Note: to remove a project: `dbdocs remove <project></project name>`
2. To convert frmo DBML back to SQL, install dbml CLI: `yarn global add @dbml/cli`
3. convert the dbml file to sql code: `dbml2sql --postgres -o docs/schema.sql docs/db.dbml`
4. Update `Makefile` by adding two commands: 
    `db_docs` - generate db docs
    `db_schema` - generate sql (postgres in this case)