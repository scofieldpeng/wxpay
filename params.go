package wxpay

import (
	"encoding/xml"
	"io"
	"strconv"
)

type Params map[string]string

// map本来已经是引用类型了，所以不需要 *Params
func (p Params) SetString(k, s string) Params {
	p[k] = s
	return p
}

func (p Params) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p Params) SetInt64(k string, i int64) Params {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p Params) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// 判断key是否存在
func (p Params) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m Params) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	start.Name.Local = "xml" // 更改xml开始标签
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		if err := e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (m *Params) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = Params{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func XmlToMap(xmlStr string) Params {

	params := make(Params)
	err := xml.Unmarshal([]byte(xmlStr), &params)
	if err != nil {
		panic(err)
	}
	return params
}

func MapToXml(params Params) string {
	buf, err := xml.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
