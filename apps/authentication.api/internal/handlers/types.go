package handlers 

type SignupRequest struct {
    Email string  `json:"email"` 
    Password string `json:"password"`
    FirstName string `json:"firstName"`
    LastName string  `json:"lastName"`
}

func (sr *SignupRequest) Validate() (map[string]string) {
    errs := make(map[string]string)
    if sr.Email == "" {
        errs["email"] = "invalid email"
    }
    return errs
}

type SignupResponse struct {}

type Response[T any] struct {
    Status uint `json:"-"` 
    Message string `json:"message"`
    Errors map[string]string `json:"errors,omitempty"`
    Data *T  `json:"data"`
}


