package email

type Email struct {
    To          string
    Subject     string
    HTMLBody    string
    TextBody    string
    Tag         string // optional analytics tag
    TrackOpens  bool
    Attachments []Attachment // optional
}

type Attachment struct {
    Name        string
    Content     []byte
    ContentType string
}
