package apig

func validateForeignKey(fields []*Field, name string) bool {
	for _, field := range fields {
		val := name + "ID"
		if field.Name == val {
			return true
		}
	}
	return false
}
