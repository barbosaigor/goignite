// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type DeleteBucketCorsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the bucket whose cors configuration is being deleted.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketCorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketCorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBucketCorsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteBucketCorsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketCorsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *DeleteBucketCorsInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *DeleteBucketCorsInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type DeleteBucketCorsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketCorsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBucketCorsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBucketCors = "DeleteBucketCors"

// DeleteBucketCorsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Deletes the cors configuration information set for the bucket.
//
// To use this operation, you must have permission to perform the s3:PutBucketCORS
// action. The bucket owner has this permission by default and can grant this
// permission to others.
//
// For information about cors, see Enabling Cross-Origin Resource Sharing (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html)
// in the Amazon Simple Storage Service Developer Guide.
//
// Related Resources:
//
//    *
//
//    * RESTOPTIONSobject
//
//    // Example sending a request using DeleteBucketCorsRequest.
//    req := client.DeleteBucketCorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/DeleteBucketCors
func (c *Client) DeleteBucketCorsRequest(input *DeleteBucketCorsInput) DeleteBucketCorsRequest {
	op := &aws.Operation{
		Name:       opDeleteBucketCors,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?cors",
	}

	if input == nil {
		input = &DeleteBucketCorsInput{}
	}

	req := c.newRequest(op, input, &DeleteBucketCorsOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBucketCorsRequest{Request: req, Input: input, Copy: c.DeleteBucketCorsRequest}
}

// DeleteBucketCorsRequest is the request type for the
// DeleteBucketCors API operation.
type DeleteBucketCorsRequest struct {
	*aws.Request
	Input *DeleteBucketCorsInput
	Copy  func(*DeleteBucketCorsInput) DeleteBucketCorsRequest
}

// Send marshals and sends the DeleteBucketCors API request.
func (r DeleteBucketCorsRequest) Send(ctx context.Context) (*DeleteBucketCorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBucketCorsResponse{
		DeleteBucketCorsOutput: r.Request.Data.(*DeleteBucketCorsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBucketCorsResponse is the response type for the
// DeleteBucketCors API operation.
type DeleteBucketCorsResponse struct {
	*DeleteBucketCorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBucketCors request.
func (r *DeleteBucketCorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
