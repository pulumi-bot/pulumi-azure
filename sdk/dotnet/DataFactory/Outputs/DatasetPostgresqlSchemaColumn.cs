// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class DatasetPostgresqlSchemaColumn
    {
        /// <summary>
        /// The description of the column.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the column.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type of the column. Valid values are `Byte`, `Byte[]`, `Boolean`, `Date`, `DateTime`,`DateTimeOffset`, `Decimal`, `Double`, `Guid`, `Int16`, `Int32`, `Int64`, `Single`, `String`, `TimeSpan`. Please note these values are case sensitive.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DatasetPostgresqlSchemaColumn(
            string? description,

            string name,

            string? type)
        {
            Description = description;
            Name = name;
            Type = type;
        }
    }
}