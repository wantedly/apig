package apig

import "testing"

func TestResolveAssociate(t *testing.T) {
	user := &Model{
		Name: "User",
		Fields: []*Field{
			&Field{
				Name: "ID",
				Type: "uint",
			},
			&Field{
				Name: "Profile",
				Type: "*Profile",
			},
			&Field{
				Name: "ProfileID",
				Type: "uint",
			},
		},
	}

	profile := &Model{
		Name: "Profile",
		Fields: []*Field{
			&Field{
				Name: "ID",
				Type: "uint",
			},
			&Field{
				Name: "Name",
				Type: "string",
			},
			&Field{
				Name: "User",
				Type: "*User",
			},
		},
	}

	modelMap := modelToMap(user, profile)

	if len(modelMap) != 2 {
		t.Fatalf("Number of models map is incorrect. expected: 2, actual: %d", len(modelMap))
	}

	resolveAssociate(user, modelMap, make(map[string]bool))

	// Profile
	result := user.Fields[1].Association.Type
	expect := AssociationBelongsTo
	if result != expect {
		t.Fatalf("Incorrect result. expected: %v, actual: %v", expect, result)
	}
}

func modelToMap(models ...*Model) map[string]*Model {
	modelMap := make(map[string]*Model)
	for _, model := range models {
		modelMap[model.Name] = model
	}
	return modelMap
}
