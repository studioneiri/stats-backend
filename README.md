# stats-backend

## Local Setup

Install Postgres.  If using a Mac, you can download Postgres via Brew:
```shell
brew install postgresql
```
Once PGAdmin is installed, create a new user for yourself for development.  In PGAdmin, right click on 
`Login/Group Roles` -> `Create` -> `Login/Group Roles`.  Make sure to add `create databases` permission to your user

Create database named `stats`:
```shell
createdb stats
```

Run the script `create_tables.sql` to create `locations` and `languages` tables.
```shell
psql -U [user] -d stats -a -f scripts/create_tables.sql
```

### Notes on Example Data
Locations and languages are accurate, but the `estimated_speakers` for each language is not accurate and just a Proof-of-Concept