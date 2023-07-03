package errorpkg

type HTTPError struct {
	Code int
	Err  error
}
