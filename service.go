package oscrud

// Service :
type Service interface {
	Find(Context) Context
	Create(Context) Context
	Get(string, Context) Context
	Update(string, Context) Context
	Patch(string, Context) Context
	Delete(string, Context) Context
}

// ServiceOptions :
type ServiceOptions struct {
	DisableFind   bool
	DisableCreate bool
	DisableGet    bool
	DisableUpdate bool
	DisablePatch  bool
	DisableDelete bool
}

// ServiceAction :
type ServiceAction string

// ServiceActions :
var (
	ServiceActionCreate ServiceAction = "CREATE"
	ServiceActionFind   ServiceAction = "FIND"
	ServiceActionGet    ServiceAction = "GET"
	ServiceActionUpdate ServiceAction = "UPDATE"
	ServiceActionPatch  ServiceAction = "PATCH"
	ServiceActionDelete ServiceAction = "DELETE"
)

// ServiceModel :
type ServiceModel interface {
	ToResult(ServiceAction) (interface{}, error)
	ToQuery(ServiceAction) (interface{}, error)
	ToCreate() error
	ToDelete() error
	ToPatch(ServiceModel) error
	ToUpdate(ServiceModel) error
}

// transforms $id endpoint to proper oscrud handler
func serviceHandler(handler func(string, Context) Context) Handler {
	return func(ctx Context) Context {
		var i struct {
			ID string `param:"id"`
		}

		ctx.Bind(&i)
		return handler(i.ID, ctx)
	}
}
