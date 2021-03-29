package encoding

import (
	"log"
	"testing"
)

func TestEncode64(t *testing.T) {
	var c = Encode64("赵云涛")
	log.Println(c)
	c, _ = Decode64(c)
	log.Println(c)
}
