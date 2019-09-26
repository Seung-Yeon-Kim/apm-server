// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateListenerInput
type CreateListenerInput struct {
	_ struct{} `type:"structure"`

	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN
	// but do not set IsDefault.
	//
	// To create a certificate list for the listener, use AddListenerCertificates.
	Certificates []Certificate `type:"list"`

	// The actions for the default rule. The rule must include one forward action
	// or one or more fixed-response actions.
	//
	// If the action type is forward, you specify a target group. The protocol of
	// the target group must be HTTP or HTTPS for an Application Load Balancer.
	// The protocol of the target group must be TCP, TLS, UDP, or TCP_UDP for a
	// Network Load Balancer.
	//
	// [HTTPS listeners] If the action type is authenticate-oidc, you authenticate
	// users through an identity provider that is OpenID Connect (OIDC) compliant.
	//
	// [HTTPS listeners] If the action type is authenticate-cognito, you authenticate
	// users through the user pools supported by Amazon Cognito.
	//
	// [Application Load Balancer] If the action type is redirect, you redirect
	// specified client requests from one URL to another.
	//
	// [Application Load Balancer] If the action type is fixed-response, you drop
	// specified client requests and return a custom HTTP response.
	//
	// DefaultActions is a required field
	DefaultActions []Action `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`

	// The port on which the load balancer is listening.
	//
	// Port is a required field
	Port *int64 `min:"1" type:"integer" required:"true"`

	// The protocol for connections from clients to the load balancer. For Application
	// Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP.
	//
	// Protocol is a required field
	Protocol ProtocolEnum `type:"string" required:"true" enum:"true"`

	// [HTTPS and TLS listeners] The security policy that defines which ciphers
	// and protocols are supported. The default is the current predefined security
	// policy.
	SslPolicy *string `type:"string"`
}

// String returns the string representation
func (s CreateListenerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateListenerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateListenerInput"}

	if s.DefaultActions == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultActions"))
	}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if s.Port == nil {
		invalidParams.Add(aws.NewErrParamRequired("Port"))
	}
	if s.Port != nil && *s.Port < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Port", 1))
	}
	if len(s.Protocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}
	if s.DefaultActions != nil {
		for i, v := range s.DefaultActions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DefaultActions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateListenerOutput
type CreateListenerOutput struct {
	_ struct{} `type:"structure"`

	// Information about the listener.
	Listeners []Listener `type:"list"`
}

// String returns the string representation
func (s CreateListenerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateListener = "CreateListener"

// CreateListenerRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Creates a listener for the specified Application Load Balancer or Network
// Load Balancer.
//
// To update a listener, use ModifyListener. When you are finished with a listener,
// you can delete it using DeleteListener. If you are finished with both the
// listener and the load balancer, you can delete them both using DeleteLoadBalancer.
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple listeners with the same settings, each
// call succeeds.
//
// For more information, see Listeners for Your Application Load Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)
// in the Application Load Balancers Guide and Listeners for Your Network Load
// Balancers (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html)
// in the Network Load Balancers Guide.
//
//    // Example sending a request using CreateListenerRequest.
//    req := client.CreateListenerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/CreateListener
func (c *Client) CreateListenerRequest(input *CreateListenerInput) CreateListenerRequest {
	op := &aws.Operation{
		Name:       opCreateListener,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateListenerInput{}
	}

	req := c.newRequest(op, input, &CreateListenerOutput{})
	return CreateListenerRequest{Request: req, Input: input, Copy: c.CreateListenerRequest}
}

// CreateListenerRequest is the request type for the
// CreateListener API operation.
type CreateListenerRequest struct {
	*aws.Request
	Input *CreateListenerInput
	Copy  func(*CreateListenerInput) CreateListenerRequest
}

// Send marshals and sends the CreateListener API request.
func (r CreateListenerRequest) Send(ctx context.Context) (*CreateListenerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateListenerResponse{
		CreateListenerOutput: r.Request.Data.(*CreateListenerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateListenerResponse is the response type for the
// CreateListener API operation.
type CreateListenerResponse struct {
	*CreateListenerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateListener request.
func (r *CreateListenerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}