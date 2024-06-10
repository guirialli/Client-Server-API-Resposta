package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getBind(ctx context.Context) (string, error) {
	req, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
		return getBind(ctx)
	}
	defer req.Body.Close()

	bind, _ := io.ReadAll(req.Body)

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("client error timeout excepiton")
	default:
		return string(bind), nil
	}
}

func Request() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	bind, err := getBind(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	formatedBind := fmt.Sprintf("Dólar: {%s}", strings.Split(bind, "\n")[0])
	file.Write([]byte(formatedBind))
}
