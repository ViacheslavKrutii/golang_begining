package myinterfaces

type content struct {
}

type Parcel interface {
	send(from, to string) Parcel
}

type envelope struct {
	addressOfSender    string
	addressOfRecipient string
	content            content
}

func (e envelope) send(from, to string) Parcel {
	e.addressOfSender = from
	e.addressOfRecipient = to
	return e
}

type box struct {
	addressOfSender    string
	addressOfRecipient string
}

func (b box) send(from, to string) Parcel {
	b.addressOfSender = from
	b.addressOfRecipient = to
	return b
}

func sortSection(p Parcel, b, e chan Parcel) {
	p, ok := p.(box)
	if ok {
		b <- p
		return
	}
	p, ok = p.(envelope)
	if ok {
		e <- p
		return
	}
}
