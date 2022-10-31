package goconvertstruct

type DummyStructTest struct {
	Field1  map[string]string `json:"field_1"`
	Field2  int               `custom:"field_2"`
	Field3  []int
	field4  string
	field5  *DummyStructTest
	Field6  *DummyStructTest           `json:"field_6"`
	Field7  *int                       `json:"field_7"`
	Field8  []DummyStructTest          `json:"field_8"`
	Field9  map[string]DummyStructTest `json:"field_9"`
	Field10 interface{}                `json:"field_10"`
}
