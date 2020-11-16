docker run --rm --name pg-docker -e POSTGRES_PASSWORD=$1 -d -p 5432:5432 -v ~/datadir/postgres:/var/lib/postgresql/data postgres:13.0-alpine
