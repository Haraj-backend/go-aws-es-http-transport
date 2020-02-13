package transport

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

// AWSESTransport is used for providing signed http request
// for accessing AWS Elasticsearch Cluster using IAM
type AWSESTransport struct {
	signer *v4.Signer
	region string
}

// NewAWSESTransport returns new instance of AWSESTransport
func NewAWSESTransport(sess *session.Session) *AWSESTransport {
	region := *sess.Config.Region
	if len(region) == 0 {
		// try to resolve region from metadata service
		meta := ec2metadata.New(sess)
		region, _ = meta.Region()
	}
	return &AWSESTransport{
		signer: v4.NewSigner(sess.Config.Credentials),
		region: region,
	}
}

// RoundTrip is used for execute single HTTP transaction returning response
func (t *AWSESTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	now := time.Now()
	serviceName := "es"
	if req.Body != nil {
		b, _ := ioutil.ReadAll(req.Body)
		t.signer.Sign(req, bytes.NewReader(b), serviceName, t.region, now)
	} else {
		t.signer.Sign(req, nil, serviceName, t.region, now)
	}
	return http.DefaultTransport.RoundTrip(req)
}
