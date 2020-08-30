package controllers

import (
	"github.com/JingdaMai/bookstore_items-api/domain/items"
	"github.com/JingdaMai/bookstore_items-api/services"
	"github.com/JingdaMai/bookstore_oauth-go/oauth"
	"log"
	"net/http"
)

var ItemsController itemsControllerInterface = &itemsController{}

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// todo: return error to the caller.
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// todo: return error json to the caller
	}

	log.Println(result)
	// todo: return created item as json with http status 201 - created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	return
}
