


package main
import (
	"log"		
)

func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
// Println writes to the standard logger.
	log.Println(" TODO ")

}
