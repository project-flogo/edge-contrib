module main

go 1.12

require (
	github.com/project-flogo/contrib/activity/log v0.9.0
	github.com/project-flogo/contrib/trigger/timer v0.9.0
	github.com/project-flogo/core v0.9.4-0.20190829220729-31eb91f2b8a7
	github.com/project-flogo/edge-contrib/activity/mqtt v0.0.0
	github.com/project-flogo/edge-contrib/connection/mqtt v0.0.0
	github.com/project-flogo/flow v0.9.3
	github.com/stretchr/testify v1.4.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/project-flogo/edge-contrib/connection/mqtt => ../../../connection/mqtt

replace github.com/project-flogo/edge-contrib/activity/mqtt v0.0.0 => ../../../activity/mqtt
