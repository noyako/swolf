package swolf

// Config for database manager.
type Config struct {
	Driver         string
	Connection     string         // Connection string with user, password and connection address
	MasterDatabase string         // Name of the master database
	MasterTable    string         // Name of the master table in master database
	Template       string         // Database template, works only for postgresql
	MasterData     MasterResolver // Names of control columns
	Mapper         FMapper        // Mapping functions
}

// MasterResolver is used to configure key and tenant fields.
type MasterResolver interface {
	getKey() string
	getTenant() string
}

type masterDataResolver struct {
	data interface{}
}

type masterFieldResolver struct {
	key    string
	tenant string
}

// MasterFieldResolver sets key and dependent fields from string.
func MasterFieldResolver(k, t string) MasterResolver {
	return &masterFieldResolver{
		key:    k,
		tenant: t,
	}
}

// MasterDataResolver is used to configure key and tenant fields from given type.
// Given struct must contain tags:
// `sw:"key"` for key
// `sw:"tenant"` for tenant database
func MasterDataResolver(in interface{}) MasterResolver {
	return &masterDataResolver{
		data: in,
	}
}

func (mdr *masterDataResolver) getKey() string {
	// TODO
	return ""
}

func (mdr *masterDataResolver) getTenant() string {
	// TODO
	return ""
}

func (mfr *masterFieldResolver) getKey() string {
	return mfr.key
}

func (mfr *masterFieldResolver) getTenant() string {
	return mfr.tenant
}
