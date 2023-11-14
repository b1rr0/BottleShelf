package controllers

import (
	"context"
	"errors"
	"net/http"

	"users_service/models"
	"users_service/persistence"
	"users_service/serialization"

	"github.com/google/uuid"
)

type OrgController struct {
	persister persistence.Persister
}

func (controller *OrgController) CreateOrg(ctx context.Context, req models.CreateOrgRequest) (resp models.CreateOrgResponse, err error) {
	res, err := controller.persister.SearchUserId(ctx, req.OwnerId)
	if err != nil {
		return
	}
	if res {
		err = errors.New("Owner does not exist")
		return
	}
	org := models.Org{}
	org.Id = uuid.New()
	org.Name = req.Name
	id, err := controller.persister.CreateOrg(ctx, org, req.OwnerId)
	if err != nil {
		return
	}
	resp = models.CreateOrgResponse{Id: id}
	err = nil
	return
}

func (controller *OrgController) GetAllOrgs(ctx context.Context) (resp models.GetOrgsResponse, err error) {
	orgs, err := controller.persister.GetAllOrgs(ctx)
	if err != nil {
		return
	}
	orgnames := []string{}
	for _, org := range orgs {
		orgnames = append(orgnames, org.Name)
	}
	resp = models.GetOrgsResponse{Orgnames: orgnames}
	err = nil
	return
}

func NewOrgController(persister persistence.Persister) *OrgController {
	controller := new(OrgController)
	controller.persister = persister
	return controller
}

/*
API ENDPOINTS
*/

//  @Summary      Create a new organization
//  @Accept       json
//  @Produce      json
//  @Param        user    body    models.CreateOrgRequest    true     "Create user"
//  @Failure      400     body    nil                                 "Bad request"
//  @Success      200     body    models.CreateOrgResponse            ""
//  @Router       /orgs/create [post]
func (controller *OrgController) ServeCreateOrg(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CreateOrgRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := controller.CreateOrg(req.Context(), request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}

//  @Summary      Get all usernames
//  @Produce      json
//  @Failure      400     body    nil                                "Bad request"
//  @Success      200
//  @Router       /orgs [get]
func (controller *OrgController) ServeAllOrgnames(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CheckUserRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := controller.GetAllOrgs(req.Context())
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}
