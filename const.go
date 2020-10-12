package lightweight_api

const (
	EnvDriverName = "DB_DRIVER_NAME"
	HintDriverNameNotRecognized = "lightweight_api did not recognize the environment variable DB_DRIVER_NAME: %s you provide, please set Connector by yourself"
)

const (
	ResourceAlreadyExists = "%s already exists. guidColName: %s, guidValue: %+v"
	ResourceMustExists = "%s must exists."
	ResourceMustNotExistsExceptSelf = "%s must not exists except self."

	GuidTagMustNotBeEmpty = "guid tag must not be empty"
	GuidFieldMustNotExists = "guidField %s must not exists"
	FieldCannotFind = "cannot find field: %s"
)

const (
	MsgResourceCreateSuccess = "%s %v create success"
)
