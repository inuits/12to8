// Code generated by go-swagger; DO NOT EDIT.

package companies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CompaniesReadReader is a Reader for the CompaniesRead structure.
type CompaniesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompaniesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompaniesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompaniesReadOK creates a CompaniesReadOK with default headers values
func NewCompaniesReadOK() *CompaniesReadOK {
	return &CompaniesReadOK{}
}

/*CompaniesReadOK handles this case with default header values.

CompaniesReadOK companies read o k
*/
type CompaniesReadOK struct {
}

func (o *CompaniesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/companies/{id}/][%d] companiesReadOK ", 200)
}

func (o *CompaniesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
