module main

go 1.24.2

require server v0.0.0

require service v0.0.0

require morse v0.0.0

require handlers v0.0.0 // indirect

replace server => ../internal/server

replace handlers => ../internal/handlers

replace service => ../internal/service

replace morse => ../pkg/morse
