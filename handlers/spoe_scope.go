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

package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	client_native "github.com/haproxytech/client-native/v2"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/spoe"
)

// SpoeCreateSpoeScopeHandlerImpl implementation of the SpoeCreateSpoeScopeHandler interface
type SpoeCreateSpoeScopeHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (s *SpoeCreateSpoeScopeHandlerImpl) Handle(params spoe.CreateSpoeScopeParams, principal interface{}) middleware.Responder {
	ss, err := s.Client.Spoe.GetSingleSpoe(params.Spoe)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewCreateSpoeScopeDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	v := int64(0)
	if params.Version != nil {
		v = *params.Version
	}
	err = ss.CreateScope(&params.Data, t, v)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewCreateSpoeScopeDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	return spoe.NewCreateSpoeScopeCreated().WithPayload(spoe.NewCreateSpoeScopeCreated().Payload)
}

//SpoeDeleteSpoeScopeHandlerImpl implementation of the SpoeDeleteSpoeScopeHandler interface
type SpoeDeleteSpoeScopeHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (s *SpoeDeleteSpoeScopeHandlerImpl) Handle(params spoe.DeleteSpoeScopeParams, principal interface{}) middleware.Responder {
	ss, err := s.Client.Spoe.GetSingleSpoe(params.Spoe)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewDeleteSpoeScopeDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	v := int64(0)
	if params.Version != nil {
		v = *params.Version
	}
	err = ss.DeleteScope(params.Name, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return spoe.NewDeleteSpoeScopeDefault(int(*e.Code)).WithPayload(e)
	}
	return spoe.NewDeleteSpoeScopeNoContent()
}

// SpoeGetSpoeScopesHandlerImpl implementation of the SpoeGetSpoeScopesHandler interface
type SpoeGetSpoeScopesHandlerImpl struct {
	Client *client_native.HAProxyClient
}

// SpoeGetAllSpoeFilesHandlerImpl implementation of the SpoeGetAllSpoeFilesHandler
func (s *SpoeGetSpoeScopesHandlerImpl) Handle(params spoe.GetSpoeScopesParams, principal interface{}) middleware.Responder {
	ss, err := s.Client.Spoe.GetSingleSpoe(params.Spoe)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewGetAllSpoeFilesDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	_, scopes, err := ss.GetScopes(t)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewGetAllSpoeFilesDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	return spoe.NewGetSpoeScopesOK().WithPayload(&spoe.GetSpoeScopesOKBody{Data: scopes})
}

// SpoeGetSpoeScopeHandlerImpl implementation of the SpoeGetSpoeScopeHandler interface
type SpoeGetSpoeScopeHandlerImpl struct {
	Client *client_native.HAProxyClient
}

func (s *SpoeGetSpoeScopeHandlerImpl) Handle(params spoe.GetSpoeScopeParams, principal interface{}) middleware.Responder {
	ss, err := s.Client.Spoe.GetSingleSpoe(params.Spoe)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewGetSpoeScopeDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	v, scope, err := ss.GetScope(params.Name, t)
	if err != nil {
		status := misc.GetHTTPStatusFromErr(err)
		return spoe.NewGetSpoeScopeDefault(status).WithPayload(misc.SetError(status, err.Error()))
	}
	if scope == nil {
		return spoe.NewGetSpoeScopeNotFound()
	}
	return spoe.NewGetSpoeScopeOK().WithPayload(&spoe.GetSpoeScopeOKBody{Version: v, Data: *scope})
}
