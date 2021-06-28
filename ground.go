package bass

import (
	"embed"
)

//go:embed std/*.bass
var std embed.FS

var ground = NewEnv()

func init() {
	for _, pred := range primPreds {
		ground.Set(pred.name, Func(string(pred.name), pred.check))
	}

	ground.Set("ground", ground, `ground environment please ignore`,
		`This value is only here to aid in developing prior to first release.`,
		`Fetching this binding voids your warranty.`)

	ground.Set("cons",
		Func("cons", func(a, d Value) Value {
			return Pair{a, d}
		}),
		`construct a pair from the given values`)

	ground.Set("wrap",
		Func("wrap", func(c Combiner) Applicative {
			return Applicative{c}
		}),
		`construct an applicative from a combiner (typically an operative)`)

	ground.Set("unwrap",
		Func("unwrap", func(a Applicative) Combiner {
			return a.Underlying
		}),
		`access an applicative's underlying combiner`)

	ground.Set("op",
		Op("op", func(env *Env, formals, eformal, body Value) *Operative {
			return &Operative{
				Env:     env,
				Formals: formals,
				Eformal: eformal,
				Body:    body,
			}
		}),
		`a primitive operative constructor`,
		`op is redefined later, so no one should see this comment.`)

	ground.Set("eval",
		Func("eval", func(val Value, env *Env) (Value, error) {
			return val.Eval(env)
		}),
		`evaluate a value in an env`)

	ground.Set("make-env",
		Func("make-env", func(envs ...*Env) *Env {
			return NewEnv(envs...)
		}),
		`construct an env with the given parents`)

	ground.Set("def",
		Op("def", func(env *Env, formals, val Value) (Value, error) {
			res, err := val.Eval(env)
			if err != nil {
				return nil, err
			}

			err = env.Define(formals, res)
			if err != nil {
				return nil, err
			}

			return formals, nil
		}),
		`bind symbols to values in the current env`)

	ground.Set("doc",
		Op("doc", PrintDocs),
		`print docs for symbols`,
		`Prints the documentation for the given symbols resolved from the current environment.`,
		`With no arguments, prints the commentary for the current environment.`)

	ground.Set("if",
		Op("if", func(env *Env, cond, yes, no Value) (Value, error) {
			cond, err := cond.Eval(env)
			if err != nil {
				return nil, err
			}

			var res bool
			err = cond.Decode(&res)
			if err != nil {
				return yes.Eval(env)
			}

			if !res {
				return no.Eval(env)
			}

			return yes.Eval(env)
		}),
		`if then else (branching logic)`,
		`Evaluates a condition. If nil or false, evaluates the third operand. Otherwise, evaluates the second operand.`)

	ground.Set("+",
		Func("+", func(nums ...int) int {
			sum := 0
			for _, num := range nums {
				sum += num
			}

			return sum
		}),
		`sum the given numbers`)

	ground.Set("*",
		Func("*", func(nums ...int) int {
			mul := 1
			for _, num := range nums {
				mul *= num
			}

			return mul
		}),
		`multiply the given numbers`)

	ground.Set("-",
		Func("-", func(num int, nums ...int) int {
			if len(nums) == 0 {
				return -num
			}

			sub := num
			for _, num := range nums {
				sub -= num
			}

			return sub
		}),
		`subtract ys from x`,
		`If only x is given, returns the negation of x.`)

	ground.Set("max", Func("max", func(num int, nums ...int) int {
		max := num
		for _, num := range nums {
			if num > max {
				max = num
			}
		}

		return max
	}),
		`largest number given`)

	ground.Set("min",
		Func("min", func(num int, nums ...int) int {
			min := num
			for _, num := range nums {
				if num < min {
					min = num
				}
			}

			return min
		}),
		`smallest number given`)

	ground.Set("=?",
		Func("=?", func(cur int, nums ...int) bool {
			for _, num := range nums {
				if num != cur {
					return false
				}
			}

			return true
		}),
		`numeric equality`,
	)

	ground.Set(">?",
		Func(">?", func(num int, nums ...int) bool {
			min := num
			for _, num := range nums {
				if num >= min {
					return false
				}

				min = num
			}

			return true
		}),
		`descending order`)

	ground.Set(">=?",
		Func(">=?", func(num int, nums ...int) bool {
			max := num
			for _, num := range nums {
				if num > max {
					return false
				}

				max = num
			}

			return true
		}),
		`descending or equal order`)

	ground.Set("<?",
		Func("<?", func(num int, nums ...int) bool {
			max := num
			for _, num := range nums {
				if num <= max {
					return false
				}

				max = num
			}

			return true
		}),
		`increasing order`)

	ground.Set("<=?",
		Func("<=?", func(num int, nums ...int) bool {
			max := num
			for _, num := range nums {
				if num < max {
					return false
				}

				max = num
			}

			return true
		}),
		`increasing or equal order`)
	for _, lib := range []string{
		"std/root.bass",
	} {
		file, err := std.Open(lib)
		if err != nil {
			panic(err)
		}

		_, err = EvalReader(ground, file)
		if err != nil {
			panic(err)
		}
	}
}

type primPred struct {
	name  Symbol
	check func(Value) bool
}

// basic predicates built in to the language.
//
// these are also used in (doc) to show which predicates a value satisfies.
var primPreds = []primPred{
	// primitive type checks
	{"null?", func(val Value) bool {
		var x Null
		return val.Decode(&x) == nil
	}},
	{"boolean?", func(val Value) bool {
		var x Bool
		return val.Decode(&x) == nil
	}},
	{"number?", func(val Value) bool {
		var x Int
		return val.Decode(&x) == nil
	}},
	{"string?", func(val Value) bool {
		var x String
		return val.Decode(&x) == nil
	}},
	{"symbol?", func(val Value) bool {
		var x Symbol
		return val.Decode(&x) == nil
	}},
	{"env?", func(val Value) bool {
		var x *Env
		return val.Decode(&x) == nil
	}},
	{"list?", func(val Value) bool {
		return IsList(val)
	}},
	{"pair?", func(val Value) bool {
		var x Pair
		return val.Decode(&x) == nil
	}},
	{"applicative?", func(val Value) bool {
		var x Applicative
		return val.Decode(&x) == nil
	}},
	{"operative?", func(val Value) bool {
		var b *Builtin
		if val.Decode(&b) == nil {
			return b.Operative
		}

		var o *Operative
		return val.Decode(&o) == nil
	}},
	{"combiner?", func(val Value) bool {
		var x Combiner
		return val.Decode(&x) == nil
	}},
	{"empty?", func(val Value) bool {
		var empty Empty
		if err := val.Decode(&empty); err == nil {
			return true
		}

		var str string
		if err := val.Decode(&str); err == nil {
			return str == ""
		}

		var nul Null
		if err := val.Decode(&nul); err == nil {
			return true
		}

		return false
	}},
}
