target = main.uf2
disk_label = RPI-RP2

.PHONY: build
build: ${target}

.PHONY: ${target}
${target}:
	tinygo build --target=badger2040-w --size full --print-allocs=. -o "${target}" --monitor --stack-size=8kb

.PHONY: flash
flash:
	udisksctl mount -b "/dev/disk/by-label/${disk_label}"
	cp "${target}" "/media/${USER}/${disk_label}/"
	#udisksctl unmount -b "/media/${USER}/${disk_label}/"
