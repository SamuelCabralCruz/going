package pkg

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model1"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model2"
	"github.com/samber/mo"
	"time"
)

var ConfigProvider1 = func() mo.Result[model1.Config] {
	fmt.Println("should be called every time a resolve is performed")
	return mo.Ok(model1.Config{
		A: "coucou1",
		B: time.Now().Nanosecond(),
		C: true,
	})
}

var ConfigProvider2 = func() mo.Result[model1.Config] {
	return mo.Ok(model1.Config{
		A: "coucou2",
		B: 8888,
		C: false,
	})
}

var ConfigProvider3 = func() mo.Result[model2.Config] {
	return mo.Ok(model2.Config{
		A: "coucou3",
		B: 5467,
		C: false,
	})
}

var ConfigProvider4 = func() mo.Result[model1.Config] {
	fmt.Println("singleton provider invoked for model1.Config - configProvider4")
	return mo.Ok(model1.Config{
		A: "coucou4",
		B: 0,
		C: true,
	})
}
