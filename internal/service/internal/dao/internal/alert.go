// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-17 21:12:32
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlertDao is the data access object for table a_alert.
type AlertDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns AlertColumns // columns contains all the column names of Table for convenient usage.
}

// AlertColumns defines and stores column names for table a_alert.
type AlertColumns struct {
	Alert        string // 报警名称，主键
	Desc         string // 描述
	Threshold    string // 阈值
	Default      string // 
	State        string //
	Info         string //
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

//  alertColumns holds the columns for table a_alert.
var alertColumns = AlertColumns{
	Alert:        "alert",
	Desc:         "desc",
	Threshold:    "threshold",
	Default:      "default_",
	State:        "state",
	Info:         "info",
	CreatedAt:    "createdAt",
	UpdatedAt:    "updatedAt",
}

// NewAlertDao creates and returns a new DAO object for table data access.
func NewAlertDao() *AlertDao {
	return &AlertDao{
		group:   "sqlite",
		table:   "a_alert",
		columns: alertColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AlertDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AlertDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AlertDao) Columns() AlertColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AlertDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AlertDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AlertDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
