package pdf

import (
	"fmt"
	"lgo/test/unipdf/contentstream"
	"lgo/test/unipdf/core"
	"lgo/test/unipdf/model"
)

var colorspaces = map[string]int{}
var filters = map[string]int{}

// List images and properties of a PDF specified by inputPath.
func ListImages(inputPath string) error {
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	if err != nil {
		return err
	}
	defer f.Close()

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}
	fmt.Printf("PDF Num Pages: %d\n", numPages)

	for i := 0; i < numPages; i++ {
		fmt.Printf("-----\nPage %d:\n", i+1)

		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return err
		}

		// List images on the page.
		err = listImagesOnPage(page)
		if err != nil {
			return err
		}
		break
	}
	fmt.Println(filters)
	fmt.Println(colorspaces)
	return nil
}

func listImagesOnPage(page *model.PdfPage) error {
	contents, err := page.GetAllContentStreams()
	if err != nil {
		return err
	}

	return listImagesInContentStream(contents, page.Resources)
}
func listImagesInContentStream(contents string, resources *model.PdfPageResources) error {
	cstreamParser := contentstream.NewContentStreamParser(contents)
	operations, err := cstreamParser.Parse()
	if err != nil {
		return err
	}

	processedXObjects := map[string]bool{}

	for _, op := range *operations {
		// if op.Operand == "BI" && len(op.Params) == 1 {
		if op.Operand == "BI" {
			// Inline image.

			iimg, ok := op.Params[0].(*contentstream.ContentStreamInlineImage)
			if !ok {
				continue
			}

			img, err := iimg.ToImage(resources)
			if err != nil {
				return err
			}

			cs, err := iimg.GetColorSpace(resources)
			if err != nil {
				return err
			}

			encoder, err := iimg.GetEncoder()
			if err != nil {
				return err
			}
			fmt.Println("BI")
			fmt.Printf(" Inline image\n")
			fmt.Printf("  Filter: %s\n", encoder.GetFilterName())
			fmt.Printf("  Width: %d\n", img.Width)
			fmt.Printf("  Height: %d\n", img.Height)
			fmt.Printf("  Color components: %d\n", img.ColorComponents)
			fmt.Printf("  ColorSpace: %s\n", cs.String())
			//fmt.Printf("  ColorSpace: %+v\n", cs)
			fmt.Printf("  BPC: %d\n", img.BitsPerComponent)

			// Log filter use globally.
			filter := encoder.GetFilterName()
			if _, has := filters[filter]; has {
				filters[filter]++
			} else {
				filters[filter] = 1
			}
			// Log colorspace use globally.
			csName := "?"
			if cs != nil {
				csName = cs.String()
			}
			if _, has := colorspaces[csName]; has {
				colorspaces[csName]++
			} else {
				colorspaces[csName] = 1
			}
		// } else if op.Operand == "Do" && len(op.Params) == 1 {
		} else if op.Operand == "Do"  {
			// XObject.
			name := op.Params[0].(*core.PdfObjectName)

			// Only process each one once.
			_, has := processedXObjects[string(*name)]
			if has {
				continue
			}
			processedXObjects[string(*name)] = true

			_, xtype := resources.GetXObjectByName(*name)
			if xtype == model.XObjectTypeImage {
				fmt.Printf(" XObject Image: %s\n", *name)

				ximg, err := resources.GetXObjectImageByName(*name)
				if err != nil {
					return err
				}
				img, err := ximg.ToImage()
				if err != nil {
					return err
				}
				fmt.Println("DO")
				fmt.Printf("  Filter: %#v\n", ximg.Filter)
				fmt.Printf("  Width: %v\n", *ximg.Width)
				fmt.Printf("  Height: %d\n", *ximg.Height)
				fmt.Printf("  Color components: %d\n", img.ColorComponents)
				fmt.Printf("  ColorSpace: %s\n", ximg.ColorSpace.String())
				fmt.Printf("  ColorSpace: %#v\n", ximg.ColorSpace)
				fmt.Printf("  BPC: %v\n", *ximg.BitsPerComponent)

				// Log filter use globally.
				filter := ximg.Filter.GetFilterName()
				if _, has := filters[filter]; has {
					filters[filter]++
				} else {
					filters[filter] = 1
				}
				// Log colorspace use globally.
				cs := ximg.ColorSpace.String()
				if _, has := colorspaces[cs]; has {
					colorspaces[cs]++
				} else {
					colorspaces[cs] = 1
				}
			} else if xtype == model.XObjectTypeForm {
				// Go through the XObject Form content stream.
				fmt.Printf("--> XObject Form: %s\n", *name)
				xform, err := resources.GetXObjectFormByName(*name)
				if err != nil {
					return err
				}

				formContent, err := xform.GetContentStream()
				if err != nil {
					return err
				}
				fmt.Printf("xform: %#v\n", xform)
				fmt.Printf("xform res: %#v\n", xform.Resources)
				fmt.Printf("Content: %s\n", formContent)

				// Process the content stream in the Form object too:
				// XXX/TODO: Use either form resources (priority) and fall back to page resources alternatively if not found.
				if xform.Resources != nil {
					err = listImagesInContentStream(string(formContent), xform.Resources)
				} else {
					err = listImagesInContentStream(string(formContent), resources)
				}
				if err != nil {
					return err
				}
				fmt.Printf("<-- XObject Form: %s\n", *name)
			}
		}
	}

	return nil
}
