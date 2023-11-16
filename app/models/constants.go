package models

const (
	AuthorizationHeader          = "Authorization"
	BearerPrefixTitleCase        = "Bearer"
	BearerPrefixLowerCase        = "bearer"
	BindingError                 = "BINDING"
	ValidationError              = "VALIDATION"
	DataError                    = "DATA"
	ResourceNotFound             = "RESOURCE_NOT_FOUND"
	RequestLogMessage            = "Request => %s"
	ResponseLogMessage           = "Response => %s"
	BookResourcePathID           = "/book/:id"
	BookResourcePath             = "/books"
	UpdateErrorMessage           = "Unable to update record on the database"
	InsertErrorMessage           = "Unable to insert record on the database"
	ResourceNotFoundErrorMessage = "Unable to locate resource"
	DefaultPageText              = "page"
	DefaultSizeText              = "size"
	DefaultPage                  = "1"
	DefaultPageSize              = "50"
	NotDeleted                   = 0
	DefaultMinPageSize           = 1
	DefaultMaxPageSize           = 1000
)
