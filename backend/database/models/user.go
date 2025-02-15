package models

import (
	"time"

	"github.com/guregu/null"
)

type AuthMethod string

const (
	AuthMethodGoogle   AuthMethod = "google"
	AuthMethodGithub   AuthMethod = "github"
	AuthMethodPassword AuthMethod = "password"
)

type User struct {
	ID                       string      `json:"id" db:"id"`
	Email                    string      `json:"email" db:"email"`
	EmailVerified            bool        `json:"email_verified" db:"email_verified"`
	Name                     string      `json:"name" db:"given_name"`
	DisplayName              string      `json:"display_name" db:"display_name"`
	AvatarUrl                string      `json:"avatar_url" db:"avatar_url"`
	AuthenticationMethod     AuthMethod  `json:"authentication_method" db:"authentication_method"`
	Password                 null.String `json:"password" db:"password"`
	SubscriptionPlan         string      `json:"subscription_plan" db:"subscription_plan"`
	StripeCustomerId         null.String `json:"-" db:"stripe_customer_id"`
	StripeSubscriptionId     null.String `json:"-" db:"stripe_subscription_id"`
	StripeSubscriptionStatus null.String `json:"-" db:"stripe_subscription_status"`
	Metadata                 interface{} `json:"metadata"`
	CreatedAt                time.Time   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt                time.Time   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	DeletedAt                null.Time   `json:"deleted_at,omitempty" db:"deleted_at"`
}
