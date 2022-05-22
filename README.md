# Messages app

## How to use:
```bash
go run app.go
```
## API:

`POST /messages`
### Request body:
```json
{
  "sender_name": "name1",
  "receiver_name": "name2",
  "message_body": "This is a message about something"
}
```

`GET /messages`
### Response body:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2022-05-22T11:45:43.607994+03:00",
        "UpdatedAt": "2022-05-22T13:28:06.655854+03:00",
        "DeletedAt": null,
        "sender_id": 3,
        "receiver_id": 2,
        "message_body": "Hello world"
    },
    {
        "ID": 2,
        "CreatedAt": "2022-05-22T17:36:19.796904+03:00",
        "UpdatedAt": "2022-05-22T17:36:19.796904+03:00",
        "DeletedAt": null,
        "sender_id": 3,
        "receiver_id": 1,
        "message_body": "Another hello world"
    }
]
```
`PATCH /messages/:id`
### Request body:
```json
{
  "message_body": "Updated text of the message"
}
```

`DELETE /messages/:id`
### Deletes the message

`POST /users`
### Request body:
```json
{
  "name": "John Wick"
}
```
`GET /users`
### Response body:
```json
[
  {
    "ID": 1,
    "CreatedAt": "2022-05-21T11:03:56.40738+03:00",
    "UpdatedAt": "2022-05-22T13:01:17.328622+03:00",
    "DeletedAt": null,
    "name": "John Cena"
  },
  {
    "ID": 2,
    "CreatedAt": "2022-05-21T11:04:07.864172+03:00",
    "UpdatedAt": "2022-05-22T17:37:25.474983+03:00",
    "DeletedAt": null,
    "name": "John Smith"
  },
  {
    "ID": 3,
    "CreatedAt": "2022-05-21T11:04:14.966229+03:00",
    "UpdatedAt": "2022-05-21T11:04:14.966229+03:00",
    "DeletedAt": null,
    "name": "John Johnson"
  }
]
```
`PATCH /users/:id`
### Request body:
```json
{
  "name": "Not a John"
}
```
`DELETE /users/:id`
### Deletes the user