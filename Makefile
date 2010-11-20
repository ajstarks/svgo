include $(GOROOT)/src/Make.inc

ifeq ($(GOOS), darwin)
OPENER=open -a /Applications/Safari.app
else
OPENER=firefox
endif

SVGLIB=svg

CLIENTS=svgdef flower funnel imfade lewitt planets randcomp\
		richter rl vismem android gradient bubtrail svgopher pmap paths

TCLIENTS=svgdef flower funnel imfade lewitt planets randcomp\
		richter rl vismem android gradient bubtrail paths svgopher

all:	$(CLIENTS)

$(SVGLIB).$(O):	$(SVGLIB).go
	$(GC) $(SVGLIB).go
	
svgdef:	svgdef.go svg.$(O)
	$(GC) -I. svgdef.go
	$(LD) -L. -o svgdef svgdef.$(O)

flower:	flower.go svg.$(O)
	$(GC) -I. flower.go
	$(LD) -L. -o flower flower.$(O)
	
funnel:	funnel.go svg.$(O)
	$(GC) -I. funnel.go
	$(LD) -L. -o funnel funnel.$(O)

imfade:	imfade.go svg.$(O)
	$(GC) -I. imfade.go
	$(LD) -L. -o imfade imfade.$(O)
	
lewitt:	lewitt.go svg.$(O)
	$(GC) -I. lewitt.go
	$(LD) -L. -o lewitt lewitt.$(O)
	
planets:	planets.go svg.$(O)
	$(GC) -I. planets.go
	$(LD) -L. -o planets planets.$(O)

randcomp:	randcomp.go svg.$(O)
	$(GC) -I. randcomp.go
	$(LD) -L. -o randcomp randcomp.$(O)
	
richter:	richter.go svg.$(O)
	$(GC) -I. richter.go
	$(LD) -L. -o richter richter.$(O)
	
rl:	rl.go svg.$(O)
	$(GC) -I. rl.go
	$(LD) -L. -o rl rl.$(O)
	
vismem:	vismem.go svg.$(O)
	$(GC) -I. vismem.go
	$(LD) -L. -o vismem vismem.$(O)

android:	android.go svg.$(O)
	$(GC) -I. android.go
	$(LD) -L. -o android android.$(O)

gradient:	gradient.go svg.$(O)
	$(GC) -I. gradient.go
	$(LD) -L. -o gradient gradient.$(O)
	
bubtrail:	bubtrail.go svg.$(O)
	$(GC) -I. bubtrail.go
	$(LD) -L. -o bubtrail bubtrail.$(O)

svgopher:	svgopher.go svg.$(O)
	$(GC) -I. svgopher.go
	$(LD) -L. -o svgopher svgopher.$(O)

pmap:	pmap.go svg.$(O)
	$(GC) -I. pmap.go
	$(LD) -L. -o pmap pmap.$(O)

paths:	paths.go svg.$(O)
	$(GC) -I. paths.go
	$(LD) -L. -o paths paths.$(O)
	
test:	$(TCLIENTS)
	for c in $(TCLIENTS); do ./$$c > $$c.svg; $(OPENER) $$c.svg; done