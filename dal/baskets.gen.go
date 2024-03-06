// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"Food-delivery/domain"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newBasket(db *gorm.DB, opts ...gen.DOOption) basket {
	_basket := basket{}

	_basket.basketDo.UseDB(db, opts...)
	_basket.basketDo.UseModel(&domain.Basket{})

	tableName := _basket.basketDo.TableName()
	_basket.ALL = field.NewAsterisk(tableName)
	_basket.ID = field.NewUint(tableName, "id")
	_basket.CreatedAt = field.NewTime(tableName, "created_at")
	_basket.UpdatedAt = field.NewTime(tableName, "updated_at")
	_basket.DeletedAt = field.NewTime(tableName, "deleted_at")
	_basket.UserID = field.NewUint(tableName, "user_id")
	_basket.PromotionID = field.NewUint(tableName, "promotion_id")
	_basket.BasketProduct = basketHasManyBasketProduct{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("BasketProduct", "domain.BasketProduct"),
	}

	_basket.User = basketBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "domain.User"),
	}

	_basket.Promotion = basketBelongsToPromotion{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Promotion", "domain.Promotion"),
		PromotionProduct: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Promotion.PromotionProduct", "domain.PromotionProduct"),
		},
	}

	_basket.fillFieldMap()

	return _basket
}

type basket struct {
	basketDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Time
	UserID        field.Uint
	PromotionID   field.Uint
	BasketProduct basketHasManyBasketProduct

	User basketBelongsToUser

	Promotion basketBelongsToPromotion

	fieldMap map[string]field.Expr
}

func (b basket) Table(newTableName string) *basket {
	b.basketDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b basket) As(alias string) *basket {
	b.basketDo.DO = *(b.basketDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *basket) updateTableName(table string) *basket {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewUint(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewTime(table, "deleted_at")
	b.UserID = field.NewUint(table, "user_id")
	b.PromotionID = field.NewUint(table, "promotion_id")

	b.fillFieldMap()

	return b
}

func (b *basket) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *basket) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 9)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["promotion_id"] = b.PromotionID

}

func (b basket) clone(db *gorm.DB) basket {
	b.basketDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b basket) replaceDB(db *gorm.DB) basket {
	b.basketDo.ReplaceDB(db)
	return b
}

type basketHasManyBasketProduct struct {
	db *gorm.DB

	field.RelationField
}

func (a basketHasManyBasketProduct) Where(conds ...field.Expr) *basketHasManyBasketProduct {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a basketHasManyBasketProduct) WithContext(ctx context.Context) *basketHasManyBasketProduct {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a basketHasManyBasketProduct) Session(session *gorm.Session) *basketHasManyBasketProduct {
	a.db = a.db.Session(session)
	return &a
}

func (a basketHasManyBasketProduct) Model(m *domain.Basket) *basketHasManyBasketProductTx {
	return &basketHasManyBasketProductTx{a.db.Model(m).Association(a.Name())}
}

type basketHasManyBasketProductTx struct{ tx *gorm.Association }

func (a basketHasManyBasketProductTx) Find() (result []*domain.BasketProduct, err error) {
	return result, a.tx.Find(&result)
}

func (a basketHasManyBasketProductTx) Append(values ...*domain.BasketProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a basketHasManyBasketProductTx) Replace(values ...*domain.BasketProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a basketHasManyBasketProductTx) Delete(values ...*domain.BasketProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a basketHasManyBasketProductTx) Clear() error {
	return a.tx.Clear()
}

func (a basketHasManyBasketProductTx) Count() int64 {
	return a.tx.Count()
}

type basketBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a basketBelongsToUser) Where(conds ...field.Expr) *basketBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a basketBelongsToUser) WithContext(ctx context.Context) *basketBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a basketBelongsToUser) Session(session *gorm.Session) *basketBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a basketBelongsToUser) Model(m *domain.Basket) *basketBelongsToUserTx {
	return &basketBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type basketBelongsToUserTx struct{ tx *gorm.Association }

func (a basketBelongsToUserTx) Find() (result *domain.User, err error) {
	return result, a.tx.Find(&result)
}

func (a basketBelongsToUserTx) Append(values ...*domain.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a basketBelongsToUserTx) Replace(values ...*domain.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a basketBelongsToUserTx) Delete(values ...*domain.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a basketBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a basketBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type basketBelongsToPromotion struct {
	db *gorm.DB

	field.RelationField

	PromotionProduct struct {
		field.RelationField
	}
}

func (a basketBelongsToPromotion) Where(conds ...field.Expr) *basketBelongsToPromotion {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a basketBelongsToPromotion) WithContext(ctx context.Context) *basketBelongsToPromotion {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a basketBelongsToPromotion) Session(session *gorm.Session) *basketBelongsToPromotion {
	a.db = a.db.Session(session)
	return &a
}

func (a basketBelongsToPromotion) Model(m *domain.Basket) *basketBelongsToPromotionTx {
	return &basketBelongsToPromotionTx{a.db.Model(m).Association(a.Name())}
}

type basketBelongsToPromotionTx struct{ tx *gorm.Association }

func (a basketBelongsToPromotionTx) Find() (result *domain.Promotion, err error) {
	return result, a.tx.Find(&result)
}

func (a basketBelongsToPromotionTx) Append(values ...*domain.Promotion) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a basketBelongsToPromotionTx) Replace(values ...*domain.Promotion) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a basketBelongsToPromotionTx) Delete(values ...*domain.Promotion) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a basketBelongsToPromotionTx) Clear() error {
	return a.tx.Clear()
}

func (a basketBelongsToPromotionTx) Count() int64 {
	return a.tx.Count()
}

type basketDo struct{ gen.DO }

type IBasketDo interface {
	gen.SubQuery
	Debug() IBasketDo
	WithContext(ctx context.Context) IBasketDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBasketDo
	WriteDB() IBasketDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBasketDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBasketDo
	Not(conds ...gen.Condition) IBasketDo
	Or(conds ...gen.Condition) IBasketDo
	Select(conds ...field.Expr) IBasketDo
	Where(conds ...gen.Condition) IBasketDo
	Order(conds ...field.Expr) IBasketDo
	Distinct(cols ...field.Expr) IBasketDo
	Omit(cols ...field.Expr) IBasketDo
	Join(table schema.Tabler, on ...field.Expr) IBasketDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBasketDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBasketDo
	Group(cols ...field.Expr) IBasketDo
	Having(conds ...gen.Condition) IBasketDo
	Limit(limit int) IBasketDo
	Offset(offset int) IBasketDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBasketDo
	Unscoped() IBasketDo
	Create(values ...*domain.Basket) error
	CreateInBatches(values []*domain.Basket, batchSize int) error
	Save(values ...*domain.Basket) error
	First() (*domain.Basket, error)
	Take() (*domain.Basket, error)
	Last() (*domain.Basket, error)
	Find() ([]*domain.Basket, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Basket, err error)
	FindInBatches(result *[]*domain.Basket, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Basket) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBasketDo
	Assign(attrs ...field.AssignExpr) IBasketDo
	Joins(fields ...field.RelationField) IBasketDo
	Preload(fields ...field.RelationField) IBasketDo
	FirstOrInit() (*domain.Basket, error)
	FirstOrCreate() (*domain.Basket, error)
	FindByPage(offset int, limit int) (result []*domain.Basket, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBasketDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b basketDo) Debug() IBasketDo {
	return b.withDO(b.DO.Debug())
}

func (b basketDo) WithContext(ctx context.Context) IBasketDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b basketDo) ReadDB() IBasketDo {
	return b.Clauses(dbresolver.Read)
}

func (b basketDo) WriteDB() IBasketDo {
	return b.Clauses(dbresolver.Write)
}

func (b basketDo) Session(config *gorm.Session) IBasketDo {
	return b.withDO(b.DO.Session(config))
}

func (b basketDo) Clauses(conds ...clause.Expression) IBasketDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b basketDo) Returning(value interface{}, columns ...string) IBasketDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b basketDo) Not(conds ...gen.Condition) IBasketDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b basketDo) Or(conds ...gen.Condition) IBasketDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b basketDo) Select(conds ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b basketDo) Where(conds ...gen.Condition) IBasketDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b basketDo) Order(conds ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b basketDo) Distinct(cols ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b basketDo) Omit(cols ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b basketDo) Join(table schema.Tabler, on ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b basketDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBasketDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b basketDo) RightJoin(table schema.Tabler, on ...field.Expr) IBasketDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b basketDo) Group(cols ...field.Expr) IBasketDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b basketDo) Having(conds ...gen.Condition) IBasketDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b basketDo) Limit(limit int) IBasketDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b basketDo) Offset(offset int) IBasketDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b basketDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBasketDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b basketDo) Unscoped() IBasketDo {
	return b.withDO(b.DO.Unscoped())
}

func (b basketDo) Create(values ...*domain.Basket) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b basketDo) CreateInBatches(values []*domain.Basket, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b basketDo) Save(values ...*domain.Basket) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b basketDo) First() (*domain.Basket, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Basket), nil
	}
}

func (b basketDo) Take() (*domain.Basket, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Basket), nil
	}
}

func (b basketDo) Last() (*domain.Basket, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Basket), nil
	}
}

func (b basketDo) Find() ([]*domain.Basket, error) {
	result, err := b.DO.Find()
	return result.([]*domain.Basket), err
}

func (b basketDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Basket, err error) {
	buf := make([]*domain.Basket, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b basketDo) FindInBatches(result *[]*domain.Basket, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b basketDo) Attrs(attrs ...field.AssignExpr) IBasketDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b basketDo) Assign(attrs ...field.AssignExpr) IBasketDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b basketDo) Joins(fields ...field.RelationField) IBasketDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b basketDo) Preload(fields ...field.RelationField) IBasketDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b basketDo) FirstOrInit() (*domain.Basket, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Basket), nil
	}
}

func (b basketDo) FirstOrCreate() (*domain.Basket, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Basket), nil
	}
}

func (b basketDo) FindByPage(offset int, limit int) (result []*domain.Basket, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b basketDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b basketDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b basketDo) Delete(models ...*domain.Basket) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *basketDo) withDO(do gen.Dao) *basketDo {
	b.DO = *do.(*gen.DO)
	return b
}