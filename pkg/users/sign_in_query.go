package users

type SignInQuery struct {
	ReturnURL string `uri:"returnUrl,omitempty" url:"returnUrl,omitempty"`
}
