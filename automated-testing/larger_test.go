package compare

func TestFirstLarger(t *testinf.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Error()
	}
}
