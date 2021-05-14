package sqlh

import (
	"github.com/hnit-acm/hfunc/basich"
	"reflect"
	"time"
)

type StructFormatter func() (basich.StringFormatFunc, CacheFunc, LayerSplitFunc)

type LayerSplitFunc func() string

type CacheFunc func() (basich.GetFunc, basich.SetFunc)

// DefaultGetFieldsArray 默认方法
var DefaultGetFieldsArray, DefaultGetFieldsString = NewGetFields(
	func() (basich.StringFormatFunc, CacheFunc, LayerSplitFunc) {
		LayerSplitFunc := LayerSplitFunc(func() string {
			return ","
		})

		CacheFunc := CacheFunc(func() (basich.GetFunc, basich.SetFunc) {
			get, set, _ := basich.NewHashMapFunc(1024)
			return get, set
		})

		return basich.SnakeCasedStringFormat,
			CacheFunc,
			LayerSplitFunc
	},
)

type GetFieldsArrayFunc func(p interface{}, cacheNo string) basich.ArrayString

type GetFieldsStringFunc func(p interface{}, splitChar, cacheNo string) string

// NewGetFields 新建域获取func
//参数：
//	p 要取出的结构体
//	cacheNo 用于多func内重名结构体时唯一性确定
//	针对于path，因为path所用的字符串内存是提前分配好的，不需要重新分配，不需要使用builder
//tag:
//	alias: 字段别名，对应sql别名，
//	expr: 表达式，用于支持使用count，sum等函数

func NewGetFields(structFormatter StructFormatter) (GetFieldsArrayFunc, GetFieldsStringFunc) {
	format, cache, split := structFormatter()
	get, set := cache()
	GetFieldsArrayFunc := GetFieldsArrayFunc(func(p interface{}, cacheNo string) basich.ArrayString {
		t := reflect.ValueOf(p).Elem()
		path := cacheNo + t.Type().PkgPath() + "." + t.Type().String()
		//key := crc32.ChecksumIEEE([]byte(path))
		if val, ok := get(path); ok {
			return val.(basich.ArrayString)
		}
		all := make(basich.ArrayString, 0)
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
				all = layerHandle(tf, all, format(tt.Name), format, split)
				continue
			}
			name := tt.Name
			// expression tag
			if val, ok := tt.Tag.Lookup("expr"); ok {
				name = val
			}
			// alias tag
			if val, ok := tt.Tag.Lookup("alias"); ok {
				all = append(all, format(name)+" as "+val)
				continue
			}
			all = append(all, format(name))
		}
		set(path, all)
		return all
	})
	GetFieldsStringFunc := GetFieldsStringFunc(func(p interface{}, splitChar string, cacheNo string) string {
		t := reflect.ValueOf(p).Elem()
		pathStr := cacheNo + t.Type().PkgPath() + "." + t.Type().String() + "-string-" + splitChar
		// look up string cache.
		val, ok := get(pathStr)
		if ok {
			return val.(string)
		}
		// look up struct cache.
		pathStruct := cacheNo + t.Type().PkgPath() + "." + t.Type().String()
		val, ok = get(pathStruct)
		if ok {
			array := val.(basich.ArrayString)
			all := array.GetFunc().ToString(splitChar)
			//all := ArrayStringToStringBuilder(val.([]string), splitChar)
			set(pathStr, all)
			return all
		}
		// no cache.
		array := GetFieldsArrayFunc(p, cacheNo)
		all := array.GetFunc().ToString(splitChar)

		//all := ArrayStringToStringBuilder(arrayh, splitChar)
		// store cache
		set(pathStr, all)
		return all
	})
	return GetFieldsArrayFunc, GetFieldsStringFunc

}

func layerHandle(t reflect.Value, all []string, prefix string, format basich.StringFormatFunc, split LayerSplitFunc) []string {
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
			all = layerHandle(tf, all, prefix+split()+format(tt.Name), format, split)

			continue
		}

		name := tt.Name
		// 表达式tag
		if val, ok := tt.Tag.Lookup("expr"); ok {
			name = val
			// 别名tag
			if val, ok := tt.Tag.Lookup("alias"); ok {
				all = append(all, format(name)+" as "+val)
				continue
			}
			all = append(all, format(name))
			continue
		}
		// 别名tag
		if val, ok := tt.Tag.Lookup("alias"); ok {
			all = append(all, prefix+split()+format(name)+" as "+val)
			continue
		}
		all = append(all, prefix+split()+format(name))
	}
	return all
}

var DefaultStructToMap = NewStructToMap(
	func() (basich.StringFormatFunc, CacheFunc, LayerSplitFunc) {
		return basich.SnakeCasedStringFormat, nil, nil
	},
)

type FilterFunc func(v interface{}) bool

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
func NewStructToMap(structFormatter StructFormatter) func(p interface{}, zeroFilter ...FilterFunc) []map[string]interface{} {
	format, _, _ := structFormatter()
	return func(p interface{}, zeroFilter ...FilterFunc) (all []map[string]interface{}) {
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
			name := format(t.Field(i).Name)
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
