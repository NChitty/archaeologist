package character

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ActionError struct {
	CooldownExpiration time.Time
	errorMessage       string
}

func (actionError *ActionError) Error() string {
	return actionError.errorMessage
}

type ActionResult struct {
	CooldownRemaining time.Duration
	Success           bool
}

func (character *Character) Move(x int, y int) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	moveResp, err := character.client.ActionMoveMyNameActionMovePostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.DestinationSchema{X: x, Y: y},
	)
	if err != nil {
		slog.Error("Could not move character:", err)
		return nil, err
	}

	if moveResp.StatusCode() == 499 {
		cooldownError := make(chan error)
		go character.WaitCooldown(cooldownError)
		slog.Debug(
			"Waiting for cooldown",
			"expiration", character.Character.CooldownExpiration,
			"timeRemaining", character.Character.CooldownExpiration.Sub(time.Now()),
		)
		cooldownError <- err
		if err != nil {
			return nil, err
		}
		return character.Move(x, y)
	}

	if moveResp.StatusCode() != 200 {
		return nil, &ActionError{*character.Character.CooldownExpiration, string(moveResp.Body)}
	}

	character.Character = &moveResp.JSON200.Data.Character

	return &ActionResult{
		CooldownRemaining: time.Duration(moveResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (character *Character) Fight() (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fightResp, err := character.client.ActionFightMyNameActionFightPostWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		slog.Error("Could not fight:", err)
		return nil, err
	}

	success := fightResp.StatusCode() == 200
	return &ActionResult{
		CooldownRemaining: time.Duration(fightResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           success,
	}, nil
}

func (character *Character) Gather() (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gatherResp, err := character.client.ActionGatheringMyNameActionGatheringPostWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		slog.Error("Could not gather:", err)
		return nil, err
	}

	if gatherResp.StatusCode() == 499 {
		cooldownError := make(chan error)
		go character.WaitCooldown(cooldownError)
		cooldownError <- err
		if err != nil {
			return nil, err
		}
		return character.Gather()
	}


	if gatherResp.StatusCode() != 200 {
		return nil, &ActionError{*character.Character.CooldownExpiration, string(gatherResp.Body)}
	}

	character.Character = &gatherResp.JSON200.Data.Character

	return &ActionResult{
		CooldownRemaining: time.Duration(gatherResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (character *Character) Rest() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, _ = character.client.ActionFightMyNameActionFightPostWithResponse(
		ctx,
		character.Name,
	)
}
