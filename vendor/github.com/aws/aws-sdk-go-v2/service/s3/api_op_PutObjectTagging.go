// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutObjectTaggingInput struct {
	_ struct{} `type:"structure" payload:"Tagging"`

	// The bucket name containing the object.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Name of the tag.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Container for the TagSet and Tag elements
	//
	// Tagging is a required field
	Tagging *Tagging `locationName:"Tagging" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The versionId of the object that the tag-set will be added to.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s PutObjectTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Tagging == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tagging"))
	}
	if s.Tagging != nil {
		if err := s.Tagging.Validate(); err != nil {
			invalidParams.AddNested("Tagging", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutObjectTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.Tagging != nil {
		v := s.Tagging

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "Tagging", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *PutObjectTaggingInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutObjectTaggingInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutObjectTaggingOutput struct {
	_ struct{} `type:"structure"`

	// The versionId of the object the tag-set was added to.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s PutObjectTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutObjectTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	return nil
}

const opPutObjectTagging = "PutObjectTagging"

// PutObjectTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the supplied tag-set to an object that already exists in a bucket
//
// A tag is a key-value pair. You can associate tags with an object by sending
// a PUT request against the tagging subresource that is associated with the
// object. You can retrieve tags by sending a GET request. For more information,
// see GetObjectTagging.
//
// For tagging-related restrictions related to characters and encodings, see
// Tag Restrictions (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html).
// Note that Amazon S3 limits the maximum number of tags to 10 tags per object.
//
// To use this operation, you must have permission to perform the s3:PutObjectTagging
// action. By default, the bucket owner has this permission and can grant this
// permission to others.
//
// To put tags of any other version, use the versionId query parameter. You
// also need permission for the s3:PutObjectVersionTagging action.
//
// For information about the Amazon S3 object tagging feature, see Object Tagging
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html).
//
// Special Errors
//
//    * Code: InvalidTagError Cause: The tag provided was not a valid tag. This
//    error can occur if the tag did not pass input validation. For more information,
//    see Object Tagging (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html).
//
//    * Code: MalformedXMLError Cause: The XML provided does not match the schema.
//
//    * Code: OperationAbortedError Cause: A conflicting conditional operation
//    is currently in progress against this resource. Please try again.
//
//    * Code: InternalError Cause: The service was unable to apply the provided
//    tag to the object.
//
// Related Resources
//
//    * GetObjectTagging
//
//    // Example sending a request using PutObjectTaggingRequest.
//    req := client.PutObjectTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectTagging
func (c *Client) PutObjectTaggingRequest(input *PutObjectTaggingInput) PutObjectTaggingRequest {
	op := &aws.Operation{
		Name:       opPutObjectTagging,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}?tagging",
	}

	if input == nil {
		input = &PutObjectTaggingInput{}
	}

	req := c.newRequest(op, input, &PutObjectTaggingOutput{})
	return PutObjectTaggingRequest{Request: req, Input: input, Copy: c.PutObjectTaggingRequest}
}

// PutObjectTaggingRequest is the request type for the
// PutObjectTagging API operation.
type PutObjectTaggingRequest struct {
	*aws.Request
	Input *PutObjectTaggingInput
	Copy  func(*PutObjectTaggingInput) PutObjectTaggingRequest
}

// Send marshals and sends the PutObjectTagging API request.
func (r PutObjectTaggingRequest) Send(ctx context.Context) (*PutObjectTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectTaggingResponse{
		PutObjectTaggingOutput: r.Request.Data.(*PutObjectTaggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectTaggingResponse is the response type for the
// PutObjectTagging API operation.
type PutObjectTaggingResponse struct {
	*PutObjectTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectTagging request.
func (r *PutObjectTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
