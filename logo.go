package main

// 112x96
var logochip = []uint8{
	0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x6E, 0x80, 0x0, 0xB7, 0xC0, 0x1, 0xFF, 0x80, 0x5, 0xB4, 0x0, 0x3, 0xFF, 0x80, 0x0, 0xFF, 0xC0, 0x1, 0xFF, 0x80, 0x7, 0xFF, 0x0,
	0x3, 0xFF, 0xC0, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0x7, 0xFF, 0x0, 0x3, 0xFF, 0x80, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x0, 0x7, 0xFF, 0xC0, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x0, 0x7, 0xFF, 0xC0, 0x1, 0xFF, 0xF0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x0, 0x7, 0xFF, 0xC0, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x80, 0x7, 0xFF, 0xE0, 0x3, 0xFF, 0xF0, 0x3, 0xFF, 0xE0, 0xF, 0xFF, 0x80, 0xF, 0xFF, 0xC0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0x1F, 0xFF, 0x80, 0xF, 0xFF, 0xE0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0xF, 0xFF, 0xC0, 0xF, 0xFF, 0xE0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0x1F, 0xFF, 0xC0, 0x1F, 0xFF, 0xEA, 0x57, 0xFF, 0xFA, 0xAF, 0xFF, 0xF5, 0x5F, 0xFF, 0xC0, 0x3F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF0, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xF5, 0x48, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2F, 0xF8, 0xFF, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xF8, 0x7F, 0x80, 0x0, 0x0, 0x3B, 0xA0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xF8, 0x7F, 0x0, 0x0, 0x0, 0xFF, 0xFC, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x80, 0x0, 0x3, 0xFD, 0xFF, 0x25, 0x55, 0x55, 0x55, 0xB7, 0xF8, 0x7F, 0x0, 0x0, 0x7, 0x80, 0xF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0x0, 0x3E, 0x1F, 0x0, 0x3, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0x0, 0x7F, 0xDC, 0x0, 0x0, 0xF0, 0x0, 0x0, 0x0, 0x1, 0xF8, 0x7F, 0x1, 0xF7, 0xF8, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xC1, 0xF0, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x1, 0xC0, 0xF0, 0x0, 0x0, 0x1C, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xDD, 0xE0, 0x0, 0x0, 0x1C, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x1, 0xE7, 0xC0, 0x0, 0x0, 0xC, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xE3, 0xC0, 0x0, 0x0, 0xC, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0x7E, 0xCF, 0x80, 0x0, 0xC, 0x0, 0x0, 0x1F, 0x0, 0xF8, 0x7F, 0x0, 0x7F, 0x9F, 0x80, 0x0, 0xE, 0x0, 0x0, 0x3F, 0x80, 0xF8, 0x7F, 0x0, 0x1C, 0xDF, 0xC0, 0x0, 0xC, 0x0, 0x0, 0x3F, 0xC0, 0xF8, 0x7F, 0x0, 0x19, 0xDF, 0xC0, 0x0, 0xC, 0x0, 0x0, 0x61, 0xC0, 0xF8, 0x7F, 0x0, 0x39, 0x93, 0xC0, 0x0, 0xE, 0x0, 0x0, 0x70, 0xC0, 0xF8, 0x7F, 0x0, 0x39, 0xCF, 0xC0, 0x0, 0xC, 0x0, 0x0, 0x7F, 0xC0, 0xF8, 0x7F, 0x0, 0x71, 0x87, 0x80, 0x0, 0x1C, 0x0, 0x0, 0x3F, 0x80, 0xF8, 0x7F, 0x0, 0x71, 0xE0, 0x0, 0x0, 0x1C, 0x0, 0x0, 0xF, 0x0, 0xF8, 0x7F, 0x0, 0x60, 0xC0, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xE0, 0xE0, 0x0, 0x0, 0x38, 0x0, 0x1, 0x9F, 0x80, 0xF8, 0x7F, 0x1, 0xC0, 0x70, 0x0, 0x0, 0x70, 0x0, 0x3, 0x8F, 0xC0, 0xF8, 0x7F, 0x0, 0xC0, 0x78, 0x0, 0x0, 0xE0, 0x0, 0x1, 0x9F, 0xC0, 0xF8, 0x7F, 0x1, 0xC0, 0x3C, 0x0, 0x1, 0xE0, 0x0, 0x3, 0x1C, 0xC0, 0xF8, 0x7F, 0x1, 0xC0, 0x1E, 0x0, 0x7, 0xC0, 0x0, 0x3, 0x88, 0xC0, 0xF8, 0x7F, 0x1, 0xC0, 0xF, 0x80, 0x1F, 0x0, 0x0, 0x3, 0x81, 0xC0, 0xF8, 0x7F, 0x1, 0x80, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x1, 0xEB, 0xC0, 0xF8, 0x7F, 0x1, 0x80, 0x0, 0xFF, 0xF8, 0xA0, 0x0, 0x1, 0xFF, 0x80, 0xF8, 0x1F, 0x3, 0x80, 0x0, 0x1F, 0x81, 0xF0, 0x0, 0x0, 0xFF, 0x1, 0xF8, 0x1F, 0x3, 0x80, 0x0, 0x0, 0x3, 0x18, 0x0, 0x0, 0x3C, 0x0, 0xF8, 0xF, 0x3, 0x80, 0x0, 0x0, 0x6, 0x1A, 0x0, 0x0, 0x20, 0x0, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x1E, 0x1F, 0x80, 0x0, 0x7E, 0x0, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x1E, 0x1A, 0xC0, 0x0, 0x7F, 0xC1, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x1F, 0x18, 0x40, 0x0, 0x7, 0xF0, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x3F, 0xF, 0xC0, 0x0, 0x3, 0xF8, 0xF8, 0x7, 0x3, 0x0, 0x0, 0x0, 0x1F, 0x18, 0x40, 0x0, 0x3F, 0xDC, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x1F, 0x1E, 0xC0, 0x0, 0x7F, 0x98, 0xF8, 0x7, 0x3, 0x80, 0x0, 0x0, 0x1E, 0xF, 0x80, 0x0, 0x74, 0x1C, 0xF8, 0xF, 0x3, 0x80, 0x0, 0x0, 0x6, 0x18, 0x0, 0x0, 0x0, 0x0, 0xF8, 0xF, 0x3, 0x80, 0x0, 0x3E, 0x83, 0x98, 0x0, 0x0, 0x3F, 0xC0, 0xF8, 0x1F, 0x3, 0x80, 0x1, 0xFF, 0xF9, 0xF8, 0x0, 0x0, 0x3F, 0xC0, 0xF8, 0x3F, 0x1, 0x80, 0x7, 0xFF, 0xFE, 0x0, 0x0, 0x0, 0x7F, 0x80, 0xF8, 0x7F, 0x3, 0x80, 0xF, 0x80, 0x1F, 0x0, 0x0, 0x0, 0x70, 0x0, 0xF8, 0x7F, 0x1, 0xC0, 0x1E, 0x0, 0x7, 0xC0, 0x0, 0x0, 0x3A, 0x80, 0xF8, 0x7F, 0x1, 0x80, 0x3C, 0x0, 0x1, 0xC0, 0x0, 0x0, 0x7F, 0xC0, 0xF8, 0x7F, 0x1, 0xC0, 0x78, 0x0, 0x0, 0xE0, 0x0, 0x0, 0x7F, 0xC0, 0xF8, 0x7F, 0x1, 0xC0, 0xF0, 0x0, 0x0, 0x70, 0x0, 0x0, 0x0, 0x80, 0xF8, 0x7F, 0x1, 0xC0, 0xE0, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xC0, 0xC0, 0x0, 0x0, 0x38, 0x0, 0x2, 0x3F, 0xC0, 0xF8, 0x7F, 0x0, 0xE1, 0xC0, 0x0, 0x0, 0x18, 0x0, 0x3, 0xFF, 0xC0, 0xF8, 0x7F, 0x0, 0xE1, 0xC0, 0x0, 0x0, 0x1C, 0x0, 0x3, 0xBF, 0xC0, 0xF8, 0x7F, 0x0, 0x71, 0x86, 0x0, 0x0, 0x1C, 0x0, 0x0, 0x70, 0x0, 0xF8, 0x7F, 0x0, 0x71, 0xCF, 0xC0, 0x0, 0xC, 0x0, 0x1, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0x39, 0x9F, 0xC0, 0x0, 0xE, 0x0, 0x3, 0x80, 0x0, 0xF8, 0x7F, 0x0, 0x39, 0xDF, 0xE0, 0x0, 0xC, 0x0, 0x3, 0x80, 0x0, 0xF8, 0x7F, 0x0, 0x3D, 0x9F, 0xE0, 0x0, 0xE, 0x0, 0x1, 0xFF, 0xC0, 0xF8, 0x7F, 0x0, 0x7E, 0xD3, 0xC0, 0x0, 0xC, 0x0, 0x3, 0xFF, 0xC0, 0xF8, 0x7F, 0x0, 0xFF, 0xCF, 0xC0, 0x0, 0xE, 0x0, 0x3, 0xFF, 0xC0, 0xF8, 0x7F, 0x0, 0xE7, 0xC3, 0x0, 0x0, 0xC, 0x0, 0x1, 0x80, 0x0, 0xF8, 0x7F, 0x1, 0xDF, 0xE0, 0x0, 0x0, 0x1C, 0x0, 0x3, 0x80, 0x0, 0xF8, 0x7F, 0x1, 0x84, 0xE0, 0x0, 0x0, 0x1C, 0x0, 0x1, 0x0, 0x0, 0xF8, 0x7F, 0x3, 0xCC, 0xF0, 0x0, 0x0, 0x38, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x1, 0xC1, 0xF8, 0x0, 0x0, 0x78, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x1, 0xEF, 0xFC, 0x0, 0x0, 0xF0, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0xFF, 0x9E, 0x0, 0x3, 0xF0, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x0, 0x34, 0xF, 0xC0, 0xF, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xF8, 0x7F, 0x0, 0x0, 0x3, 0xFA, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0x0, 0x0, 0x0, 0xFF, 0xFC, 0x1, 0x7F, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0x0, 0x0, 0x0, 0x1F, 0xA0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xF8, 0x7F, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xF8, 0x7F, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xF8, 0x7F, 0xC0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xF8, 0x7F, 0xFF, 0xFB, 0xBB, 0xFB, 0xDE, 0xED, 0xB7, 0x77, 0x6E, 0xEF, 0xF8, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xF8, 0x1F, 0xFF, 0xFF, 0xF7, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xE0, 0x1F, 0xFF, 0xE0, 0x7, 0xFF, 0xF0, 0xF, 0xFF, 0xE0, 0x1F, 0xFF, 0x80, 0x1F, 0xFF, 0xE0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0x1F, 0xFF, 0x80, 0xF, 0xFF, 0xC0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0x1F, 0xFF, 0x80, 0xF, 0xFF, 0xC0, 0x3, 0xFF, 0xE0, 0x7, 0xFF, 0xE0, 0x1F, 0xFF, 0x80, 0xF, 0xFF, 0xC0, 0x3, 0xFF, 0xF0, 0x7, 0xFF, 0xE0, 0xF, 0xFF, 0x80, 0x7, 0xFF, 0xC0, 0x3, 0xFF, 0xE0, 0x7, 0xFF, 0xC0, 0x1F, 0xFF, 0x0, 0x7, 0xFF, 0x80, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x0, 0x7, 0xFF, 0xC0, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0xC0, 0xF, 0xFF, 0x0, 0x7, 0xFF, 0x80, 0x1, 0xFF, 0xE0, 0x3, 0xFF, 0x80, 0xF, 0xFF, 0x0, 0x3, 0xFF, 0x80, 0x0, 0xFF, 0xC0, 0x1, 0xFF, 0xC0, 0xF, 0xFE, 0x0, 0x3, 0xFF, 0x80, 0x1, 0xFF, 0xC0, 0x3, 0xFF, 0x80, 0x7, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xAB, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
}
