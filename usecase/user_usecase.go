package usecase

import (
	"context"
	"encoding/hex"
	"errors"
	"test_tracker_backend/domain"
	"test_tracker_backend/internal/tokkenutil"
	"test_tracker_backend/internal/mailutil"
	"time"
	"crypto/rand"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
	jwtSecret      string
	jwtExpiryHours int
}

func NewUserUsecase(repo domain.UserRepository, timeout time.Duration, secret string, expiry int) domain.UserUsecase {
	return &userUsecase{
		userRepo:       repo,
		contextTimeout: timeout,
		jwtSecret:      secret,
		jwtExpiryHours: expiry,
	}
}

func (u *userUsecase) Register(ctx context.Context , req *domain.RegisterRequest)error{
	ctx, cancel := context.WithTimeout(ctx ,u.contextTimeout)
	defer cancel()

	if req.Email == "" || req.Name == "" || req.Password == "" {
		return errors.New("All the fields are mandatory")
	}

	existingUser, _ := u.userRepo.GetByEmail(ctx , req.Email)
	if existingUser != nil{
		return errors.New("User with same email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password),12)
	if err != nil {
		return errors.New("security system failed to process credentials")
	}
	
	user := &domain.User{
		ID: uuid.New().String(),
		Email: req.Email,
		PasswordHash: string(hashedPassword),
		Name: req.Name,
	}

	return u.userRepo.Create(ctx, user)

}

func (u *userUsecase) Login (ctx context.Context ,req  *domain.LoginRequest)(string ,error){
	ctx,cancel := context.WithTimeout(ctx,u.contextTimeout)
	defer cancel()

	if req.Email == "" || req.Password == ""{
		return "", errors.New("Required feilds are missing")
	}

	user,err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil{
		return "", errors.New("Invalid Email or password configuration")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(req.Password))
	if err != nil{
		return "", errors.New("Invalid email or password configurations")
	}

	token, err := tokkenutil.CreateAccessToken(user ,u.jwtSecret, u.jwtExpiryHours)
	if err != nil {
		return "", errors.New("Token authentication engine configuration breakdown")
	}

	return token, nil
}

func (u *userUsecase) ResetPassword (ctx context.Context, req *domain.ResetPasswordRequest) error{
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("Identity mapping lookup failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),[]byte(req.OldPassword))
	if err != nil {
		return errors.New("Current credentials varification failed")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword),12)
	if err != nil{
		return errors.New("CryptoGraphic mutation process failure") // a statement of error yet interesting thing to read ,read later 
	}

	return u.userRepo.UpdatePassword(ctx, user.ID, string(newHashedPassword))

}
func generateSecureToken (length int)(string, error){
	bytes := make([]byte,length)
	if _,error := rand.Read(bytes); error != nil {
		return "",error
	}
	return hex.EncodeToString(bytes) , nil
}



func (u *userUsecase) ForgotPassword (ctx context.Context , email string) error{
	ctx , cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user,err := u.userRepo.GetByEmail(ctx, email)
	if err != nil{
		// return errors.New("No user registered with this id")
		return nil
	}

	token, err := generateSecureToken(12)
	if err != nil {
		return errors.New("Couldn't generate token")
	}

	expiry := time.Now().Add(15 * time.Minute)

	err = u.userRepo.UpdateResetToken(ctx,user.ID,token,expiry)
	if err != nil {
		return errors.New("Couldn't complete the verification")
	}

	mailConfig :=mailutil.EmailConfig{
		SMTPHost: "smtp.gmail.com",
		SMTPPort: "587",
		Sender:  "sdejew",
		Password: "something",
	}

	return mailutil.SendResetTokenEmail(&mailConfig,token ,user.Email)
}

func (u *userUsecase) ConfirmPasswordReset(ctx context.Context, req *domain.ConfirmPasswordResetRequest) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("invalid verification context parameters")
	}

	if user.ResetToken == nil || *user.ResetToken != req.RandomToken {
		return errors.New("invalid or incorrect verification token sequence")
	}

	if user.ResetTokenExpiry == nil || time.Now().After(*user.ResetTokenExpiry) {
		return errors.New("verification token lifecycle has expired")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 12)
	if err != nil {
		return errors.New("cryptographic pipeline failure")
	}

	err = u.userRepo.UpdatePassword(ctx, user.ID, string(newHashedPassword))
	if err != nil {
		return errors.New("failed to save updated credentials")
	}

	return u.userRepo.ClearResetToken(ctx, user.ID)
}

func (u *userUsecase) UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return errors.New("security cryptographic pipeline failure")
	}

	return u.userRepo.UpdatePassword(ctx, userID, string(hashedPassword))
}
