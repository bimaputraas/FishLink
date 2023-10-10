package consumer

type Consumer interface {
	ConsumeQueuedMessage(queueName string)
}
