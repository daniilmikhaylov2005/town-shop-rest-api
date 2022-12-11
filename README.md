# town-shop-rest-api

You must have installed go and postgresql

## Install dependencies

go mod tidy

## Run the server

go run main.go

## Important

1) Create .env file

```yaml
SIGN_KEY = YOU_SIGNING_KEY
```

2) Change the db url in function getConnection from repository/common.go

```yaml
"postgres://user:password@localhost:5432/dbname?sslmode=disable"
```

## How to use migrations

this article describes in detail

<https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g>

# Endpoint

# *Work with users*

## /api/goods/:category

**Method: GET**

**return**

```yaml
[
    {
        "id": 3,
        "name": "book1",
        "description": "lorem lorem lorem",
        "image": "url/to/book1.jpg",
        "category": "books"
    },
    {
        "id": 4,
        "name": "book2",
        "description": "lorem2 lorem2 lorem2",
        "image": "url/to/bool2.jpg",
        "category": "books"
    }
]
```

## /api/goods/:category/:id

**Method: GET**

**return**

```yaml
{
        "id": 3,
        "name": "book1",
        "description": "lorem lorem lorem",
        "image": "url/to/book1.jpg",
        "category": "books"
}
```

# *Work with admin panel*

## /auth/signup

**Method: POST**

**take**

```yaml
{
  "name": "name",
  "username": "username",
  "password": "password"
}
```

**return**

```yaml
    "status": "id",
```

## /auth/signin

**Method: POST**

**take**

```yaml
{
  "username": "username",
  "password": "password"
}
```

**return**

```yaml
{
  "token": "token.token.token"
}
```

## /admin/good

**Method: POST**

To create a good, the category must be in the category table

**take**

```yaml
{
  "name": "name",
  "category": "category",
  "description": "description"
}
```

**return**

```yaml
{ 
  "id": 3, 
  "name": "name",
  "description": "description",
  "image": "url/to/image.jpg",
  "category": "category"
}
```

## /admin/good/:id

**Method: PUT**



**take**

```yaml
{
  "name?": "name",
  "category?": "category",
  "description?": "description"
}
```

**return**

```yaml
{ 
  "id": 3, 
  "name": "name",
  "description": "description",
  "image": "url/to/image.jpg",
  "category": "category"
}
```

## /admin/good/:id

**Method: DELETE**

**take**

Only id in param

**return**

```yaml
{
  "status": "good with id # deleted"
}
```
