package conv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 定义一个测试用的结构体
type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// TestTrans 测试 Trans 函数的正确性
func TestTrans(t *testing.T) {
	src := TestStruct{Name: "Alice", Age: 30}
	dst, err := Trans[TestStruct](src)

	assert.NoError(t, err)
	assert.Equal(t, src, dst, "源对象和目标对象应当相等")

	// 测试错误情况，比如使用不支持序列化的类型
	_, err = Trans[chan int](make(chan int))
	assert.Error(t, err, "使用不支持序列化的类型应当返回错误")
}

// TestMustTrans 测试 MustTrans 函数的正确性
func TestMustTrans(t *testing.T) {
	src := TestStruct{"Bob", 25}
	dst := MustTrans[TestStruct](src)

	assert.NotNil(t, dst)
	assert.Equal(t, src, *dst, "源对象和目标对象应当相等")

	// 测试 nil 输入的情况
	result := MustTrans[TestStruct](nil)
	assert.Nil(t, result, "输入 nil 应当返回 nil")
}

// BenchmarkTrans 性能测试 Trans 函数
func BenchmarkTrans(b *testing.B) {
	src := TestStruct{"Benchmark", 100}

	for i := 0; i < b.N; i++ {
		_, _ = Trans[TestStruct](src)
	}
}

// BenchmarkMustTrans 性能测试 MustTrans 函数
func BenchmarkMustTrans(b *testing.B) {
	src := TestStruct{"Benchmark", 100}

	for i := 0; i < b.N; i++ {
		_ = MustTrans[TestStruct](src)
	}
}
