package handlers

import (
	"log"
	"net/http"

	"go-scaled-api/services/api/models"
	"go-scaled-api/services/api/utils"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func GetDomainHandler(c echo.Context) error {
	db := utils.GetDbConnection()
	if db == nil {
		log.Fatal("database connection is nil")
	}
	domainName := c.Param("domainName")
	response, err := getDomain(db, domainName)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, response)
}

func getDomain(db *sqlx.DB, domainName string) (models.DomainResponse, error) {
	var domain models.Domain
	var isCatchall string
	var response models.DomainResponse

	err := db.QueryRowx("SELECT * FROM domains WHERE name=$1", domainName).StructScan(&domain)
	if err != nil {
		return response, err
	}
	if domain.Bounced == true {
		isCatchall = "not catch-all"
	} else {
		if domain.Events >= 1000 {
			isCatchall = "catch-all"
		} else {
			isCatchall = "unknown"
		}
	}

	response.Name = domain.Name
	response.IsCatchall = isCatchall
	response.Events = domain.Events
	response.Bounced = domain.Bounced

	return response, nil
}
