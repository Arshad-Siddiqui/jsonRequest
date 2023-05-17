# jsonRequest

## Description

This package was made for quickly and easily creating json requests for use in testing (in particular with [GoFiber](github.com/gofiber/fiber/v2)).

## Usage

### Creating a request

```go
person := Person{
    Name: "John Doe",
    Age:  42,
}

// First value is the path, second is the body
req, _ := jsonRequest.PostJsonReq("/", person)
app.test(req)
Expect(resp.StatusCode).To(Equal(http.StatusCreated))
```

> Note: [gomega](github.com/onsi/gomega) is used to compare status codes and app is an instance of a fiber app.
