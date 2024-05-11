package session

import (
	"reflect"

	"github.com/MysticalDevil/sorm/log"
)

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

type IBeforeQuery interface {
	BeforeQuery(s *Session) error
}

type IAfterQuery interface {
	AfterQuery(s *Session) error
}

type IBeforeUpdate interface {
	BeforeUpdate(s *Session) error
}

type IAfterUpdate interface {
	AfterUpdate(s *Session) error
}

type IBeforeDelete interface {
	BeforeDelete(s *Session) error
}

type IAfterDelete interface {
	AfterDelete(s *Session) error
}

type IBeforeInsert interface {
	BeforeInsert(s *Session) error
}

type IAfterInsert interface {
	AfterInsert(s *Session) error
}

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	var hookFunc reflect.Value

	// Check if value implements any of the hook interfaces
	switch method {
	case BeforeQuery:
		if hook, ok := value.(IBeforeQuery); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case AfterQuery:
		if hook, ok := value.(IAfterQuery); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case BeforeUpdate:
		if hook, ok := value.(IBeforeUpdate); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case AfterUpdate:
		if hook, ok := value.(IAfterUpdate); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case BeforeDelete:
		if hook, ok := value.(IBeforeDelete); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case AfterDelete:
		if hook, ok := value.(IAfterDelete); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case BeforeInsert:
		if hook, ok := value.(IBeforeInsert); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	case AfterInsert:
		if hook, ok := value.(IAfterInsert); ok {
			hookFunc = reflect.ValueOf(hook).MethodByName(method)
		}
	}

	// If hookFunc is found, call it with the session as parameter
	if hookFunc.IsValid() {
		param := []reflect.Value{reflect.ValueOf(s)}
		if v := hookFunc.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
}
