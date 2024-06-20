// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineConfigPools implements MachineConfigPoolInterface
type FakeMachineConfigPools struct {
	Fake *FakeMachineconfigurationV1
}

var machineconfigpoolsResource = v1.SchemeGroupVersion.WithResource("machineconfigpools")

var machineconfigpoolsKind = v1.SchemeGroupVersion.WithKind("MachineConfigPool")

// Get takes name of the machineConfigPool, and returns the corresponding machineConfigPool object, and an error if there is any.
func (c *FakeMachineConfigPools) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MachineConfigPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machineconfigpoolsResource, name), &v1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.MachineConfigPool), err
}

// List takes label and field selectors, and returns the list of MachineConfigPools that match those selectors.
func (c *FakeMachineConfigPools) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MachineConfigPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machineconfigpoolsResource, machineconfigpoolsKind, opts), &v1.MachineConfigPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.MachineConfigPoolList{ListMeta: obj.(*v1.MachineConfigPoolList).ListMeta}
	for _, item := range obj.(*v1.MachineConfigPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineConfigPools.
func (c *FakeMachineConfigPools) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(machineconfigpoolsResource, opts))
}

// Create takes the representation of a machineConfigPool and creates it.  Returns the server's representation of the machineConfigPool, and an error, if there is any.
func (c *FakeMachineConfigPools) Create(ctx context.Context, machineConfigPool *v1.MachineConfigPool, opts metav1.CreateOptions) (result *v1.MachineConfigPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machineconfigpoolsResource, machineConfigPool), &v1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.MachineConfigPool), err
}

// Update takes the representation of a machineConfigPool and updates it. Returns the server's representation of the machineConfigPool, and an error, if there is any.
func (c *FakeMachineConfigPools) Update(ctx context.Context, machineConfigPool *v1.MachineConfigPool, opts metav1.UpdateOptions) (result *v1.MachineConfigPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machineconfigpoolsResource, machineConfigPool), &v1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.MachineConfigPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineConfigPools) UpdateStatus(ctx context.Context, machineConfigPool *v1.MachineConfigPool, opts metav1.UpdateOptions) (*v1.MachineConfigPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(machineconfigpoolsResource, "status", machineConfigPool), &v1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.MachineConfigPool), err
}

// Delete takes name of the machineConfigPool and deletes it. Returns an error if one occurs.
func (c *FakeMachineConfigPools) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(machineconfigpoolsResource, name, opts), &v1.MachineConfigPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineConfigPools) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(machineconfigpoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.MachineConfigPoolList{})
	return err
}

// Patch applies the patch and returns the patched machineConfigPool.
func (c *FakeMachineConfigPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineConfigPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machineconfigpoolsResource, name, pt, data, subresources...), &v1.MachineConfigPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.MachineConfigPool), err
}