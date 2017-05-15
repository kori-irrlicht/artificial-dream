package core

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
	Add(Scene)
	Change(string)
}

// sceneManager is a basic implementation of SceneManager
type sceneManager struct {
}

// NewSceneManager returns a basic implementation of SceneManager
func NewSceneManager() SceneManager {
	return &sceneManager{}
}

func (sm *sceneManager) Update()            {}
func (sm *sceneManager) Input()             {}
func (sm *sceneManager) Render()            {}
func (sm *sceneManager) Name() string       { return "" }
func (sm *sceneManager) Add(scene Scene)    {}
func (sm *sceneManager) Change(name string) {}
