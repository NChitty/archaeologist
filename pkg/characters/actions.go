package characters

import (
	"context"
	"errors"
	"time"

	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ActionError struct {
	CooldownExpiration time.Time
	errorMessage       string
}

func (actionError *ActionError) Error() string {
	return actionError.errorMessage
}

func (service *CharacterService) Move(character *character.Character, x int, y int) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	moveResp, err := service.client.ActionMoveMyNameActionMovePostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.DestinationSchema{X: x, Y: y},
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		service.logger.Error("Could not move character", "error", err)
		return nil, err
	}

	if moveResp.StatusCode() == 499 {
		cancel()
		err = service.WaitCooldown(character)
		service.logger.Debug(
			"Waiting for cooldown",
			"expiration", character.CooldownExpiration,
			"timeRemaining", character.CooldownExpiration.Sub(time.Now().UTC()),
		)
		if err != nil {
			return nil, err
		}
		return service.Move(character, x, y)
	}
	if moveResp.StatusCode() == 490 {
		return &actors.ActionResult{
			CooldownRemaining: 0 * time.Second,
			Success:           true,
		}, nil
	}

	if moveResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(moveResp.Body)}
	}

	character.FromSchema(moveResp.JSON200.Data.Character)

	service.logger.Info(
		"Completed move",
		"cooldownExpiration",
		character.CooldownExpiration.Local(),
		"cooldown",
		time.Now().UTC().Sub(*character.CooldownExpiration),
		"location",
		struct {
			X int
			Y int
		}{character.X, character.Y},
	)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(moveResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Fight(character *character.Character) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fightResp, err := service.client.ActionFightMyNameActionFightPostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		service.logger.Error("Could not fight", "error", err)
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if fightResp.StatusCode() == 499 {
		cancel()
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Fight(character)
	}
	if fightResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(fightResp.Body)}
	}

  service.logger.Info("Finished fight", "fightResult", fightResp.JSON200.Data.Fight)
	character.FromSchema(fightResp.JSON200.Data.Character)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(fightResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Gather(character *character.Character) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gatherResp, err := service.client.ActionGatheringMyNameActionGatheringPostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		service.logger.Error("Could not gather", "error", err)
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if gatherResp.StatusCode() == 499 {
		cancel()
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Gather(character)
	}
	if gatherResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(gatherResp.Body)}
	}

	character.FromSchema(gatherResp.JSON200.Data.Character)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(gatherResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Craft(character *character.Character, itemCode string, quantity int) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	craftResp, err := service.client.ActionCraftingMyNameActionCraftingPostWithResponse(
		ctx,
		character.Name,
		*&artifactsmmo.CraftingSchema{
			Code:     itemCode,
			Quantity: &quantity,
		},
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		service.logger.Error("Could not craft", "error", err)
		return nil, err
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if craftResp.StatusCode() == 499 {
		cancel()
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Craft(character, itemCode, quantity)
	}

	if craftResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(craftResp.Body)}
	}

	character.FromSchema(craftResp.JSON200.Data.Character)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(craftResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Rest(character *character.Character) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	restResp, err := service.client.ActionRestMyNameActionRestPostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		return nil, err
	}
	if restResp.StatusCode() == 499 {
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Rest(character)
	}

	if restResp.StatusCode() != 200 {
		service.logger.Error("Unexpected status code", "statusCode", restResp.StatusCode(), "body", string(restResp.Body))
		return nil, errors.New(string(restResp.Body))
	}

	character.FromSchema(restResp.JSON200.Data.Character)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(restResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Use(character *character.Character, item *artifactsmmo.SimpleItemSchema) (*actors.ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	service.logger.Info("Using item", "request", *item)
	useResp, err := service.client.ActionUseItemMyNameActionUsePostWithResponse(
		ctx,
		character.Name,
		*item,
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		return nil, err
	}
	if useResp.StatusCode() == 499 {
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Use(character, item)
	}

	if useResp.StatusCode() != 200 {
		service.logger.Error("Unexpected status code", "statusCode", useResp.StatusCode(), "body", string(useResp.Body))
		return nil, errors.New(string(useResp.Body))
	}

	character.FromSchema(useResp.JSON200.Data.Character)

	return &actors.ActionResult{
		CooldownRemaining: time.Duration(useResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}
