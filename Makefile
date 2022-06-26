SERVICE := compiler60
DESTDIR ?= dist_root
SERVICEDIR ?= /srv/$(SERVICE)
SYSTEMD_DIR ?= $(DESTDIR)/etc/systemd/system
SYSTEMD_WANT_DIR ?= $(SYSTEMD_DIR)/faustctf.target.wants

.PHONY: dist_root

gui.html: $(shell find website/)
	python website/build.py $@

dist_root: gui.html
# systemd conf
	mkdir -p $(SYSTEMD_WANT_DIR)
	cp compiler60.service $(SYSTEMD_DIR)
	ln -s /etc/systemd/system/compiler60.service $(SYSTEMD_WANT_DIR)/compiler60.service
# base service files
	mkdir -p $(DESTDIR)$(SERVICEDIR)
	cp docker-compose.yml $(DESTDIR)$(SERVICEDIR)
	cp README.md $(DESTDIR)$(SERVICEDIR)
	cp gui.html $(DESTDIR)$(SERVICEDIR)
# compiler service
	cp -r compiler/ $(DESTDIR)$(SERVICEDIR)
	rm -rf $(DESTDIR)$(SERVICEDIR)/compiler/src/test/
	cp $(wildcard ./compiler/src/test/resources/programs/example*.a60) $(DESTDIR)$(SERVICEDIR)/compiler/
# executor service
	mkdir -p $(DESTDIR)$(SERVICEDIR)/executor/
	cp executor/LICENSE $(DESTDIR)$(SERVICEDIR)/executor/
	cp executor/syscall-filter.kafel $(DESTDIR)$(SERVICEDIR)/executor/
