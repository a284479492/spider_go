package persist

import (
	"crawler/types"
	"log"

	"github.com/olivere/elastic/v7"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item types.Item, result *string) error {
	err := Save(item, s.Client, s.Index)
	log.Printf("Item %v saved.\n", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item: %v.\n", item)
	}
	return err
}
