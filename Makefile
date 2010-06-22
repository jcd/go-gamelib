include $(GOROOT)/src/Make.$(GOARCH)

TARG=main
GOFILES= \
	main.go \
	cmdline.go \
	video.go \
	draw.go \
	draw_font.go \
	cmd/rope-01.go \
	cmd/sim-grav.go \
	cmd/chain-reaction.go \

DIRS= \
	common \
	sched \
	phys \

DEPS=common sched phys

include $(GOROOT)/src/Make.cmd

