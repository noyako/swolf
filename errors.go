package swolf

const (
	errorTypeNotStruct    = "swolf: MasterData is not a struct! Got: %s"
	errorCannotDelete     = "Could not delete for id=%s (key=%s)"
	errorCannotInsert     = "Could not insert for id=%s (key=%s, value=%s)"
	errorTenantNotFound   = "Tenant for id=%s (key=%s) not found"
	panicTagInWrongFormat = "Tag %s is in wrong format"
	panicTagNotFound      = "No tags are found"
)
