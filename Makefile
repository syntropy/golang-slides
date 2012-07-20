SLIDES:=$(wildcard *.slides)
HTML_SLIDES:=$(patsubst %.slides,%.html,$(SLIDES))

HTML_TEMPLATE:=template.html
SLIDE_DIR:=slides

RENDER:=render

all: $(RENDER) $(HTML_SLIDES)

$(RENDER): $(RENDER).go
	go build -o $@

%.html: %.slides $(RENDER)
	./$(RENDER) -o $@ -t $(HTML_TEMPLATE) -S $(SLIDE_DIR) $<

clean:
	$(RM) $(HTML_SLIDES)

distclean: clean
	$(RM) $(RENDER)

.PHONY: clean
