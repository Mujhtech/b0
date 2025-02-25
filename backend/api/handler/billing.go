package handler

import (
	"net/http"

	"github.com/guregu/null"
	"github.com/mujhtech/b0/api/dto"
	"github.com/mujhtech/b0/api/middleware"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/request"
	"github.com/mujhtech/b0/internal/pkg/response"
)

func (h *Handler) GetUsage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	usage, err := h.store.AIUsageRepo.GetTotalUsage(ctx, store.TotalAIUsageFilter{
		OwnerID: session.User.ID,
		Range:   store.TotalAIUsageFilterRangeMonth,
	})

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	_ = response.Ok(w, r, "usage retrieved", usage)
}

func (h *Handler) UpgradePlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	session, ok := middleware.GetAuthSession(ctx)

	if !ok {
		_ = response.Unauthorized(w, r, nil)
		return
	}

	dst := new(dto.UpgradePlanDto)

	if err := request.ReadBody(r, dst); err != nil {
		_ = response.BadRequest(w, r, err)
		return
	}

	if session.User.SubscriptionPlan == dst.Plan {
		_ = response.BadRequest(w, r, nil)
		return
	}

	user := session.User

	if !session.User.StripeCustomerId.Valid || session.User.StripeCustomerId.String == "" {
		newStripeCustomer, err := h.billing.CreateCustomer(session.User.Email, session.User.Name)
		if err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}

		user.StripeCustomerId = null.NewString(newStripeCustomer.ID, true)

		if err := h.store.UserRepo.UpdateUser(ctx, user); err != nil {
			_ = response.InternalServerError(w, r, err)
			return
		}
	}

	subscription, err := h.billing.CreateCustomerSubscription(session.User.StripeCustomerId.String, dst.Plan)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	user.StripeSubscriptionId = null.NewString(subscription, true)
	user.EnablePayAsYouGo = true

	if err := h.store.UserRepo.UpdateUser(ctx, user); err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	portalLink, err := h.billing.Portal(session.User.StripeCustomerId.String)

	if err != nil {
		_ = response.InternalServerError(w, r, err)
		return
	}

	// if err := response.Redirect(w, r, portalLink, http.StatusTemporaryRedirect, true); err != nil {
	// 	log.Error().Err(err).Msg("failed to redirect")
	// 	return
	// }
	_ = response.Ok(w, r, "upgrade plan", map[string]string{
		"portal_link": portalLink,
	})
}
