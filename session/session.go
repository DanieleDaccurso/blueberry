package session

type SaverInterface interface {
	Save(string, map[string]string) error
	Load(string) (map[string]string, error)
	Delete(string) error
}

type Session struct {
	kv    map[string]string
	saver SaverInterface
}

func (s *Session) Set(key string, value interface{}) error {

}

func (s *Session) GetString(key string) string {

}

func (s *Session) GetInt(key string) int {

}

func (s *Session) GetFloat32(key string) float32 {

}

func (s *Session) GetFloat64(key string) float64 {

}
