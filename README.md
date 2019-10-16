# go-core
My core stuff with types, code, etc that I use almost everywhere

## core.SendRequest

This func allows to send HTTP request to REST servers and takes care of payloads, JSON, result collection.

Examples:

```go
res, err := core.SendRequest(&core.RequestOptions{
    URL: myURL,
}, nil)
if err != nil {
    return err
}
data := struct{Data string}{}
err := res.UnmarshalContentJSON(&data)
```
Here we send an HTTP GET request and unmarshal the response (a `ContentReader`).

It is also possible to let SendRequest do the unmarshal for us:

```go
data := struct{Data string}{}
_, err := core.SendRequest(&core.RequestOptions{
    URL: myURL,
}, &data)
if err != nil {
    return err
}
```

Authorization can be stored in the `RequestOptions.Authorization`:

```go
payload := struct{Key string}{}
data := struct{Data string}{}
_, err := core.SendRequest(&core.RequestOptions{
    URL:     myURL,
    Authorization: "Basic sdfgsdfgsdfgdsfgw42agoi0s9ix"
}, &data)
if err != nil {
    return err
}
```

Objects can be sent as payloads:

```go
payload := struct{Key string}{}
data := struct{Data string}{}
_, err := core.SendRequest(&core.RequestOptions{
    URL:     myURL,
    Payload: payload,
}, &data)
if err != nil {
    return err
}
```

A payload will induce an HTTP POST unless mentioned.

```go
payload := struct{Key string}{}
data := struct{Data string}{}
_, err := core.SendRequest(&core.RequestOptions{
    URL:     myURL,
    Payload: payload,
}, &data)
if err != nil {
    return err
}
```

So, to send an `HTTP UPDATE`, simply:

```go
payload := struct{Key string}{}
data := struct{Data string}{}
_, err := core.SendRequest(&core.RequestOptions{
    Method:  http.MethodUPDAE,
    URL:     myURL,
    Payload: payload,
}, &data)
if err != nil {
    return err
}
```

if the PayloadType is not mentioned, it is calculated when processing the Payload.

if the payload is a `ContentReader` or a `Content`, it is used directly.

if the payload is a `map[string]string`

if the payload is a struct{}, this func will send the body as `application/json` and will marshal it ino