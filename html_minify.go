// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// Source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package html // import "aahframe.work/minify/html"

import (
	"io"
	"regexp"

	"aahframe.work"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

const htmlContentType = "text/html"
const cssContentType = "text/css"

var m *minify.M

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Global methods
//___________________________________

// HTML method calls the minify library func to minify content. It reads data
// from reader and writes it to the given writer.
func HTML(contentType string, w io.Writer, r io.Reader) error {
	return m.Minify(htmlContentType, w, r)
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Unexported methods
//___________________________________

// addingHTMLMinifer event func to initilize the minify instance and
// adding minifier into framework
func addingHTMLMinifer(e *aah.Event) {
	m = minify.New()
	cfg := aah.AppConfig()

	m.Add(htmlContentType, &html.Minifier{
		KeepConditionalComments: cfg.BoolDefault("render.minify.html.keep.conditional_comments", true),
		KeepDocumentTags:        cfg.BoolDefault("render.minify.html.keep.document_tags", true),
		KeepWhitespace:          cfg.BoolDefault("render.minify.html.keep.whitespace", false),
		KeepDefaultAttrVals:     cfg.BoolDefault("render.minify.html.keep.default_attr_vals", false),
		KeepEndTags:             cfg.BoolDefault("render.minify.html.keep.end_tags", false),
	})

	if !cfg.BoolDefault("render.minify.html.keep.inline_javascript_asis", false) {
		m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	}

	if !cfg.BoolDefault("render.minify.html.keep.inline_css_asis", false) {
		m.Add(cssContentType, &css.Minifier{
			// number of decimals to preserve for numbers
			Decimals: cfg.IntDefault("render.minify.html.css.decimals", -1),
			// prohibits using CSS3 syntax (such as exponents in numbers, or rgba( → rgb(
			KeepCSS2: cfg.BoolDefault("render.minify.html.keep.css2", false),
		})
	}

	// set `HTML` minify func.
	aah.SetMinifier(HTML)
}

func init() {
	// register into aah framework
	aah.OnInit(addingHTMLMinifer)
}
