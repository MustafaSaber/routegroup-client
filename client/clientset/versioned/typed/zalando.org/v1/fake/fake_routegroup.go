//lint:file-ignore ST1005 Generated code

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteGroups implements RouteGroupInterface
type FakeRouteGroups struct {
	Fake *FakeZalandoV1
	ns   string
}

var routegroupsResource = v1.SchemeGroupVersion.WithResource("routegroups")

var routegroupsKind = v1.SchemeGroupVersion.WithKind("RouteGroup")

// Get takes name of the routeGroup, and returns the corresponding routeGroup object, and an error if there is any.
func (c *FakeRouteGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routegroupsResource, c.ns, name), &v1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RouteGroup), err
}

// List takes label and field selectors, and returns the list of RouteGroups that match those selectors.
func (c *FakeRouteGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RouteGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routegroupsResource, routegroupsKind, c.ns, opts), &v1.RouteGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RouteGroupList{ListMeta: obj.(*v1.RouteGroupList).ListMeta}
	for _, item := range obj.(*v1.RouteGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeGroups.
func (c *FakeRouteGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routegroupsResource, c.ns, opts))

}

// Create takes the representation of a routeGroup and creates it.  Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *FakeRouteGroups) Create(ctx context.Context, routeGroup *v1.RouteGroup, opts metav1.CreateOptions) (result *v1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routegroupsResource, c.ns, routeGroup), &v1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RouteGroup), err
}

// Update takes the representation of a routeGroup and updates it. Returns the server's representation of the routeGroup, and an error, if there is any.
func (c *FakeRouteGroups) Update(ctx context.Context, routeGroup *v1.RouteGroup, opts metav1.UpdateOptions) (result *v1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routegroupsResource, c.ns, routeGroup), &v1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RouteGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouteGroups) UpdateStatus(ctx context.Context, routeGroup *v1.RouteGroup, opts metav1.UpdateOptions) (*v1.RouteGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routegroupsResource, "status", c.ns, routeGroup), &v1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RouteGroup), err
}

// Delete takes name of the routeGroup and deletes it. Returns an error if one occurs.
func (c *FakeRouteGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routegroupsResource, c.ns, name, opts), &v1.RouteGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RouteGroupList{})
	return err
}

// Patch applies the patch and returns the patched routeGroup.
func (c *FakeRouteGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RouteGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routegroupsResource, c.ns, name, pt, data, subresources...), &v1.RouteGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RouteGroup), err
}
