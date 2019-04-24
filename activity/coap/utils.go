package coap

import (
	"strconv"

	"github.com/dustin/go-coap"
)

const (
	methodGET    = "GET"
	methodPOST   = "POST"
	methodPUT    = "PUT"
	methodDELETE = "DELETE"

	typeCON = "CONFIRMABLE"
	typeNON = "NONCONFIRMABLE"
	typeACK = "ACKNOWLEDGEMENT"
	typeRST = "RESET"
)

func toCoapCode(method string) coap.COAPCode {

	var code coap.COAPCode

	switch method {
	case methodGET:
		code = coap.GET
	case methodPOST:
		code = coap.POST
	case methodPUT:
		code = coap.PUT
	case methodDELETE:
		code = coap.DELETE
	}

	return code
}

func toCoapType(typeStr string) coap.COAPType {

	var ctype coap.COAPType

	switch typeStr {
	case typeCON:
		ctype = coap.Confirmable
	case typeNON:
		ctype = coap.NonConfirmable
	case typeACK:
		ctype = coap.Acknowledgement
	case typeRST:
		ctype = coap.Reset
	}

	return ctype
}

func toOption(name string, value string) (coap.OptionID, interface{}) {

	var opID coap.OptionID
	var val interface{}

	val = value

	switch name {
	case "IFMATCH":
		opID = coap.IfMatch
	case "URIHOST":
		opID = coap.URIHost
	case "ETAG":
		opID = coap.ETag
	//case "IFNONEMATCH":
	//	opID = coap.IfNoneMatch
	case "OBSERVE":
		opID = coap.Observe
		val, _ = strconv.Atoi(value)
	case "URIPORT":
		opID = coap.URIPort
		val, _ = strconv.Atoi(value)
	case "LOCATIONPATH":
		opID = coap.LocationPath
	case "URIPATH":
		opID = coap.URIPath
	case "CONTENTFORMAT":
		opID = coap.ContentFormat
		val, _ = strconv.Atoi(value)
	case "MAXAGE":
		opID = coap.MaxAge
		val, _ = strconv.Atoi(value)
	case "URIQUERY":
		opID = coap.URIQuery
	case "ACCEPT":
		opID = coap.IfMatch
		val, _ = strconv.Atoi(value)
	case "LOCATIONQUERY":
		opID = coap.LocationQuery
	case "PROXYURI":
		opID = coap.ProxyURI
	case "PROXYSCHEME":
		opID = coap.ProxyScheme
	case "SIZE1":
		opID = coap.Size1
		val, _ = strconv.Atoi(value)
	default:
		opID = 0
		val = nil
	}

	return opID, val
}
