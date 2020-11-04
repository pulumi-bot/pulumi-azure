// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: ZipBlob resource is deprecated in the 2.0 version of the provider. Use Blob resource instead.
type ZipBlob struct {
	pulumi.CustomResourceState

	AccessTier           pulumi.StringOutput    `pulumi:"accessTier"`
	Content              pulumi.ArchiveOutput   `pulumi:"content"`
	ContentType          pulumi.StringPtrOutput `pulumi:"contentType"`
	Metadata             pulumi.StringMapOutput `pulumi:"metadata"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Parallelism          pulumi.IntPtrOutput    `pulumi:"parallelism"`
	Size                 pulumi.IntPtrOutput    `pulumi:"size"`
	SourceContent        pulumi.StringPtrOutput `pulumi:"sourceContent"`
	SourceUri            pulumi.StringPtrOutput `pulumi:"sourceUri"`
	StorageAccountName   pulumi.StringOutput    `pulumi:"storageAccountName"`
	StorageContainerName pulumi.StringOutput    `pulumi:"storageContainerName"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	Url                  pulumi.StringOutput    `pulumi:"url"`
}

// NewZipBlob registers a new resource with the given unique name, arguments, and options.
func NewZipBlob(ctx *pulumi.Context,
	name string, args *ZipBlobArgs, opts ...pulumi.ResourceOption) (*ZipBlob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ZipBlob
	err := ctx.RegisterResource("azure:storage/zipBlob:ZipBlob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZipBlob gets an existing ZipBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZipBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZipBlobState, opts ...pulumi.ResourceOption) (*ZipBlob, error) {
	var resource ZipBlob
	err := ctx.ReadResource("azure:storage/zipBlob:ZipBlob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZipBlob resources.
type zipBlobState struct {
	AccessTier           *string           `pulumi:"accessTier"`
	Content              pulumi.Archive    `pulumi:"content"`
	ContentType          *string           `pulumi:"contentType"`
	Metadata             map[string]string `pulumi:"metadata"`
	Name                 *string           `pulumi:"name"`
	Parallelism          *int              `pulumi:"parallelism"`
	Size                 *int              `pulumi:"size"`
	SourceContent        *string           `pulumi:"sourceContent"`
	SourceUri            *string           `pulumi:"sourceUri"`
	StorageAccountName   *string           `pulumi:"storageAccountName"`
	StorageContainerName *string           `pulumi:"storageContainerName"`
	Type                 *string           `pulumi:"type"`
	Url                  *string           `pulumi:"url"`
}

type ZipBlobState struct {
	AccessTier           pulumi.StringPtrInput
	Content              pulumi.ArchiveInput
	ContentType          pulumi.StringPtrInput
	Metadata             pulumi.StringMapInput
	Name                 pulumi.StringPtrInput
	Parallelism          pulumi.IntPtrInput
	Size                 pulumi.IntPtrInput
	SourceContent        pulumi.StringPtrInput
	SourceUri            pulumi.StringPtrInput
	StorageAccountName   pulumi.StringPtrInput
	StorageContainerName pulumi.StringPtrInput
	Type                 pulumi.StringPtrInput
	Url                  pulumi.StringPtrInput
}

func (ZipBlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*zipBlobState)(nil)).Elem()
}

type zipBlobArgs struct {
	AccessTier           *string           `pulumi:"accessTier"`
	Content              pulumi.Archive    `pulumi:"content"`
	ContentType          *string           `pulumi:"contentType"`
	Metadata             map[string]string `pulumi:"metadata"`
	Name                 *string           `pulumi:"name"`
	Parallelism          *int              `pulumi:"parallelism"`
	Size                 *int              `pulumi:"size"`
	SourceContent        *string           `pulumi:"sourceContent"`
	SourceUri            *string           `pulumi:"sourceUri"`
	StorageAccountName   string            `pulumi:"storageAccountName"`
	StorageContainerName string            `pulumi:"storageContainerName"`
	Type                 string            `pulumi:"type"`
}

// The set of arguments for constructing a ZipBlob resource.
type ZipBlobArgs struct {
	AccessTier           pulumi.StringPtrInput
	Content              pulumi.ArchiveInput
	ContentType          pulumi.StringPtrInput
	Metadata             pulumi.StringMapInput
	Name                 pulumi.StringPtrInput
	Parallelism          pulumi.IntPtrInput
	Size                 pulumi.IntPtrInput
	SourceContent        pulumi.StringPtrInput
	SourceUri            pulumi.StringPtrInput
	StorageAccountName   pulumi.StringInput
	StorageContainerName pulumi.StringInput
	Type                 pulumi.StringInput
}

func (ZipBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zipBlobArgs)(nil)).Elem()
}
