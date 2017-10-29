package bbcodehtml

import (
	"strings"

	"github.com/MJKWoolnough/bbcode"
)

type table struct{}

func (table) Name() string {
	return "table"
}

var (
	tableOpen      = []byte("<table>")
	tableClose     = []byte("</table>")
	tableHeadOpen  = []byte("<thead>")
	tableHeadClose = []byte("</thead>")
	tableFootOpen  = []byte("<tfoot>")
	tableFootClose = []byte("</tfoot>")
	trOpen         = []byte("<tr>")
	trClose        = []byte("</tr>")
	tdOpen         = []byte("<td>")
	tdClose        = []byte("</td>")
	thOpen         = []byte("<th>")
	thClose        = []byte("</th>")
)

func (table) Handle(p *bbcode.Processor, attr string) {
	p.Write(tableOpen)
	tableHandle(p)
	p.Write(tableClose)
}

func tableHandle(p *bbcode.Processor) {
	var thead, tfoot bool
Loop:
	for {
		switch t := p.Get().(type) {
		case bbcode.Text:
		case bbcode.OpenTag:
			if strings.EqualFold(t.Name, "thead") {
				if !thead {
					p.Write(tableHeadOpen)
					tableIHandle(p, "thead")
					p.Write(tableHeadClose)
					thead = true
				}
			} else if strings.EqualFold(t.Name, "tfoot") {
				if !tfoot {
					p.Write(tableFootOpen)
					tableIHandle(p, "tfoot")
					p.Write(tableFootClose)
					tfoot = true
				}
			} else if strings.EqualFold(t.Name, "row") || strings.EqualFold(t.Name, "tr") {
				p.Write(trOpen)
				tableRow(p, t.Name)
				p.Write(trClose)
			}
		case bbcode.CloseTag:
			if strings.EqualFold(t.Name, "table") {
				break Loop
			}
		default:
			break Loop
		}
	}
}

func tableIHandle(p *bbcode.Processor, tagName string) {
Loop:
	for {
		switch t := p.Get().(type) {
		case bbcode.Text:
		case bbcode.OpenTag:
			if strings.EqualFold(t.Name, "row") || strings.EqualFold(t.Name, "tr") {
				p.Write(trOpen)
				tableRow(p, t.Name)
				p.Write(trClose)
			}
		case bbcode.CloseTag:
			if t.Name == tagName {
				break Loop
			}
		default:
			break Loop
		}
	}
}

func tableRow(p *bbcode.Processor, tagName string) {
Loop:
	for {
		switch t := p.Get().(type) {
		case bbcode.Text:
		case bbcode.OpenTag:
			if strings.EqualFold(t.Name, "col") || strings.EqualFold(t.Name, "td") {
				p.Write(tdOpen)
				p.Process(t.Name)
				p.Write(tdClose)
			} else if strings.EqualFold(t.Name, "th") {
				p.Write(thOpen)
				p.Process(t.Name)
				p.Write(thClose)
			}
		case bbcode.CloseTag:
			if t.Name == tagName {
				break Loop
			}
		default:
			break Loop
		}
	}
}
