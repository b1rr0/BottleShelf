// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MembershipsColumns holds the columns for the "memberships" table.
	MembershipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "role", Type: field.TypeString},
		{Name: "membership_member", Type: field.TypeUUID, Nullable: true},
		{Name: "membership_org", Type: field.TypeUUID, Nullable: true},
	}
	// MembershipsTable holds the schema information for the "memberships" table.
	MembershipsTable = &schema.Table{
		Name:       "memberships",
		Columns:    MembershipsColumns,
		PrimaryKey: []*schema.Column{MembershipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "memberships_users_member",
				Columns:    []*schema.Column{MembershipsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "memberships_orgs_org",
				Columns:    []*schema.Column{MembershipsColumns[3]},
				RefColumns: []*schema.Column{OrgsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrgsColumns holds the columns for the "orgs" table.
	OrgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// OrgsTable holds the schema information for the "orgs" table.
	OrgsTable = &schema.Table{
		Name:       "orgs",
		Columns:    OrgsColumns,
		PrimaryKey: []*schema.Column{OrgsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MembershipsTable,
		OrgsTable,
		UsersTable,
	}
)

func init() {
	MembershipsTable.ForeignKeys[0].RefTable = UsersTable
	MembershipsTable.ForeignKeys[1].RefTable = OrgsTable
}
