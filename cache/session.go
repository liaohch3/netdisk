package cache

var sessionMap map[string]string

func InitSessionMap() {
	sessionMap = make(map[string]string)
}

func UpdateSessionMap(name, session string) {
	//sessionMap[name] = session
	client.Set(name, session, 0)
}

func GetSession(name string) (string, error) {
	//session, ok := sessionMap[name]
	//return session, ok
	return client.Get(name).Result()
}
