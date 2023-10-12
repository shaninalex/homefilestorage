// Code generated by templ@v0.2.364 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type State struct {
	LoggedIn  bool
	CSRFToken string
	Error     string
}

func Login(state State) templ.Component {
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
		_, err = templBuffer.WriteString("<form action=\"/login\" method=\"POST\"><input type=\"hidden\" name=\"authenticity_token\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(state.CSRFToken))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><div class=\"mb-3\"><input type=\"email\" name=\"email\" placeholder=\"Email\" class=\"form-control form-control-sm\"></div><div class=\"mb-3\"><input type=\"password\" name=\"password\" placeholder=\"Password\" class=\"form-control form-control-sm\"></div><div><button type=\"submit\" class=\"btn btn-info btn-sm\">")
		if err != nil {
			return err
		}
		var_2 := `Login`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div>")
		if err != nil {
			return err
		}
		if state.Error != "" {
			_, err = templBuffer.WriteString("<div class=\"alert alert-danger mt-4\" role=\"alert\">")
			if err != nil {
				return err
			}
			var var_3 string = state.Error
			_, err = templBuffer.WriteString(templ.EscapeString(var_3))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

// TODO: Check if user logged in

func Home(state State) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>")
		if err != nil {
			return err
		}
		var_5 := `Home | HFS`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN\" crossorigin=\"anonymous\"></head><body><div class=\"container py-3\">")
		if err != nil {
			return err
		}
		if state.LoggedIn {
			_, err = templBuffer.WriteString("<div>")
			if err != nil {
				return err
			}
			var_6 := `Welcome back!`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div> <a href=\"/logout\">")
			if err != nil {
				return err
			}
			var_7 := `Logout`
			_, err = templBuffer.WriteString(var_7)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
		} else {
			err = Login(state).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL\" crossorigin=\"anonymous\">")
		if err != nil {
			return err
		}
		var_8 := ``
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
