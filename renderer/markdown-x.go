// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package renderer

import (
	"fmt"
	"strings"

	"github.com/elastic/crd-ref-docs/config"
)

type MarkdownXRenderer struct {
	*MarkdownRenderer
}

func NewMarkdownXRenderer(conf *config.Config) (*MarkdownXRenderer, error) {
	markdownRenderer, err := NewMarkdownRenderer(conf)
	if err != nil {
		return nil, err
	}
	return &MarkdownXRenderer{markdownRenderer}, nil
}

func (m *MarkdownXRenderer) RenderLocalLink(text string) string {
	anchor := strings.ToLower(
		strings.NewReplacer(
			" ", "-",
			".", "",
			"/", "",
			"(", "",
			")", "",
		).Replace(text),
	)

	label := strings.NewReplacer("{", "\\{").Replace(text)

	return fmt.Sprintf("[%s](#%s)", label, anchor)
}

func (m *MarkdownXRenderer) RenderFieldDoc(text string) string {
	out := text

	// escape inlined markup
	out = strings.ReplaceAll(out, "<", "&lt;")
	out = strings.ReplaceAll(out, ">", "&gt;")

	// Pass to the inner renderer for further processing
	return m.MarkdownRenderer.RenderFieldDoc(out)
}
