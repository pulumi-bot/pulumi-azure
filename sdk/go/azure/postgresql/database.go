// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a PostgreSQL Database within a PostgreSQL Server
type Database struct {
	pulumi.CustomResourceState

	// Specifies the Charset for the PostgreSQL Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Changing this forces a new resource to be created.
	Charset pulumi.StringOutput `pulumi:"charset"`
	// Specifies the Collation for the PostgreSQL Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Note that Microsoft uses different [notation](https://msdn.microsoft.com/library/windows/desktop/dd373814.aspx) - en-US instead of en_US. Changing this forces a new resource to be created.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// Specifies the name of the PostgreSQL Database, which needs [to be a valid PostgreSQL identifier](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS). Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.Charset == nil {
		return nil, errors.New("missing required argument 'Charset'")
	}
	if args == nil || args.Collation == nil {
		return nil, errors.New("missing required argument 'Collation'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("azure:postgresql/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure:postgresql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Specifies the Charset for the PostgreSQL Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Changing this forces a new resource to be created.
	Charset *string `pulumi:"charset"`
	// Specifies the Collation for the PostgreSQL Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Note that Microsoft uses different [notation](https://msdn.microsoft.com/library/windows/desktop/dd373814.aspx) - en-US instead of en_US. Changing this forces a new resource to be created.
	Collation *string `pulumi:"collation"`
	// Specifies the name of the PostgreSQL Database, which needs [to be a valid PostgreSQL identifier](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS). Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
}

type DatabaseState struct {
	// Specifies the Charset for the PostgreSQL Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Changing this forces a new resource to be created.
	Charset pulumi.StringPtrInput
	// Specifies the Collation for the PostgreSQL Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Note that Microsoft uses different [notation](https://msdn.microsoft.com/library/windows/desktop/dd373814.aspx) - en-US instead of en_US. Changing this forces a new resource to be created.
	Collation pulumi.StringPtrInput
	// Specifies the name of the PostgreSQL Database, which needs [to be a valid PostgreSQL identifier](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS). Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Specifies the Charset for the PostgreSQL Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Changing this forces a new resource to be created.
	Charset string `pulumi:"charset"`
	// Specifies the Collation for the PostgreSQL Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Note that Microsoft uses different [notation](https://msdn.microsoft.com/library/windows/desktop/dd373814.aspx) - en-US instead of en_US. Changing this forces a new resource to be created.
	Collation string `pulumi:"collation"`
	// Specifies the name of the PostgreSQL Database, which needs [to be a valid PostgreSQL identifier](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS). Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Specifies the Charset for the PostgreSQL Database, which needs [to be a valid PostgreSQL Charset](https://www.postgresql.org/docs/current/static/multibyte.html). Changing this forces a new resource to be created.
	Charset pulumi.StringInput
	// Specifies the Collation for the PostgreSQL Database, which needs [to be a valid PostgreSQL Collation](https://www.postgresql.org/docs/current/static/collation.html). Note that Microsoft uses different [notation](https://msdn.microsoft.com/library/windows/desktop/dd373814.aspx) - en-US instead of en_US. Changing this forces a new resource to be created.
	Collation pulumi.StringInput
	// Specifies the name of the PostgreSQL Database, which needs [to be a valid PostgreSQL identifier](https://www.postgresql.org/docs/current/static/sql-syntax-lexical.html#SQL-SYNTAX-IDENTIFIERS). Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the PostgreSQL Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}
