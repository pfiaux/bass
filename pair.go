package bass

import (
	"fmt"
)

type Pair struct {
	A Value
	D Value
}

var _ Value = Pair{}

func (value Pair) String() string {
	return formatList(value, "(", ")")
}

func (value Pair) Decode(dest interface{}) error {
	switch x := dest.(type) {
	case *Pair:
		*x = value
		return nil
	case *List:
		*x = value
		return nil
	}

	return DecodeError{
		Source:      value,
		Destination: dest,
	}
}

var _ List = Pair{}

func (value Pair) First() Value {
	return value.A
}

func (value Pair) Rest() Value {
	return value.D
}

// Pair combines the first operand with the second operand.
//
// If the first value is not a Combiner, an error is returned.
func (value Pair) Eval(env *Env) (Value, error) {
	f, err := value.A.Eval(env)
	if err != nil {
		return nil, err
	}

	var combiner Combiner
	err = f.Decode(&combiner)
	if err != nil {
		return nil, fmt.Errorf("apply %s: %w", f, err)
	}

	return combiner.Call(value.D, env)
}

func formatList(list List, odelim, cdelim string) string {
	out := odelim

	for list != (Empty{}) {
		out += list.First().String()

		var empty Empty
		err := list.Rest().Decode(&empty)
		if err == nil {
			break
		}

		err = list.Rest().Decode(&list)
		if err != nil {
			out += " . "
			out += list.Rest().String()
			break
		}

		out += " "
	}

	out += cdelim

	return out
}
