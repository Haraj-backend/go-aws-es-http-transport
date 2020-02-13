# Go AWS ElasticSearch HTTP Transport

Have you ever face confusion when trying to establish IAM access to AWS ElasticSearch Domain using Go?

The solution is actually very simple, we just need to sign every http request to our ElasticSearch Domain with correct [AWS Signature V4](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). This library help you to do just that. 😉

## How to Use

The library usage is very simple. We just need to set it up as our HTTP Client transport:

```go
...
client := &http.Client{
    Transport: transport.NewAWSESTransport(session.New()),
}
...
```

Or if you are using [Official ElasticSearch Client for Go](https://github.com/elastic/go-elasticsearch), you could just do like following:

```go
...
client, err := es6.NewClient(es6.Config{
    Addresses: []string{endpoint},
    Transport: transport.NewAWSESTransport(session.New()),
})
...
```

So simple, right? 😉

For more details please check out our [examples](./examples/) directory.

## Acknowledgement

This library is inspired by following resources:

- https://github.com/elastic/go-elasticsearch/issues/78
- https://gist.github.com/jriquelme/d11e0bb2e4523523a4ee507282ba6184
- https://github.com/aws/aws-sdk-go/issues/676
- https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-request-signing.html#es-request-signing-go

We are just trying to make everything simpler and cleaner. 😉
