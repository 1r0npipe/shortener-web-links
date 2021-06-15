import "flag"

type Flag struct {
	IpAddr string
	Port string
	Timeout string
}
var (
	FileConfig = flag.String("fileConfig", "", "define the file with config")
	IpAddr = flag.String("methodType", "online", "Choose online or offline, online by default")
	Port   = flag.String("unitType", "metric", "Choose unit time, Celcius by default")
	Timeout   = flag.String("city", "", "Choose city")
)

func Init() {
	flag.Parse()

}