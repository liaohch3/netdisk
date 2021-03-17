package cache

var sessionMap map[string]string

func InitSessionMap() {
	sessionMap = make(map[string]string)
}

func UpdateSessionMap(name, session string) {
	sessionMap[name] = session
}

func GetSession(name string) (string, bool) {
	session, ok := sessionMap[name]
	return session, ok
}
