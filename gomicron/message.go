package gomicron

type Polarity int

const (
	Positive Polarity = iota
	Negative
)

//type ActualFormatter[T any, U any] func(expected T, observed U) string
//
//type Reporter[T any, U any] struct {
//	Expected ActualFormatter[T, U]
//	To       ActualFormatter[T, U]
//	But      ActualFormatter[T, U]
//}
//
//func NewReporter[T any, U any]() Reporter[T, U] {
//	return Reporter[T, U]{
//		Expected: func(_ T, observed U) string {
//			return fmt.Sprintf(`Expected "%+v"`, observed)
//		},
//		To: func(expected T, _ U) string {
//			return fmt.Sprintf(`to match "%+v"`, expected)
//		},
//		But: func(_ T, observed U) string {
//			return ""
//		},
//	}
//}
