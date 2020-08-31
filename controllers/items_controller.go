package controllers

import (
	"encoding/json"
	"github.com/JingdaMai/bookstore_items-api/domain/items"
	"github.com/JingdaMai/bookstore_items-api/services"
	"github.com/JingdaMai/bookstore_items-api/utils/http_utils"
	"github.com/JingdaMai/bookstore_oauth-go/oauth"
	"github.com/JingdaMai/bookstore_utils-go/rest_errors"
	"io/ioutil"
	"net/http"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	// oauth
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}

	// check the json body
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	// unmarshal body to Item struct
	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
