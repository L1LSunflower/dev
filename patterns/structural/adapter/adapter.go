package adapter

type Adapter interface {
	ToMap() map[string]any
}

type Request struct {
	username   string
	password   string
	permission int
}

func (r *Request) ToMap() map[string]any {
	return map[string]any{
		"username":   r.username,
		"password":   r.password,
		"permission": r.permission,
	}
}

type Response struct {
	id         int
	token      string
	permission int
}

func (r *Response) ToMap() map[string]any {
	return map[string]any{
		"id":         r.id,
		"token":      r.token,
		"permission": r.permission,
	}
}
