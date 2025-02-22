package characters

import (
	"context"
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

	character = FromSchema(moveResp.JSON200.Data.Character)

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

	character= FromSchema(fightResp.JSON200.Data.Character)

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

	character = FromSchema(gatherResp.JSON200.Data.Character)

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

	service.logger.Debug("Updating character", "character", craftResp.JSON200.Data.Character)
	character = FromSchema(craftResp.JSON200.Data.Character)

	return &ActionResult{
		CooldownRemaining: time.Duration(craftResp.JSON200.Data.Cooldown.RemainingSeconds) * time.Second,
		Success:           true,
	}, nil
}
