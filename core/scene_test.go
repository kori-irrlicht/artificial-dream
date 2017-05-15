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
				So(sm.scenes[s1.Name()], ShouldEqual, &s1)
				Convey("And changing to it", func() {
					sm.Change(s1.Name())
					So(sm.current, ShouldEqual, &s1)
				})
				Convey("Adding another scene", func() {
					s2 := scene{"s2"}
					sm.Add(&s2)
					So(sm.scenes[s2.Name()], ShouldEqual, &s2)
					Convey("And changing to it", func() {
						sm.Change(s2.Name())
						So(sm.current, ShouldEqual, &s2)
					})
				})
			})
		})
	})
}
