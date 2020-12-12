// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Data Lake Store File.
//
// > **Note:** If you want to change the data in the remote file without changing the `localFilePath`, then
// taint the resource so the `datalake.StoreFile` gets recreated with the new data.
//
// ## Import
//
// Data Lake Store File's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datalake/storeFile:StoreFile txt
// ```
type StoreFile struct {
	pulumi.CustomResourceState

	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringOutput `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringOutput `pulumi:"remoteFilePath"`
}

// NewStoreFile registers a new resource with the given unique name, arguments, and options.
func NewStoreFile(ctx *pulumi.Context,
	name string, args *StoreFileArgs, opts ...pulumi.ResourceOption) (*StoreFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.LocalFilePath == nil {
		return nil, errors.New("invalid value for required argument 'LocalFilePath'")
	}
	if args.RemoteFilePath == nil {
		return nil, errors.New("invalid value for required argument 'RemoteFilePath'")
	}
	var resource StoreFile
	err := ctx.RegisterResource("azure:datalake/storeFile:StoreFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoreFile gets an existing StoreFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoreFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoreFileState, opts ...pulumi.ResourceOption) (*StoreFile, error) {
	var resource StoreFile
	err := ctx.ReadResource("azure:datalake/storeFile:StoreFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoreFile resources.
type storeFileState struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName *string `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath *string `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath *string `pulumi:"remoteFilePath"`
}

type StoreFileState struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringPtrInput
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringPtrInput
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringPtrInput
}

func (StoreFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFileState)(nil)).Elem()
}

type storeFileArgs struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName string `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath string `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath string `pulumi:"remoteFilePath"`
}

// The set of arguments for constructing a StoreFile resource.
type StoreFileArgs struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringInput
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringInput
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringInput
}

func (StoreFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFileArgs)(nil)).Elem()
}

type StoreFileInput interface {
	pulumi.Input

	ToStoreFileOutput() StoreFileOutput
	ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput
}

func (*StoreFile) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFile)(nil))
}

func (i *StoreFile) ToStoreFileOutput() StoreFileOutput {
	return i.ToStoreFileOutputWithContext(context.Background())
}

func (i *StoreFile) ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFileOutput)
}

func (i *StoreFile) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return i.ToStoreFilePtrOutputWithContext(context.Background())
}

func (i *StoreFile) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFilePtrOutput)
}

type StoreFilePtrInput interface {
	pulumi.Input

	ToStoreFilePtrOutput() StoreFilePtrOutput
	ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput
}

type StoreFileOutput struct {
	*pulumi.OutputState
}

func (StoreFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFile)(nil))
}

func (o StoreFileOutput) ToStoreFileOutput() StoreFileOutput {
	return o
}

func (o StoreFileOutput) ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput {
	return o
}

type StoreFilePtrOutput struct {
	*pulumi.OutputState
}

func (StoreFilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreFile)(nil))
}

func (o StoreFilePtrOutput) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return o
}

func (o StoreFilePtrOutput) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StoreFileOutput{})
	pulumi.RegisterOutputType(StoreFilePtrOutput{})
}
