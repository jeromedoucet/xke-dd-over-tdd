// black box testing
package converter_test
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"xke-dd-over-tdd/converter"
)

func Test_Should_convert_A_To_0(t *testing.T) {
	number, _ := converter.Convert('A')
	assert.Equal(t, 0, number)
}

func Test_Should_convert_a_To_0(t *testing.T) {
	number, _ := converter.Convert('a')
	assert.Equal(t, 0, number)
}


func Test_Should_convert_b_To_1(t *testing.T) {
	number, _ := converter.Convert('b')
	assert.Equal(t, 1, number)
}

func Test_Should_convert_B_To_1(t *testing.T) {
	number, _ := converter.Convert('B')
	assert.Equal(t, 1, number)
}

func Test_Should_convert_C_To_2(t *testing.T) {
	number, _ := converter.Convert('C')
	assert.Equal(t, 2, number)
}

func Test_Should_convert_9_To_9(t *testing.T) {
	number, _ := converter.Convert('9')
	assert.Equal(t, 9, number)
}

func Test_Should_Get_Error(t *testing.T) {
	_, err := converter.Convert('ù')
	assert.Equal(t, err.Error(), "Invalid value")
}
