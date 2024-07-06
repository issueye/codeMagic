package main

import "context"

// App struct
type BaseApp struct {
	ctx context.Context
}

// startup is called at application startup
func (a *BaseApp) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *BaseApp) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *BaseApp) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *BaseApp) shutdown(ctx context.Context) {
	// Perform your teardown here
}
