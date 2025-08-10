package auth

type jwtService struct {
	secretKey string
}

// NewJwtService creates a new JwtService with the given secret key
func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey: secretKey,
	}
}

type JwtService interface {
	GenerateToken() (string, error)
	DecodeToken(token string) (string, error)
}

func (s *jwtService) GenerateToken() (string, error) {
	return "", nil
}

func (s *jwtService) DecodeToken(token string) (string, error) {
	return "", nil
}
