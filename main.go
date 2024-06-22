package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	waktu := time.Now()
	UserId := 20
	value, err := getUser(context.Background(), UserId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("print value", +value)
	fmt.Println("took: ", +time.Since(waktu))

}

type Response struct {
	Value int
	Err   error
}

func getUser(ctx context.Context, UserId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*150)
	defer cancel()
	reschnl := make(chan Response)
	go func() {
		val, err := GetUserIdToLong()
		reschnl <- Response{
			Value: val,
			Err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("ambil data dari therd party lain terlalu lama")
		case res := <-reschnl:
			return res.Value, res.Err
		}
	}
}

func GetUserIdToLong() (int, error) {
	time.Sleep(time.Millisecond * 140)
	return 666, nil
}
