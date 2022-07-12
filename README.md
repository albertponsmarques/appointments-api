# appointments-api
golang API to create appointments **by albert pons marques**.

##summary
This api is made to store and create appointments, and it's deployed in heroku.

##endpoints
- GET("/") this is the api homepage, it does nothing.
- GET("/appointments") this one return a list of all the appointments in the db. 
- GET("/appointments/{id}") this one has an ID parameter, and return the appointment associated with thhat ID.
- POST("/appointments") this one creates a record passed with the POST.
- DELETE("/appointments/{id}") this one deletes the appointment with the passed id.
- PUT("/appointments/{id}") this last one updates the appointment that has the parametered id with the passed JSON.
