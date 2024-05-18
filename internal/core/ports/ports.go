package ports

type LoggerService interface {
	Info(message string)
	Warning(message string)
	Error(message string)
}
