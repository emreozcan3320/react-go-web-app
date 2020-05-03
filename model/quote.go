package model

// Quote Struct (Model)
type Quote struct{
	Key string `datastore:"key" json:"key"`
	Quote	string `datastore:"quote" json:"quote"`
	Reference string `datastore:"reference" json:"reference"`
	Owner string `datastore:"owner" json:"owner"`
	Created string `datastore:"created" json:"created"`
}