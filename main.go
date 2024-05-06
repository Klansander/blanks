package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
	"time"
)

type ErrorApp struct {
	Code    int
	Message string
}

func (e *ErrorApp) Error() string {

	return e.Message

}

func New(code int, message string) error {

	err := &ErrorApp{Code: code, Message: message}

	return errors.Wrap(err, wrapPosition(2))

}
func Wrap(err error) error {

	return errors.Wrap(err, wrapPosition(2))

}

func Cause(err error) error {

	return errors.Cause(err)

}

func wrapPosition(level int) string {

	f := getFrame(2)
	_, filename := path.Split(f.File)
	dir := strings.Split(f.Function, ".")[0]

	return fmt.Sprintf("--> %s/%s/%d", dir, filename, f.Line)

}

func getFrame(skipFrames int) runtime.Frame {

	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame

}

func simulateWork(ctx context.Context) (err error) {
	// Имитация длительной операции
	chFinish := make(chan bool)

	go func() {
		defer func() { chFinish <- true }()

		ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
		defer cancel()

		select {
		case <-time.After(2 * time.Second): // Предположим, работа занимает 2 секунды
			fmt.Println("Work completed successfully")

		case <-ctx.Done():

			err = Wrap(ctx.Err())

		}

	}()

	for {
		select {
		case <-ctx.Done():
			return Wrap(ctx.Err())

		case <-chFinish:
			return err

		}
	}

	//if ctx.Err() != nil {
	//	err = errors.New(http.StatusFailedDependency, "необходимо повторить запрос еще раз")
	//} else {
	//	err = errors.Wrap(err)
	//}

}

func main() {
	// Создание контекста с таймаутом
	// Отмена контекста при выходе из функции
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Вызов функции с контекстом
	err := simulateWork(ctx)
	fmt.Println(1)
	// Обработка ошибки, возвращаемой функцией
	if err != nil {
		fmt.Println(Cause(err), 1111, context.Cause(ctx))
		if errors.Is(context.Cause(ctx), context.DeadlineExceeded) {
			fmt.Println("Error: Task didn't finish within the deadline")
		} else {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println(2)
	for i := 0; i < 10; i++ {
		err = simulateWork(ctx)
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			if errors.Is(err, context.DeadlineExceeded) {

				err = simulateWork(ctx)
				if err != nil {
					logrus.Errorln("Ошибка при создание пакетной задачи\nВходные данные: ", i)
					continue
				}

			}

		}
		fmt.Println(111)
	}
	fmt.Println(4)
}
