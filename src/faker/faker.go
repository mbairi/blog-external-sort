package faker

import (
	"blog-external-sort/src/cio"

	"github.com/go-faker/faker/v4"
)

type Person struct {
	UUID             string            `faker:"uuid_digit"`
	Name             string            `faker:"name"`
	UserName         string            `faker:"username"`
	PhoneNumber      string            `faker:"phone_number"`
	RealAddress      faker.RealAddress `faker:"real_address"`
	CreditCardNumber string            `faker:"cc_number"`
	CreditCardType   string            `faker:"cc_type"`
	Email            string            `faker:"email"`
	PaymentMethod    string            `faker:"oneof: cc, paypal, check, money order"` // oneof will randomly pick one of the comma-separated values supplied in the tag
}

func GenerateMockFile(limit int) {
	writer, _ := cio.NewJsonWriter("mock.jsonl")
	for range limit {
		a := Person{}
		faker.FakeData(&a)
		writer.WriteLine(a)
	}
	writer.Close()
}
