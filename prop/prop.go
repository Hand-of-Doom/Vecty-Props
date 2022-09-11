package prop

import (
	"fmt"
	"github.com/hexops/vecty"
	"regexp"
	"strings"
	"time"
)

type (
	URL       = string // google.com
	EntityRef = string // id/class referred to element
)

type AcceptCase = string

const (
	AcceptCaseMedia AcceptCase = "audio/*"
	AcceptCaseVideo AcceptCase = "video/*"
	AcceptCaseImage AcceptCase = "image/*"
)

// Accept specifies the types of files that the server accepts (only for type="file")
//
// <input>
func Accept(c AcceptCase) vecty.Applyer {
	return vecty.Property("accept", c)
}

// AcceptCharset specifies the character encodings that are to be used for the form submission
//
// <form>
func AcceptCharset(values ...string) vecty.Applyer {
	return vecty.Property("accept-charset", strings.Join(values, " "))
}

// AccessKey specifies a shortcut key to activate/focus an element
//
// Global Attributes
func AccessKey(value string) vecty.Applyer {
	return vecty.Property("accesskey", value)
}

// Action specifies where to send the form-data when a form is submitted
//
// <form>
func Action(value URL) vecty.Applyer {
	return vecty.Property("action", value)
}

// Alt specifies an alternate text when the original element fails to display
//
// <area>, <img>, <input>
func Alt(value string) vecty.Applyer {
	return vecty.Property("alt", value)
}

// Async specifies that the script is executed asynchronously (only for external scripts)
// <script>
func Async(flag bool) vecty.Applyer {
	return vecty.Property("async", flag)
}

// Autocomplete specifies whether the <form> or the <input> element should have autocomplete enabled
//
// <form>, <input>
func Autocomplete(flag bool) vecty.Applyer {
	var stringFlag string
	if flag {
		stringFlag = "on"
	} else {
		stringFlag = "off"
	}

	return vecty.Property("autocomplete", stringFlag)
}

// Autofocus specifies that the element should automatically get focus when the page loads
//
// <button>, <input>, <select>, <textarea>
func Autofocus(flag bool) vecty.Applyer {
	return vecty.Property("autofocus", flag)
}

// Autoplay specifies that the audio/video will start playing as soon as it is ready
//
// <audio>, <video>
func Autoplay(flag bool) vecty.Applyer {
	return vecty.Property("autoplay", flag)
}

// Charset specifies the character encoding
//
// <meta>, <script>
func Charset(value string) vecty.Applyer {
	return vecty.Property("charset", value)
}

// Checked specifies that an <input> element should be pre-selected when the page loads (for type="checkbox" or type="radio")
//
// <input>
func Checked(flag bool) vecty.Applyer {
	return vecty.Property("checked", flag)
}

// Cite specifies a URL which explains the quote/deleted/inserted text
//
// <blockquote>, <del>, <ins>, <q>
func Cite(value URL) vecty.Applyer {
	return vecty.Property("cite", value)
}

// Class specifies one or more classnames for an element (refers to a class in a style sheet)
//
// Global Attributes
func Class(values ...string) vecty.Applyer {
	return vecty.Property("class", strings.Join(values, " "))
}

// Cols specifies the visible width of a text area
//
// <textarea>
func Cols(value uint64) vecty.Applyer {
	return vecty.Property("cols", value)
}

// Colspan specifies the number of columns a table cell should span
//
// <td>, <th>
func Colspan(value uint64) vecty.Applyer {
	return vecty.Property("colspan", value)
}

// Content gives the value associated with the http-equiv or name attribute
//
// <meta>
func Content(value interface{}) vecty.Applyer {
	return vecty.Property("content", value)
}

// ContentEditable specifies whether the content of an element is editable or not
//
// Global Attributes
func ContentEditable(flag bool) vecty.Applyer {
	return vecty.Property("contenteditable", flag)
}

// Controls specifies that audio/video controls should be displayed (such as a play/pause button etc)
//
// <audio>, <video>
func Controls(flag bool) vecty.Applyer {
	return vecty.Property("controls", flag)
}

type CoordsSet interface {
	buildTemplate() string
}

type RectCoords struct {
	XLeftTop     int64
	YLeftTop     int64
	XBottomRight int64
	YBottomRight int64
}

func (set RectCoords) buildTemplate() string {
	if set.XLeftTop >= set.XBottomRight {
		panic("the first integer must be less than the third")
	}
	if set.YLeftTop >= set.YBottomRight {
		panic("the second integer must be less than the fourth")
	}

	return fmt.Sprintf("%d,%d,%d,%d",
		set.XLeftTop, set.YLeftTop, set.XBottomRight, set.YBottomRight)
}

type CircleCoords struct {
	X      int64
	Y      int64
	Radius string
}

func (set CircleCoords) buildTemplate() string {
	radiusPattern := regexp.MustCompile(`^([0-9]+)(%|)$`)

	if !radiusPattern.MatchString(set.Radius) {
		wrongRunePattern := regexp.MustCompile(`[^0-9]`)
		wrongRune := wrongRunePattern.FindString(set.Radius)[0]

		panic("expected a digit but saw " + string(wrongRune) + " instead")
	}

	return fmt.Sprintf("%d,%d,%s",
		set.X, set.Y, set.Radius)
}

type PolyCoords [][2]int64

func (set PolyCoords) buildTemplate() string {
	if len(set) < 3 {
		panic("a polyline must have at least six comma-separated integers")
	}

	var template string
	for _, pair := range set {
		template += fmt.Sprintf("%d,%d,", pair[0], pair[1])
	}

	return strings.TrimSuffix(template, ",")
}

// Coords specifies the coordinates of the area
//
// <area>
func Coords(value CoordsSet) vecty.Applyer {
	return vecty.Property("coords", value.buildTemplate())
}

// Data specifies the URL of the resource to be used by the object
//
// <object>
func Data(value URL) vecty.Applyer {
	return vecty.Property("data", value)
}

// DataPair used to store custom data private to the page or application
//
// Global Attributes
func DataPair(key string, value interface{}) vecty.Applyer {
	return vecty.Property("data-"+key, value)
}

// Datetime specifies the date and time
//
// <del>, <ins>, <time>
func Datetime(value time.Time) vecty.Applyer {
	return vecty.Property("datetime", value.String())
}

// Default specifies that the track is to be enabled if the user's preferences do not indicate that another track would be more appropriate
//
// <track>
func Default(flag bool) vecty.Applyer {
	return vecty.Property("default", flag)
}

// Defer specifies that the script is executed when the page has finished parsing (only for external scripts)
//
// <script>
func Defer(flag bool) vecty.Applyer {
	return vecty.Property("defer", flag)
}

type DirCase = string

const (
	DirCaseLTR  DirCase = "ltr"
	DirCaseRTL  DirCase = "rtl"
	DirCaseAuto DirCase = "auto"
)

// Dir specifies the text direction for the content in an element
//
// Global Attributes
func Dir(c DirCase) vecty.Applyer {
	return vecty.Property("dir", c)
}

// Dirname specifies that the text direction will be submitted
//
// <input>, <textarea>
func Dirname(value string) vecty.Applyer {
	return vecty.Property("dirname", value+".dir")
}

// Disabled specifies that the specified element/group of elements should be disabled
//
// <button>, <fieldset>, <input>, <optgroup>, <option>, <select>, <textarea>
func Disabled(flag bool) vecty.Applyer {
	return vecty.Property("disabled", flag)
}

// Download specifies that the target will be downloaded when a user clicks on the hyperlink
//
// <a>, <area>
func Download(flag bool) vecty.Applyer {
	return vecty.Property("download", flag)
}

// DownloadWithFilename specifies that the target will be downloaded when a user clicks on the hyperlink
//
// <a>, <area>
func DownloadWithFilename(filename string) vecty.Applyer {
	return vecty.Property("download", filename)
}

// Draggable specifies whether an element is draggable or not
//
// Global Attributes
func Draggable(flag bool) vecty.Applyer {
	return vecty.Property("draggable", flag)
}

type EnctypeCase = string

const (
	EnctypeCaseFormUrlencoded    EnctypeCase = "application/x-www-form-urlencoded"
	EnctypeCaseMultipartFormData EnctypeCase = "multipart/form-data"
	EnctypeCasePlainText         EnctypeCase = "text/plain"
)

// Enctype specifies how the form-data should be encoded when submitting it to the server (only for method="post")
//
// <form>
func Enctype(c EnctypeCase) vecty.Applyer {
	return vecty.Property("enctype", c)
}

// For specifies which form element(s) a label/calculation is bound to
//
// <label>, <output>
func For(value EntityRef) vecty.Applyer {
	return vecty.Property("htmlFor", value)
}

// Form specifies the name of the form the element belongs to
//
// <button>, <fieldset>, <input>, <label>, <meter>, <object>, <output>, <select>, <textarea>
func Form(value EntityRef) vecty.Applyer {
	return vecty.Property("form", value)
}

// FormAction specifies where to send the form-data when a form is submitted. Only for type="submit"
//
// <button>, <input>
func FormAction(value URL) vecty.Applyer {
	return vecty.Property("formaction", value)
}

// Headers specifies one or more headers cells a cell is related to
//
// <td>, <th>
func Headers(value EntityRef) vecty.Applyer {
	return vecty.Property("headers", value)
}

// Height specifies the height of the element
//
// <canvas>, <embed>, <iframe>, <img>, <input>, <object>, <video>
func Height(value uint64) vecty.Applyer {
	return vecty.Property("height", value)
}

// Hidden specifies that an element is not yet, or is no longer, relevant
//
// Global Attributes
func Hidden(flag bool) vecty.Applyer {
	return vecty.Property("hidden", flag)
}

// High specifies the range that is considered to be a high value
//
// <meter>
func High(value int64) vecty.Applyer {
	return vecty.Property("high", value)
}

// Href specifies the URL of the page the link goes to
//
// <a>, <area>, <base>, <link>
func Href(value URL) vecty.Applyer {
	return vecty.Property("href", value)
}

// HrefLang specifies the language of the linked document
//
// <a>, <area>, <link>
func HrefLang(value string) vecty.Applyer {
	return vecty.Property("hreflang", value)
}

type httpEquivCase = string

const (
	httpEquivCaseSecurityPolice httpEquivCase = "content-security-policy"
	httpEquivContentTypeCase    httpEquivCase = "content-type"
	httpEquivCaseDefaultStyle   httpEquivCase = "default-style"
	httpEquivCaseRefresh        httpEquivCase = "refresh"
)

// HttpEquiv provides an HTTP header for the information/value of the content attribute
//
// <meta>
func HttpEquiv(c httpEquivCase) vecty.Applyer {
	return vecty.Property("http-equiv", c)
}

// ID specifies a unique id for an element
//
// Global Attributes
func ID(value EntityRef) vecty.Applyer {
	return vecty.Property("id", value)
}

// IsMap specifies an image as a server-side image map
//
// <img>
func IsMap(flag bool) vecty.Applyer {
	return vecty.Property("ismap", flag)
}

type KindCase = string

const (
	KindCaseCaptions     KindCase = "captions"
	KindCaseChapters     KindCase = "chapters"
	KindCaseDescriptions KindCase = "descriptions"
	KindCaseMetadata     KindCase = "metadata"
	KindCaseSubtitles    KindCase = "subtitles"
)

// Kind specifies the kind of text track
//
// <track>
func Kind(c KindCase) vecty.Applyer {
	return vecty.Property("kind", c)
}

// Label specifies the title of the text track
//
// <track>, <option>, <optgroup>
func Label(value string) vecty.Applyer {
	return vecty.Property("label", value)
}

// Lang specifies the language of the element's content
//
// Global Attributes
func Lang(value string) vecty.Applyer {
	return vecty.Property("lang", value)
}

// List refers to a <datalist> element that contains pre-defined options for an <input> element
//
// <input>
func List(value EntityRef) vecty.Applyer {
	return vecty.Property("list", value)
}

// Loop specifies that the audio/video will start over again, every time it is finished
//
// <audio>, <video>
func Loop(flag bool) vecty.Applyer {
	return vecty.Property("loop", flag)
}

// Low specifies the range that is considered to be a low value
//
// <meter>
func Low(value int64) vecty.Applyer {
	return vecty.Property("low", value)
}

// Max specifies the maximum value
//
// <input>, <meter>, <progress>
func Max(value string) vecty.Applyer {
	return vecty.Property("max", value)
}

// MaxLength specifies the maximum number of characters allowed in an element
//
// <input>, <textarea>
func MaxLength(value uint64) vecty.Applyer {
	return vecty.Property("maxlength", value)
}

type MediaQuery string

func (b *MediaQuery) And() *MediaQuery {
	*b += "and "

	return b
}

func (b *MediaQuery) Comma() *MediaQuery {
	*b += ", "

	return b
}

func (b *MediaQuery) Not() *MediaQuery {
	*b += "not "

	return b
}

func (b *MediaQuery) All() *MediaQuery {
	*b += "all "

	return b
}

func (b *MediaQuery) Aural() *MediaQuery {
	*b += "aural "

	return b
}

func (b *MediaQuery) Braille() *MediaQuery {
	*b += "braille "

	return b
}

func (b *MediaQuery) Handheld() *MediaQuery {
	*b += "handheld "

	return b
}

func (b *MediaQuery) Projection() *MediaQuery {
	*b += "projection "

	return b
}

func (b *MediaQuery) Print() *MediaQuery {
	*b += "print "

	return b
}

func (b *MediaQuery) Screen() *MediaQuery {
	*b += "screen "

	return b
}

func (b *MediaQuery) TTY() *MediaQuery {
	*b += "tty "

	return b
}

func (b *MediaQuery) TV() *MediaQuery {
	*b += "tv "

	return b
}

func (b *MediaQuery) Width(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(width: %dpx) ", value))

	return b
}

func (b *MediaQuery) Height(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(heigth: %dpx) ", value))

	return b
}

func (b *MediaQuery) DeviceWidth(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(device-width: %dpx) ", value))

	return b
}

func (b *MediaQuery) DeviceHeight(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(device-height: %dpx) ", value))

	return b
}

type OrientationCase = string

const (
	OrientationCaseLandscape = "landscape"
	OrientationCasePortrait  = "portrait"
)

func (b *MediaQuery) Orientation(t OrientationCase) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(orientation: %s) ", t))

	return b
}

func (b *MediaQuery) AspectRatio(width, height int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(aspect-ratio: %d/%d) ", width, height))

	return b
}

func (b *MediaQuery) DeviceAspectRatio(width, height int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(device-aspect-ratio: %d/%d) ", width, height))

	return b
}

func (b *MediaQuery) Color(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(color: %d) ", value))

	return b
}

func (b *MediaQuery) ColorIndex(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(color-index: %d) ", value))

	return b
}

func (b *MediaQuery) Monochrome(value int64) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(monochrome: %d) ", value))

	return b
}

func (b *MediaQuery) Resolution(value string) *MediaQuery {
	resolutionPattern := regexp.MustCompile(`^([0-9]+)(dpi|dpcm)$`)
	if !resolutionPattern.MatchString(value) {
		panic("unknown dimension")
	}

	*b += MediaQuery(fmt.Sprintf("(resolution: %s) ", value))

	return b
}

type ScanCase = string

const (
	ScanCaseProgressive = "progressive"
	ScanCaseInterlace   = "interlace"
)

func (b *MediaQuery) Scan(t ScanCase) *MediaQuery {
	*b += MediaQuery(fmt.Sprintf("(scan: %s) ", t))

	return b
}

func (b *MediaQuery) Grid(value bool) *MediaQuery {
	var intValue int
	if value {
		intValue = 1
	}

	*b += MediaQuery(fmt.Sprintf("(scan: %d) ", intValue))

	return b
}

func NewMediaQuery() *MediaQuery {
	return new(MediaQuery)
}

// Media specifies what media/device the linked document is optimized for
//
// <a>, <area>, <link>, <source>, <style>
func Media(value *MediaQuery) vecty.Applyer {
	return vecty.Property("media", string(*value))
}

type MethodCase = string

const (
	MethodCaseGET  MethodCase = "GET"
	MethodCasePOST MethodCase = "POST"
)

// Method specifies the HTTP method to use when sending form-data
//
// <form>
func Method(c MethodCase) vecty.Applyer {
	return vecty.Property("method", c)
}

// Min specifies a minimum value
//
// <input>, <meter>
func Min(value string) vecty.Applyer {
	return vecty.Property("min", value)
}

// Multiply specifies that a user can enter more than one value
//
// <input>, <select>
func Multiply(flag bool) vecty.Applyer {
	return vecty.Property("multiply", flag)
}

// Muted specifies that the audio output of the video should be muted
//
// <video>, <audio>
func Muted(flag bool) vecty.Applyer {
	return vecty.Property("muted", flag)
}

// Name specifies the name of the element
//
// <button>, <fieldset>, <form>, <iframe>, <input>, <map>, <meta>, <object>, <output>, <param>, <select>, <textarea>
func Name(value EntityRef) vecty.Applyer {
	return vecty.Property("name", value)
}

// Novalidate specifies that the form should not be validated when submitted
//
// <form>
func Novalidate(flag bool) vecty.Applyer {
	return vecty.Property("novalidate", flag)
}

// Open specifies that the details should be visible (open) to the user
//
// <details>
func Open(flag bool) vecty.Applyer {
	return vecty.Property("open", flag)
}

// Optimum specifies what value is the optimal value for the gauge
//
// <meter>
func Optimum(value int64) vecty.Applyer {
	return vecty.Property("optimum", value)
}

// Pattern specifies a regular expression that an <input> element's value is checked against
//
// <input>
func Pattern(value *regexp.Regexp) vecty.Applyer {
	return vecty.Property("pattern", value.String())
}

// Placeholder specifies a short hint that describes the expected value of the element
//
// <input>, <textarea>
func Placeholder(value string) vecty.Applyer {
	return vecty.Property("placeholder", value)
}

// Poster specifies an image to be shown while the video is downloading, or until the user hits the play button
//
// <video>
func Poster(value URL) vecty.Applyer {
	return vecty.Property("poster", value)
}

type PreloadCase = string

const (
	PreloadCaseAuto     PreloadCase = "auto"
	PreloadCaseMetadata PreloadCase = "metadata"
	PreloadCaseNone     PreloadCase = "none"
)

// Preload specifies if and how the author thinks the audio/video should be loaded when the page loads
//
// <audio>, <video>
func Preload(c PreloadCase) vecty.Applyer {
	return vecty.Property("preload", c)
}

// Readonly specifies that the element is read-only
//
// <input>, <textarea>
func Readonly(flag bool) vecty.Applyer {
	return vecty.Property("readonly", flag)
}

type RelCase = string

const (
	RelCaseAlternate  RelCase = "alternate"
	RelCaseAuthor     RelCase = "author"
	RelCaseBookmark   RelCase = "bookmark"
	RelCaseExternal   RelCase = "external"
	RelCaseHelp       RelCase = "help"
	RelCaseLicence    RelCase = "licence"
	RelCaseNext       RelCase = "next"
	RelCaseNofollow   RelCase = "nofollow"
	RelCaseNoOpener   RelCase = "noopener"
	RelCaseNoReferrer RelCase = "noreferrer"
	RelCasePrev       RelCase = "prev"
	RelCaseSearch     RelCase = "search"
	RelCaseTag        RelCase = "tag"
)

// Rel specifies the relationship between the current document and the linked document
//
// <a>, <area>, <form>, <link>
func Rel(c RelCase) vecty.Applyer {
	return vecty.Property("rel", c)
}

// Required specifies that the element must be filled out before submitting the form
//
// <input>, <select>, <textarea>
func Required(flag bool) vecty.Applyer {
	return vecty.Property("required", flag)
}

// Reversed specifies that the list order should be descending (9,8,7...)
//
// <ol>
func Reversed(flag bool) vecty.Applyer {
	return vecty.Property("reversed", flag)
}

// Rows specifies the visible number of lines in a text area
//
// <textarea>
func Rows(value uint64) vecty.Applyer {
	return vecty.Property("rows", value)
}

// RowSpan specifies the number of rows a table cell should span
//
// <td>, <th>
func RowSpan(value uint64) vecty.Applyer {
	return vecty.Property("rowspan", value)
}

// Sandbox enables an extra set of restrictions for the content in an <iframe>
//
// <iframe>
func Sandbox(flag bool) vecty.Applyer {
	return vecty.Property("sandbox", flag)
}

type ScopeCase = string

const (
	ScopeCaseCol      ScopeCase = "col"
	ScopeCaseRow      ScopeCase = "row"
	ScopeCaseColGroup ScopeCase = "colgroup"
	ScopeCaseRowGroup ScopeCase = "rowgroup"
)

// Scope specifies whether a header cell is a header for a column, row, or group of columns or rows
//
// <th>
func Scope(c ScopeCase) vecty.Applyer {
	return vecty.Property("scope", c)
}

// Selected specifies that an option should be pre-selected when the page loads
//
// <option>
func Selected(flag bool) vecty.Applyer {
	return vecty.Property("selected", flag)
}

type ShapeCase = string

const (
	ShapeCaseDefault ShapeCase = "default"
	ShapeCaseRect    ShapeCase = "rect"
	ShapeCaseCircle  ShapeCase = "circle"
	ShapeCasePoly    ShapeCase = "poly"
)

// Shape specifies the shape of the area
//
// <area>
func Shape(c ShapeCase) vecty.Applyer {
	return vecty.Property("shape", c)
}

// Size specifies the width, in characters (for <input>) or specifies the number of visible options (for <select>)
//
// <input>, <select>
func Size(value uint64) vecty.Applyer {
	return vecty.Property("size", value)
}

type SizesSet interface {
	buildTemplate() string
}

// MediaQuerySize applies to <img> <source>
type MediaQuerySize struct {
	conditions []string
	size       string
}

func (b *MediaQuerySize) MinWidth(value string) *MediaQuerySize {
	b.conditions = append(b.conditions, fmt.Sprintf("(min-width: %s)", value))

	return b
}

func (b *MediaQuerySize) MaxWidth(value string) *MediaQuerySize {
	b.conditions = append(b.conditions, fmt.Sprintf("(max-width: %s)", value))

	return b
}

func (b *MediaQuerySize) And() *MediaQuerySize {
	b.conditions = append(b.conditions, "and")

	return b
}

func (b *MediaQuerySize) Or() *MediaQuerySize {
	b.conditions = append(b.conditions, "or")

	return b
}

func (b *MediaQuerySize) build() string {
	tpl := strings.Join(b.conditions, " ")

	if len(b.conditions) > 1 {
		tpl = fmt.Sprintf("(%s)", tpl)
	}
	tpl += fmt.Sprintf(" %s", b.size)

	return tpl
}

func NewMediaQuerySize(size string) *MediaQuerySize {
	return &MediaQuerySize{
		size: size,
	}
}

type ImageSizes struct {
	sizes []string
}

func (b *ImageSizes) Group(size *MediaQuerySize) *ImageSizes {
	b.sizes = append(b.sizes, size.build())

	return b
}

func (b *ImageSizes) Default(size string) *ImageSizes {
	b.sizes = append(b.sizes, size)

	return b
}

func (b *ImageSizes) buildTemplate() string {
	return strings.Join(b.sizes, ", ")
}

func NewImageSizes() *ImageSizes {
	return &ImageSizes{}
}

type LinkSizes struct {
	sizes [][2]uint64
}

func (b *LinkSizes) Pair(width uint64, height uint64) *LinkSizes {
	b.sizes = append(b.sizes, [2]uint64{width, height})

	return b
}

func (b *LinkSizes) buildTemplate() string {
	var tpl string

	for _, size := range b.sizes {
		width, height := size[0], size[1]

		if width == 0 || height == 0 {
			tpl = "any"
			break
		}

		tpl += fmt.Sprintf("%dx%d ", width, height)
	}

	return tpl
}

func NewLinkSizes() *LinkSizes {
	return &LinkSizes{}
}

// Sizes specifies the size of the linked resource
//
// <img>, <link>, <source>
func Sizes(value SizesSet) vecty.Applyer {
	return vecty.Property("sizes", value.buildTemplate())
}

// Span specifies the number of columns to span
//
// <col>, <colgroup>
func Span(value uint64) vecty.Applyer {
	return vecty.Property("span", value)
}

// SpellCheck specifies whether the element is to have its spelling and grammar checked or not
//
// Global Attributes
func SpellCheck(flag bool) vecty.Applyer {
	return vecty.Property("spellcheck", flag)
}

// Src specifies the URL of the media file
//
// <audio>, <embed>, <iframe>, <img>, <input>, <script>, <source>, <track>, <video>
func Src(value URL) vecty.Applyer {
	return vecty.Property("src", value)
}

type Attr struct {
	key,
	value string
}

func (b *Attr) build() string {
	return fmt.Sprintf(`%s="%s"`, b.key, b.value)
}

func NewAttr(key, value string) *Attr {
	return &Attr{
		key:   key,
		value: value,
	}
}

type FakeDOM interface {
	buildTree() string
}

type RawNode struct {
	tree string
}

func (b *RawNode) Include(nodes ...FakeDOM) *RawNode {
	for _, node := range nodes {
		b.tree += node.buildTree()
	}

	return b
}

func (b *RawNode) buildTree() string {
	return b.tree
}

func NewRawNode(html string) *RawNode {
	return &RawNode{
		tree: html,
	}
}

func NewEmptyNode(nodes ...Node) *RawNode {
	var tree string
	for _, node := range nodes {
		tree += node.buildTree()
	}

	return &RawNode{
		tree: tree,
	}
}

type Node struct {
	name  string
	attrs []*Attr
	nodes []FakeDOM
}

func (b *Node) Include(nodes ...FakeDOM) *Node {
	b.nodes = append(b.nodes, nodes...)

	return b
}

func (b *Node) single() bool {
	singleTags := "area base br col command embed hr img input keygen link meta param source track wbr"

	return strings.Contains(b.name, singleTags)
}

func (b *Node) buildTree() string {
	tpl := fmt.Sprintf("<%s", b.name)

	if len(b.attrs) != 0 {
		attrs := " "
		for _, attr := range b.attrs {
			attrs += attr.build()
		}

		tpl += attrs
	}
	tpl += ">"

	for _, node := range b.nodes {
		tpl += node.buildTree()
	}

	if !b.single() {
		tpl += fmt.Sprintf("</%s>", b.name)
	}

	return tpl
}

func NewNode(name string, attrs ...*Attr) *Node {
	return &Node{
		name:  name,
		attrs: attrs,
	}
}

// SrcDoc specifies the HTML content of the page to show in the <iframe>
//
// <iframe>
func SrcDoc(value FakeDOM) vecty.Applyer {
	return vecty.Property("srcdoc", value.buildTree())
}

// SrcLang specifies the language of the track text data (required if kind="subtitles")
//
// <track>
func SrcLang(value string) vecty.Applyer {
	return vecty.Property("srclang", value)
}

type SrcsetPair struct {
	url      string
	template string
}

func (b *SrcsetPair) Width(value uint64) *SrcsetPair {
	b.template = fmt.Sprintf("%s %dw", b.url, value)

	return b
}

func (b *SrcsetPair) PixelDensity(value uint64) *SrcsetPair {
	b.template = fmt.Sprintf("%s %dx", b.url, value)

	return b
}

func (b *SrcsetPair) build() string {
	if b.template == "" {
		panic(fmt.Sprintf("Bad value %s for attribute srcset on element source: Must contain one or more image candidate strings.",
			b.template))
	}

	return b.template
}

func NewSrcsetPair(url string) *SrcsetPair {
	return &SrcsetPair{
		url: url,
	}
}

// Srcset specifies the URL of the image to use in different situations
//
// <img>, <source>
func Srcset(values ...*SrcsetPair) vecty.Applyer {
	pairs := make([]string, 0, len(values))

	for _, pair := range values {
		pairs = append(pairs, pair.build())
	}
	tpl := strings.Join(pairs, ", ")

	return vecty.Property("srcset", tpl)
}

// Start specifies the start value of an ordered list
//
// <ol>
func Start(value int64) vecty.Applyer {
	return vecty.Property("start", value)
}

// Step specifies the legal number intervals for an input field
//
// <input>
func Step(value uint64) vecty.Applyer {
	var stringValue string
	if value == 0 {
		stringValue = "any"
	} else {
		stringValue = fmt.Sprintf("%d", value)
	}

	return vecty.Property("step", stringValue)
}

// Style specifies an inline CSS style for an element
//
// Global Attributes

// Vecty has a style package so this function panics if I set the style with vecty.Property
/*
func Style(values CSS) vecty.Applyer {
	return vecty.Property("style", values.FormatRaw())
}
*/

// TabIndex specifies the tabbing order of an element
//
// Global Attributes
func TabIndex(value int64) vecty.Applyer {
	return vecty.Property("tabindex", value)
}

type TargetCase = string

const (
	TargetCaseBlank  TargetCase = "_blank"
	TargetCaseSelf   TargetCase = "_self"
	TargetCaseParent TargetCase = "_parent"
	TargetCaseTop    TargetCase = "_top"
)

// Target specifies the target for where to open the linked document or where to submit the form
//
// <a>, <area>, <base>, <form>
func Target(c TargetCase) vecty.Applyer {
	return vecty.Property("target", c)
}

// Title specifies extra information about an element
//
// Global Attributes
func Title(value string) vecty.Applyer {
	return vecty.Property("title", value)
}

// Translate specifies whether the content of an element should be translated or not
//
// Global Attributes
func Translate(flag bool) vecty.Applyer {
	return vecty.Property("translate", flag)
}

type TypeCase = string

const (
	TypeCaseButton        TypeCase = "button"
	TypeCaseCheckbox      TypeCase = "checkbox"
	TypeCaseColor         TypeCase = "color"
	TypeCaseDate          TypeCase = "date"
	TypeCaseDatetime      TypeCase = "datetime"
	TypeCaseDatetimeLocal TypeCase = "datetime-local"
	TypeCaseEmail         TypeCase = "email"
	TypeCaseFile          TypeCase = "file"
	TypeCaseHidden        TypeCase = "hidden"
	TypeCaseImage         TypeCase = "image"
	TypeCaseMonth         TypeCase = "month"
	TypeCaseNumber        TypeCase = "number"
	TypeCasePassword      TypeCase = "password"
	TypeCaseRadio         TypeCase = "radio"
	TypeCaseRange         TypeCase = "range"
	TypeCaseMin           TypeCase = "min"
	TypeCaseMax           TypeCase = "max"
	TypeCaseValue         TypeCase = "value"
	TypeCaseStep          TypeCase = "step"
	TypeCaseReset         TypeCase = "reset"
	TypeCaseSearch        TypeCase = "search"
	TypeCaseSubmit        TypeCase = "submit"
	TypeCaseTel           TypeCase = "tel"
	TypeCaseText          TypeCase = "text"
	TypeCaseTime          TypeCase = "time"
	TypeCaseURL           TypeCase = "url"
	TypeCaseWeek          TypeCase = "week"
	TypeCaseList          TypeCase = "list"
	TypeCaseContext       TypeCase = "context"
	TypeCaseToolbar       TypeCase = "toolbar"
	TypeCaseModule        TypeCase = "module"
)

// Type specifies the type of element
//
// <a>, <button>, <embed>, <input>, <link>, <menu>, <object>, <script>, <source>, <style>
func Type(c TypeCase) vecty.Applyer {
	return vecty.Property("type", c)
}

// UseMap specifies an image as a client-side image map
//
// <img>, <object>
func UseMap(value EntityRef) vecty.Applyer {
	return vecty.Property("usemap", "#"+value)
}

// Value specifies the value of the element
//
// <button>, <input>, <li>, <option>, <meter>, <progress>, <param>
func Value(propValue interface{}) vecty.Applyer {
	return vecty.Property("value", propValue)
}

// Width specifies the width of the element
//
// <canvas>, <embed>, <iframe>, <img>, <input>, <object>, <video>
func Width(value uint64) vecty.Applyer {
	return vecty.Property("width", value)
}

type WrapCase = string

const (
	WrapCaseSoft WrapCase = "soft"
	WrapCaseHard WrapCase = "hard"
)

// Wrap specifies how the text in a text area is to be wrapped when submitted in a form.
//
// <textarea>
func Wrap(c WrapCase) vecty.Applyer {
	return vecty.Property("wrap", c)
}

// Vecty has an event package so this function does literally nothing
/*
func On(event string, rawJS string) vecty.Applyer {
	return vecty.Property("on"+event, rawJS)
}
*/
