package core

import "errors"

// Scene defines a scene, which can be displayed by the SceneManager
type Scene interface {
	Update()
	Input()
	Render()
	Name() string
}

// SceneManager containts the Scenes and displays the current one
type SceneManager interface {
	Scene

	// Add adds a scene to the manager
	Add(Scene)

	// Change changes the current scene to the scene with the given name
	Change(string) error
}

// sceneManager is a basic implementation of SceneManager
type sceneManager struct {
	scenes  map[string]Scene
	current Scene
}

// NewSceneManager returns a basic implementation of SceneManager
func NewSceneManager() SceneManager {
	sm := &sceneManager{}
	sm.scenes = make(map[string]Scene)
	return sm
}

func (sm *sceneManager) Update()      {}
func (sm *sceneManager) Input()       {}
func (sm *sceneManager) Render()      {}
func (sm *sceneManager) Name() string { return "" }
func (sm *sceneManager) Add(scene Scene) {
	sm.scenes[scene.Name()] = scene
}
func (sm *sceneManager) Change(name string) error {
	res, ok := sm.scenes[name]
	if !ok {
		return errors.New("Unknown scene name")
	}
	sm.current = res
	return nil
}
