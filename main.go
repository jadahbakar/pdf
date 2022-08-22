package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jadahbakar/pdf/data"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	// begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHHeading(m)
	buildFruitList(m)

	err := m.OutputFileAndClose("pdf/sample.pdf")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
		os.Exit(1)
	}
	fmt.Println("PDF file saved successfully.")
}

func buildHHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/1656673786690.jpeg", props.Rect{
					Center:  true,
					Percent: 75,
				})
				if err != nil {
					log.Printf("error loading image: %v", err)
				}
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func buildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	// contents := [][]string{{"Apple", "Red & Juicy", "2.00"}, {"Orange", "Orange and Juicy", "3.00"}}
	contents := data.FruitList(20)
	lightPurpleColor := getLightPurpleColor()
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
	})

}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}
