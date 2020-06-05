package lightweight_api

const (
	ResourceAlreadyExists = "%s already exists. guidColName: %s, guidValue: %+v"
	ResourceMustExists = "%s must exists."
	ResourceMustNotExistsExceptSelf = "%s must not exists except self."

	GuidTagMustNotBeEmpty = "guid tag must not be empty"
	GuidFieldMustNotExists = "guidField %s must not exists"
	FieldCannotFind = "cannot find field: %s"
)
