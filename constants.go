package modiniter

const (
	HttpModeReadOnly  = "readOnly"
	HttpModeFull      = "full"
	HttpModeAdminOnly = "adminOnly"
	HttpModeDsp       = "dsp"

	JWTSigningKey          = "3fNE49BD49hZ58mXzYX4DVCG"
	JWTFieldIss            = "iss"
	JWTFieldIssValuePrefix = "auth."
	JWTFieldAud            = "aud"
	JWTFieldExp            = "exp"
	CurrentUserKey         = "currentUser"

	UserIdKey            = "id"
	UserRolesKey         = "roles"
	FirstNameKey         = "firstName"
	SecondNameKey        = "secondName"
	LastNameKey          = "lastName"
	DepartmentIdKey      = "departmentId"
	DepartmentLabelKey   = "departmentLabel"
	OrganizationIdKey    = "organizationId"
	OrganizationLabelKey = "organizationLabel"

	InitModeSA               = "SA"
	InitModePkg              = "PKG"
	InitModeStorelessService = "StorelessService"
)
