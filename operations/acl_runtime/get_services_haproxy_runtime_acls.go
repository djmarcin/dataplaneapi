// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package acl_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetServicesHaproxyRuntimeAclsHandlerFunc turns a function with the right signature into a get services haproxy runtime acls handler
type GetServicesHaproxyRuntimeAclsHandlerFunc func(GetServicesHaproxyRuntimeAclsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetServicesHaproxyRuntimeAclsHandlerFunc) Handle(params GetServicesHaproxyRuntimeAclsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetServicesHaproxyRuntimeAclsHandler interface for that can handle valid get services haproxy runtime acls params
type GetServicesHaproxyRuntimeAclsHandler interface {
	Handle(GetServicesHaproxyRuntimeAclsParams, interface{}) middleware.Responder
}

// NewGetServicesHaproxyRuntimeAcls creates a new http.Handler for the get services haproxy runtime acls operation
func NewGetServicesHaproxyRuntimeAcls(ctx *middleware.Context, handler GetServicesHaproxyRuntimeAclsHandler) *GetServicesHaproxyRuntimeAcls {
	return &GetServicesHaproxyRuntimeAcls{Context: ctx, Handler: handler}
}

/*GetServicesHaproxyRuntimeAcls swagger:route GET /services/haproxy/runtime/acls ACL Runtime getServicesHaproxyRuntimeAcls

Return an array of all ACL files

Returns all ACL files using the runtime socket.

*/
type GetServicesHaproxyRuntimeAcls struct {
	Context *middleware.Context
	Handler GetServicesHaproxyRuntimeAclsHandler
}

func (o *GetServicesHaproxyRuntimeAcls) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetServicesHaproxyRuntimeAclsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}