package types

type Response[T any] struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
	Body   T      `json:"body,omitempty"`
}
