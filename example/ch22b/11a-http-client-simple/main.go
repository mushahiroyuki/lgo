package main
import(
	"fmt"
	"io"
	"net/http"
	"log"
)

func main() {
	url := "https://www.marlin-arms.com/support/learning-go/example-jpn/simple-page.html"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Printf("%s", body)
}
