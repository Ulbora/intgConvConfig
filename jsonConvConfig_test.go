package intgconvconf

import (
	"fmt"
	"testing"
)

func TestGetSpfConversion(t *testing.T) {
	var c SpfConvConf
	var jc JSONConvConf
	c = &jc
	res := c.GetSpfConversion("./confFileTest")
	if len(*res) <= 0{
		t.Fail()
	}
	fmt.Println("res: ", res)
	e1 := (*res)["test1"]
	fmt.Println(" e1.CfDirectory: ",  e1.CfDirectory)
	if e1.CfDirectory != "test1"{
		t.Fail()
	}
}