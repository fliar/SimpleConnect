package main
import(
	"fmt"
	"net"
	"flag"
)

const NET_PROTOCOL 			= "tcp"
const SERVER_IP				= "192.168.202.78"
const PRIMARY_PORT			= ":8013"
const SERVER_PRIMARY_ADDR	= SERVER_IP + PRIMARY_PORT
func main(){
	remoteAddr := flag.String("r", SERVER_PRIMARY_ADDR, "remote address")
	flag.Parse()
	conn, err := net.Dial(NET_PROTOCOL, *remoteAddr)
	if err != nil {
		panic(err)
	}
	localAddr	:= conn.LocalAddr().String()
	conn.Close()
	fmt.Println(localAddr)
}
