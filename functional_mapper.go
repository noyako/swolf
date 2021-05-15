package swolf

// FMapper is used to convert ids to names
type FMapper struct {
	IDToField       func(string) string // Convert tenant ID to key field in master database
	FieldToDatabase func(string) string // Convert dependent field to database name
	KeyToValue      func(string) string // Convert key field to dependent in master database
}

func (fm *FMapper) defaultArgs() {
	if fm.IDToField == nil {
		fm.IDToField = defaultIDToFIledFunction
	}

	if fm.FieldToDatabase == nil {
		fm.FieldToDatabase = defaultFieldToDatabase
	}

	if fm.KeyToValue == nil {
		fm.KeyToValue = defaultKeyToValue
	}
}

func defaultIDToFIledFunction(id string) string {
	return id
}

func defaultFieldToDatabase(field string) string {
	return field
}

func defaultKeyToValue(field string) string {
	return field
}
