package ch5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TooShortFileNameShouldLogError(t *testing.T) {
	fw := &FakeWebService{}  // 建立假物件
	la := NewLogAnalyzer(fw)

	tooShortFileName := "abc.txt"
	la.Analyze(tooShortFileName)
	
	expected := "File name too short: " + "abc.txt"
	actual := fw.lastError

	assert.Equal(t, expected, actual, "they should be equal")  // 把假物件當模擬物件來使用並驗證
}