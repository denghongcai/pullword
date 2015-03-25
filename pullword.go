package pullword

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type request struct {
	source string
	param1 float32
	param2 uint
}

func NewRequest(source string, threshold float32, debug bool) request {
	var param2 uint
	if debug {
		param2 = 1
	} else {
		param2 = 0
	}

	return request{
		source: source,
		param1: threshold,
		param2: param2,
	}
}

func (req request) Do() ([]string, error) {
	conn, err := net.Dial("tcp", "api.pullword.com:2015")
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	_, err = writer.WriteString(fmt.Sprintf("%s\t%1.2f\t%d]\r\n", req.source, req.param1, req.param2))
	if err != nil {
		return nil, err
	}
	writer.Flush()
	if err != nil {
		return nil, err
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.Trim(line, "\r\n"), " "), nil
}
