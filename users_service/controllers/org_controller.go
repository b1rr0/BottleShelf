package controllers

import (
	"context"
	"net/http"

	"users_service/models"
	"users_service/persistence"
	"users_service/resources"
	"users_service/serialization"

	"github.com/google/uuid"
)

type OrgController struct {
	persister persistence.Persister
}

func NewOrgController(persister persistence.Persister) *OrgController {
	controller := new(OrgController)
	controller.persister = persister
	return controller
}

func (controller *OrgController) CreateOrg(ctx context.Context, req models.CreateOrgRequest) (resp models.CreateOrgResponse, serr models.ServiceError) {
	res, err := controller.persister.SearchUserId(ctx, req.OwnerId)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	if !res {
		return resp, models.NewServiceError(http.StatusNotFound, resources.OwnerNotExist)
	}
	org := models.Org{}
	org.Id = uuid.New()
	org.Name = req.Name
	id, err := controller.persister.CreateOrg(ctx, org, req.OwnerId)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	return models.CreateOrgResponse{Id: id}, models.NoError()
}

func (controller *OrgController) GetAllOrgs(ctx context.Context) (resp models.GetOrgsResponse, serr models.ServiceError) {
	orgs, err := controller.persister.GetAllOrgs(ctx)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	orgnames := []string{}
	for _, org := range orgs {
		orgnames = append(orgnames, org.Name)
	}
	return models.GetOrgsResponse{Orgnames: orgnames}, models.NoError()
}

/*
API ENDPOINTS
*/

//  @Summary      Create a new organization
//  @Accept       json
//  @Produce      json
//  @Param        user    body        models.CreateOrgRequest    true     "Create user"
//  @Failure      400     {object}    models.ErrorResponse
//  @Success      200     {object}    models.CreateOrgResponse
//  @Router       /orgs/create [post]
func (controller *OrgController) ServeCreateOrg(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CreateOrgRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		serialization.SerializeError(writer, status, resources.SerializationFailed)
		return
	}
	// Call controller
	resp, serr := controller.CreateOrg(req.Context(), request)
	if !serr.IsOk() {
		serialization.SerializeServiceError(writer, serr)
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}

//  @Summary      Get all organization names
//  @Produce      json
//  @Failure      400     {object}    models.ErrorResponse
//  @Success      200     {object}    models.GetOrgsResponse
//  @Router       /orgs [get]
func (controller *OrgController) ServeAllOrgnames(writer http.ResponseWriter, req *http.Request) {
	// Call controller
	resp, serr := controller.GetAllOrgs(req.Context())
	if !serr.IsOk() {
		serialization.SerializeServiceError(writer, serr)
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}
