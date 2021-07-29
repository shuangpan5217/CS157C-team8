# This repo is created for final project of CS157C taught by Prof. Ching-seh Wu in San Jose State University.

Project Name:
Secret Box

Team Members: 
1. Yunlin Xie
2. Shuang Pan

### Run APIs locally
From the main directory which contains `main.go`
```
1. go build .
2. ./CS157C-TEAM8
```
Then the API service is listening to the `localhost:3000`

### Dockerize Golang APIs and Cassaandra

Make sure you have docker installed and `make` command on your laptop.
Run the following command from the same directly like above
```
make docker
```
Then run `docker ps`, you will see something like 
```
CONTAINER ID   IMAGE              COMMAND                  CREATED          STATUS          PORTS                                                                          NAMES
47bd82b19097   secretbox-api:v1   "./CS157C-TEAM8"         3 seconds ago    Up 3 seconds    0.0.0.0:3000->3000/tcp, :::3000->3000/tcp                                      relaxed_blackwell
74655547da58   cassandra          "docker-entrypoint.s…"   28 minutes ago   Up 34 seconds   7000-7001/tcp, 7199/tcp, 9160/tcp, 0.0.0.0:9042->9042/tcp, :::9042->9042/tcp   cassandra
```
If you are not able to see `secretbox-api:v1`, run the following command and run `docker ps` again. You will see that.
```
docker run -p 3000:3000 -d -e CASSANDRA_URL=cassandra:9042 --link=cassandra secretbox-api:v1
```
If it still doesn't work. Run the following commands one by one.
```
1. docker run -p 9042:9042 -d --name cassandra cassandra
2. docker restart cassandra
3. docker build -t secretbox-api:v1 .
4. docker run -p 3000:3000 -d -e CASSANDRA_URL=cassandra:9042 --link=cassandra secretbox-api:v1
```
After that, you can use API services from `localhost:3000`

### Introduction to APIs
#### User API
- User signup request body `POST 127.0.0.1:3000/login?signup=true`
```
{
    "username": "Tian",
    "password": "pass",
    "nickname": "nonono", // Optional
    "description": "" // Optional
}
```
Successful Response
```
{
    "Message": "success",
    "SuccessMessage": "Successfully create an user.",
    "StatusCode": 201,
    "Body": {
        "description": "",
        "nickname": "nonono",
        "username": "Tian"
    }
}
```
username and password are required. If not provided, errors will be returned.
```
{
    "Message": "failure",
    "Error": "password is required in request Body",
    "StatusCode": 400
}
```
or
```
{
    "Message": "failure",
    "Error": "username is required in request Body",
    "StatusCode": 400
}
```
- User login request body `POST 127.0.0.1:3000/login`
```
{
    "username": "Tian",
    "password": "pass"
}
```
Successful response
```
{
    "Message": "success",
    "SuccessMessage": "Successfully log in.",
    "StatusCode": 200,
    "Body": {
        "description": "",
        "nickname": "nonono",
        "username": "Tian"
    }
}
```
If username or password is wrong, error response
```
{
    "Message": "failure",
    "Error": "username or password is not correct.",
    "StatusCode": 401
}
```
or 
```
{
    "Message": "failure",
    "Error": "password is required in request Body",
    "StatusCode": 400
}
```
- Update user's nickname or description `PATCH 127.0.0.1:3000/updateuser`

request body: password is not required.
```
{
    "username": "Tian", // required
    "nickname": "new nickname",
    "description": "new description"
}
```
Successful response
```
{
    "Message": "success",
    "SuccessMessage": "Successfully updated.",
    "StatusCode": 200,
    "Body": {
        "description": "new description",
        "nickname": "new nickname",
        "username": "Tian"
    }
}
```
The only required field in the request body is `username`. If no username is present or a non-existing username,
```
{
    "Message": "failure",
    "Error": "user doesn't exist.",
    "StatusCode": 401
}
```
#### Secret API
- Add a secret `POST 127.0.0.1:3000/secret`
```
{
    "username": "pan", // required
    "nickname": "nonono", // required
    "content": "content"
}
```
Successful response
```
{
    "Message": "success",
    "SuccessMessage": "Your secret have been added to the secret box!",
    "StatusCode": 201,
    "Body": {
        "secret_id": "cf6a3dce-bd4d-4809-bc00-5b1e158c12d2"
    }
}
```
Error response if secret content is empty
```
{
    "Message": "failure",
    "Error": "Empty secret content is not allowed.",
    "StatusCode": 400
}
```
username and nickname don't match
```
{
    "Message": "failure",
    "Error": "username or nickname is not correct.",
    "StatusCode": 400
}
```
- Get a secret `GET 127.0.0.1:3000/secret?username=pan` 

The `username=pan` is the user who wants to get a secret.
The user will not get the secret that he or she posted.

Successful response
```
{
    "Message": "success",
    "SuccessMessage": "Successfully get a secret",
    "StatusCode": 200,
    "Body": {
        "content": "content",
        "created_time": "2021-07-16 05:03:13.983 +0000 UTC",
        "nickname": "nonono",
        "secret_id": "c9784de0-fd7b-432b-8f02-ffc687e4d113",
        "username": "shuang"
    }
}
```
If there is no secret in the database besides himself or herself.

Error response
```
{
    "Message": "failure",
    "Error": "No more secrets, please try again later.",
    "StatusCode": 500
}
```
- Delete a secret `DELETE 127.0.0.1:3000/secret?username=pan&&secret_id=fe5c05bf-fae3-445b-aa64-d8c3cebdaa26` 
username is the owner of the secret.
(This api will be called when the user doesn't want to throw it back to secretBox after getting a secret or after save)

Successful response
```
{
    "Message": "success",
    "SuccessMessage": "The secret is successfully deleted.",
    "StatusCode": 200,
    "Body": {
        "secret_id": "fe5c05bf-fae3-445b-aa64-d8c3cebdaa26"
    }
}
```
Error response
```
{
    "Message": "failure",
    "Error": "secret doesn't exist.",
    "StatusCode": 404
}
```
#### SavedSecret API
When a user fetch a secret from SecretBox , he or she may want to save it to favorite list.
At the same time, the secret will be removed from the SecretBox.

- Save a secret (All fields are required) `POST 127.0.0.1:3000/savedsecret`
```
{
    "secret_id": "57fd26f7-ae2f-43cc-a440-b3ba2f95b581",
    "secret_owner": "shuang",
    "nickname": "nonono",
    "username": "pan",
    // front-end will take care of this field because it is not the same with the response
    // need some format conversion
    "created_time": "2021-07-13T06:07:33.214Z",
    "content": "what is content"
}
```
Successful response (Body doesn't need those fields, may be modified in the future).
```
{
    "Message": "success",
    "SuccessMessage": "Successfully added to your favorite list.",
    "StatusCode": 200,
    "Body": {
        "content": "what is content",
        "created_time": "2021-07-13 06:07:33.214 +0000 UTC",
        "nickname": "nonono",
        "secret_id": "57fd26f7-ae2f-43cc-a440-b3ba2f95b581",
        "secret_owner": "shuang",
        "username": "pan"
    }
}
```
Sample error response
```
{
    "Message": "failure",
    "Error": "secret owner's username or nickname is not correct.",
    "StatusCode": 400
}
```
- Show all saved secrets in the list. `GET 127.0.0.1:3000/savedsecret?username=pan`

Successful response. The body is an array.
```
{
    "Message": "success",
    "SuccessMessage": "All saved secrets",
    "StatusCode": 200,
    "Body": {
        "saved_secrets": [
            {
                "secret_id": "57fd26f7-ae2f-43cc-a440-b3ba2f95b581",
                "secret_owner": "",
                "username": "pan",
                "content": "what is content",
                "nickname": "nonono",
                "created_time": "2021-07-13T06:07:33.214Z"
            }
        ]
    }
}
```
Error response
```
{
    "Message": "failure",
    "Error": "username doesn't exist",
    "StatusCode": 404
}
```
- Remove saved secrets 

One day, if users wants to remove it from the favorite list, they have two options.

1.Delete it.

`DELETE 127.0.0.1:3000/savedsecret?username=pan&&secret_id=9041de96-0874-47f0-8de3-1ae2c40d59d9`
Successful response
```
{
    "Message": "success",
    "SuccessMessage": "Successfully removed from your favorite list.",
    "StatusCode": 200,
    "Body": {
        "secret_id": "9041de96-0874-47f0-8de3-1ae2c40d59d9"
    }
}
```
Error response
```
{
    "Message": "failure",
    "Error": "Saved secret Not found",
    "StatusCode": 404
}
```

2.Throw back to the SecretBox
`DELETE 127.0.0.1:3000/savedsecret?throwback=true&&username=pan&&secret_id=9041de96-0874-47f0-8de3-1ae2c40d59d9` 

Request Body and response will be the same like above, but the secret will be added to SecretBos again.
We can see this from Cassandra.
![image](https://user-images.githubusercontent.com/24786635/125909765-72ea1e33-d552-4f26-8c9c-5c14008f8956.png)

#### Comment API (TO BE DECIDED)
