package render

import (
	"net/http"
	"strings"
)

// autoHTML renders an HTML template, making some assumptions
func autoHTML(req *http.Request) Renderer {
	var names []string

	path := strings.TrimLeft(req.URL.Path, "/")
	splitPath := strings.Split(path, "/")

	if len(splitPath) > 0 {
		if strings.EqualFold(req.Method, "POST") {
			// POST /resource
			names = append(names, splitPath[0]+"/new.html")
		} else if strings.EqualFold(req.Method, "PUT") {
			// PUT /resource/ID
			names = append(names, splitPath[0]+"/edit.html")
		} else if strings.EqualFold(req.Method, "DELETE") {
			// DELETE /resource/ID
			names = append(names, splitPath[0]+"/show.html")
		} else if strings.EqualFold(req.Method, "GET") {
			if strings.HasSuffix(path, "/new") {
				// GET /resource/new
				names = append(names, splitPath[0]+"/new.html")
			} else if strings.HasSuffix(path, "/edit") {
				// GET /resource/ID/edit
				names = append(names, splitPath[0]+"/edit.html")
			} else if len(splitPath) > 1 {
				// GET /resource/ID
				names = append(names, splitPath[0]+"/show.html")
			} else {
				// GET /resource
				names = append(names, splitPath[0]+"/index.html")
			}
		}
	}

	e := New(Options{})

	if e.HTMLLayout != "" {
		names = append(names, e.HTMLLayout)
	}
	hr := templateRenderer{
		Engine:      e,
		contentType: "text/html",
		names:       names,
	}
	return hr
}

// Auto renders the value using the content type
// detected in "Content-Type" header. If the header doesn't exist
// it tries to fetch it from the "format" URL query argument.
// However, if it fails to detect the content type, json is provided
// as a fallback.
func Auto(v interface{}, req *http.Request) Renderer {
	if contentType, found := req.Header["Content-Type"]; found {
		// Try to read content type from Content-Type HTTP header
		if strings.EqualFold(contentType[0], "text/html") {
			return autoHTML(req)
		} else if strings.EqualFold(contentType[0], "application/json") {
			return jsonRenderer{value: v}
		} else if strings.EqualFold(contentType[0], "application/xml") {
			return xmlRenderer{value: v}
		}
	} else {
		// Try to get content type as a query argument
		format := req.URL.Query().Get("format")
		if len(format) > 0 {
			if strings.EqualFold(format, "html") {
				return autoHTML(req)
			} else if strings.EqualFold(format, "json") {
				return jsonRenderer{value: v}
			} else if strings.EqualFold(format, "xml") {
				return xmlRenderer{value: v}
			}
		}
	}

	// jsonRenderer as fallback
	return jsonRenderer{value: v}
}

// Auto renders the value using the content type
// detected in "Content-Type" header. If the header doesn't exist
// it tries to fetch it from the "format" URL query argument.
// However, if it fails to detect the content type, json is provided
// as a fallback.
func (e *Engine) Auto(v interface{}, req *http.Request) Renderer {
	return Auto(v, req)
}
