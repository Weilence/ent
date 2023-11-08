// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/pc"
)

// PC is the model entity for the PC schema.
type PC struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PC) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pc.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PC fields.
func (_pc *PC) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pc.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_pc.ID = int(value.Int64)
		default:
			_pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PC.
// This includes values selected through modifiers, order, etc.
func (_pc *PC) Value(name string) (ent.Value, error) {
	return _pc.selectValues.Get(name)
}

// Update returns a builder for updating this PC.
// Note that you need to call PC.Unwrap() before calling this method if this PC
// was returned from a transaction, and the transaction was committed or rolled back.
func (_pc *PC) Update() *PCUpdateOne {
	return NewPCClient(_pc.config).UpdateOne(_pc)
}

// Unwrap unwraps the PC entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_pc *PC) Unwrap() *PC {
	_tx, ok := _pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PC is not a transactional entity")
	}
	_pc.config.driver = _tx.drv
	return _pc
}

// String implements the fmt.Stringer.
func (_pc *PC) String() string {
	var builder strings.Builder
	builder.WriteString("PC(")
	builder.WriteString(fmt.Sprintf("id=%v", _pc.ID))
	builder.WriteByte(')')
	return builder.String()
}

// PCs is a parsable slice of PC.
type PCs []*PC
