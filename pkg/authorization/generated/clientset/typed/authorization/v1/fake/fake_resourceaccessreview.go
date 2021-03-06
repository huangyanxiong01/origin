package fake

import (
	v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeResourceAccessReviews implements ResourceAccessReviewInterface
type FakeResourceAccessReviews struct {
	Fake *FakeAuthorizationV1
}

var resourceaccessreviewsResource = schema.GroupVersionResource{Group: "authorization.openshift.io", Version: "v1", Resource: "resourceaccessreviews"}

var resourceaccessreviewsKind = schema.GroupVersionKind{Group: "authorization.openshift.io", Version: "v1", Kind: "ResourceAccessReview"}

// Create takes the representation of a resourceAccessReview and creates it.  Returns the server's representation of the resourceAccessReview, and an error, if there is any.
func (c *FakeResourceAccessReviews) Create(resourceAccessReview *v1.ResourceAccessReview) (result *v1.ResourceAccessReview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(resourceaccessreviewsResource, resourceAccessReview), &v1.ResourceAccessReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceAccessReview), err
}
