// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"fmt"
)

// WhereBuilder holds multiple where conditions in a group.
// 在一个builder组保存一个where条件
type WhereBuilder struct {
	model       *Model        // A WhereBuilder should be bound to certain Model.
	whereHolder []WhereHolder // Condition strings for where operation.
}

// WhereHolder is the holder for where condition preparing.
type WhereHolder struct {
	Type     string        // Type of this holder. 持有类型
	Operator int           // Operator for this holder. 操作
	Where    interface{}   // Where parameter, which can commonly be type of string/map/struct.
	Args     []interface{} // Arguments for where parameter.
	Prefix   string        // Field prefix, eg: "user.", "order.".
}

// Builder creates and returns a WhereBuilder.
// 获取一个WhereBuilder
func (m *Model) Builder() *WhereBuilder {
	b := &WhereBuilder{
		model:       m,
		whereHolder: make([]WhereHolder, 0),
	}
	return b
}

// getBuilder creates and returns a cloned WhereBuilder of current WhereBuilder if `safe` is true,
// or else it returns the current WhereBuilder.
// 返回当前的wherebuilder克隆对象，如果是链式安全，非链式安全就返回当前wherebuild
func (b *WhereBuilder) getBuilder() *WhereBuilder {
	return b.Clone()
}

// Clone clones and returns a WhereBuilder that is a copy of current one.
// 克隆一个wherebuilder
func (b *WhereBuilder) Clone() *WhereBuilder {
	newBuilder := b.model.Builder()
	// 创建一个和原来whereholder一样长度的切片，并且拷贝一个值过去
	newBuilder.whereHolder = make([]WhereHolder, len(b.whereHolder))
	copy(newBuilder.whereHolder, b.whereHolder)
	return newBuilder
}

// Build builds current WhereBuilder and returns the condition string and parameters.
func (b *WhereBuilder) Build() (conditionWhere string, conditionArgs []interface{}) {
	var (
		ctx                         = b.model.GetCtx()
		autoPrefix                  = b.model.getAutoPrefix()
		tableForMappingAndFiltering = b.model.tables
	)
	if len(b.whereHolder) > 0 {
		for _, holder := range b.whereHolder {
			if holder.Prefix == "" {
				holder.Prefix = autoPrefix
			}
			switch holder.Operator {
			case whereHolderOperatorWhere, whereHolderOperatorAnd:
				newWhere, newArgs := formatWhereHolder(ctx, b.model.db, formatWhereHolderInput{
					WhereHolder: holder,
					OmitNil:     b.model.option&optionOmitNilWhere > 0,
					OmitEmpty:   b.model.option&optionOmitEmptyWhere > 0,
					Schema:      b.model.schema,
					Table:       tableForMappingAndFiltering,
				})
				if len(newWhere) > 0 {
					if len(conditionWhere) == 0 {
						conditionWhere = newWhere
					} else if conditionWhere[0] == '(' {
						conditionWhere = fmt.Sprintf(`%s AND (%s)`, conditionWhere, newWhere)
					} else {
						conditionWhere = fmt.Sprintf(`(%s) AND (%s)`, conditionWhere, newWhere)
					}
					conditionArgs = append(conditionArgs, newArgs...)
				}

			case whereHolderOperatorOr:
				newWhere, newArgs := formatWhereHolder(ctx, b.model.db, formatWhereHolderInput{
					WhereHolder: holder,
					OmitNil:     b.model.option&optionOmitNilWhere > 0,
					OmitEmpty:   b.model.option&optionOmitEmptyWhere > 0,
					Schema:      b.model.schema,
					Table:       tableForMappingAndFiltering,
				})
				if len(newWhere) > 0 {
					if len(conditionWhere) == 0 {
						conditionWhere = newWhere
					} else if conditionWhere[0] == '(' {
						conditionWhere = fmt.Sprintf(`%s OR (%s)`, conditionWhere, newWhere)
					} else {
						conditionWhere = fmt.Sprintf(`(%s) OR (%s)`, conditionWhere, newWhere)
					}
					conditionArgs = append(conditionArgs, newArgs...)
				}
			}
		}
	}
	return
}

// convertWhereBuilder converts parameter `where` to condition string and parameters if `where` is also a WhereBuilder.
func (b *WhereBuilder) convertWhereBuilder(where interface{}, args []interface{}) (newWhere interface{}, newArgs []interface{}) {
	var builder *WhereBuilder
	switch v := where.(type) {
	case WhereBuilder:
		builder = &v

	case *WhereBuilder:
		builder = v
	}
	// TODO 什么样的情况下builder != nil ?
	if builder != nil {
		conditionWhere, conditionArgs := builder.Build()
		if conditionWhere != "" && len(b.whereHolder) == 0 {
			conditionWhere = "(" + conditionWhere + ")"
		}
		return conditionWhere, conditionArgs
	}
	return where, args
}
