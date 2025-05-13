package request

type SceneUpdateEnable struct {
	ID     int  `json:"id"`
	Enable bool `json:"enable"`
}
