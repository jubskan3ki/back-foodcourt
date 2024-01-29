package handlers

import (
	"foodcourt/app/api/response"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

//----------------------------------------------------customer----------------------------------------------------------

func GetAllOpenRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetAllOpenRestaurant()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   "Failed to retrieve restaurants",
			},
		})
	}

	responseRestaurant := make([]response.GetAllRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetAllRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Open:        restaurant.Open,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}

//----------------------------------------------------seller------------------------------------------------------------

//----------------------------------------------------admin-------------------------------------------------------------

func GetAllRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetAllRestaurant()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   "Failed to retrieve restaurants",
			},
		})
	}

	responseRestaurant := make([]response.GetAllRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetAllRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Open:        restaurant.Open,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}
func GetDraftRestaurant(c fiber.Ctx, stores *stores.Store) error {
	restaurants, err := stores.GetDraftRestaurant()

	if err != nil {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"error":   err,
			},
		})
	}
	responseRestaurant := make([]response.GetDraftRestaurantResponseType, len(restaurants))
	for i, restaurant := range restaurants {
		responseRestaurant[i] = response.GetDraftRestaurantResponseType{
			Id:          restaurant.Id,
			Name:        restaurant.Name,
			Email:       restaurant.Email,
			Picture:     restaurant.Picture,
			Description: restaurant.Description,
			CategoryId:  restaurant.CategoryId,
			Draft:       restaurant.Draft,
		}
	}

	err = c.JSON(fiber.Map{
		"data": responseRestaurant,
	})
	return err

}

func PatchDraftRestaurant(c fiber.Ctx, stores *stores.Store, id int) error {
	err := stores.UpdateDraftRestaurant(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error": err,
			},
		})
	}
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"success": true,
		},
	})
}

func DeleteRestaurant(c fiber.Ctx, stores *stores.Store, id int) error {
	err := stores.DeleteRestaurant(id)

	if err != nil {
		err := c.JSON(fiber.Map{
			"data": fiber.Map{
				"success": false,
				"erreur":  err,
			},
		})
		return err
	}

	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"success": true,
		},
	})

	return err

}

func CreateNewRestaurantCategory(c fiber.Ctx, stores *stores.Store, item model.RestaurantCategoryItem) error {

	if item.Name == "" {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   "name is empty",
				"success": false,
			},
		})
	}

	verify, err := stores.GetOneCategoryByName(item.Name)

	if verify.Id != 0 {
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   "category already exist",
				"success": false,
			},
		})
	}

	res, err := stores.CreateCategory(item)

	if err != nil {
		err = c.JSON(fiber.Map{
			"data": fiber.Map{
				"error":   err,
				"success": false,
			},
		})
		return err
	}
	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"id": res,
		},
	})
	return nil
}