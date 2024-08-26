# KP API

## TODO List

- [ ] Add authentication and authorization for account access
- [ ] Implement CRUD management for Admin
- [ ] Implement firebase for storing images
- [ ] Implement full unit testing
- [ ] Add swagger for testing API
- [ ] Implement multi limit logic with percentage, refer: [Tiket.com Paylater Multi Limit](https://www.tiket.com/info/paylater-multi-limit)

## Requirements
- Go >= 1.20
- [Docker](https://docs.docker.com/get-started/get-docker/)
- Docker Compose
- [Goose](https://github.com/pressly/goose) (for database migration)
- [Make](https://www.gnu.org/software/make/#download) (for running scripts)

## How to Run

Using Docker Compose

```bash
# Run
docker compose up -d

# Stop and delete volumes
docker compose down -v
```

Manual

> Make sure `docker` is installed for mysql to be running\
> Make sure `goose` is installed for mysql migration
 
```bash
# Run mysql server using credentials admin:admin/database on port 3306
make mysql

# Check migration status
make migrate-status

# Start migration
make migrate-up

# Reset migration
make migrate-reset

# Run api
make run
```

## Database Diagram

![db diagram](./docs/db-diagram.png)

## Application Flow

![app flow](./docs/app-flow.png)