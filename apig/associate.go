package apig

import "strings"

func resolveAssociate(model *Model, modelMap map[string]*Model, parents map[string]bool) {
	parents[model.Name] = true

	for i, field := range model.Fields {
		if field.Association != nil && field.Association.Type != AssociationNone {
			continue
		}

		str := strings.Trim(field.Type, "[]*")
		if modelMap[str] != nil && parents[str] != true {
			resolveAssociate(modelMap[str], modelMap, parents)

			var assoc int
			switch string([]rune(field.Type)[0]) {
			case "[":
				if validateForeignKey(modelMap[str].Fields, model.Name) {
					assoc = AssociationHasMany
					break
				}
				assoc = AssociationBelongsTo
			default:
				if validateForeignKey(modelMap[str].Fields, model.Name) {
					assoc = AssociationHasOne
					break
				}
				assoc = AssociationBelongsTo
			}
			model.Fields[i].Association = &Association{Type: assoc, Model: modelMap[str]}
		} else {
			model.Fields[i].Association = &Association{Type: AssociationNone}
		}
	}
}
