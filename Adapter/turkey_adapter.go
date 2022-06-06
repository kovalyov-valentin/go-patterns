package main 

type turkeyAdapter struct {
	turkey turkey
}

func newTurkeyAdapter(t turkey) *turkeyAdapter {
	return &turkeyAdapter{
		turkey: t,
	}
}

func (t *turkeyAdapter) quack() {
	t.turkey.gobble()
}

func (t *turkeyAdapter) fly() {
	for i := 1; i <=5; i++ {
		t.turkey.fly()
	}
}