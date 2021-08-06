package builder

import "testing"

func TestBuilder(t *testing.T) {
	t.Run("builder1", func(t *testing.T) {
		builder1 := &Builder1{}
		dirctor := NewDirector(builder1)
		dirctor.Construct()

		get := builder1.GetResult()
		want := "123"
		if get != want {
			t.Fatalf("Builder1 fail expect %s acture %s", want, get)
		}
	})

	t.Run("builder2", func(t *testing.T) {
		builder1 := &Builder2{}
		dirctor := NewDirector(builder1)
		dirctor.Construct()

		get := builder1.GetResult()
		want := 6
		if get != want {
			t.Fatalf("Builder1 fail expect %d acture %d", want, get)
		}
	})
}
