# PostgreSQL and Go

This repo contains a small example of using [PostgreSQL](https://www.postgresql.org/) and the [go programming language](https://go.dev/) together.

The [pgx](https://github.com/jackc/pgx) driver is required for connecting to a local instance of PostgreSQL.

The sample data for this repo is from [Microsoft's](https://learn.microsoft.com/en-us/sql/samples/adventureworks-install-configure?view=sql-server-ver16&tabs=ssms) AdventureWorks OLTP sample database.

The AdventureWorks `.bak` files are backups for Microsoft SQL Server which we are replacing with PostgreSQL. In order to do this we can follow [lorint repo's](https://github.com/lorint/AdventureWorks-for-Postgres) instructions for converting the files and inserting them into PostgreSQL tables.

From here we can execute raw SQL statements in our go programs.