package service

import (
	"errors"
	"fmt"
	"lgo/tools/gorm_server"
	"lgo/tools/gorm_server/model"
	"lgo/tools/gorm_server/utils"
	"net/http"
)

type CreatePeopleReq struct {
	*model.People
}

func CreateData(m *gorm_server.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreatePeopleReq
		var err error
		err = utils.UnmarshalBody(r.Body, &req)
		if err != nil {
			err = errors.New(fmt.Sprintf("request param err:%s", err.Error()))
		}
		//fmt.Printf("body: %+v\r\n", req)
		var p model.People
		p.Name = req.Name
		p.Password = req.Password
		fmt.Printf("p: %+v\r\n", p)
		err = m.PeopleRepo.CreateData(&p)
	})

}

type UpdateReq struct {
	*CreatePeopleReq
}

func UpdatePeople(m *gorm_server.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req UpdateReq
		var err error
		err = utils.UnmarshalBody(r.Body, &req)
		if err != nil {
			err = errors.New(fmt.Sprintf("request param err:%s", err.Error()))
		}
		err = m.PeopleRepo.UpdateData(req.People)
	})
}

type DeleteReq struct {
	*CreatePeopleReq
}

func DeletePeople(m *gorm_server.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var mr map[string]interface{}
		err = utils.UnmarshalMap(r.Body, &mr)
		if err != nil {
			err = errors.New(fmt.Sprintf("request param err:%s", err.Error()))
		}
		err = m.PeopleRepo.DeletePeople(mr)
	})

}

func GetPeople(m *gorm_server.Manager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Println("id: ", id)
	})
}
