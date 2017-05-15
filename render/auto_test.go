package render_test

import (
	"bytes"
	"net/http"
	"strings"
	"testing"

	"github.com/gobuffalo/buffalo/render"
	"github.com/stretchr/testify/require"
)

func Test_AutoJsonXml(t *testing.T) {
	r := require.New(t)

	type ji func(v interface{}, req *http.Request) render.Renderer

	type user struct {
		Name string
	}

	table := []ji{
		render.Auto,
		render.New(render.Options{}).Auto,
	}

	for _, j := range table {

		// Test fallback on JSON
		req, _ := http.NewRequest("GET", "http://localhost/", nil)
		re := j(map[string]string{"hello": "world"}, req)
		r.Equal("application/json", re.ContentType())
		bb := &bytes.Buffer{}
		err := re.Render(bb, nil)
		r.NoError(err)
		r.Equal(`{"hello":"world"}`, strings.TrimSpace(bb.String()))

		// Test format query argument JSON
		req, _ = http.NewRequest("GET", "http://localhost/?format=json", nil)
		re = j(map[string]string{"hello": "world"}, req)
		r.Equal("application/json", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)
		r.NoError(err)
		r.Equal(`{"hello":"world"}`, strings.TrimSpace(bb.String()))

		// Test Content-Type header JSON
		req, _ = http.NewRequest("GET", "http://localhost/", nil)
		req.Header.Set("Content-Type", "application/json")
		re = j(map[string]string{"hello": "world"}, req)
		r.Equal("application/json", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)
		r.NoError(err)
		r.Equal(`{"hello":"world"}`, strings.TrimSpace(bb.String()))

		// Test format query argument XML
		req, _ = http.NewRequest("GET", "http://localhost/?format=xml", nil)
		re = j(user{Name: "mark"}, req)
		r.Equal("application/xml", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)
		r.NoError(err)
		r.Equal("<user>\n  <Name>mark</Name>\n</user>", strings.TrimSpace(bb.String()))

		// Test Content-Type header XML
		req, _ = http.NewRequest("GET", "http://localhost/", nil)
		req.Header.Set("Content-Type", "application/xml")
		re = j(user{Name: "mark"}, req)
		r.Equal("application/xml", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)
		r.NoError(err)
		r.Equal("<user>\n  <Name>mark</Name>\n</user>", strings.TrimSpace(bb.String()))
	}
}

func Test_AutoHtml(t *testing.T) {
	r := require.New(t)

	type ji func(v interface{}, req *http.Request) render.Renderer

	type user struct {
		Name string
	}

	table := []ji{
		render.Auto,
		render.New(render.Options{}).Auto,
	}

	for _, j := range table {

		// Test GET /resource HTML
		req, _ := http.NewRequest("GET", "http://localhost/resource", nil)
		req.Header.Set("Content-Type", "text/html")
		re := j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb := &bytes.Buffer{}
		err := re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/index.html: no such file or directory")

		// Test GET /resource/ID HTML
		req, _ = http.NewRequest("GET", "http://localhost/resource/1", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/show.html: no such file or directory")

		// Test GET /resource/ID/edit HTML
		req, _ = http.NewRequest("GET", "http://localhost/resource/1/edit", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/edit.html: no such file or directory")

		// Test GET /resource/new HTML
		req, _ = http.NewRequest("GET", "http://localhost/resource/new", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/new.html: no such file or directory")

		// Test POST /resource HTML
		req, _ = http.NewRequest("POST", "http://localhost/resource", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/new.html: no such file or directory")

		// Test PUT /resource/ID HTML
		req, _ = http.NewRequest("PUT", "http://localhost/resource/1", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/edit.html: no such file or directory")

		// Test DELETE /resource/ID HTML
		req, _ = http.NewRequest("DELETE", "http://localhost/resource/1", nil)
		req.Header.Set("Content-Type", "text/html")
		re = j(nil, req)
		r.Equal("text/html", re.ContentType())
		bb = &bytes.Buffer{}
		err = re.Render(bb, nil)

		// No need to re-test rendering, just ensure the right file is picked
		r.EqualError(err, "open resource/show.html: no such file or directory")
	}
}