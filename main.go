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
	remoteAddr := flag.String("r", "127.0.0.1:1234", "remote address")
	flag.Parse()
	conn, err := net.Dial(NET_PROTOCOL, *remoteAddr)
	if err != nil {
		panic(err)
	}
	localAddr	:= conn.LocalAddr().String()
	conn.Close()
	fmt.Println(localAddr)
}
