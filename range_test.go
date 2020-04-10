package tddbc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("下端点 3, 上端点 8 を渡すと整数閉区間オブジェクト[3,8]を返す", func(t *testing.T) {
		expected := Range{
			lower: 3,
			upper: 8,
		}

		actual, err := Create(3, 8)
		assert.NoError(t, err)
		assert.Equal(t, expected.lower, actual.lower)
		assert.Equal(t, expected.upper, actual.upper)

	})

	t.Run("下端点 8, 上端点 3 の整数閉区間を渡すとエラーを返す", func(t *testing.T) {
		expectedError := "エラー：上端点より下端点が大きい閉区間を作ることはできない"

		_, err := Create(8, 3)
		assert.Equal(t, expectedError, err.Error())

	})

}

func TestToString(t *testing.T) {
	t.Run("整数閉区間オブジェクトは文字列表現を返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		expected := "[3,8]"

		actual := r.ToString()

		assert.Equal(t, expected, actual)
	})
}

func TestIncludeNum(t *testing.T) {
	t.Run("整数閉区間オブジェクトは指定した整数を含むときTrueを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		n := int64(4)
		actual, err := r.IncludeNum(n)
		assert.NoError(t, err)
		assert.True(t, actual)
	})

	t.Run("整数閉区間オブジェクトは指定した整数を含まないときFalseを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		n := int64(9)
		actual, err := r.IncludeNum(n)
		assert.NoError(t, err)
		assert.False(t, actual)
	})

}

func TestEqualRange(t *testing.T) {
	t.Run("下端点 3, 上端点 8, 下端点 3, 上端点 8 を渡すとTrueを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		cmp := &Range{
			lower: 3,
			upper: 8,
		}

		actual, err := r.EqualRange(cmp)
		assert.NoError(t, err)
		assert.True(t, actual)
	})

	t.Run("下端点 3, 上端点 8, 下端点 2, 上端点 8 を渡すとTrueを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		cmp := &Range{
			lower: 2,
			upper: 8,
		}

		actual, err := r.EqualRange(cmp)
		assert.NoError(t, err)
		assert.False(t, actual)
	})
}

func TestIncludeRange(t *testing.T) {
	t.Run("下端点 3, 上端点 8, 下端点 4, 上端点 7 を渡すとTrueを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		cmp := &Range{
			lower: 4,
			upper: 7,
		}
		actual, err := r.IncludeRange(cmp)
		assert.NoError(t, err)
		assert.True(t, actual)
	})

	t.Run("下端点 3, 上端点 8, 下端点 3, 上端点 9 を渡すとFalseを返す", func(t *testing.T) {
		r := &Range{
			lower: 3,
			upper: 8,
		}

		cmp := &Range{
			lower: 4,
			upper: 9,
		}
		actual, err := r.IncludeRange(cmp)
		assert.NoError(t, err)
		assert.False(t, actual)
	})

}
