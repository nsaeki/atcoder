import (
	"bufio"
	"os"
	"strconv"
)

type AcScanner struct {
	scanner *bufio.Scanner
}

func NewScanner() *AcScanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	// MaxScanTokenSize = 64 * 1024
	// buf := make([]byte, 10000)
	// scanner.Buffer(buf, 1000000)
	return &AcScanner{scanner}
}

func (s *AcScanner) Scan() bool {
	if res := s.scanner.Scan(); res {
		return true
	}
	if err := s.scanner.Err(); err != nil {
		panic(err)
	}
	return false
}

func (s *AcScanner) ScanInt() int {
	s.Scan()
	t := s.scanner.Text()
	v, err := strconv.Atoi(t)
	if err != nil {
		return v
	} else {
		panic(err)
	}
}

func (s *AcScanner) ScanInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		x := s.ScanInt()
		ret[i] = x
	}
	return ret
}

func (s *AcScanner) ScanString() string {
	s.Scan()
	return s.scanner.Text()
}

func (s *AcScanner) ScanStrings(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		t := s.ScanString()
		ret[i] = t
	}
	return ret
}

func (s *AcScanner) ScanBytes() []byte {
	s.Scan()
	return []byte(s.scanner.Text())
}
