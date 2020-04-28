package validator

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// DefaultValidator 验证器
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

// ValidateStruct 如果接收到的类型是一个结构体或指向结构体的指针，则执行验证。
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		//如果传递不合规则的值，则返回InvalidValidationError，否则返回nil。
		///如果返回err != nil，可通过err.(validator.ValidationErrors)来访问错误数组。
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine 返回支持`StructValidator`实现的底层验证引擎
func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("validate")
		// v8版本，v8版本使用"binding"
		// v.validate.SetTagName("binding")
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
