package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected{
		t.Errorf("got %q want %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i:=0; i<b.N; i++{
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	result := Repeat("Z",3)
	fmt.Println(result)
	//output: ZZZ
}