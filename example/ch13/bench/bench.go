package bench

import "os"

func FileLen(f string, bufsize int) (int, error) { //liststart
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	count := 0
	for {
		buf := make([]byte, bufsize)
		num, err := file.Read(buf)
		count += num
		if err != nil {
			break
		}
	}
	return count, nil
} //listend

