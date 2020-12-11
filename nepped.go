// 2020 © GoldenXelenium

package gonepnep

import (
    "encoding/json"
    "fmt"
    "strconv"
)

// Endpoints
const (
    ImageAPI string = "https://neppedcord.top/api/images/"
    SHARPAPI string = "https://neppedcord.top/api/sharp/"
)

// IsValidCategory category validation
func IsValidCategory(category string) bool {
    switch category {
    case
        "baka",
        "cry",
        "cuddle",
        "happy",
        "hug",
        "kiss",
        "sad",
        "wag",
        "pat",
        "poke",
        "dance",
        "smug":
        return true
    }
    return false
}

// Request to the Images API
func (s NeppedSession) GetImages(category string) (*NeppedResponse, error) {
    var resp NeppedResponse

    if !IsValidCategory(category) {
        err := fmt.Errorf("[goNep] category \"%s\" is not valid", category)
        return nil, err
    }

    req := request{
        url:   ImageAPI + category,
        token: s.Token,
    }

    body, err := req.getResponse()
    if err != nil {
        return nil, err
    }

    json.Unmarshal(body, &resp)

    return &resp, nil
}

// SHARP API | Get info about given user
func (ss SHARPSession) Check(ID int64) (*SHARPResponse, error) {
    var resp SHARPResponse

    req := request{
        url:   SHARPAPI + strconv.FormatInt(ID, 10),
        token: ss.session.Token,
    }

    body, err := req.getResponse()
    if err != nil {
        return nil, err
    }

    json.Unmarshal(body, &resp)

    return &resp, nil
}

// SHARP API | Ban given user
func (ss SHARPSession) Ban(ID int64) (*SHARPResponse, error) {
    var resp SHARPResponse

    req := request{
        url:   SHARPAPI + strconv.FormatInt(ID, 10) + "/ban",
        token: ss.session.Token,
    }

    body, err := req.getResponse()
    if err != nil {
        return nil, err
    }

    json.Unmarshal(body, &resp)

    return &resp, nil
}

// SHARP API | Unban given user
func (ss SHARPSession) Unban(ID int64) (*SHARPResponse, error) {
    var resp SHARPResponse

    req := request{
        url:   SHARPAPI + strconv.FormatInt(ID, 10) + "/unban",
        token: ss.session.Token,
    }

    body, err := req.getResponse()
    if err != nil {
        return nil, err
    }

    json.Unmarshal(body, &resp)

    return &resp, nil
}
