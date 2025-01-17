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

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/client-native/v4/models"
)

// GetBackendSwitchingRulesOKCode is the HTTP code returned for type GetBackendSwitchingRulesOK
const GetBackendSwitchingRulesOKCode int = 200

/*GetBackendSwitchingRulesOK Successful operation

swagger:response getBackendSwitchingRulesOK
*/
type GetBackendSwitchingRulesOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetBackendSwitchingRulesOKBody `json:"body,omitempty"`
}

// NewGetBackendSwitchingRulesOK creates GetBackendSwitchingRulesOK with default headers values
func NewGetBackendSwitchingRulesOK() *GetBackendSwitchingRulesOK {

	return &GetBackendSwitchingRulesOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get backend switching rules o k response
func (o *GetBackendSwitchingRulesOK) WithConfigurationVersion(configurationVersion string) *GetBackendSwitchingRulesOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get backend switching rules o k response
func (o *GetBackendSwitchingRulesOK) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get backend switching rules o k response
func (o *GetBackendSwitchingRulesOK) WithPayload(payload *GetBackendSwitchingRulesOKBody) *GetBackendSwitchingRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend switching rules o k response
func (o *GetBackendSwitchingRulesOK) SetPayload(payload *GetBackendSwitchingRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendSwitchingRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBackendSwitchingRulesDefault General Error

swagger:response getBackendSwitchingRulesDefault
*/
type GetBackendSwitchingRulesDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion string `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackendSwitchingRulesDefault creates GetBackendSwitchingRulesDefault with default headers values
func NewGetBackendSwitchingRulesDefault(code int) *GetBackendSwitchingRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBackendSwitchingRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) WithStatusCode(code int) *GetBackendSwitchingRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) WithConfigurationVersion(configurationVersion string) *GetBackendSwitchingRulesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) SetConfigurationVersion(configurationVersion string) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) WithPayload(payload *models.Error) *GetBackendSwitchingRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend switching rules default response
func (o *GetBackendSwitchingRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendSwitchingRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := o.ConfigurationVersion
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
