# golang-mini-exercise

## Prerequisite
- install go
- install mysql

## How to run
- execute db.sql
- inside directory, run ``cp .env.example .env`` on terminal
- set your env
- after cloning this repo, run ``go build``
- run ``go run main.go``

## Notes
- after hitting endpoint ``api/v1/init``, stick to use token from response payload for api token authorization header
- error response still not match the expected
