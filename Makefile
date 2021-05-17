Fuck.app/Contents/MacOS/Fuck: fuck.m fuck.h main.go
	@mkdir -p $(@D)
	go build -o $@ .

install: Fuck.app/Contents/MacOS/Fuck
	cp -a Fuck.app /Applications

.PHONY: install
