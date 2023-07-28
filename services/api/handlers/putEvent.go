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

func DeliveredEventHandler(c echo.Context) error {
	db := utils.GetDbConnection()
	domainName := c.Param("domainName")
	err := putEvent(db, domainName, false)
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK, "Success")
}

func BouncedEventHandler(c echo.Context) error {
	db := utils.GetDbConnection()
	domainName := c.Param("domainName")
	err := putEvent(db, domainName, true)
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK, "Success")
}

func putEvent(db *sqlx.DB, domainName string, bounced bool) error {
	domain := models.Domain{
		Name:    domainName,
		Events:  1,
		Bounced: bounced,
	}
	query := "INSERT INTO domains (name, events, bounced) VALUES (:name, :events, :bounced) ON CONFLICT (name) DO UPDATE SET events = domains.events+1, bounced = (CASE WHEN domains.bounced = TRUE THEN TRUE ELSE EXCLUDED.bounced END);"
	_, err := db.NamedExec(query, domain)
	return err
}
