package validator

// var (
//  EmailX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// )

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, msg string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = msg
	}
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddError(key, msg)
	}
}

func In(v string, list ...string) bool {
	for i := range list {
		if v == list[i] {
			return true
		}
	}
	return false
}

func Unique(vs []string) bool {
	unique_vs := make(map[string]bool)

	for _, v := range vs {
		unique_vs[v] = true
	}

	return len(vs) == len(unique_vs)
}

// func Match(v string, x *regexp.Regexp) bool {
// 	return x.MatchString(v)
// }
