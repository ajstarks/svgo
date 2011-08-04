include $(GOROOT)/src/Make.inc

ifeq ($(GOOS), darwin)
OPENER=open -a /Applications/Safari.app
OPENER=open -a /Applications/Google\ Chrome.app
else
OPENER=firefox
endif

SVGLIB=svg

CLIENTS=android bulletgraph bubtrail svgdef flower funnel gradient imfade lewitt planets paths pmap randcomp\
		richter rl skewabc svgopher vismem webfonts

SCLIENTS=android  bubtrail svgdef flower funnel gradient imfade lewitt planets paths randcomp\
		richter rl skewabc svgopher vismem webfonts

all:	$(CLIENTS)

$(SVGLIB).$(O):	$(SVGLIB).go
	$(GC) $(SVGLIB).go
	
android:	android.go svg.$(O)
	$(GC) android.go
	$(LD)  -o android android.$(O)
	
bubtrail:	bubtrail.go svg.$(O)
	$(GC) bubtrail.go
	$(LD)  -o bubtrail bubtrail.$(O)
	
bulletgraph:	bulletgraph.go svg.$(O)
	$(GC) bulletgraph.go
	$(LD)  -o bulletgraph bulletgraph.$(O)
	
flower:	flower.go svg.$(O)
	$(GC) flower.go
	$(LD)  -o flower flower.$(O)
	
	
funnel:	funnel.go svg.$(O)
	$(GC) funnel.go
	$(LD)  -o funnel funnel.$(O)
	
	
gradient:	gradient.go svg.$(O)
	$(GC) gradient.go
	$(LD)  -o gradient gradient.$(O)
	
imfade:	imfade.go svg.$(O)
	$(GC) imfade.go
	$(LD)  -o imfade imfade.$(O)
	
	
lewitt:	lewitt.go svg.$(O)
	$(GC) lewitt.go
	$(LD)  -o lewitt lewitt.$(O)
	
	
paths:	paths.go svg.$(O)
	$(GC) paths.go
	$(LD)  -o paths paths.$(O)
	
	
planets:	planets.go svg.$(O)
	$(GC) planets.go
	$(LD)  -o planets planets.$(O)
	
	
pmap:	pmap.go svg.$(O)
	$(GC) pmap.go
	$(LD)  -o pmap pmap.$(O)
	
randcomp:	randcomp.go svg.$(O)
	$(GC) randcomp.go
	$(LD)  -o randcomp randcomp.$(O)
	
	
richter:	richter.go svg.$(O)
	$(GC) richter.go
	$(LD)  -o richter richter.$(O)
	
	
rl:	rl.go svg.$(O)
	$(GC) rl.go
	$(LD)  -o rl rl.$(O)
	
	
svgdef:	svgdef.go svg.$(O)
	$(GC) svgdef.go
	$(LD) -o svgdef svgdef.$(O)
	
	
skewabc:	skewabc.go svg.$(O)
	$(GC) skewabc.go
	$(LD)  -o skewabc skewabc.$(O)
	
svgopher:	svgopher.go svg.$(O)
	$(GC) svgopher.go
	$(LD)  -o svgopher svgopher.$(O)
	
vismem:	vismem.go svg.$(O)
	$(GC) vismem.go
	$(LD)  -o vismem vismem.$(O)
	
webfonts:	webfonts.go svg.$(O)
	$(GC) webfonts.go
	$(LD)  -o webfonts webfonts.$(O)
	
	
defs:	svgdef
	./svgdef > svgdef.svg
	svg2pdf svgdef.svg
	svg2png svgdef.svg
	tidy -xml -indent -modify svgdef.svg

fmt:	$(CLIENTS)
	for c in $(CLIENTS); do gofmt -w $$c.go; done
	
pdf:	$(SCLIENTS)
	for c in $(SCLIENTS); do ./$$c > $$c.svg; svg2pdf $$c.svg; done
	./pmap pmaptest.xml > pmaptest.svg
	./bulletgraph bg.xml > bg.svg
	svg2pdf pmaptest.svg bg.svg

test:	$(CLIENTS)
	gofmt -w svg.go
	for c in $(SCLIENTS); do ./$$c > $$c.svg; $(OPENER) $$c.svg; done
	./pmap pmaptest.xml > pmaptest.svg
	./bulletgraph bg.xml > bg.svg
	$(OPENER) pmaptest.svg bg.svg
	
clean:
	rm -f svg.$(O) 
	for c in $(CLIENTS); do rm -f $$c.$(O) $$c; done