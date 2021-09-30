### Features

- support of counter for each user;
- support of counter;
- list of counter for each users;
- used REST-API;
- used secure login with token;
- show all counters/teams ;
- add and remove counters/users on fly;

# Counters. Test for ApsisRecruitment.

# API
```
POST		/login
GET			/logout
GET       	/counter/
PUT       	/counter/
PUT       	/counter/:id/:userId
DELETE 	/counter/:id
DELETE 	/counter/:id/:userId
GET       	/counter/:id
GET       	/users/
PUT       	/users/
GET       	/users/:id/increate
DELETE 	/users/:id
GET      	/users/:id
```
### Login-Logout
```golang
func (l Login) Route(router *gin.Engine) {
	
	var loginController controller.LoginController = controller.LoginHandler()

	router.Use(sessions.Sessions("token", sessions.NewCookieStore([]byte("secret"))))
	
	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			session := sessions.Default(ctx)
			session.Set(USERKEY, token)
			if err := session.Save(); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})


	router.GET("/logout", logout)

}


func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(USERKEY)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(USERKEY)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

```

#### Auth in code
````golang
private := router.Group("/" + ID)
	private.Use(AuthRequired)
	{
		private.GET("/", us.getUsers)
		....
````
#### Parameters
##### Login
````
POST /login
form-data example:
email=NGreen@gmail.com
password=test2pass
````
#### Logout
````
GET /logout
````
#### Counter list
````
GET /counter/
````

#### Create counter
````
PUT /counter/
{"ID": "4", "Name": "Brown"},
````
#### Delete counter
````
DELETE /counter/4
````
#### Add user to counter
````
PUT /counter/1/2 #add user with id 2 to counter with id 1
````
#### Remove user from counter
````
DELETE /counter/1/2 #remove user with id 2 from counter with id 1
````
#### Create user
````
PUT /users/
{"ID": "4", "Name": "Tim Sailor ", "email": "TSailor@gmail.com","password": "test4pass" }
````
#### Delete user
````
DELETE /users/4
````
#### Increate counter
````
GET /users/4/increate
````
# Additional questions
## Persistence
### How would you add a persistent storage layer such that the app could be restarted without losing counter states?
I can add a database (MySQL, PostgreSQL) or storage (clickhouse). 
### What storage technology would you suggest?
Depending on the conditions. In some cases, you can use Elasticsearch for cases where there is a very large amount of data processing.
## Fault tolerance
### How would you design the app in order to make the functionality be available even if some parts of the underlying hardware systems were to fail?
### How would you design the app in order to ensure that it wouldn’t slow down or fail if usage increased by many orders of magnitude? what if this turned into a global contest with 10x, 100x, 1000x the number of teams and traffic?
I would use cluster solutions like Elasticsearch as frondend. For the backend, i want use Mysql (or PostgreSQL) clusters.
### Does your choice of persistence layer change in this scenario?
Yes. For high-load systems, data denormalization is required. Data normalization is commonly used. Because the technologies are very different. 
## Authentication
### How would you ensure that only authorised users can submit and/or retrieve data?
To provide authentication I use middleware for all private token and session based operations. An example is shown in this document above.
### How would you then add support to allow different users to only update specific counters? Or perform only specific operations?
To perform only your own operations, it is enough to use the userId check in the token.
To separate the rights to groups of operations, I would introduce roles for different groups of users.

# Others
(copiright) Not for commercial use only
