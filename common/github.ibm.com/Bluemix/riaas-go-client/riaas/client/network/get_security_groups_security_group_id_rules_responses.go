// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSecurityGroupsSecurityGroupIDRulesReader is a Reader for the GetSecurityGroupsSecurityGroupIDRules structure.
type GetSecurityGroupsSecurityGroupIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsSecurityGroupIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsSecurityGroupIDRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSecurityGroupsSecurityGroupIDRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupsSecurityGroupIDRulesOK creates a GetSecurityGroupsSecurityGroupIDRulesOK with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesOK() *GetSecurityGroupsSecurityGroupIDRulesOK {
	return &GetSecurityGroupsSecurityGroupIDRulesOK{}
}

/*GetSecurityGroupsSecurityGroupIDRulesOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsSecurityGroupIDRulesOK struct {
	Payload *GetSecurityGroupsSecurityGroupIDRulesOKBody
}

func (o *GetSecurityGroupsSecurityGroupIDRulesOK) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules][%d] getSecurityGroupsSecurityGroupIdRulesOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSecurityGroupsSecurityGroupIDRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsSecurityGroupIDRulesInternalServerError creates a GetSecurityGroupsSecurityGroupIDRulesInternalServerError with default headers values
func NewGetSecurityGroupsSecurityGroupIDRulesInternalServerError() *GetSecurityGroupsSecurityGroupIDRulesInternalServerError {
	return &GetSecurityGroupsSecurityGroupIDRulesInternalServerError{}
}

/*GetSecurityGroupsSecurityGroupIDRulesInternalServerError handles this case with default header values.

error
*/
type GetSecurityGroupsSecurityGroupIDRulesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsSecurityGroupIDRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/rules][%d] getSecurityGroupsSecurityGroupIdRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSecurityGroupsSecurityGroupIDRulesOKBody SecurityGroupRuleCollection
//
// Collection of rules in a security group
swagger:model GetSecurityGroupsSecurityGroupIDRulesOKBody
*/
type GetSecurityGroupsSecurityGroupIDRulesOKBody struct {

	// Array of rules
	// Required: true
	Rules []*models.SecurityGroupRule `json:"rules"`
}

// Validate validates this get security groups security group ID rules o k body
func (o *GetSecurityGroupsSecurityGroupIDRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityGroupsSecurityGroupIDRulesOKBody) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("getSecurityGroupsSecurityGroupIdRulesOK"+"."+"rules", "body", o.Rules); err != nil {
		return err
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getSecurityGroupsSecurityGroupIdRulesOK" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityGroupsSecurityGroupIDRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityGroupsSecurityGroupIDRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityGroupsSecurityGroupIDRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}