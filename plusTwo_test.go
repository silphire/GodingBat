package array1

func TestPlusTwo(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	actual := plusTwo([]int{1, 2}, []int{3, 4})
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("value mismatch")
		}
	}
}