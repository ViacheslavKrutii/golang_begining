package myinterfaces

type Parcel interface {
	send()
}

type envelope struct {
	addressOfSender    string
	addressOfRecipient string
}

func (e envelope) send() {

}

type box struct {
	addressOfSender    string
	addressOfRecipient string
}

func (b box) send() {

}
