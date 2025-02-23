package characters

import (
	"context"
	"errors"
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

func (service *CharacterService) Move(character *CharacterWrapper, x int, y int) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	moveResp, err := service.client.ActionMoveMyNameActionMovePostWithResponse(
		ctx,
		character.Name,
		artifactsmmo.DestinationSchema{X: x, Y: y},
	)
	if err != nil {
		service.logger.Error("Could not move character", "error", err)
		return nil, err
	}

	if moveResp.StatusCode() == 499 {
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
		return &ActionResult{
			CooldownRemaining: 0 * time.Second,
			Success:           true,
		}, nil
	}

	if moveResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(moveResp.Body)}
	}

	*character = *FromSchema(moveResp.JSON200.Data.Character)

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

	return &ActionResult{
		CooldownRemaining: time.Duration(moveResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Fight(character *CharacterWrapper) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fightResp, err := service.client.ActionFightMyNameActionFightPostWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		service.logger.Error("Could not fight", "error", err)
		return nil, err
	}

	if fightResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(fightResp.Body)}
	}

	*character = *FromSchema(fightResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(fightResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Gather(character *CharacterWrapper) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gatherResp, err := service.client.ActionGatheringMyNameActionGatheringPostWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		service.logger.Error("Could not gather", "error", err)
		return nil, err
	}

	if gatherResp.StatusCode() == 499 {
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Gather(character)
	}

	if gatherResp.StatusCode() != 200 {
		return nil, &ActionError{*character.CooldownExpiration, string(gatherResp.Body)}
	}

	*character = *FromSchema(gatherResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(gatherResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Craft(character *CharacterWrapper, itemCode string, quantity int) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	craftResp, err := service.client.ActionCraftingMyNameActionCraftingPostWithResponse(
		ctx,
		character.Name,
		*&artifactsmmo.CraftingSchema{
			Code:     itemCode,
			Quantity: &quantity,
		},
	)
	if err != nil {
		service.logger.Error("Could not craft", "error", err)
		return nil, err
	}

	if craftResp.StatusCode() == 499 {
		err = service.WaitCooldown(character)
		if err != nil {
			return nil, err
		}
		return service.Craft(character, itemCode, quantity)
	}

	service.logger.Debug("Craft response", "status", craftResp.StatusCode(), "body", string(craftResp.Body))

	*character = *FromSchema(craftResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(craftResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Rest(character *CharacterWrapper) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	restResp, err := service.client.ActionRestMyNameActionRestPostWithResponse(ctx, character.Name)
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

	*character = *FromSchema(restResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(restResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}

func (service *CharacterService) Use(character *CharacterWrapper, item *artifactsmmo.SimpleItemSchema) (*ActionResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

  service.logger.Info("Using item", "request", *item)
	useResp, err := service.client.ActionUseItemMyNameActionUsePostWithResponse(ctx, character.Name, *item)
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

	*character = *FromSchema(useResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(useResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}
