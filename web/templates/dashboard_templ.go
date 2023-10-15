// Code generated by templ@v0.2.364 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"strconv"
)

func Dashboard(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		err = Header(state).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"container\">")
		if err != nil {
			return err
		}
		if state.ActiveRoute == "files" {
			err = FilesComponent(state).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL\" crossorigin=\"anonymous\">")
		if err != nil {
			return err
		}
		var_2 := ``
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Header(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<nav class=\"navbar navbar-expand-lg navbar-dark bg-dark mb-3\"><div class=\"container\"><a class=\"navbar-brand\" href=\"#\">")
		if err != nil {
			return err
		}
		var_4 := `HFS`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button><div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\"><ul class=\"navbar-nav me-auto\"><li class=\"nav-item\"><a class=\"nav-link active\" aria-current=\"page\" href=\"/\">")
		if err != nil {
			return err
		}
		var_5 := `Files`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul><ul class=\"navbar-nav\"><li class=\"nav-item\"><a class=\"nav-link active\" aria-current=\"page\" href=\"/logout\">")
		if err != nil {
			return err
		}
		var_6 := `Logout`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div></div></nav>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FilesComponent(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_7 := templ.GetChildren(ctx)
		if var_7 == nil {
			var_7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		err = FilesFormUpload(state).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = FilesList(state).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FilesFormUpload(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form action=\"/files/upload\" enctype=\"multipart/form-data\" method=\"post\" class=\"mb-3 d-flex justify-content-between align-items-center gap-3\"><input type=\"hidden\" name=\"authenticity_token\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(state.CSRFToken))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><input name=\"files\" class=\"form-control\" type=\"file\" multiple required><button type=\"submit\" class=\"btn btn-primary\">")
		if err != nil {
			return err
		}
		var_9 := `Upload`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></form>")
		if err != nil {
			return err
		}
		if state.Error != "" {
			_, err = templBuffer.WriteString("<div class=\"alert alert-danger mt-4\" role=\"alert\">")
			if err != nil {
				return err
			}
			var var_10 string = state.Error
			_, err = templBuffer.WriteString(templ.EscapeString(var_10))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FilesList(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"row\">")
		if err != nil {
			return err
		}
		for _, file := range state.Files {
			err = FileItem(file, state).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FileItem(file database.File, state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_12 := templ.GetChildren(ctx)
		if var_12 == nil {
			var_12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"col-md-4\"><div class=\"card mb-3\"><div class=\"card-header text-body-secondary\"><a href=\"")
		if err != nil {
			return err
		}
		var var_13 templ.SafeURL = templ.URL(file.PreviewPath())
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_13)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var var_14 string = file.Name
		_, err = templBuffer.WriteString(templ.EscapeString(var_14))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a> ")
		if err != nil {
			return err
		}
		var_15 := `(`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<b>")
		if err != nil {
			return err
		}
		var_16 := `ID: `
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		var var_17 string = strconv.FormatInt(file.ID, 10)
		_, err = templBuffer.WriteString(templ.EscapeString(var_17))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</b>")
		if err != nil {
			return err
		}
		var_18 := `)`
		_, err = templBuffer.WriteString(var_18)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"card-body\">")
		if err != nil {
			return err
		}
		var_19 := `SIZE: `
		_, err = templBuffer.WriteString(var_19)
		if err != nil {
			return err
		}
		var var_20 string = file.FormatSize()
		_, err = templBuffer.WriteString(templ.EscapeString(var_20))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" <br> ")
		if err != nil {
			return err
		}
		var_21 := `Created at: `
		_, err = templBuffer.WriteString(var_21)
		if err != nil {
			return err
		}
		var var_22 string = file.FormatTime()
		_, err = templBuffer.WriteString(templ.EscapeString(var_22))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</div><div class=\"card-footer text-body-secondary\"><div class=\"d-flex gap-3\"><a class=\"btn btn-primary btn-sm btn-default\" href=\"")
		if err != nil {
			return err
		}
		var var_23 templ.SafeURL = templ.URL(file.DownloadPath())
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_23)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" class=\"bi bi-download\" viewBox=\"0 0 16 16\"><path d=\"M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5z\"></path><path d=\"M7.646 11.854a.5.5 0 0 0 .708 0l3-3a.5.5 0 0 0-.708-.708L8.5 10.293V1.5a.5.5 0 0 0-1 0v8.793L5.354 8.146a.5.5 0 1 0-.708.708l3 3z\"></path></svg></a><form action=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(fmt.Sprintf("/files/%d/delete", file.ID)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" method=\"POST\"><input type=\"hidden\" name=\"authenticity_token\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(state.CSRFToken))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><button type=\"submit\" class=\"btn btn-primary btn-sm btn-danger\">")
		if err != nil {
			return err
		}
		var_24 := `X`
		_, err = templBuffer.WriteString(var_24)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></form></div></div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}