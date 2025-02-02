package character

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

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

	success := moveResp.StatusCode() == 200
	return &ActionResult{
		CooldownRemaining: time.Duration(moveResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           success,
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

	success := gatherResp.StatusCode() == 200
  if !success {

  }
	slog.Debug("Gather response", "StatusCode", gatherResp.StatusCode(), "GatherResult", gatherResp.JSON200.Data.Details, "Cooldown", gatherResp.JSON200.Data.Cooldown)
	return &ActionResult{
		CooldownRemaining: time.Duration(gatherResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           success,
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
