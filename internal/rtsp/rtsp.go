package rtsp

import (
	"net"

	"github.com/Lyyyttooon/zeusmedia/utils"
)

type RTSP struct {
	rawURL string
	urlCtx utils.URLContext

	conn net.Conn
}

// NewPullSession 新建RTSP结构体
func NewPullSession(url string) RTSP {
	return RTSP{
		rawURL: url,
	}
}

// Conn 连接
func (r *RTSP) Conn() {
	errChan := make(chan error, 1)

	go func() {
		if err := r.connect(); err != nil {
			errChan <- err
		}
	}()
}

// connect 建立连接
func (r *RTSP) connect() (err error) {
	r.urlCtx, err = utils.ParseRTSPURL(r.rawURL)
	if err != nil {
		return err
	}
	utils.SLogger.Debugf("[%s] > tcp connect.", r.rawURL)

	// 建立连接
	conn, err := net.Dial("tcp", r.urlCtx.HostWithPort)
	if err != nil {
		return err
	}
	r.conn = conn

	return nil
}
