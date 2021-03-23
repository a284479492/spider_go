package persist

import (
	"context"
	"crawler/types"
	"errors"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan types.Item, error) {
	out := make(chan types.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return out, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: get item #%d:%v\n", itemCount, item)
			itemCount++

			err = Save(item, client, index)
			if err != nil {
				log.Printf("ItemSaver ERROR saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}

func Save(item types.Item, client *elastic.Client, index string) error {
	if item.Type == "" {
		return errors.New("Must supply Type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
