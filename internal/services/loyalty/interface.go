package loyalty

type loyaltyStorage interface {
	GetUser(username string) (Loyalty, error)
	GetStatus(username string) (string, error)
	IncrementCounter(username string) error
	DecrementCounter(username string) error
}
