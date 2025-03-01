package ssg

func NewProps(props ...string) map[string]string {
	m := make(map[string]string)
	for i := 0; i < len(props); i += 2 {
		if i+1 >= len(props) {
			panic("props must be in key-value pairs")
		}
		m[props[i]] = props[i+1]
	}
	return m
}
