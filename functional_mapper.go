package swolf

type FMapper struct {
	IDToField       func(string) string
	FieldToDatabase func(string) string
}

func (fm *FMapper) defaultArgs() {
	if fm.IDToField == nil {
		fm.IDToField = defaultIDToFIledFunction
	}

	if fm.FieldToDatabase == nil {
		fm.FieldToDatabase = defaultFieldToDatabase
	}
}

func defaultIDToFIledFunction(id string) string {
	return id
}

func defaultFieldToDatabase(field string) string {
	return field
}
