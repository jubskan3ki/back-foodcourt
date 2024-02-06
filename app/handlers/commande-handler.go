package handlers

import (
	"fmt"
	"foodcourt/app/model"
	"foodcourt/app/stores"

	"github.com/gofiber/fiber/v3"
)

// creer une commande

func CreateCommande(c fiber.Ctx , commandeStore *stores.Store , commande model.CommandeItem) error {
	_, err := commandeStore.CreateCommande(commande)
	fmt.Println("ceci est l'erreur", err)

	if err != nil {
      err = c.JSON(fiber.Map{
		"data" : fiber.Map{
			"error" : err,
			"succes" : false,
		},
	  })

	  return err
	}

	err = c.JSON(fiber.Map{
		"data": fiber.Map{
			"succes": true,
		},
	})

	return err

}

// GET commande by id
func GetCommandeById(c fiber.Ctx, CommandeStore *stores.Store, id int) error {

	res, err := CommandeStore.GetCommandeById(id)

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
			"succes": res,
		},
	})
	return err
}

// Get commandes by restaurantId

func GetAllCommandeByRestaurantId(c fiber.Ctx, CommandeStore *stores.Store, id int) error {

	res, err := CommandeStore.GetAllCommandeByRestaurantId(id)
	fmt.Println("ceci est l'erreur" , err)

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
			"succes": res,
		},
	})
	return err
}

// Update Commande 
func UpdateCommande(c fiber.Ctx, CommandeStore *stores.Store, id int , state int) error {

	_, err := CommandeStore.UpdateCommande(id , state)

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
			"succes": true,
		},
	})
	return err
}

