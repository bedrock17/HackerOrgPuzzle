package game

//SetDebugMode debug모드 설정
func SetDebugMode(mode bool) {
	gDebugMode = mode
}

//SetGoTestMode 단위테스트 모드 설정
func SetGoTestMode(mode bool) {
	goTestMode = mode
}
