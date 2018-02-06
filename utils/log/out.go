package log

type discardOut struct {}

func (discardOut) Write(p []byte) (n int, err error) {
    return 0, nil
}

