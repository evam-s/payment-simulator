package uploads

func init() {
	go ConsumeFromBanksTopic()
	go ConsumeFromAccountsTopic()
}
