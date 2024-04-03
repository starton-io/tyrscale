package monitoring

type IExternalMonitoring interface {
	Check(testCmd string) (bool, error)
}
