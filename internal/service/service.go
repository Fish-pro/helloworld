package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

// DemoSet is demo service providers
var DemoSet = wire.NewSet(NewDemoService)
