// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securitycenter

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the subscription's Security Center Contact.
//
// > **NOTE:** Owner access permission is required.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/securitycenter"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = securitycenter.NewContact(ctx, "example", &securitycenter.ContactArgs{
// 			AlertNotifications: pulumi.Bool(true),
// 			AlertsToAdmins:     pulumi.Bool(true),
// 			Email:              pulumi.String("contact@example.com"),
// 			Phone:              pulumi.String("+1-555-555-5555"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Contact struct {
	pulumi.CustomResourceState

	// Whether to send security alerts notifications to the security contact.
	AlertNotifications pulumi.BoolOutput `pulumi:"alertNotifications"`
	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins pulumi.BoolOutput `pulumi:"alertsToAdmins"`
	// The email of the Security Center Contact.
	Email pulumi.StringOutput `pulumi:"email"`
	// The phone number of the Security Center Contact.
	Phone pulumi.StringPtrOutput `pulumi:"phone"`
}

// NewContact registers a new resource with the given unique name, arguments, and options.
func NewContact(ctx *pulumi.Context,
	name string, args *ContactArgs, opts ...pulumi.ResourceOption) (*Contact, error) {
	if args == nil || args.AlertNotifications == nil {
		return nil, errors.New("missing required argument 'AlertNotifications'")
	}
	if args == nil || args.AlertsToAdmins == nil {
		return nil, errors.New("missing required argument 'AlertsToAdmins'")
	}
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	if args == nil {
		args = &ContactArgs{}
	}
	var resource Contact
	err := ctx.RegisterResource("azure:securitycenter/contact:Contact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContact gets an existing Contact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactState, opts ...pulumi.ResourceOption) (*Contact, error) {
	var resource Contact
	err := ctx.ReadResource("azure:securitycenter/contact:Contact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Contact resources.
type contactState struct {
	// Whether to send security alerts notifications to the security contact.
	AlertNotifications *bool `pulumi:"alertNotifications"`
	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins *bool `pulumi:"alertsToAdmins"`
	// The email of the Security Center Contact.
	Email *string `pulumi:"email"`
	// The phone number of the Security Center Contact.
	Phone *string `pulumi:"phone"`
}

type ContactState struct {
	// Whether to send security alerts notifications to the security contact.
	AlertNotifications pulumi.BoolPtrInput
	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins pulumi.BoolPtrInput
	// The email of the Security Center Contact.
	Email pulumi.StringPtrInput
	// The phone number of the Security Center Contact.
	Phone pulumi.StringPtrInput
}

func (ContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactState)(nil)).Elem()
}

type contactArgs struct {
	// Whether to send security alerts notifications to the security contact.
	AlertNotifications bool `pulumi:"alertNotifications"`
	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins bool `pulumi:"alertsToAdmins"`
	// The email of the Security Center Contact.
	Email string `pulumi:"email"`
	// The phone number of the Security Center Contact.
	Phone *string `pulumi:"phone"`
}

// The set of arguments for constructing a Contact resource.
type ContactArgs struct {
	// Whether to send security alerts notifications to the security contact.
	AlertNotifications pulumi.BoolInput
	// Whether to send security alerts notifications to subscription admins.
	AlertsToAdmins pulumi.BoolInput
	// The email of the Security Center Contact.
	Email pulumi.StringInput
	// The phone number of the Security Center Contact.
	Phone pulumi.StringPtrInput
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactArgs)(nil)).Elem()
}
