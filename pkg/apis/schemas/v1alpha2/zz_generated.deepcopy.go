// +build !ignore_autogenerated

// autogenerated by controller-gen object, do not modify manually

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Migration) DeepCopyInto(out *Migration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Migration.
func (in *Migration) DeepCopy() *Migration {
	if in == nil {
		return nil
	}
	out := new(Migration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Migration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationList) DeepCopyInto(out *MigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Migration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationList.
func (in *MigrationList) DeepCopy() *MigrationList {
	if in == nil {
		return nil
	}
	out := new(MigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationSpec) DeepCopyInto(out *MigrationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationSpec.
func (in *MigrationSpec) DeepCopy() *MigrationSpec {
	if in == nil {
		return nil
	}
	out := new(MigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationStatus) DeepCopyInto(out *MigrationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationStatus.
func (in *MigrationStatus) DeepCopy() *MigrationStatus {
	if in == nil {
		return nil
	}
	out := new(MigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableColumn) DeepCopyInto(out *SQLTableColumn) {
	*out = *in
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = new(SQLTableColumnConstraints)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableColumn.
func (in *SQLTableColumn) DeepCopy() *SQLTableColumn {
	if in == nil {
		return nil
	}
	out := new(SQLTableColumn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableColumnConstraints) DeepCopyInto(out *SQLTableColumnConstraints) {
	*out = *in
	if in.NotNull != nil {
		in, out := &in.NotNull, &out.NotNull
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableColumnConstraints.
func (in *SQLTableColumnConstraints) DeepCopy() *SQLTableColumnConstraints {
	if in == nil {
		return nil
	}
	out := new(SQLTableColumnConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableForeignKey) DeepCopyInto(out *SQLTableForeignKey) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.References.DeepCopyInto(&out.References)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableForeignKey.
func (in *SQLTableForeignKey) DeepCopy() *SQLTableForeignKey {
	if in == nil {
		return nil
	}
	out := new(SQLTableForeignKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableForeignKeyReferences) DeepCopyInto(out *SQLTableForeignKeyReferences) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableForeignKeyReferences.
func (in *SQLTableForeignKeyReferences) DeepCopy() *SQLTableForeignKeyReferences {
	if in == nil {
		return nil
	}
	out := new(SQLTableForeignKeyReferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableIndex) DeepCopyInto(out *SQLTableIndex) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableIndex.
func (in *SQLTableIndex) DeepCopy() *SQLTableIndex {
	if in == nil {
		return nil
	}
	out := new(SQLTableIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLTableSchema) DeepCopyInto(out *SQLTableSchema) {
	*out = *in
	if in.PrimaryKey != nil {
		in, out := &in.PrimaryKey, &out.PrimaryKey
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForeignKeys != nil {
		in, out := &in.ForeignKeys, &out.ForeignKeys
		*out = make([]*SQLTableForeignKey, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SQLTableForeignKey)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Indexes != nil {
		in, out := &in.Indexes, &out.Indexes
		*out = make([]*SQLTableIndex, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SQLTableIndex)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]*SQLTableColumn, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SQLTableColumn)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLTableSchema.
func (in *SQLTableSchema) DeepCopy() *SQLTableSchema {
	if in == nil {
		return nil
	}
	out := new(SQLTableSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableList) DeepCopyInto(out *TableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Table, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableList.
func (in *TableList) DeepCopy() *TableList {
	if in == nil {
		return nil
	}
	out := new(TableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSchema) DeepCopyInto(out *TableSchema) {
	*out = *in
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = new(SQLTableSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.Mysql != nil {
		in, out := &in.Mysql, &out.Mysql
		*out = new(SQLTableSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSchema.
func (in *TableSchema) DeepCopy() *TableSchema {
	if in == nil {
		return nil
	}
	out := new(TableSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpec) DeepCopyInto(out *TableSpec) {
	*out = *in
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(TableSchema)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpec.
func (in *TableSpec) DeepCopy() *TableSpec {
	if in == nil {
		return nil
	}
	out := new(TableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableStatus) DeepCopyInto(out *TableStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableStatus.
func (in *TableStatus) DeepCopy() *TableStatus {
	if in == nil {
		return nil
	}
	out := new(TableStatus)
	in.DeepCopyInto(out)
	return out
}
