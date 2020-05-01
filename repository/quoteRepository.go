package repository

import(
	"log"
	"strconv"
	"context"
	"cloud.google.com/go/datastore"
	"github.com/emre/react-golang-web-app/model"
)

// GetAllQuotes returns all the quotes in ascending order of creation time.
func GetAllQuotes(ctx context.Context, client *datastore.Client) ([]*model.Quote, error) {
	var quotes []*model.Quote
	query := datastore.NewQuery("Quote").Order("created")
	keys, err := client.GetAll(ctx, query, &quotes)
	if err != nil {
		return nil, err
	}
	// Set the id field on each Task from the corresponding key.
	for i, key := range keys{
		quotes[i].Key = strconv.FormatInt(key.ID, 10)
	}	
	return quotes, nil
}

//GetSingleQuote returns a quote 
func GetSingleQuote(ctx context.Context, client *datastore.Client, stringQuoteID string) (*model.Quote, error) {
	var quote model.Quote	
	int64QuoteID, convErr := strconv.ParseInt(stringQuoteID, 10, 64)
	if convErr != nil {
		log.Println(convErr)
		return nil, convErr
	}
	quoteKey := datastore.IDKey("Quote", int64QuoteID, nil)
	if err := client.Get(ctx, quoteKey, &quote); err != nil {
		return nil, err
	}
	quote.Key = stringQuoteID
	return  &quote, nil
}

// CreateQuote adds a quote with the given description to the datastore
// returning the key of the newly created entity.
func CreateQuote(ctx context.Context, client *datastore.Client, data model.Quote) (*datastore.Key, error) {
	quoteKey := datastore.IncompleteKey("Quote", nil)
	key, err := client.Put(ctx, quoteKey, &data)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// UpdateQuote adds a quote with the given description to the datastore
// returning the key of the newly created entity.
func UpdateQuote(ctx context.Context, client *datastore.Client, data model.Quote) (*datastore.Key, error) {
	int64QuoteID, parseErr := strconv.ParseInt(data.Key, 10, 64)
	if parseErr != nil {
		log.Println(parseErr)
		return nil, parseErr
	}
	quoteKey := datastore.IDKey("Quote", int64QuoteID, nil)
	key, err := client.Put(ctx, quoteKey, &data)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// DeleteTask deletes the task with the given ID.
func DeleteTask(ctx context.Context, client *datastore.Client, stringQuoteID string) error {
	int64QuoteID, parseErr := strconv.ParseInt(stringQuoteID, 10, 64)
	if parseErr != nil {
		log.Println(parseErr)
		return parseErr
	}
	return client.Delete(ctx, datastore.IDKey("Quote", int64QuoteID, nil))
}