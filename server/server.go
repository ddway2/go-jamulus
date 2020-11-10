package server

import (
	"fmt"
	"sync"

	"github.com/ddway2/go-jamulus/cli"
	"github.com/ddway2/opus"
)

const (
	MAX_SERVER_CHANNEL = 150

	SYSTEM_SAMPLE_RATE_HZ            = 48000
	SYSTEM_FRAME_SIZE_SAMPLES        = 64
	DOUBLE_SYSTEM_FRAME_SIZE_SAMPLES = SYSTEM_FRAME_SIZE_SAMPLES * 2
)

// Info contains server informations
type Info struct {
}

type OpusChannel struct {
	opusMode          *opus.OpusMode
	opusEncoderMono   *opus.EncoderCustom
	opusDecoderModo   *opus.DecoderCustom
	opusEncoderStereo *opus.EncoderCustom
	opusDecoderStereo *opus.DecoderCustom

	opus64Mode          *opus.OpusMode
	opus64EncoderMono   *opus.EncoderCustom
	opus64DecoderModo   *opus.DecoderCustom
	opus64EncoderStereo *opus.EncoderCustom
	opus64DecoderStereo *opus.DecoderCustom
}

// Server is the main application
type Server struct {
	mu   sync.Mutex
	done chan bool

	Conf   *cli.Config
	socket *Socket

	// OPUS encoder/decoder
	channels [MAX_SERVER_CHANNEL]OpusChannel
}

// NewServer initialize server from Config
func NewServer(conf *cli.Config) (*Server, error) {
	var s Server
	if err := s.Init(conf); err != nil {
		return nil, err
	}
	return &s, nil
}

// Init initialized constructor
func (s *Server) Init(conf *cli.Config) error {
	s.Conf = conf

	sc, err := NewSocket(s)
	if err != nil {
		return fmt.Errorf("Unable to create socket")
	}

	s.socket = sc

	for _, v := range s.channels {
		v.opusMode, _ = opus.NewOpusMode(SYSTEM_SAMPLE_RATE_HZ, DOUBLE_SYSTEM_FRAME_SIZE_SAMPLES)
		v.opus64Mode, _ = opus.NewOpusMode(SYSTEM_SAMPLE_RATE_HZ, SYSTEM_FRAME_SIZE_SAMPLES)

		v.opus64DecoderModo, _ = opus.NewDecoderCustom(1, v.opus64Mode)
		v.opus64DecoderStereo, _ = opus.NewDecoderCustom(2, v.opus64Mode)
		v.opus64EncoderMono, _ = opus.NewEncoderCustom(1, v.opus64Mode)
		v.opus64EncoderStereo, _ = opus.NewEncoderCustom(2, v.opus64Mode)

		v.opusDecoderModo, _ = opus.NewDecoderCustom(1, v.opusMode)
		v.opusDecoderStereo, _ = opus.NewDecoderCustom(2, v.opusMode)
		v.opusEncoderMono, _ = opus.NewEncoderCustom(1, v.opusMode)
		v.opusEncoderStereo, _ = opus.NewEncoderCustom(2, v.opusMode)

		// we require a constant bit rate
		v.opus64EncoderMono.SetVBR(false)
		v.opus64EncoderStereo.SetVBR(false)
		v.opusEncoderMono.SetVBR(false)
		v.opusEncoderStereo.SetVBR(false)

		// for 64 samples frame size we have to adjust the PLC behavior to avoid loud artifacts
		v.opus64EncoderMono.SetPacketLossPerc(35)
		v.opus64EncoderStereo.SetPacketLossPerc(35)

		// we want as low delay as possible
		v.opus64EncoderMono.SetApplication(opus.AppRestrictedLowdelay)
		v.opus64EncoderStereo.SetApplication(opus.AppRestrictedLowdelay)
		v.opusEncoderMono.SetApplication(opus.AppRestrictedLowdelay)
		v.opusEncoderStereo.SetApplication(opus.AppRestrictedLowdelay)

		// set encoder low complexity for legacy 128 samples frame size
		v.opusEncoderMono.SetComplexity(1)
		v.opusEncoderStereo.SetComplexity(1)
	}

	return nil
}

// Start server
func (self *Server) Start() {

}

// Shutdown server if available
func (s *Server) Shutdown() {
	for _, v := range s.channels {
		v.opus64DecoderModo.Close()
		v.opus64DecoderStereo.Close()
		v.opus64EncoderMono.Close()
		v.opus64EncoderStereo.Close()

		v.opusDecoderModo.Close()
		v.opusDecoderStereo.Close()
		v.opusEncoderMono.Close()
		v.opusEncoderStereo.Close()

		v.opusMode.Close()
		v.opus64Mode.Close()
	}
}

// MixEncodeTransmitData preare data from other clients
func (self *Server) mixEncodeTransmitData( /*Channel count, num client*/ ) {

}
