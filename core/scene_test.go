package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type scene struct {
	name string
}

func (s *scene) Update()      {}
func (s *scene) Input()       {}
func (s *scene) Render()      {}
func (s *scene) Name() string { return s.name }

func TestSceneManager(t *testing.T) {
	Convey("A new sceneManager", t, func() {
		sm := NewSceneManager().(*sceneManager)
		Convey("And a new scene", func() {
			s1 := scene{"s1"}
			Convey("Adding the scene", func() {
				sm.Add(&s1)
				So(sm.scenes[0], ShouldEqual, &s1)
			})
		})
	})
}
