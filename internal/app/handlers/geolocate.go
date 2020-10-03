package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/domgolonka/kashoo-geo/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// GeoLocate
// @Summary Gets the location data of an IPv4 or Ipv6.
// @Description Requests an IP address and returns the location data using the ipgeolocation.io API.
// @Tags geolocation
// @ID get-geo-location-from-ip
// @Produce json
// @Param ip path string true "IP Address"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /geolocate/{ip} [GET]
func GeoLocate(c echo.Context) error {

	ip := c.Param("ip")

	var ipresults map[string]interface{}

	ipaddr := strings.TrimSpace(ip)
	if net.ParseIP(ipaddr) == nil {
		return BadRequestError(c, fmt.Sprintf("IP is invalid: %s ", ipaddr))
	}

	u, err := url.Parse("https://api.ipgeolocation.io/ipgeo")

	if err != nil {
		return InternalError(c, fmt.Sprintf("Invalid GeoLocation URL: %s ", err))
	}

	token := os.Getenv("TOKEN")

	helpers.AddQ(u, "ip", ipaddr)

	helpers.AddQ(u, "apiKey", token)

	res, err := http.Get(u.String())
	if err != nil {
		return InternalError(c, fmt.Sprintf("Ipgeolocation get error: %s ", err))
	}

	switch res.StatusCode {
	case 400:
		return BadRequestError(c, fmt.Sprintf("Invalid API Key"))
	case 403:
		return BadRequestError(c, fmt.Sprintf("Invalid IP4 or Ip6 Address"))
	}

	err = json.NewDecoder(res.Body).Decode(&ipresults)
	if err != nil {
		return InternalError(c, fmt.Sprintf("IPGeoLocate error: %s ", err))
	}

	return c.JSON(200, ipresults)
}
