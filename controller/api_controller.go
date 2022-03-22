package controller

import(
	"encoding/json"
	"net/http"
	"strconv"
	"path"

	"github.com/takkiiiiiiiii/rest-api/controller/dto"
	"github.com/takkiiiiiiiii/rest-api/model/repository"
	"github.com/takkiiiiiiiii/rest-api/model/entity"
	
)

type ApiContoller interface {
     GetApi(w http.ResponseWriter, r *http.Request)
	 PostApi(w http.ResponseWriter, r *http.Request)
	 UpdateApi(w http.ResponseWriter, r *http.Request)
	 DeleteApi(w http.ResponseWriter, r *http.Request)
}

type apiController struct {
     tr repository.ApiRepository
}


func NewApiController(tr repository.ApiRepository) *apiController {
	return &apiController{tr}
}

func (tc *apiController) GetApi(w http.ResponseWriter, r *http.Request) {
	data, err := tc.tr.Get_Api()
	if err != nil {
		w.WriteHeader(500)
		return
	} 
	var apiResponse []dto.ApiResponse
	for _, v := range data {
		apiResponse = append(apiResponse, dto.ApiResponse{Id: v.Id, Name: v.Name, Contents: v.Contents, Created: v.Created})
	}
	var apisResponse dto.ApisResponse
	apisResponse.Api = apiResponse
	output, _ := json.MarshalIndent(apisResponse.Api, "", "\t")

	w.Header().Set("Content-Type","application/json")
	w.Write(output)
}

func (tc *apiController) PostApi(w http.ResponseWriter, r *http.Request){
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var apiRequest dto.ApiRequest
	json.Unmarshal(body, &apiRequest)

	datum := entity.ApiEntity{Name: apiRequest.Name, Contents: apiRequest.Contents}
	id, err := tc.tr.Insert_Api(datum)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	Id := strconv.FormatInt(id, 10)
	w.Header().Set("Location", r.Host + r.URL.Path + Id) // r.URL.Path > /api/users/
	w.WriteHeader(201)
}

func (tc *apiController) UpdateApi(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.Atoi(path.Base(r.URL.Path)) // r.URL.Path > ex:/api/users/2   path.Base(r.URL.Path) > 2
	if err != nil {
        w.WriteHeader(400)
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var apiRequest dto.ApiRequest
	json.Unmarshal(body, &apiRequest)

	datum := entity.ApiEntity{Id: Id, Name: apiRequest.Name, Contents: apiRequest.Contents}
	err = tc.tr.Update_Api(datum)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func (tc *apiController) DeleteApi(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = tc.tr.Delete_Api(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}