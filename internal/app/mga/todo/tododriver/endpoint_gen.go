// Code generated by mga tool. DO NOT EDIT.
package tododriver

import (
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
)

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateTodo endpoint.Endpoint
	ListTodos  endpoint.Endpoint
	MarkAsDone endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service todo.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Chain(middleware...)

	return Endpoints{
		CreateTodo: mw(MakeCreateTodoEndpoint(service)),
		ListTodos:  mw(MakeListTodosEndpoint(service)),
		MarkAsDone: mw(MakeMarkAsDoneEndpoint(service)),
	}
}

// TraceEndpoints returns an Endpoints struct where each endpoint is wrapped with a tracing middleware.
func TraceEndpoints(endpoints Endpoints) Endpoints {
	return Endpoints{
		CreateTodo: kitoc.TraceEndpoint("todo.CreateTodo")(endpoints.CreateTodo),
		ListTodos:  kitoc.TraceEndpoint("todo.ListTodos")(endpoints.ListTodos),
		MarkAsDone: kitoc.TraceEndpoint("todo.MarkAsDone")(endpoints.MarkAsDone),
	}
}