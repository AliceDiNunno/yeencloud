package rfc

//TODO: Implement DNS Domain Slug functions

// This file is used to generate and validate RFC 1123 name strings
// RFC 1123 names are specified using those restrictions:
//	- Must contain only lowercase alphanumeric characters and characters '-'
//	- Must start and end with an alphanumeric character
//	- Must be between 1 and 63 characters long

func IsValidRFC1123Name(string) bool {
	return true
}

func GenerateRFC1123Name(name string) string {
	return name
}

// DNS Domain Slug are a subset of RFC 1123 names
// They are specified using those restrictions:
//	- Must contain only lowercase alphanumeric characters and characters '-' and '.'
//	- Must start and end with an alphanumeric character
//	- Must be between 1 and 253 characters long

func IsValidDNSDomainName(string) bool {
	return true
}

func GenerateDNSDomainName(name string) string {
	return name
}
