package api

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	route{
		"ProjectList",
		"GET",
		"/api/v1/project",
		ProjectList,
	},
	route{
		"ProjectCreate",
		"POST",
		"/api/v1/project/{appId}",
		ProjectCreate,
	},
	route{
		"ProjectShow",
		"GET",
		"/api/v1/project/{appId}",
		ProjectShow,
	},
	route{
		"ProjectDefinition",
		"GET",
		"/api/v1/project/{appId}/definition",
		ProjectDefinition,
	},
	route{
		"ProjectUp",
		"POST",
		"/api/v1/project/{appId}/up",
		ProjectUp,
	},
	route{
		"ProjectStop",
		"POST",
		"/api/v1/project/{appId}/stop",
		ProjectStop,
	},
	route{
		"ProjectKill",
		"POST",
		"/api/v1/project/{appId}/kill",
		ProjectKill,
	},
	route{
		"ProjectDelete",
		"POST",
		"/api/v1/project/{appId}/delete",
		ProjectDelete,
	},
	route{
		"ProjectRemove",
		"POST",
		"/api/v1/project/{appId}/remove",
		ProjectRemove,
	},
	route{
		"ServiceScale",
		"POST",
		"/api/v1/project/{appId}/scale",
		ServiceScale,
	},
	route{
		"ServiceAutoScaling",
		"POST",
		"/api/v1/project/{appId}/autoscaling",
		ServiceAutoScaling,
	},
	route{
		"ServiceRemoveAutoScaling",
		"DELETE",
		"/api/v1/project/{appId}/autoscaling",
		ServiceRemoveAutoScaling,
	},
}
