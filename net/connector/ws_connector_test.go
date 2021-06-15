package cherryConnector

import (
	cherryFacade "github.com/cherry-game/cherry/facade"
	"testing"
)

// websocket client http://www.websocket-test.com/
func TestNewWSConnector(t *testing.T) {
	connector := NewWS(":9071")

	connector.OnConnect(func(conn cherryFacade.INetConn) {
		//s := cherrySession.NewSession(cherrySession.NextSID(), "", conn, nil)
		//cherryLogger.Infof("new session sid = %d, address = %s", s.SID(), s.INetConn().RemoteAddr().String())
		//
		//s.OnMessage(func(bytes []byte) (err error) {
		//	text := string(bytes)
		//	cherryLogger.Info(text)
		//
		//	if len(text) == 1 && text[0] == 99 {
		//		s.Closed()
		//	}
		//
		//	return nil
		//})
		//
		//s.Run()
	})

	connector.OnStart()
}