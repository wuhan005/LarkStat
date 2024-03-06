// Copyright 2024 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const (
	width  = 700
	height = 330
)

type DrawImageOptions struct {
	Writer http.ResponseWriter
	Count  int
}

func DrawImage(ctx context.Context, options DrawImageOptions) {
	w := options.Writer

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "cache-control: no-cache,max-age=0,no-store,s-maxage=0,proxy-revalidate")

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 255}}, image.Point{}, draw.Src)

	label := strconv.Itoa(options.Count)
	labelWidth := len(label) * 120
	x := (width - labelWidth) / 2
	y := (height + 160) / 2

	labelColor := color.RGBA{R: 20, G: 86, B: 240, A: 255}
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	face, _ := loadFontFace(100)
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(labelColor),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)

	if err := png.Encode(w, img); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("Failed to encode image")
	}
}

func loadFontFace(size float64) (font.Face, error) {
	ttf, err := opentype.Parse(gomono.TTF)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    size,
		DPI:     144,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}

	return face, nil
}
