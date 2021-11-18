package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/danic95/things-service/internal/api/v1/dto"
	"github.com/danic95/things-service/internal/things/entity"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	"github.com/sanservices/apilogger/v2"
)

var ErrDataValidation error = errors.New("invalid data in thing body")

func (h *Handler) getThings(c echo.Context) error {
	/*queryThingID := c.Param("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil && queryThingID != "" {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}*/

	things, err := h.service.GetThings(c.Request().Context())

	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError,
			errors.New("something went wrong while requesting the things"))
	}
	return helper.RespondOk(c, things)
}

func (h *Handler) getThing(c echo.Context) error {
	queryThingID := c.Param("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil && queryThingID != "" {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}

	err = nil
	var thing interface{}
	thing, err = h.service.GetThing(c.Request().Context(), uint(thingID))
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError,
			errors.New("something went wrong while requesting the thing"))
	}
	return helper.RespondOk(c, thing)
}

func (h *Handler) createThing(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.CreateThing{}

	// Decoding the request
	err := helper.DecodeBody(c, &params.Body)
	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatTypeConv, err)
		return helper.RespondError(c, http.StatusBadRequest, err)
	}

	// validating the request
	err = h.validator.Struct(params.Body)
	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatInputValidation, err)
		return helper.RespondError(c, http.StatusBadRequest, ErrDataValidation)
	}

	// Creating the thing
	t, _ := time.Parse("2006-01-02", params.Body.Birth_date)
	thing := &entity.Thing{
		Name:       params.Body.Name,
		Email:      params.Body.Email,
		Birth_date: t,
		Phone:      params.Body.Phone,
		Start_day:  params.Body.Start_day}

	err = h.service.CreateThing(ctx, thing)

	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatServiceOutput, err)
		return helper.RespondError(c, http.StatusBadRequest, err)
	}

	return helper.RespondOk(c, thing)
}
