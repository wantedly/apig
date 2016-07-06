package main

const routerTmpl = `//{{ .Name }} API
		api.GET("/{{ pluralize (tolower .Name) }}", controllers.Get{{ pluralize .Name }})
		api.GET("/{{ pluralize (tolower .Name) }}/:id", controllers.Get{{ .Name }})
		api.POST("/{{ pluralize (tolower .Name) }}", controllers.Create{{ .Name }})
		api.PUT("/{{ pluralize (tolower .Name) }}/:id", controllers.Update{{ .Name }})
		api.DELETE("/{{ pluralize (tolower .Name) }}/:id", controllers.Delete{{ .Name }})
`
