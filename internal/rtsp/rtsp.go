package rtsp

import "github.com/Lyyyttooon/zeusmedia/utils"

type RTSP struct {
	rawURL string
	urlCtx utils.URLContext
}

// NewPullSession 新建
func NewPullSession(url string) RTSP {
	return RTSP{
		rawURL: url,
	}
}

func (r *RTSP) Conn() {
	errChan := make(chan error, 1)

	go func() {
		if err := r.connect(); err != nil {
			errChan <- err
		}
	}()
}

func (r *RTSP) connect() (err error) {
	r.urlCtx, err = utils.ParseRTSPURL(r.rawURL)
	if err != nil {
		return err
	}

	return nil
}
