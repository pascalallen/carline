package security_token

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/carline/internal/carline/domain/crypto"
	"github.com/pascalallen/carline/internal/carline/domain/user"
	"time"
)

type Service struct {
	SecurityTokenRepository Repository
}

func NewService(repository Repository) Service {
	return Service{repository}
}

func (s *Service) FetchById(id ulid.ULID) (*SecurityToken, error) {
	return s.SecurityTokenRepository.GetById(id)
}

func (s *Service) FetchToken(crypto crypto.Crypto) (*SecurityToken, error) {
	return s.SecurityTokenRepository.GetByCrypto(crypto)
}

func (s *Service) FetchTokensForUser(user user.User) (*[]SecurityToken, error) {
	return s.SecurityTokenRepository.GetAllForUser(user)
}

func (s *Service) AddToken(token *SecurityToken) error {
	return s.SecurityTokenRepository.Add(token)
}

func (s *Service) GenerateActivationToken(user user.User, expiresAt time.Time) (*SecurityToken, error) {
	securityToken := GenerateActivation(ulid.Make(), user.Id, expiresAt)
	err := s.SecurityTokenRepository.Add(securityToken)
	if err != nil {
		return nil, err
	}

	return securityToken, nil
}

func (s *Service) GenerateRefreshToken(user user.User, expiresAt time.Time) (*SecurityToken, error) {
	securityToken := GenerateRefresh(ulid.Make(), user.Id, expiresAt)
	err := s.SecurityTokenRepository.Add(securityToken)
	if err != nil {
		return nil, err
	}

	return securityToken, nil
}

func (s *Service) GenerateResetToken(user user.User, expiresAt time.Time) (*SecurityToken, error) {
	securityToken := GenerateReset(ulid.Make(), user.Id, expiresAt)
	err := s.SecurityTokenRepository.Add(securityToken)
	if err != nil {
		return nil, err
	}

	return securityToken, nil
}

func (s *Service) ClearTokensForUser(user user.User) error {
	return s.SecurityTokenRepository.ClearAllForUser(user)
}
