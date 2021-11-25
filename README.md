# students.go

## Introduction

Save, Load, Delete and Update students list with Go

## Requests

- create new student

```sh
curl -vvv -X POST -d '{ "first_name": "Parham", "id": "98243032", "last_name": "Alvani" }' -H 'Content-Type: application/json' 127.0.0.1:1373/student
```

- get students list

```sh
curl -vvv 127.0.0.1:1373/student
```

- get student by id

```sh
curl -vvv 127.0.0.1:1373/student/98243032
```
