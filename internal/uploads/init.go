package uploads

func init() {
	go ConsumeFromAccountsTopic()
}
