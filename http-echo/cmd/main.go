package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"wax/wax-samples/http-echo/internal"
	"wax/wax-samples/http-echo/views"

	"github.com/michal-laskowski/wax"
	wax_echo "github.com/michal-laskowski/wax-libs/echo"
	"github.com/michal-laskowski/wax-libs/gots"
	"github.com/michal-laskowski/wax-libs/livereload"
)

var isDevRun = flag.Bool("isDev", false, `Use this flag for develop run.
    * If true - WAX we will load views from os.Fs. You can change views without restarting the application.
    Using livereload the idok will be automatically reloaded.
    * When false embedded views will be used`)

func main() {
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())

	appViewUtils := AppViewUtils{
		IsDev: *isDevRun,
	}

	// Setup WAX renderer for echo

	if *isDevRun {
		fmt.Printf("In DEV run")

		out, err := os.Create("./views/model.generated.d.ts")
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// Generate type definition for views
		gots.GenerateTypeDefinition(out, "", "", internal.PageModel{}, appViewUtils)

		// when you  - we will use FS view resolver
		e.Renderer = wax_echo.NewWaxEchoRenderer(wax.NewFsViewResolver(os.DirFS("./views/")), wax.WithGlobalObject("viewUtils", appViewUtils))

		// Use livereload
		go livereload.StartLiveReload(livereload.LiveReloadConfig{})
	} else {
		fmt.Printf("WAX will use embedded views. If you want load them from FS use -isDev flag")
		e.Renderer = wax_echo.NewWaxEchoRenderer(wax.NewFsViewResolver(views.EmbeddedViews), wax.WithGlobalObject("viewUtils", appViewUtils))
	}

	serverStartedOn := time.Now().UTC().Format(time.RFC3339Nano)
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "HelloPage", internal.PageModel{
			ServerStartedOn:  serverStartedOn,
			ServerRenderedOn: time.Now().UTC().Format(time.RFC3339Nano),
			SomeString:       c.QueryString(),
			ServerDate:       time.Now().GoString(),
			Additional: internal.SomeOtherInfo{
				IntValue: 445566,
			},
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}

type AppViewUtils struct {
	IsDev bool
}
