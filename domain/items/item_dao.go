package items

import (
	"errors"
	"github.com/JingdaMai/bookstore_items-api/clients/elasticsearch"
	"github.com/JingdaMai/bookstore_utils-go/rest_errors"
)

const indexItem = "items"

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
