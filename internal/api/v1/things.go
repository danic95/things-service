package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/danic95/things-service/internal/api/v1/dto"
	"github.com/danic95/things-service/internal/things/entity"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/helper"
	"github.com/sanservices/apilogger/v2"
)

func (h *Handler) getThings(c echo.Context) error {
	/*queryThingID := c.Param("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil && queryThingID != "" {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}*/

	thing, err := h.service.GetThings(c.Request().Context())

	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, errors.New("something went wrong"))
	}
	return helper.RespondOk(c, thing)
}

func (h *Handler) getThing(c echo.Context) error {
	queryThingID := c.Param("id")
	thingID, err := strconv.Atoi(queryThingID)
	if err != nil && queryThingID != "" {
		return helper.RespondError(c, http.StatusBadRequest, errors.New("id query param must be unsigned integer"))
	}

	err = nil
	var thing interface{}
	if queryThingID == "" {
		thing, err = h.service.GetThings(c.Request().Context())
	} else {
		thing, err = h.service.GetThing(c.Request().Context(), uint(thingID))
	}
	if err != nil {
		return helper.RespondError(c, http.StatusInternalServerError, errors.New("something went wrong"))
	}
	return helper.RespondOk(c, thing)
}

func (h *Handler) createThing(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.CreateThing{}
	err := helper.DecodeBody(c, &params.Body)

	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatTypeConv, err)
		return helper.RespondError(c, http.StatusBadRequest, err)
	}

	name := &entity.Thing{Name: params.Body.Name}
	err = h.service.CreateThing(ctx, name)

	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatServiceOutput, err)
		return helper.RespondError(c, http.StatusBadRequest, err)
	}

	return helper.RespondOk(c, name)
}
