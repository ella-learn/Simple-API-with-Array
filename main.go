package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserFormatter struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FormatUser(user User) UserFormatter {
	var userFormatter UserFormatter

	userFormatter.Id = user.Id
	userFormatter.Name = user.Name
	userFormatter.Email = user.Password
	userFormatter.Password = user.Password

	return userFormatter
}

func FormatUsers(user []User) []UserFormatter {
	usersFormatter := []UserFormatter{}

	for _, user := range user {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}

var users []User

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		res := Response{}
		je := json.NewEncoder(w)

		switch r.Method {
		case "GET":
			var formatData []UserFormatter

			if len(users) == 0 {
				res.Data = nil
				res.Message = "users not found"
				res.Status = http.StatusNotFound

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			formatData = FormatUsers(users)

			res.Data = formatData
			res.Message = "success get users"
			res.Status = http.StatusOK

			w.Header().Set("Content-Type", "application/json")
			je.Encode(res)
			return
		}
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
		jd := json.NewDecoder(r.Body)

		res := Response{}
		je := json.NewEncoder(w)

		switch r.Method {
		case "GET":
			formatData := UserFormatter{}

			if id == 0 {
				res.Data = nil
				res.Message = "can't get the id"
				res.Status = http.StatusBadRequest

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			if len(users) == 0 || id > len(users)-1 {
				res.Data = nil
				res.Message = "user not found"
				res.Status = http.StatusNotFound

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			formatData = FormatUser(users[id-1])

			res.Data = formatData
			res.Message = "success get user"
			res.Status = http.StatusOK

			w.Header().Set("Content-Type", "application/json")
			je.Encode(res)
			return

		case "PUT":
			var input User
			formatData := UserFormatter{}

			if id == 0 {
				res.Data = nil
				res.Message = "can't get the id"
				res.Status = http.StatusBadRequest

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			if err := jd.Decode(&input); err != nil {
				res.Data = nil
				res.Message = "can't get the id"
				res.Status = http.StatusBadRequest

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			if len(users) == 0 || id > len(users)-1 {
				res.Data = nil
				res.Message = "user not found"
				res.Status = http.StatusNotFound

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			users[id-1].Name = input.Name
			users[id-1].Email = input.Email
			users[id-1].Password = input.Password

			formatData = FormatUser(users[id-1])

			res.Data = formatData
			res.Message = "success update user"
			res.Status = http.StatusOK

			w.Header().Set("Content-Type", "application/json")
			je.Encode(res)
			return

		case "DELETE":
			var formatData []UserFormatter

			if id == 0 {
				res.Data = nil
				res.Message = "can't get the id"
				res.Status = http.StatusBadRequest

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			if len(users) == 0 || id > len(users)-1 {
				res.Data = nil
				res.Message = "user not found"
				res.Status = http.StatusNotFound

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			users = append(users[:id-1], users[id:]...)

			formatData = FormatUsers(users)

			res.Data = formatData
			res.Message = "success delete user"
			res.Status = http.StatusOK

			w.Header().Set("Content-Type", "application/json")
			je.Encode(res)
			return
		}

	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		jd := json.NewDecoder(r.Body)
		formatData := UserFormatter{}
		je := json.NewEncoder(w)
		res := Response{}

		switch r.Method {
		case "POST":
			var input User

			if err := jd.Decode(&input); err != nil {
				res.Data = nil
				res.Message = "can't get the input"
				res.Status = http.StatusBadRequest

				w.Header().Set("Content-Type", "application/json")
				je.Encode(res)
				return
			}

			if len(users) == 0 {
				input.Id = 1
			} else {
				input.Id = users[len(users)-1].Id + 1
			}

			users = append(users, input)

			formatData = FormatUser(users[len(users)-1])

			res.Data = formatData
			res.Message = "success create user"
			res.Status = http.StatusOK

			w.Header().Set("Content-Type", "application/json")
			je.Encode(res)
			return
		}
	})

	fmt.Println("server started!")
	http.ListenAndServe(":8000", nil)
}
