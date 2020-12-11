// 2020 © GoldenXelenium

package gonepnep

// NeppedSession processes all requests to NeppedAPI
type NeppedSession struct {
    Token string
}

// SHARPSession processes all requests to SHARP API
type SHARPSession struct {
    session *NeppedSession
}

// SHARPResponse typical response from SHARP API
type SHARPResponse map[string]interface{}

// NeppedResponse typical response from NeppedAPI
type NeppedResponse struct {
    URL string
}

// NewSession initializes the NeppedAPI session
func NewSession(token string) *NeppedSession {
    return &NeppedSession{
        Token: token,
    }
}

// NewSession initializes the SHARP API session over the NeppedAPI
func NewSHARPSession(session NeppedSession) *SHARPSession {
    return &SHARPSession{
        session: &session,
    }
}
