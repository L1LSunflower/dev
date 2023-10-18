package builder

import (
	"bytes"
	"net/http"
)

// TODO: REWORK. Make more real world case

type BuildProcess interface {
	SetMethod() BuildProcess
	SetUrl() BuildProcess
	SetBody() BuildProcess
	GetRequest() http.Request
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetMethod().SetUrl().SetBody()
}

type HttpPost struct {
	Method string
	Url    string
	Body   []byte
}

func (p *HttpPost) SetMethod() BuildProcess {
	p.Method = http.MethodPost
	return p
}

func (p *HttpPost) SetUrl() BuildProcess {
	p.Url = "http://test_method_post.com"
	return p
}

func (p *HttpPost) SetBody() BuildProcess {
	p.Body = []byte("{\"data\": \"hello_post\"}")
	return p
}

func (p *HttpPost) GetRequest() http.Request {
	request, _ := http.NewRequest(p.Method, p.Url, bytes.NewReader(p.Body))
	return *request
}

type HttpGet struct {
	Method string
	Url    string
	Body   []byte
}

func (p *HttpGet) SetMethod() BuildProcess {
	p.Method = http.MethodGet
	return p
}

func (p *HttpGet) SetUrl() BuildProcess {
	p.Url = "http://test_method_get.com"
	return p
}

func (p *HttpGet) SetBody() BuildProcess {
	p.Body = []byte("{\"data\": \"hello_get\"}")
	return p
}

func (p *HttpGet) GetRequest() http.Request {
	request, _ := http.NewRequest(p.Method, p.Url, bytes.NewReader(p.Body))
	return *request
}
