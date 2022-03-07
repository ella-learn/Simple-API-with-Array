## Simple API with Array

1. This is simple API with array as data storage and use `net/http` golang package to build the API. 
2. Run this project with: 
    ```go
    go run main.go
    ```
3. List endpoints : 
    - `POST /user` to create new user data
        <br />
        This endpoints will return a 400 http response status code if there is no input data. 
        <br />
        This endpoints also return a 200 http response status code if the endpoints success input the data in the array.
    - `GET /users` to get all user data 
        <br />
        This endpoints will return a 404 http response status code if there is no data in the array. 
        <br />
        This endpoints also return 200 http response statuscode and data if there is data in the array.
    - `GET /user/{id}` to get spesific user data 
        <br />
        This endpoints will return a 404 http response status code if there is no data with the associate id. 
        <br />
        This endpoints also return 200 http response status code and the user data with the associate if there is data in the array
    - `PUT /user/{id}` to update spesific user data 
        <br />
        This endpoints will return a 404 http response status code if there is no data with the associate id and there is no input.
        <br />
        This endpoints also return a 200 http response status code if the endpoints success update the user data with the associate id in the array.
    - `DELETE /user/{id}` to delete spesific user data
        <br />
        This endpoints will return a 404 http response status code if there is no data in the array
        <br />
        This endpoints also return a 200 http response status code if the endpoints success delete the user data with the associate id in the array
3. Acknowledgement:
    - [Dasar Pemrograman Golang Noval Agung - Golang HTTP Method: GET & POST](https://dasarpemrogramangolang.novalagung.com/B-http-method-basic.html)