// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bitmagnet-io/bitmagnet/internal/model"
)

func newTorrentHint(db *gorm.DB, opts ...gen.DOOption) torrentHint {
	_torrentHint := torrentHint{}

	_torrentHint.torrentHintDo.UseDB(db, opts...)
	_torrentHint.torrentHintDo.UseModel(&model.TorrentHint{})

	tableName := _torrentHint.torrentHintDo.TableName()
	_torrentHint.ALL = field.NewAsterisk(tableName)
	_torrentHint.InfoHash = field.NewField(tableName, "info_hash")
	_torrentHint.ContentType = field.NewString(tableName, "content_type")
	_torrentHint.ContentSource = field.NewString(tableName, "content_source")
	_torrentHint.ContentID = field.NewString(tableName, "content_id")
	_torrentHint.Title = field.NewField(tableName, "title")
	_torrentHint.ReleaseYear = field.NewField(tableName, "release_year")
	_torrentHint.Languages = field.NewField(tableName, "languages")
	_torrentHint.Episodes = field.NewField(tableName, "episodes")
	_torrentHint.VideoResolution = field.NewField(tableName, "video_resolution")
	_torrentHint.VideoSource = field.NewField(tableName, "video_source")
	_torrentHint.VideoCodec = field.NewField(tableName, "video_codec")
	_torrentHint.Video3d = field.NewField(tableName, "video_3d")
	_torrentHint.VideoModifier = field.NewField(tableName, "video_modifier")
	_torrentHint.ReleaseGroup = field.NewField(tableName, "release_group")
	_torrentHint.CreatedAt = field.NewTime(tableName, "created_at")
	_torrentHint.UpdatedAt = field.NewTime(tableName, "updated_at")

	_torrentHint.fillFieldMap()

	return _torrentHint
}

type torrentHint struct {
	torrentHintDo

	ALL             field.Asterisk
	InfoHash        field.Field
	ContentType     field.String
	ContentSource   field.String
	ContentID       field.String
	Title           field.Field
	ReleaseYear     field.Field
	Languages       field.Field
	Episodes        field.Field
	VideoResolution field.Field
	VideoSource     field.Field
	VideoCodec      field.Field
	Video3d         field.Field
	VideoModifier   field.Field
	ReleaseGroup    field.Field
	CreatedAt       field.Time
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (t torrentHint) Table(newTableName string) *torrentHint {
	t.torrentHintDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t torrentHint) As(alias string) *torrentHint {
	t.torrentHintDo.DO = *(t.torrentHintDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *torrentHint) updateTableName(table string) *torrentHint {
	t.ALL = field.NewAsterisk(table)
	t.InfoHash = field.NewField(table, "info_hash")
	t.ContentType = field.NewString(table, "content_type")
	t.ContentSource = field.NewString(table, "content_source")
	t.ContentID = field.NewString(table, "content_id")
	t.Title = field.NewField(table, "title")
	t.ReleaseYear = field.NewField(table, "release_year")
	t.Languages = field.NewField(table, "languages")
	t.Episodes = field.NewField(table, "episodes")
	t.VideoResolution = field.NewField(table, "video_resolution")
	t.VideoSource = field.NewField(table, "video_source")
	t.VideoCodec = field.NewField(table, "video_codec")
	t.Video3d = field.NewField(table, "video_3d")
	t.VideoModifier = field.NewField(table, "video_modifier")
	t.ReleaseGroup = field.NewField(table, "release_group")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *torrentHint) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *torrentHint) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 16)
	t.fieldMap["info_hash"] = t.InfoHash
	t.fieldMap["content_type"] = t.ContentType
	t.fieldMap["content_source"] = t.ContentSource
	t.fieldMap["content_id"] = t.ContentID
	t.fieldMap["title"] = t.Title
	t.fieldMap["release_year"] = t.ReleaseYear
	t.fieldMap["languages"] = t.Languages
	t.fieldMap["episodes"] = t.Episodes
	t.fieldMap["video_resolution"] = t.VideoResolution
	t.fieldMap["video_source"] = t.VideoSource
	t.fieldMap["video_codec"] = t.VideoCodec
	t.fieldMap["video_3d"] = t.Video3d
	t.fieldMap["video_modifier"] = t.VideoModifier
	t.fieldMap["release_group"] = t.ReleaseGroup
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t torrentHint) clone(db *gorm.DB) torrentHint {
	t.torrentHintDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t torrentHint) replaceDB(db *gorm.DB) torrentHint {
	t.torrentHintDo.ReplaceDB(db)
	return t
}

type torrentHintDo struct{ gen.DO }

type ITorrentHintDo interface {
	gen.SubQuery
	Debug() ITorrentHintDo
	WithContext(ctx context.Context) ITorrentHintDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITorrentHintDo
	WriteDB() ITorrentHintDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITorrentHintDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITorrentHintDo
	Not(conds ...gen.Condition) ITorrentHintDo
	Or(conds ...gen.Condition) ITorrentHintDo
	Select(conds ...field.Expr) ITorrentHintDo
	Where(conds ...gen.Condition) ITorrentHintDo
	Order(conds ...field.Expr) ITorrentHintDo
	Distinct(cols ...field.Expr) ITorrentHintDo
	Omit(cols ...field.Expr) ITorrentHintDo
	Join(table schema.Tabler, on ...field.Expr) ITorrentHintDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITorrentHintDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITorrentHintDo
	Group(cols ...field.Expr) ITorrentHintDo
	Having(conds ...gen.Condition) ITorrentHintDo
	Limit(limit int) ITorrentHintDo
	Offset(offset int) ITorrentHintDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITorrentHintDo
	Unscoped() ITorrentHintDo
	Create(values ...*model.TorrentHint) error
	CreateInBatches(values []*model.TorrentHint, batchSize int) error
	Save(values ...*model.TorrentHint) error
	First() (*model.TorrentHint, error)
	Take() (*model.TorrentHint, error)
	Last() (*model.TorrentHint, error)
	Find() ([]*model.TorrentHint, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TorrentHint, err error)
	FindInBatches(result *[]*model.TorrentHint, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TorrentHint) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITorrentHintDo
	Assign(attrs ...field.AssignExpr) ITorrentHintDo
	Joins(fields ...field.RelationField) ITorrentHintDo
	Preload(fields ...field.RelationField) ITorrentHintDo
	FirstOrInit() (*model.TorrentHint, error)
	FirstOrCreate() (*model.TorrentHint, error)
	FindByPage(offset int, limit int) (result []*model.TorrentHint, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITorrentHintDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t torrentHintDo) Debug() ITorrentHintDo {
	return t.withDO(t.DO.Debug())
}

func (t torrentHintDo) WithContext(ctx context.Context) ITorrentHintDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t torrentHintDo) ReadDB() ITorrentHintDo {
	return t.Clauses(dbresolver.Read)
}

func (t torrentHintDo) WriteDB() ITorrentHintDo {
	return t.Clauses(dbresolver.Write)
}

func (t torrentHintDo) Session(config *gorm.Session) ITorrentHintDo {
	return t.withDO(t.DO.Session(config))
}

func (t torrentHintDo) Clauses(conds ...clause.Expression) ITorrentHintDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t torrentHintDo) Returning(value interface{}, columns ...string) ITorrentHintDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t torrentHintDo) Not(conds ...gen.Condition) ITorrentHintDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t torrentHintDo) Or(conds ...gen.Condition) ITorrentHintDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t torrentHintDo) Select(conds ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t torrentHintDo) Where(conds ...gen.Condition) ITorrentHintDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t torrentHintDo) Order(conds ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t torrentHintDo) Distinct(cols ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t torrentHintDo) Omit(cols ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t torrentHintDo) Join(table schema.Tabler, on ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t torrentHintDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t torrentHintDo) RightJoin(table schema.Tabler, on ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t torrentHintDo) Group(cols ...field.Expr) ITorrentHintDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t torrentHintDo) Having(conds ...gen.Condition) ITorrentHintDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t torrentHintDo) Limit(limit int) ITorrentHintDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t torrentHintDo) Offset(offset int) ITorrentHintDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t torrentHintDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITorrentHintDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t torrentHintDo) Unscoped() ITorrentHintDo {
	return t.withDO(t.DO.Unscoped())
}

func (t torrentHintDo) Create(values ...*model.TorrentHint) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t torrentHintDo) CreateInBatches(values []*model.TorrentHint, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t torrentHintDo) Save(values ...*model.TorrentHint) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t torrentHintDo) First() (*model.TorrentHint, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TorrentHint), nil
	}
}

func (t torrentHintDo) Take() (*model.TorrentHint, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TorrentHint), nil
	}
}

func (t torrentHintDo) Last() (*model.TorrentHint, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TorrentHint), nil
	}
}

func (t torrentHintDo) Find() ([]*model.TorrentHint, error) {
	result, err := t.DO.Find()
	return result.([]*model.TorrentHint), err
}

func (t torrentHintDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TorrentHint, err error) {
	buf := make([]*model.TorrentHint, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t torrentHintDo) FindInBatches(result *[]*model.TorrentHint, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t torrentHintDo) Attrs(attrs ...field.AssignExpr) ITorrentHintDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t torrentHintDo) Assign(attrs ...field.AssignExpr) ITorrentHintDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t torrentHintDo) Joins(fields ...field.RelationField) ITorrentHintDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t torrentHintDo) Preload(fields ...field.RelationField) ITorrentHintDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t torrentHintDo) FirstOrInit() (*model.TorrentHint, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TorrentHint), nil
	}
}

func (t torrentHintDo) FirstOrCreate() (*model.TorrentHint, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TorrentHint), nil
	}
}

func (t torrentHintDo) FindByPage(offset int, limit int) (result []*model.TorrentHint, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t torrentHintDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t torrentHintDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t torrentHintDo) Delete(models ...*model.TorrentHint) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *torrentHintDo) withDO(do gen.Dao) *torrentHintDo {
	t.DO = *do.(*gen.DO)
	return t
}
