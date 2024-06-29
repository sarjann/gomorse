package sound

import (
	// "fmt"
	"github.com/ebitengine/oto/v3"
	"io"
	"time"
)

var SAMPLE_RATE = 44100
var otoCtx *oto.Context
var otoSet = false


// type AudioReader struct {
// 	buffer  []byte
// 	readIdx int64
// }
//
// func (r *AudioReader) Read(p []byte) (n int, err error) {
// 	if r.readIdx >= int64(len(r.buffer)) {
// 		return 0, io.EOF
// 	}
// 	copy(p, r.buffer[r.readIdx:])
//
// 	r.readIdx += int64(n)
// 	if n != 0 {
// 		fmt.Println(n)
// 	}
// 	return 0, nil
// }
//
// func (r *AudioReader) Reset() {
// 	r.readIdx = 0
// }
//
//
// type AudioContext struct {
// }

func getCtx() *oto.Context  {
	op := &oto.NewContextOptions{}
	op.SampleRate = SAMPLE_RATE
	op.ChannelCount = 1
	op.Format = oto.FormatUnsignedInt8

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("Context creation failed")
	}
	<-readyChan
    return otoCtx
}


func PlaySound(audioReader io.Reader) {
    if !otoSet {
        otoCtx = getCtx()
        otoSet = true
    }

	player := otoCtx.NewPlayer(audioReader)
	player.SetVolume(0.2)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

    err := player.Close()
	if err != nil {
		panic(err)
	}
}

func SquareBuffer(frequency int, duration int) []byte {
	BUFFER_SIZE := duration * SAMPLE_RATE

	byteBuffer := make([]byte, BUFFER_SIZE)

	var cur int = 0
	for i := 0; i < BUFFER_SIZE; i++ {
		frac := (SAMPLE_RATE / frequency) / 2
		if (i % frac) == 0 {
			if cur == 0 {
				cur = 255
			} else {
				cur = 0
			}
		}
		byteBuffer[i] = byte(cur)

	}
	return byteBuffer
}

func SilentBuffer(duration int) []byte {
	BUFFER_SIZE := duration * SAMPLE_RATE

	byteBuffer := make([]byte, BUFFER_SIZE)
	for i := 0; i < BUFFER_SIZE; i++ {
		byteBuffer[i] = 0
	}
	return byteBuffer
}
