// Code generated by templ@v0.2.364 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/shaninalex/homefilestorage/pkg/database"
)

type State struct {
	LoggedIn    bool
	CSRFToken   string
	Error       string
	ActiveRoute string
	Files       []database.File
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
		_, err = templBuffer.WriteString("<style>")
		if err != nil {
			return err
		}
		var_2 := `
        html,
        body {
            height: 100%;
        }

        body {
            display: flex;
            align-items: center;
            padding-top: 40px;
            padding-bottom: 40px;
            background-color: #f5f5f5;
        }

        .form-signin {
            width: 100%;
            max-width: 330px;
            padding: 15px;
            margin: auto;
        }

        .form-signin .checkbox {
            font-weight: 400;
        }

        .form-signin .form-floating:focus-within {
            z-index: 2;
        }

        .form-signin input[type="email"] {
            margin-bottom: -1px;
            border-bottom-right-radius: 0;
            border-bottom-left-radius: 0;
        }

        .form-signin input[type="password"] {
            margin-bottom: 10px;
            border-top-left-radius: 0;
            border-top-right-radius: 0;
        }

    `
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</style><main class=\"form-signin\"><form action=\"/login\" method=\"POST\"><input type=\"hidden\" name=\"authenticity_token\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(state.CSRFToken))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><h1 class=\"h3 mb-3 fw-normal\">")
		if err != nil {
			return err
		}
		var_3 := `Please sign in`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><div class=\"form-floating\"><input type=\"email\" name=\"email\" class=\"form-control\" id=\"floatingInput\" placeholder=\"Email\" value=\"email@email.com\"><label for=\"floatingInput\">")
		if err != nil {
			return err
		}
		var_4 := `Email address`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label></div><div class=\"form-floating\"><input type=\"password\" name=\"password\" class=\"form-control\" id=\"floatingPassword\" placeholder=\"Password\" value=\"password\"><label for=\"floatingPassword\">")
		if err != nil {
			return err
		}
		var_5 := `Password`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label></div><button class=\"w-100 btn btn-lg btn-primary\" type=\"submit\">")
		if err != nil {
			return err
		}
		var_6 := `Sign in`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button>")
		if err != nil {
			return err
		}
		if state.Error != "" {
			_, err = templBuffer.WriteString("<div class=\"alert alert-danger mt-4\" role=\"alert\">")
			if err != nil {
				return err
			}
			var var_7 string = state.Error
			_, err = templBuffer.WriteString(templ.EscapeString(var_7))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</form></main>")
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
		var_8 := templ.GetChildren(ctx)
		if var_8 == nil {
			var_8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><title>")
		if err != nil {
			return err
		}
		var_9 := `Home | HFS`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN\" crossorigin=\"anonymous\"></head><body>")
		if err != nil {
			return err
		}
		if state.LoggedIn {
			err = Dashboard(state).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		} else {
			err = Login(state).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
