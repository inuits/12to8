// Code generated by go-swagger; DO NOT EDIT.

package my_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// MyPerformancesActivityCreateReader is a Reader for the MyPerformancesActivityCreate structure.
type MyPerformancesActivityCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MyPerformancesActivityCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewMyPerformancesActivityCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMyPerformancesActivityCreateCreated creates a MyPerformancesActivityCreateCreated with default headers values
func NewMyPerformancesActivityCreateCreated() *MyPerformancesActivityCreateCreated {
	return &MyPerformancesActivityCreateCreated{}
}

/*MyPerformancesActivityCreateCreated handles this case with default header values.

MyPerformancesActivityCreateCreated my performances activity create created
*/
type MyPerformancesActivityCreateCreated struct {
}

func (o *MyPerformancesActivityCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/my_performances/activity/][%d] myPerformancesActivityCreateCreated ", 201)
}

func (o *MyPerformancesActivityCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*MyPerformancesActivityCreateBody my performances activity create body
swagger:model MyPerformancesActivityCreateBody
*/

type MyPerformancesActivityCreateBody struct {

	// contract
	// Required: true
	Contract *string `json:"contract"`

	// contract role
	ContractRole string `json:"contract_role,omitempty"`

	// day
	// Required: true
	Day *int64 `json:"day"`

	// description
	Description string `json:"description,omitempty"`

	// duration
	Duration float64 `json:"duration,omitempty"`

	// performance type
	// Required: true
	PerformanceType *string `json:"performance_type"`

	// redmine id
	RedmineID string `json:"redmine_id,omitempty"`

	// timesheet
	// Required: true
	Timesheet *string `json:"timesheet"`
}

/* polymorph MyPerformancesActivityCreateBody contract false */

/* polymorph MyPerformancesActivityCreateBody contract_role false */

/* polymorph MyPerformancesActivityCreateBody day false */

/* polymorph MyPerformancesActivityCreateBody description false */

/* polymorph MyPerformancesActivityCreateBody duration false */

/* polymorph MyPerformancesActivityCreateBody performance_type false */

/* polymorph MyPerformancesActivityCreateBody redmine_id false */

/* polymorph MyPerformancesActivityCreateBody timesheet false */

// MarshalBinary interface implementation
func (o *MyPerformancesActivityCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MyPerformancesActivityCreateBody) UnmarshalBinary(b []byte) error {
	var res MyPerformancesActivityCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
