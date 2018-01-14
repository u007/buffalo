package buffalo

import (
	"net/http"
	"net/url"
)

func CustomContext(context Context, response http.ResponseWriter, request *http.Request,
	params url.Values,
	logger Logger,
	session *Session,
	contentType string,
	data map[string]interface{},
	flash *Flash) DefaultContext {

	return DefaultContext{
		Context:     context,
		response:    response,
		request:     request,
		params:      params,
		logger:      logger,
		session:     session,
		contentType: contentType,
		data:        data,
		flash:       flash,
	}
}
