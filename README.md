# pokemon-tinygo

A Pokemon-style game for the [Badger 2040 W](https://shop.pimoroni.com/products/badger-2040-w) e-ink badge, built with [TinyGo](https://tinygo.org/).

Choose from three Pokemon and battle a random opponent on the e-ink display, with optional Bluetooth advertising.

## Building

Requires [TinyGo](https://tinygo.org/getting-started/install/) and a Badger 2040 W.

```bash
make build
```

## Flashing

Connect the Badger 2040 W via USB, hold `BOOTSEL` + `reset` for 5 seconds, release `reset` first:

```bash
make flash
```

## Image Converter

The `cmd/gopherbadgeimg` tool converts images to the bitmap format used by the e-ink display. See its [README](cmd/gopherbadgeimg/README.md) for details.

## Controls

| Button | Action |
|--------|--------|
| A | Select Pokemon 1 (Char Lizard) |
| B | Select Pokemon 2 (Water Turtle) |
| C | Select Pokemon 3 (Lightning Rat) |
| UP | Confirm selection and battle |
| DOWN | Toggle Bluetooth advertising |

## Troubleshooting

If you see `unable to locate any volume: [RPI-RP2]`, the device needs a reset. Hold `BOOTSEL` on the back of the badge while connecting USB, then mount the volume:

```bash
mkdir -p /media/RPI-RP2
mount /dev/sdb1 /media/RPI-RP2
```
