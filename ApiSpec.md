# twittor api specification
## GET /api/session
login

### request
#### name
specifies a user login id. up to maximum of 255.
this allows for only alphanumeric.
#### password
specifies a user login id. up to maximum of 255.
this allows for alphanumeric and some symbol(inclide *,!,#,$,%,&).

### response
nil

## DELETE /api/session
logout

### request
nil

### response
nil

## POST /api/user
create new user

### request
#### name
specifies a new user login id. up to maximum of 255.
this allows for only alphanumeric.
#### screen_name
specifies a new user screen name. up to maximum of 255
#### password
specifies a new user login id. up to maximum of 36.
this allows for alphanumeric and some symbol(inclide *,!,#,$,%,&).

### response
nil

## GET /api/statuses
### request
#### count
limit of the number of entries.

### response(array)
#### user
user name posted
#### content
status body
#### timestamp
time posted

## POST /api/statuses
### request
#### shout
content body

### response
nil

## GET /api/search
search status

### request
#### q
search query.
#### count
limit of the number of entry.

### response(array)
#### user
user name posted
#### content
status body
#### timestamp
time posted
