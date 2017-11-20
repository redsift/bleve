//  Copyright (c) 2015 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	// token maps
	_ "github.com/redsift/bleve/analysis/tokenmap"

	// fragment formatters
	_ "github.com/redsift/bleve/search/highlight/format/ansi"
	_ "github.com/redsift/bleve/search/highlight/format/html"

	// fragmenters
	_ "github.com/redsift/bleve/search/highlight/fragmenter/simple"

	// highlighters
	_ "github.com/redsift/bleve/search/highlight/highlighter/ansi"
	_ "github.com/redsift/bleve/search/highlight/highlighter/html"
	_ "github.com/redsift/bleve/search/highlight/highlighter/simple"

	// char filters
	_ "github.com/redsift/bleve/analysis/char/html"
	_ "github.com/redsift/bleve/analysis/char/regexp"
	_ "github.com/redsift/bleve/analysis/char/zerowidthnonjoiner"

	// analyzers
	_ "github.com/redsift/bleve/analysis/analyzer/custom"
	_ "github.com/redsift/bleve/analysis/analyzer/keyword"
	_ "github.com/redsift/bleve/analysis/analyzer/simple"
	_ "github.com/redsift/bleve/analysis/analyzer/standard"
	_ "github.com/redsift/bleve/analysis/analyzer/web"

	// token filters
	_ "github.com/redsift/bleve/analysis/token/apostrophe"
	_ "github.com/redsift/bleve/analysis/token/compound"
	_ "github.com/redsift/bleve/analysis/token/edgengram"
	_ "github.com/redsift/bleve/analysis/token/elision"
	_ "github.com/redsift/bleve/analysis/token/keyword"
	_ "github.com/redsift/bleve/analysis/token/length"
	_ "github.com/redsift/bleve/analysis/token/lowercase"
	_ "github.com/redsift/bleve/analysis/token/ngram"
	_ "github.com/redsift/bleve/analysis/token/shingle"
	_ "github.com/redsift/bleve/analysis/token/stop"
	_ "github.com/redsift/bleve/analysis/token/truncate"
	_ "github.com/redsift/bleve/analysis/token/unicodenorm"

	// tokenizers
	_ "github.com/redsift/bleve/analysis/tokenizer/exception"
	_ "github.com/redsift/bleve/analysis/tokenizer/regexp"
	_ "github.com/redsift/bleve/analysis/tokenizer/single"
	_ "github.com/redsift/bleve/analysis/tokenizer/unicode"
	_ "github.com/redsift/bleve/analysis/tokenizer/web"
	_ "github.com/redsift/bleve/analysis/tokenizer/whitespace"

	// date time parsers
	_ "github.com/redsift/bleve/analysis/datetime/flexible"
	_ "github.com/redsift/bleve/analysis/datetime/optional"

	// languages
	_ "github.com/redsift/bleve/analysis/lang/ar"
	_ "github.com/redsift/bleve/analysis/lang/bg"
	_ "github.com/redsift/bleve/analysis/lang/ca"
	_ "github.com/redsift/bleve/analysis/lang/cjk"
	_ "github.com/redsift/bleve/analysis/lang/ckb"
	_ "github.com/redsift/bleve/analysis/lang/cs"
	_ "github.com/redsift/bleve/analysis/lang/el"
	_ "github.com/redsift/bleve/analysis/lang/en"
	_ "github.com/redsift/bleve/analysis/lang/eu"
	_ "github.com/redsift/bleve/analysis/lang/fa"
	_ "github.com/redsift/bleve/analysis/lang/fr"
	_ "github.com/redsift/bleve/analysis/lang/ga"
	_ "github.com/redsift/bleve/analysis/lang/gl"
	_ "github.com/redsift/bleve/analysis/lang/hi"
	_ "github.com/redsift/bleve/analysis/lang/hy"
	_ "github.com/redsift/bleve/analysis/lang/id"
	_ "github.com/redsift/bleve/analysis/lang/in"
	_ "github.com/redsift/bleve/analysis/lang/it"
	_ "github.com/redsift/bleve/analysis/lang/pt"

	// kv stores
	_ "github.com/redsift/bleve/index/store/boltdb"
	_ "github.com/redsift/bleve/index/store/goleveldb"
	_ "github.com/redsift/bleve/index/store/gtreap"
	_ "github.com/redsift/bleve/index/store/moss"

	// index types
	_ "github.com/redsift/bleve/index/upsidedown"
)
