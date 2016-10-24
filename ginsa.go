package ginsa

var (
	VERSION  = "0.0.0"
	REVISION = "deadbeaf"
)

func FullVersion() string {
	return VERSION + "-" + REVISION
}
