// Code generated by go-swagger; DO NOT EDIT.

package leaves

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// LeavesListReader is a Reader for the LeavesList structure.
type LeavesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LeavesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLeavesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLeavesListOK creates a LeavesListOK with default headers values
func NewLeavesListOK() *LeavesListOK {
	return &LeavesListOK{}
}

/*LeavesListOK handles this case with default header values.

LeavesListOK leaves list o k
*/
type LeavesListOK struct {
}

func (o *LeavesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/leaves/][%d] leavesListOK ", 200)
}

func (o *LeavesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
