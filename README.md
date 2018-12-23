# Book RestAPI Go
## Implementation of Books REST API with Go Gorilla Mux router

GET : get All Books- localhost:8080/api/books<br/>
Response :
```json
[
    {
        "id": "1",
        "isbn": "448743",
        "title": "Book One",
        "author": {
            "firstname": "Avanish",
            "lastname": "Patel"
        }
    },
    {
        "id": "2",
        "isbn": "448744",
        "title": "Book Two",
        "author": {
            "firstname": "Dixit",
            "lastname": "Patel"
        }
    },
    {
        "id": "3",
        "isbn": "448745",
        "title": "Book Three",
        "author": {
            "firstname": "Tejas",
            "lastname": "Patel"
        }
    }
]
```

GET : get Book by Id - localhost:8080/api/books/{id}<br/>
id = 2<br/>
Respose : 
```json
{
    "id": "2",
    "isbn": "448744",
    "title": "Book Two",
    "author": {
        "firstname": "Dixit",
        "lastname": "Patel"
    }
}
```
POST : create Book - http://localhost:8080/api/books<br/>
Body : 
```json
{
    "id": "98081",
    "isbn": "8548239",
    "title": "Book Random",
    "author": {
        "firstname": "Random fName",
        "lastname": "Random lName"
    }
}
```
Response :
```json
{
    "id": "98081",
    "isbn": "8548239",
    "title": "Book Random",
    "author": {
        "firstname": "Random fName",
        "lastname": "Random lName"
    }
}
```

PUT : update Book by Id - localhost:8080/api/books/{id}<br/>
id = 2<br/>
Body :
```json
{
	"isbn":"8548239",
	"title":"Updated Book",
	"author": {
        "firstname":"Updated fName","lastname":"Updated lName"
    }
}
```
Response :
```json
{
    "id": "2",
    "isbn": "8548239",
    "title": "Updated Book",
    "author": {
        "firstname": "Updated fName",
        "lastname": "Updated lName"
    }
}
```
DELETE : delete book by Id - localhost:8080/api/books/{id} <br/>
id = 2 <br/>
Response : 
```json
[
    {
        "id": "1",
        "isbn": "448743",
        "title": "Book One",
        "author": {
            "firstname": "Avanish",
            "lastname": "Patel"
        }
    },
    {
        "id": "3",
        "isbn": "448745",
        "title": "Book Three",
        "author": {
            "firstname": "Tejas",
            "lastname": "Patel"
        }
    }
]
```





