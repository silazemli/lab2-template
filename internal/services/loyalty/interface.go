package loyalty

type loyaltyStorage interface {
	GetStatus(username string) (string, error)
	IncrementCounter(username string) error
	DecrementCounter(username string) error
}
