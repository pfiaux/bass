package bass

import (
	"fmt"
	"reflect"
)

type Builtin struct {
	Name      string
	Func      reflect.Value
	Operative bool
}

var _ Value = (*Builtin)(nil)

func (value *Builtin) String() string {
	return fmt.Sprintf("<builtin op: %s>", value.Name)
}

func (value *Builtin) Decode(dest interface{}) error {
	switch x := dest.(type) {
	case **Builtin:
		*x = value
		return nil
	case *Combiner:
		*x = value
		return nil
	default:
		return DecodeError{
			Source:      value,
			Destination: dest,
		}
	}
}

func (value *Builtin) Eval(*Env) (Value, error) {
	return value, nil
}

func Op(name string, f interface{}) *Builtin {
	fun := reflect.ValueOf(f)
	if fun.Kind() != reflect.Func {
		panic("Op takes a func()")
	}

	return &Builtin{
		Name:      name,
		Func:      fun,
		Operative: true,
	}
}

func Func(name string, f interface{}) Combiner {
	op := Op(name, f)
	op.Operative = false
	return Applicative{op}
}

var valType = reflect.TypeOf((*Value)(nil)).Elem()
var errType = reflect.TypeOf((*error)(nil)).Elem()

func (builtin Builtin) Call(val Value, env *Env) (Value, error) {
	args := []Value{}
	if builtin.Operative {
		args = append(args, env)
	}

	var list List
	err := val.Decode(&list)
	if err != nil {
		return nil, ErrBadSyntax
	}

	for list != (Empty{}) {
		args = append(args, list.First())

		err = list.Rest().Decode(&list)
		if err != nil {
			return nil, ErrBadSyntax
		}
	}

	ftype := builtin.Func.Type()

	argc := ftype.NumIn()
	if ftype.IsVariadic() {
		argc--

		if len(args) < argc {
			return nil, ArityError{
				Name:     builtin.Name,
				Need:     argc,
				Have:     len(args),
				Variadic: true,
			}
		}
	} else if len(args) != argc {
		return nil, ArityError{
			Name: builtin.Name,
			Need: argc,
			Have: len(args),
		}
	}

	fargs := []reflect.Value{}

	for i := 0; i < argc; i++ {
		t := ftype.In(i)

		dest := reflect.New(t)
		if t == valType { // pass Value with no conversion
			dest.Elem().Set(reflect.ValueOf(args[i]))
		} else {
			err := args[i].Decode(dest.Interface())
			if err != nil {
				return nil, err
			}
		}

		fargs = append(fargs, dest.Elem())
	}

	if ftype.IsVariadic() {
		variadic := args[argc:]
		variadicType := ftype.In(argc)

		subType := variadicType.Elem()
		for _, varg := range variadic {
			dest := reflect.New(subType)
			if subType == valType { // pass Value with no conversion
				dest.Elem().Set(reflect.ValueOf(varg))
			} else {
				err := varg.Decode(dest.Interface())
				if err != nil {
					return nil, err
				}
			}

			fargs = append(fargs, dest.Elem())
		}
	}

	result := builtin.Func.Call(fargs)

	switch ftype.NumOut() {
	case 0:
		return Null{}, nil
	case 1:
		if ftype.Out(0) == errType {
			if !result[0].IsNil() {
				return nil, result[0].Interface().(error)
			}

			return Null{}, nil
		}

		return ValueOf(result[0].Interface())
	case 2:
		if ftype.Out(1) != errType {
			return nil, fmt.Errorf("multiple return values are not supported")
		}

		if !result[1].IsNil() {
			return nil, result[1].Interface().(error)
		}

		return ValueOf(result[0].Interface())
	}

	return nil, nil
}
