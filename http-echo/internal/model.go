package internal

type PageModel struct {
	ServerStartedOn  string
	ServerRenderedOn string
	SomeString       string
	ServerDate       string
	Additional       SomeOtherInfo
}

type SomeOtherInfo struct {
	IntValue int32
}
