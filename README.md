# GoOneRoster

This project aims to implement a OneRoster compliant RESTful API
webserver in GO with a MongoDB backing for persistance

A detailed overview of the OneRoster specification can be found on the
[OneRoster site](https://www.imsglobal.org/oneroster-v11-final-specification)
, including: 
* filtering
* data structures
* endpoints
* JSON bindings

This API server attempts to extend the oneroster spec by allowing
updating/insert/upserting/PUT content to all endpoints rather than just providing
a read/GET interface for data.

## Companion projects

Assisting this project is a collection of open tools for syncing
to and from various 3rd party SIS/MIS/systems

To: 
* [Microsoft Teams](https://github.com/the-glasgow-academy/oneroster-api-to-csv-sds)
* [Apple School Manager](https://github.com/the-glasgow-academy/oneroster-api-to-csv-asm)

From:
* [WCBS PASS](https://github.com/fffnite/go-oneroster-sis-sync)

Credential Management:
* [gor-cli](https://github.com/fffnite/gor-cli)

With more to come. Further community support is welcome and encouraged.

## Download

Pre-build binaries for windows and linux x64 are available in 
the releases section as well as a pre-built docker image is 
available:

`docker pull docker.pkg.usulroster/usulroster:0.3.1`

## Setup

Start up the api server with either envs or flags

### flags

```
usulroster \
    -k "mySecretKey" \
    -a "HS256" \
    -m "mongodb://myinstance.domain.com:27017" \
    -p "443"
```

### envs

```
USULROSTER_AUTH_KEY='mySecretKey'
USULROSTER_AUTH_KEY_ALG='HS256'
USULROSTER_MONGO_URI='mongodb://myinstance.domain.com:27017'
USULROSTER_PORT='443'
```

### Upload Data

Import data using the 
[sis sync tool](https://github.com/fffnite/go-oneroster-sis-sync)
or by making POST requests.


## Query examples

Bash
```bash
# login
token=$(curl "myserver.domain.com/ims/oneroster/v1p1/login" \
    -X POST \
    -d "clientid=$ci&clientsecret=$cs")

# Remove quotes from token
t=${token:1:-1}

# Upsert user id 1
curl "myserver.domain.com/ims/oneroster/v1p1/users/1" \
    -X PUT \
    -H "Authorization: Bearer $t"
    -d '{"sourcedId": "1", "status": "active", "givenName": "bob"}'

# Get all users
curl "myserver.domain.com/ims/oneroster/v1p1/users" \
    -H "Authorization: Bearer $t"
```

PowerShell
```powershell
# login
$args = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/login"
    method = "POST"
    body = @{ "clientid" = $ci; "clientsecret" = $cs }
}
$t = Invoke-RestMethod @args

# Upsert user id 1
$upsert = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/users/1
    method = "PUT"
    headers = @{ "authorization" = "bearer $t"}
    body = "{""sourcedId"": ""1"", ""status"": ""active"", ""givenName"": ""bob""}"
}
Invoke-RestMethod @upsert

$getUsers = @{
    uri = "http://myserver.domain.com/ims/oneroster/v1p1/users"
    headers = @{ "authorization" = "bearer $t" }
    FollowRelLink = $true
}
Invoke-RestMethod @getUsers
```

## Endpoints

Listed are all the currently supported endpoints. Details of 
the fields supported/output by these endpoints can be found
on the OneRoster API 
[docs](https://www.imsglobal.org/oneroster-v11-final-specification#_Toc480452033)

```
POST /login

GET /orgs
GET /orgs/{id}
PUT /orgs/{id}

GET /academicSessions
GET /academicSessions/{id}
PUT /academicSessions/{id}

GET /courses
GET /courses/{id}
PUT /courses/{id}

GET /classes
GET /classes/{id}
PUT /classes

GET /enrollments
GET /enrollments/{id}
PUT /enrollments/{id}

GET /users
GET /users/{id}
PUT /users/{id}
PUT /users/{id}/userIds/{subId}
```

## To Do
- [x] Connect to SQL database
- [x] Read from Conf.hcl
- [x] Build DB handler
- [x] Connect packages
- [x] Output from DB
- [x] Output JSON
- [x] Output correct json from single endpoint sample
- [x] Implement RESTful parameters & operators
- [x] Implement logging
- [x] Implement error handling
- [x] Implement security
- [x] Build core endpoints
- [x] Sync DB
- [ ] Build extra endpoints
- [ ] Test test test
