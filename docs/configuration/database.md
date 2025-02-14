# Database

GoToSocial stores statuses, accounts, etc, in a database. This can be either [SQLite](https://sqlite.org/index.html) or [Postgres](https://www.postgresql.org/).

By default, GoToSocial will use Postgres, but this is easy to change.

## SQLite

SQLite, as the name implies, is the lightest database type that GoToSocial can use. It stores entries in a simple file format, usually in the same directory as the GoToSocial binary itself. SQLite is great for small instances and lower-powered machines like Raspberry Pi, where a dedicated database would be overkill.

To configure GoToSocial to use SQLite, change `db-type` to `sqlite`. The `address` setting will then be a filename instead of an address, so you might want to change it to `sqlite.db` or something similar.

Note that the `:memory:` setting will use an *in-memory database* which will be wiped when your GoToSocial instance stops running. This is for testing only and is absolutely not suitable for running a proper instance, so *don't do this*.

## Postgres

Postgres is a heavier database format, which is useful for larger instances where you need to scale performance, or where you need to run your database on a dedicated machine separate from your GoToSocial instance (or do funky stuff like run a database cluster).

GoToSocial supports connecting to Postgres using SSL/TLS. If you're running Postgres on a different machine from GoToSocial, and connecting to it via an IP address or hostname (as opposed to just running on localhost), then SSL/TLS is **CRUCIAL** to avoid leaking data all over the place!

When you're using Postgres, GoToSocial expects whatever you've set for `db-user` to already be created in the database, and to have ownership of whatever you've set for `db-database`.

For example, if you set:

```text
db:
  [...]
  user: "gotosocial"
  password: "some_really_good_password"
  database: "gotosocial"  
```

Then you should have already created database `gotosocial` in Postgres, and given ownership of it to the `gotosocial` user.

The psql commands to do this will look something like:

```psql
create database gotosocial;
create user gotosocial with password 'some_really_good_password';
grant all privileges on database gotosocial to gotosocial;
```

## Settings

```yaml
############################
##### DATABASE CONFIG ######
############################

# Config pertaining to the Gotosocial database connection

# String. Database type.
# Options: ["postgres","sqlite"]
# Default: "postgres"
db-type: "postgres"

# String. Database address or parameters.
# Examples: ["localhost","my.db.host","127.0.0.1","192.111.39.110",":memory:"]
# Default: "localhost"
db-address: "127.0.0.1"

# Int. Port for database connection.
# Examples: [5432, 1234, 6969]
# Default: 5432
db-port: 5432

# String. Username for the database connection.
# Examples: ["mydbuser","postgres","gotosocial"]
# Default: "postgres"
db-user: "postgres"

# REQUIRED
# String. Password to use for the database connection
# Examples: ["password123","verysafepassword","postgres"]
# Default: "postgres"
db-password: "postgres"

# String. Name of the database to use within the provided database type.
# Examples: ["mydb","postgres","gotosocial"]
# Default: "postgres"
db-database: "postgres"

# String. Disable, enable, or require SSL/TLS connection to the database.
# If "disable" then no TLS connection will be attempted.
# If "enable" then TLS will be tried, but the database certificate won't be checked (for self-signed certs).
# If "require" then TLS will be required to make a connection, and a valid certificate must be presented.
# Options: ["disable", "enable", "require"]
# Default: "disable"
db-tls-mode: "disable"

# String. Path to a CA certificate on the host machine for db certificate validation.
# If this is left empty, just the host certificates will be used.
# If filled in, the certificate will be loaded and added to host certificates.
# Examples: ["/path/to/some/cert.crt"]
# Default: ""
db-tls-ca-cert: ""
```
