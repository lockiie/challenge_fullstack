package utilities

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FilterParams string

const (
	tagDB     = "db"
	tagFilter = "filter"
	tagJson   = "json"
	//equal
	tagEq FilterParams = "eq"
	//between
	tagBw FilterParams = "bw"
	//like
	tagLk FilterParams = "lk"
	//different
	tagDf FilterParams = "df"
	//initial for json date
	dataIni = "_initial"
	//final for json date
	dataFin        = "_final"
	valueEq        = "="
	valueBw        = " BETWEEN "
	valueLk        = " LIKE"
	valueDf        = "<>"
	valueAnd       = " AND "
	valueParamBind = "?"
	dataFormat     = "2006-01-02T15:04:05.000Z"
	LIMIT          = "limit"
	OFFSET         = "offset"
)

type BindParam struct {
	value string
	key   string
}

type ResponseQuery struct {
	TotalRows int           `json:"returnedTotal"`
	Total     int           `json:"total"`
	ResultSet []interface{} `json:"resultSet"`
}

type Query struct {
	SqlQuery       string
	Model          interface{}
	C              *fiber.Ctx
	IsPagination   bool
	NickName       string
	Where          string
	Args           []interface{}
	SqlResult      string
	GroupBy        string
	OrderByDefault string
	Response       ResponseQuery
	Conn           interface{}
	Context        *context.Context
}

func (q *Query) Mount() {
	if q.NickName != "" {
		q.NickName += "."
	}
	q.mount()
}

func (q *Query) Count() {

	conn, ok := q.Conn.(*sql.Conn)
	if !ok {
		return
	}
	q.Response.TotalRows = len(q.Response.ResultSet)

	var sqlCount string
	sqlCount = q.SqlQuery[strings.LastIndex(q.SqlQuery, "FROM"):len(q.SqlQuery)]
	sqlCount += " " + q.Where
	sqlCount += q.GroupBy
	var args []interface{}
	if q.IsPagination {
		args = q.Args[0 : len(q.Args)-2]
	} else {
		args = q.Args
	}

	sqlCount = "SELECT COUNT(*) " + sqlCount
	row := conn.QueryRowContext(*q.Context, sqlCount, args...)
	row.Scan(&q.Response.Total)

}

//func (q *Query) addArgs(value string, typeValue reflect.Type) error {
func (q *Query) AddWhere(whereAdd string, value string, typeValue reflect.Type) {
	if typeValue != nil {
		err := q.addArgs(value, typeValue)
		if err != nil {
			return
		}
	}
	if q.Where == "" {
		q.Where = "WHERE "
	} else {
		q.Where += " AND "
	}
	q.Where += whereAdd
}

func (q *Query) mount() {
	val := reflect.ValueOf(q.Model).Elem()
	queryParams := q.C.Context().QueryArgs()

	var orderBy string
	const orderByName = "sort"

	queryParams.VisitAll(func(key, value []byte) {
		if string(value) == "" {
			return
		}
		if string(key) == LIMIT || string(key) == OFFSET {
			return
		}

		for i := 0; i < val.NumField(); i++ {
			tagField := val.Type().Field(i).Tag
			valueTagjson := tagField.Get(tagJson)

			var valueTagIni = valueTagjson + dataIni

			valueTagFilter := tagField.Get(tagFilter)

			if valueTagjson != string(key) && valueTagIni != string(key) && string(key) != orderByName {
				continue
			}
			valueTagDB := tagField.Get(tagDB)
			if string(key) == orderByName {
				if string(value) == valueTagjson {
					orderBy = "ORDER BY " + q.NickName + valueTagDB
				}
				continue
			}

			typeField := val.Type().Field(i).Type
			var err error
			switch valueTagFilter {
			case string(tagEq): //////////////
				q.AddWhere(q.NickName+valueTagDB+valueEq+valueParamBind, string(value), typeField)
				//q.where += valueAnd + q.NickName + valueTagDB + valueEq + valueParamBind + strconv.Itoa(len(q.Args)-1)
			case string(tagBw): ///////////////
				if valueTagIni != string(key) || q.C.Query(valueTagjson+dataFin) == "" {
					continue
				}
				err = q.addArgs(string(value), typeField)
				if err != nil {
					return
				}
				err = q.addArgs(q.C.Query(valueTagjson+dataFin), typeField)
				if err != nil {
					q.Args = q.Args[:len(q.Args)-1]
					return
				}

				q.AddWhere(q.NickName+valueTagDB+valueBw+valueParamBind+
					valueAnd+valueParamBind, "", nil)
			case string(tagLk):

				q.AddWhere(q.NickName+valueTagDB+` LIKE concat('%',?,'%')`, q.C.Query(valueTagjson), typeField)
			case string(tagDf):
				q.AddWhere(valueTagDB+valueDf+valueParamBind, string(value), typeField)
			default:
				break
			}
		}
	})
	q.SqlResult += q.SqlQuery + " " + q.Where
	if q.GroupBy != "" {
		q.SqlResult += " " + q.GroupBy
	}
	if orderBy == "" {
		if q.OrderByDefault != "" {
			orderBy = "ORDER BY " + q.NickName + q.OrderByDefault
		}
	}
	q.SqlResult += " " + orderBy
	if q.IsPagination {
		q.SqlResult += " " + q.pagination()
	}
}

func (q *Query) pagination() string {
	var sLimit, sOffSet string
	var iLimit, iOffSet int
	var err error
	sLimit = q.C.Query(LIMIT)
	sOffSet = q.C.Query(OFFSET)
	iOffSet, err = strconv.Atoi(sOffSet)
	if err != nil {
		iOffSet = 0
	}
	iLimit, err = strconv.Atoi(sLimit)
	if err != nil {
		iLimit = iOffSet + 20
	}
	q.Args = append(q.Args, iOffSet, iLimit)
	return " LIMIT " + valueParamBind + ", " + valueParamBind
}

func filters(tagFilter string) []string {
	return strings.Split(tagFilter, ",")
}

func isValueFilter(filters *[]string) {

}

func mountOperationEqual() {

}

func (q *Query) addArgs(value string, typeValue reflect.Type) error {
	switch typeValue.Kind() {
	case reflect.String:
		q.Args = append(q.Args, value)
	case reflect.Bool:
		valueBool, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, valueBool)
	case reflect.Int:
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, valueInt)

	case reflect.Int8:
		valueInt8, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, int8(valueInt8))

	case reflect.Int16:
		valueInt16, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, int16(valueInt16))

	case reflect.Int32:
		valueInt32, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, int32(valueInt32))

	case reflect.Int64:
		valueInt64, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, int64(valueInt64))

	case reflect.Uint:
		valueUint, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, uint(valueUint))

	case reflect.Uint8:
		valueUint8, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, uint8(valueUint8))

	case reflect.Uint16:
		valueUint16, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, uint16(valueUint16))

	case reflect.Uint32:
		valueUint32, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, uint32(valueUint32))

	case reflect.Uint64:
		valueUint64, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, uint64(valueUint64))

	case reflect.Float32:
		valueFloat32, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, float32(valueFloat32))

	case reflect.Float64:
		valueFloat64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, float64(valueFloat64))

	case reflect.Complex64:
		valueComplex64, err := strconv.ParseComplex(value, 64)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, complex64(valueComplex64))

	case reflect.Complex128:
		valueComplex128, err := strconv.ParseComplex(value, 128)
		if err != nil {
			return err
		}
		q.Args = append(q.Args, complex64(valueComplex128))

	default:
		switch typeValue {
		case reflect.TypeOf((*sql.NullBool)(nil)).Elem():
			valueBool, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			dbBool := sql.NullBool{}
			dbBool.Bool = valueBool
			dbBool.Valid = true
			q.Args = append(q.Args, dbBool)

		case reflect.TypeOf((*sql.NullTime)(nil)).Elem():
			t, err := time.Parse(dataFormat, value)
			if err != nil {
				return err
			}
			dbTime := sql.NullTime{}
			dbTime.Time = t
			dbTime.Valid = true
			q.Args = append(q.Args, dbTime)
		case reflect.TypeOf((*time.Time)(nil)).Elem():
			t, err := time.Parse(dataFormat, value)
			if err != nil {
				return err
			}
			dbTime := sql.NullTime{}
			dbTime.Time = t
			dbTime.Valid = true
			q.Args = append(q.Args, dbTime)

		case reflect.TypeOf((*sql.NullInt32)(nil)).Elem():
			valueInt32, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			dbInt32 := sql.NullInt32{}
			dbInt32.Int32 = int32(valueInt32)
			dbInt32.Valid = true
			q.Args = append(q.Args, dbInt32)

		case reflect.TypeOf((*sql.NullInt64)(nil)).Elem():
			valueInt64, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			dbInt64 := sql.NullInt64{}
			dbInt64.Int64 = int64(valueInt64)
			dbInt64.Valid = true
			q.Args = append(q.Args, dbInt64)

		case reflect.TypeOf((*sql.NullFloat64)(nil)).Elem():
			valueFloat64, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			q.Args = append(q.Args, float64(valueFloat64))
			dbFloat64 := sql.NullFloat64{}
			dbFloat64.Float64 = valueFloat64
			dbFloat64.Valid = true
			q.Args = append(q.Args, dbFloat64)
		default:
			return errors.New("Param query invalid")
		}

	}
	return nil
}
