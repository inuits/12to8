// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ServicesMonthInfoListReader is a Reader for the ServicesMonthInfoList structure.
type ServicesMonthInfoListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicesMonthInfoListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServicesMonthInfoListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServicesMonthInfoListOK creates a ServicesMonthInfoListOK with default headers values
func NewServicesMonthInfoListOK() *ServicesMonthInfoListOK {
	return &ServicesMonthInfoListOK{}
}

/*ServicesMonthInfoListOK handles this case with default header values.

ServicesMonthInfoListOK services month info list o k
*/
type ServicesMonthInfoListOK struct {
}

func (o *ServicesMonthInfoListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/services/month_info/][%d] servicesMonthInfoListOK ", 200)
}

func (o *ServicesMonthInfoListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
