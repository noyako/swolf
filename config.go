package swolf

type Config struct {
	Driver         string
	Connection     string
	MasterDatabase string
	MasterTable    string
	Template       string
	MasterData     masterResolver
	Mapper         FMapper
}

type masterResolver interface {
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

func MasterFieldResolver(k, t string) masterResolver {
	return &masterFieldResolver{
		key:    k,
		tenant: t,
	}
}

func MasterDataResolver(in interface{}) masterResolver {
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
