package sql

import (
	"github.com/hnit-acm/go-common/basic"
	"reflect"
	"sync"
	"time"
)

// StructFormatter 格式化结构
type StructFormatter interface {
	// 域字符串处理
	StrFormatHandle(str string) string
	// 嵌套层次结构体处理符
	LayerFormatHandle() string
	// 域切片缓存
	Cache() func() FormatterCache
}

// 格式化器缓存结构
type FormatterCache interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, val interface{})
}

// DefaultGetFieldsArray 默认方法
var DefaultGetFieldsArray, DefaultGetFieldsString = NewGetFields(&defaultFormatter{})

// defaultFormatter 默认结构体格式化器实现
type defaultFormatter struct {
}

func (d *defaultFormatter) StrFormatHandle(str string) string {
	return basic.String(str).SnakeCasedString()
}

func (d *defaultFormatter) LayerFormatHandle() string {
	return "."
}

// defaultCache 默认缓存实现
type defaultCache struct {
	sync.Map
}

func (d *defaultCache) Get(key interface{}) (interface{}, bool) {
	return d.Map.Load(key)
}

func (d *defaultCache) Set(key interface{}, val interface{}) {
	d.Map.Store(key, val)
}

// Cache 闭包加入缓存
func (d *defaultFormatter) Cache() func() FormatterCache {
	cache := &defaultCache{}
	return func() FormatterCache {
		return cache
	}
}

type GetFieldsArray func(p interface{}, cacheNo string) basic.ArrayString

type GetFieldsString func(p interface{}, splitChar, cacheNo string) string

// NewGetFields 新建域获取func
//参数：
//	p 要取出的结构体
//	cacheNo 用于多func内重名结构体时唯一性确定
//	针对于path，因为path所用的字符串内存是提前分配好的，不需要重新分配，不需要使用builder
//tag:
//	alias: 字段别名，对应sql别名，
//	expr: 表达式，用于支持使用count，sum等函数
//
//
//
//
//
//
//

func NewGetFields(format StructFormatter) (GetFieldsArray, GetFieldsString) {
	cache := format.Cache()()
	GetFieldsArrayFunc := func(p interface{}, cacheNo string) basic.ArrayString {
		t := reflect.ValueOf(p).Elem()
		path := cacheNo + t.Type().PkgPath() + "." + t.Type().String()
		//key := crc32.ChecksumIEEE([]byte(path))
		if val, ok := cache.Get(path); ok {
			return val.(basic.ArrayString)
		}
		all := make(basic.ArrayString, 0)
		//fmt.Println(	reflect.ValueOf(a).Elem().Field(0).Kind())
		//fmt.Println(	reflect.ValueOf(a).Elem().Type()*.Field(0).Name)
		num := t.NumField()
		for i := 0; i < num; i++ {
			tf := t.Field(i)
			tt := t.Type().Field(i)
			t := 0
			switch tf.Interface().(type) {
			case time.Time:
				t = 1
			case *time.Time:
				t = 1
			}
			// if is struct
			if tf.Kind() == reflect.Struct && t == 0 {
				all = layerHandle(tf, all, format.StrFormatHandle(tt.Name), format)
				continue
			}
			name := tt.Name
			// expression tag
			if val, ok := tt.Tag.Lookup("expr"); ok {
				name = val
			}
			// alias tag
			if val, ok := tt.Tag.Lookup("alias"); ok {
				all = append(all, format.StrFormatHandle(name)+" as "+val)
				continue
			}
			all = append(all, format.StrFormatHandle(name))
		}
		cache.Set(path, all)
		return all
	}
	GetFieldsStringFunc := func(p interface{}, splitChar string, cacheNo string) string {
		t := reflect.ValueOf(p).Elem()
		pathStr := cacheNo + t.Type().PkgPath() + "." + t.Type().String() + "-string-" + splitChar
		// look up string cache.
		val, ok := cache.Get(pathStr)
		if ok {
			return val.(string)
		}
		// look up struct cache.
		pathStruct := cacheNo + t.Type().PkgPath() + "." + t.Type().String()
		val, ok = cache.Get(pathStruct)
		if ok {
			array := val.(basic.ArrayString)
			all := array.ToString(splitChar)
			//all := ArrayStringToStringBuilder(val.([]string), splitChar)
			cache.Set(pathStr, all)
			return all
		}
		// no cache.
		array := GetFieldsArrayFunc(p, cacheNo)
		all := array.ToString(splitChar)

		//all := ArrayStringToStringBuilder(array, splitChar)
		// store cache
		cache.Set(pathStr, all)
		return all
	}
	return GetFieldsArrayFunc, GetFieldsStringFunc

}

func layerHandle(t reflect.Value, all []string, prefix string, format StructFormatter) []string {
	num := t.NumField()
	for i := 0; i < num; i++ {
		tf := t.Field(i)
		tt := t.Type().Field(i)

		t := 0
		switch tf.Interface().(type) {
		case time.Time:
			t = 1
		}
		if tf.Kind() == reflect.Struct && t == 0 {
			all = layerHandle(tf, all, prefix+format.LayerFormatHandle()+format.StrFormatHandle(tt.Name), format)

			continue
		}

		name := tt.Name
		// 表达式tag
		if val, ok := tt.Tag.Lookup("expr"); ok {
			name = val
			// 别名tag
			if val, ok := tt.Tag.Lookup("alias"); ok {
				all = append(all, format.StrFormatHandle(name)+" as "+val)
				continue
			}
			all = append(all, format.StrFormatHandle(name))
			continue
		}
		// 别名tag
		if val, ok := tt.Tag.Lookup("alias"); ok {
			all = append(all, prefix+format.LayerFormatHandle()+format.StrFormatHandle(name)+" as "+val)
			continue
		}
		all = append(all, prefix+format.LayerFormatHandle()+format.StrFormatHandle(name))
	}
	return all
}

var DefaultStructToMap = NewStructToMap(&defaultFormatter{})

type Filter func(v interface{}) bool

//参数:
//	p: 用于生成的结构体
//	filter: 字段过滤器
//tag:
//	name: 表字段名，不设置则默认，蛇形结构体字段名
//	op: 数据库操作符， =,<,>,<>,<=,>=,like等,默认为 =
//结果集：map[key]value
//	key为name+op+?,value为结构体字段值
//example:
//
//	type Sql struct{
//		FirstName string 	`op:"like"`
//		LastName string 	`op:"="`
//		FromTime time.Time 	`name:"create_time" op:">"`
//		ToTime time.Time 	`name:"create_time" op:"<"`
//	}
//	s := session.Session.Table(this.Table())
//	params := Sql{FirstName:"tom%",LastName:"tomm",FromTime:略,ToTime:time.Now()}
//	paramsMap := db.DefaultStructToMap(params, convertor.IsEmpty)
//	for _, val := range paramsMap {
//		for k, v := range val {
//			s = s.Where(k, v)
//		}
//	}
//  使用过滤器可以使用不定参数
//	使用and生成的Sql为，where first_name like ? and last_name = ? and create_time > ? and create_time < ?
//
func NewStructToMap(format StructFormatter) func(p interface{}, zeroFilter ...Filter) []map[string]interface{} {
	return func(p interface{}, zeroFilter ...Filter) (all []map[string]interface{}) {
		e := reflect.ValueOf(p)
		t := e.Type()
		all = make([]map[string]interface{}, 0)

	for1:
		for i := 0; i < e.NumField(); i++ {
			for j := range zeroFilter {
				x := e.Field(i).Interface()
				if zeroFilter[j](x) {
					continue for1
				}
			}
			op := "="
			name := format.StrFormatHandle(t.Field(i).Name)
			if val, ok := t.Field(i).Tag.Lookup("op"); ok && val != "" {
				//key = format.StrFormatHandle(t.Field(i).Name) + " " + val + " ? "
				op = val
			}

			if val, ok := t.Field(i).Tag.Lookup("name"); ok && val != "" {
				name = val
			}

			key := name + " " + op + " ? "

			item := map[string]interface{}{key: e.Field(i).Interface()}
			all = append(all, item)
		}
		return all
	}
}
