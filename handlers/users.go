package handlers

import (
	"api/handlers/models"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const userCount = 4

var users []models.User

func init() {
	ids := make([]uuid.UUID, 0, userCount)
	for i := 0; i < userCount; i++ {
		id, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
			return
		}
		ids = append(ids, id)
	}

	users = []models.User{
		{
			Id:         ids[0],
			FirstName:  "Пётр",
			Surname:    "Колесников",
			MiddleName: "Олегович",
			FIO:        "Колесников Пётр Олегович",
			Sex:        models.Male,
			Age:        35,
		},
		{
			Id:         ids[1],
			FirstName:  "Татьяна",
			Surname:    "Колесникова",
			MiddleName: "Александровна",
			FIO:        "Колесникова Татьяна Александровна",
			Sex:        models.Female,
			Age:        33,
		},
		{
			Id:         ids[2],
			FirstName:  "Михаил",
			Surname:    "Колесников",
			MiddleName: "Петрович",
			FIO:        "Колесников Михаил Петрович",
			Sex:        models.Male,
			Age:        9,
		},
		{
			Id:         ids[3],
			FirstName:  "Александр",
			Surname:    "Колесников",
			MiddleName: "Петрович",
			FIO:        "Колесников Александр Петрович",
			Sex:        models.Male,
			Age:        5,
		},
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	u, err := json.Marshal(users)
	if err == nil {
		_, err = w.Write(u)
		if err != nil {
			http.Error(w, "write body error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "marshal error", http.StatusInternalServerError)
	}
}
