package encoding

import (
	"log"
	"testing"
)

func TestEncode64(t *testing.T) {
	var c = Encode64("demo")
	log.Println(c)
	c, _ = Decode64(c)
	log.Println(c)
}
