// Code generated by go-mockgen 1.3.4; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package uploads

import (
	"context"
	"sync"

	shared "github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/shared"
	graphql "github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/transport/graphql"
)

// MockResolver is a mock implementation of the Resolver interface (from the
// package
// github.com/sourcegraph/sourcegraph/internal/codeintel/uploads/transport/graphql)
// used for unit testing.
type MockResolver struct {
	// GetUploadsByIDsFunc is an instance of a mock function object
	// controlling the behavior of the method GetUploadsByIDs.
	GetUploadsByIDsFunc *ResolverGetUploadsByIDsFunc
	// UploadsConnectionResolverFromFactoryFunc is an instance of a mock
	// function object controlling the behavior of the method
	// UploadsConnectionResolverFromFactory.
	UploadsConnectionResolverFromFactoryFunc *ResolverUploadsConnectionResolverFromFactoryFunc
}

// NewMockResolver creates a new mock of the Resolver interface. All methods
// return zero values for all results, unless overwritten.
func NewMockResolver() *MockResolver {
	return &MockResolver{
		GetUploadsByIDsFunc: &ResolverGetUploadsByIDsFunc{
			defaultHook: func(context.Context, ...int) (r0 []shared.Upload, r1 error) {
				return
			},
		},
		UploadsConnectionResolverFromFactoryFunc: &ResolverUploadsConnectionResolverFromFactoryFunc{
			defaultHook: func(shared.GetUploadsOptions) (r0 *graphql.UploadsResolver) {
				return
			},
		},
	}
}

// NewStrictMockResolver creates a new mock of the Resolver interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockResolver() *MockResolver {
	return &MockResolver{
		GetUploadsByIDsFunc: &ResolverGetUploadsByIDsFunc{
			defaultHook: func(context.Context, ...int) ([]shared.Upload, error) {
				panic("unexpected invocation of MockResolver.GetUploadsByIDs")
			},
		},
		UploadsConnectionResolverFromFactoryFunc: &ResolverUploadsConnectionResolverFromFactoryFunc{
			defaultHook: func(shared.GetUploadsOptions) *graphql.UploadsResolver {
				panic("unexpected invocation of MockResolver.UploadsConnectionResolverFromFactory")
			},
		},
	}
}

// NewMockResolverFrom creates a new mock of the MockResolver interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockResolverFrom(i graphql.Resolver) *MockResolver {
	return &MockResolver{
		GetUploadsByIDsFunc: &ResolverGetUploadsByIDsFunc{
			defaultHook: i.GetUploadsByIDs,
		},
		UploadsConnectionResolverFromFactoryFunc: &ResolverUploadsConnectionResolverFromFactoryFunc{
			defaultHook: i.UploadsConnectionResolverFromFactory,
		},
	}
}

// ResolverGetUploadsByIDsFunc describes the behavior when the
// GetUploadsByIDs method of the parent MockResolver instance is invoked.
type ResolverGetUploadsByIDsFunc struct {
	defaultHook func(context.Context, ...int) ([]shared.Upload, error)
	hooks       []func(context.Context, ...int) ([]shared.Upload, error)
	history     []ResolverGetUploadsByIDsFuncCall
	mutex       sync.Mutex
}

// GetUploadsByIDs delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockResolver) GetUploadsByIDs(v0 context.Context, v1 ...int) ([]shared.Upload, error) {
	r0, r1 := m.GetUploadsByIDsFunc.nextHook()(v0, v1...)
	m.GetUploadsByIDsFunc.appendCall(ResolverGetUploadsByIDsFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetUploadsByIDs
// method of the parent MockResolver instance is invoked and the hook queue
// is empty.
func (f *ResolverGetUploadsByIDsFunc) SetDefaultHook(hook func(context.Context, ...int) ([]shared.Upload, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetUploadsByIDs method of the parent MockResolver instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *ResolverGetUploadsByIDsFunc) PushHook(hook func(context.Context, ...int) ([]shared.Upload, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ResolverGetUploadsByIDsFunc) SetDefaultReturn(r0 []shared.Upload, r1 error) {
	f.SetDefaultHook(func(context.Context, ...int) ([]shared.Upload, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ResolverGetUploadsByIDsFunc) PushReturn(r0 []shared.Upload, r1 error) {
	f.PushHook(func(context.Context, ...int) ([]shared.Upload, error) {
		return r0, r1
	})
}

func (f *ResolverGetUploadsByIDsFunc) nextHook() func(context.Context, ...int) ([]shared.Upload, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverGetUploadsByIDsFunc) appendCall(r0 ResolverGetUploadsByIDsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverGetUploadsByIDsFuncCall objects
// describing the invocations of this function.
func (f *ResolverGetUploadsByIDsFunc) History() []ResolverGetUploadsByIDsFuncCall {
	f.mutex.Lock()
	history := make([]ResolverGetUploadsByIDsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverGetUploadsByIDsFuncCall is an object that describes an invocation
// of method GetUploadsByIDs on an instance of MockResolver.
type ResolverGetUploadsByIDsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg1 []int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []shared.Upload
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c ResolverGetUploadsByIDsFuncCall) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg1 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{c.Arg0}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverGetUploadsByIDsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// ResolverUploadsConnectionResolverFromFactoryFunc describes the behavior
// when the UploadsConnectionResolverFromFactory method of the parent
// MockResolver instance is invoked.
type ResolverUploadsConnectionResolverFromFactoryFunc struct {
	defaultHook func(shared.GetUploadsOptions) *graphql.UploadsResolver
	hooks       []func(shared.GetUploadsOptions) *graphql.UploadsResolver
	history     []ResolverUploadsConnectionResolverFromFactoryFuncCall
	mutex       sync.Mutex
}

// UploadsConnectionResolverFromFactory delegates to the next hook function
// in the queue and stores the parameter and result values of this
// invocation.
func (m *MockResolver) UploadsConnectionResolverFromFactory(v0 shared.GetUploadsOptions) *graphql.UploadsResolver {
	r0 := m.UploadsConnectionResolverFromFactoryFunc.nextHook()(v0)
	m.UploadsConnectionResolverFromFactoryFunc.appendCall(ResolverUploadsConnectionResolverFromFactoryFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// UploadsConnectionResolverFromFactory method of the parent MockResolver
// instance is invoked and the hook queue is empty.
func (f *ResolverUploadsConnectionResolverFromFactoryFunc) SetDefaultHook(hook func(shared.GetUploadsOptions) *graphql.UploadsResolver) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UploadsConnectionResolverFromFactory method of the parent MockResolver
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *ResolverUploadsConnectionResolverFromFactoryFunc) PushHook(hook func(shared.GetUploadsOptions) *graphql.UploadsResolver) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ResolverUploadsConnectionResolverFromFactoryFunc) SetDefaultReturn(r0 *graphql.UploadsResolver) {
	f.SetDefaultHook(func(shared.GetUploadsOptions) *graphql.UploadsResolver {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ResolverUploadsConnectionResolverFromFactoryFunc) PushReturn(r0 *graphql.UploadsResolver) {
	f.PushHook(func(shared.GetUploadsOptions) *graphql.UploadsResolver {
		return r0
	})
}

func (f *ResolverUploadsConnectionResolverFromFactoryFunc) nextHook() func(shared.GetUploadsOptions) *graphql.UploadsResolver {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverUploadsConnectionResolverFromFactoryFunc) appendCall(r0 ResolverUploadsConnectionResolverFromFactoryFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// ResolverUploadsConnectionResolverFromFactoryFuncCall objects describing
// the invocations of this function.
func (f *ResolverUploadsConnectionResolverFromFactoryFunc) History() []ResolverUploadsConnectionResolverFromFactoryFuncCall {
	f.mutex.Lock()
	history := make([]ResolverUploadsConnectionResolverFromFactoryFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverUploadsConnectionResolverFromFactoryFuncCall is an object that
// describes an invocation of method UploadsConnectionResolverFromFactory on
// an instance of MockResolver.
type ResolverUploadsConnectionResolverFromFactoryFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 shared.GetUploadsOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *graphql.UploadsResolver
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverUploadsConnectionResolverFromFactoryFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverUploadsConnectionResolverFromFactoryFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}