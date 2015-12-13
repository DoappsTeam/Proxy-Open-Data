package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffail/gabs"
	"github.com/parnurzeal/gorequest"
)

// Environment constants
const (
	API_KEY = "4b5bb9b55452e67464c2bef2b5bb92c47661f4da"
	PATH    = "http://api.datosabiertos.msi.gob.pe/datastreams/invoke/"
)

func makeRequest(full_path string) string {
	request := gorequest.New()
	_, body, errGet := request.Get(full_path).End()

	if errGet != nil {
		return `{"error": "has produced"}`
	}

	jsonParsed, _ := gabs.ParseJSON([]byte(body))
	main_node, errParse := jsonParsed.S("result", "fArray").Children()
	if errParse != nil {
		panic(errParse)
	}

	fRows, _ := jsonParsed.S("result", "fRows").Data().(float64)
	fCols, _ := jsonParsed.S("result", "fCols").Data().(float64)

	for i, child := range main_node {
		node := child.Data().(map[string]interface{})

		fmt.Println("------------------->", i, node["fStr"])
	}

	return body
}

func getPayload(GUID, api_key string) func(int) gin.H {
	return func(status int) gin.H {
		var full_path string

		if status >= 400 {
			full_path = "error"
		} else {
			full_path = PATH + GUID + "?auth_key=" + api_key
		}

		content := makeRequest(full_path)
		if content == `{"error": "has produced"}` {
			status = 400
		}

		return gin.H{"status": status, "content": content}
	}
}

func pathBuilder(api_key string) func(string) gin.H {
	namespace_guid := map[string]string{
		"actividades_culturales":     "ACTIV-CULTU-86315",
		"actividades_discapacitados": "ACTIV-ATENC-A-LA-PERSO",
		"actividades_gratuitas":      "ACTIV-GRATU",
		"campana_veterinaria":        "CAMPA-VETER",
		"chapa_bici":                 "CHAPA-TU-BICI-CICLO-EXIST",
		"intervenciones_serenazgo":   "CONSO-INTER-SEREN-31211",
		"jornada_salud":              "JORNA-Y-CAMPA-SALUD",
		"presupuesto_gastos":         "PRESU-GASTO-2015",
		"presupuesto_ingresos":       "PRESU-INGRE-2015-57742",
		"programa_recicla":           "HORAR-Y-RUTAS-RECOL-RESID",
	}

	return func(namespace string) gin.H {
		GUID := namespace_guid[namespace]
		payload := getPayload(GUID, api_key)

		var result gin.H

		if GUID != "" {
			result = payload(200)
		} else {
			result = payload(404)
		}

		return result
	}
}

func DataHandler(c *gin.Context) {
	namespace := c.Params.ByName("namespace")
	pb := pathBuilder(API_KEY)

	result := pb(namespace)

	c.JSON(200, result)
}

func main() {
	// API Schema
	r := gin.Default()
	base := r.Group("api")
	{
		base.GET("/data/:namespace", DataHandler)
	}
	r.Run(":8080")
}
