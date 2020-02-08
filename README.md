# Go AWS ElasticSearch HTTP Transport

Have you ever face the hardship when trying to connect to AWS ElasticSearch Domain using IAM users especially in Go?

Yet the basic idea to the solution is actually very simple, we just need to sign every http request to our ElasticSearch Domain with [AWS Signature V4](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). Essentially this library give you just that. 😉

## How to Use

The library usage is very simple. We just need to set as transport in our HTTP Client like following:

```go
...
client := &http.Client{
    Transport: transport.NewAWSESTransport(session.New()),
}
...
```

Or if you are using [Official ElasticSearch Client for Go](https://github.com/elastic/go-elasticsearch) like I usually did, you could just do following:

```go
...
client, err := es6.NewClient(es6.Config{
    Addresses: []string{endpoint},
    Transport: transport.NewAWSESTransport(session.New()),
})
...
```

So simple, right? 😉

For more details please check out [examples](./examples/) directory.

## Acknowledgement

This library is inspired by following resources:

- https://github.com/elastic/go-elasticsearch/issues/78
- https://gist.github.com/jriquelme/d11e0bb2e4523523a4ee507282ba6184
- https://github.com/aws/aws-sdk-go/issues/676
- https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-request-signing.html#es-request-signing-go

I'm just trying to make everything simpler and cleaner. 😉
