package request_structs

type StorePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
