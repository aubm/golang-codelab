# Amount converter

The goal of this exercise is to build a small web API that is able to convert a given amount from one currency to another.
You should start with the code in `$GOPATH/src/github.com/aubm/golang-codelab/exercises/workshops/amount-converter/main.go`.
Feel free to create as many sub-packages as you feel it is relevant.

## Requirements

The web API must implement one service accessible from the path `/convert`.
When using the service, the user provides the amount to convert, the base currency and the target currency.
Then the service must respond with the converted amount.

Here is an example request that illustrates the nominal case:

```http
POST /convert HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
	"currency": "EUR",
	"target_currency": "USD",
	"amount": 1
}
```

And the response:

```json
{
    "converted_amount": 1.1221
}
```

Here are the acceptance criteria:

- the service must respond with the content type `application/json`
- when the user provides an invalid JSON request, the service must respond with a `400` status code
- when the provided base currency is invalid, the service must respond with a `400` status code
- when the provided target currency is invalid, the service must respond with a `400` status code
- for any other errors, the service must respond with a `500` status code

**Hint**: conversion rates change from one day to another, to get the conversion rates for today and for a given money, you should use the Fixer API: http://api.fixer.io/latest?base=EUR.

## How do I know when I'm finished?

Let your program run in a terminal:

```bash
$ go run $GOPATH/src/github.com/aubm/golang-codelab/exercises/workshops/amount-converter/*.go
2017/06/12 21:20:25 Server started on port 8080
```

Then open a new terminal, and use the following command:

```bash
$ go test github.com/aubm/golang-codelab/exercises/workshops/amount-converter/tests
```

When the job is done, the command should output something like:

```
ok  	github.com/aubm/golang-codelab/exercises/workshops/amount-converter/tests	0.405s
```

If not, read the tests error output to get an understanding of what you might have missed.
When you are ready to test again, simply re-run the command. Just don't forget to restart the application for your
changes to take effect.
