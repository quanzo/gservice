package status

import (
	"fmt"
	"strconv"
)

type IStatus interface {
	error
	fmt.Stringer
	IsWarning() bool
	IsError() bool
	GetCode() int
	GetMessage() string
	GetMessageCode() string
}
type Status struct {
	IStatus
	code      int
	message   string
	isError   bool
	isWarning bool
}

func New(code int, message string, isError, isWarning bool) *Status {
	this := new(Status)
	this.code = code
	this.message = message
	this.isError = isError
	this.isWarning = isWarning
	return this
}

func (this *Status) Error() string {
	if this.isError {
		return this.String()
	}
	return ""
}
func (this *Status) String() (result string) {
	if this.isError {
		result = "WARNING. " + result
	}
	if this.isError {
		result = "ERROR. " + result
	}
	result = result + this.GetMessage()
	return
}
func (this *Status) Code2Message(code int) (result string) {
	return strconv.FormatInt(int64(code), 10)
}
func (this *Status) GetMessageCode() string {
	return this.Code2Message(this.code) + ": " + this.message
}
func (this *Status) IsWarning() bool {
	return this.isWarning
}
func (this *Status) IsError() bool {
	return this.isError
}
func (this *Status) GetCode() int {
	return this.code
}
func (this *Status) GetMessage() string {
	return this.message
}
