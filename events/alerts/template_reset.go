package alerts

import (
	"html/template"
)

var resetTemplate = template.Must(template.New("index").Parse(`
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>Vertice Reset password</title>
</head>
<div style="font-family: sans-serif; font-size: 14px; color: #8a9ca2;">
  <body bgcolor="#eaf0f2">
    <table cellspacing="0" cellpadding="0" border="0" width="100%">
      <tr>
        <td bgcolor="#eaf0f2" align="center">
          <table width="650px" cellspacing="0" cellpadding="3" class="container">
            <tr>
              <td align="center">
                <img style="margin: 30px;" src="{{.logo}}">
              </td>
            </tr>
            <tr>
              <td align="center">
                <h3 style="font-size: 28px; font-weight: normal;">Hey <b>{{.email}}</b>,
                  <br>Reset your password</h3>
                <table width="500px" cellspacing="0" cellpadding="3">
                  <tr style=" background-color: #fff;">
                    <td align="center" width="30%">
                      <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/lock.png">
                    </td>
                    <td align="left" style="padding: 30px;">
                      <h2 style="margin: 30px 0px;">Forgot password !</h2>
                      <p>You have requested to reset your password, in case you didn't, ignore this email </p>
                    </td>
                  </tr>
                </table>
                <table width="500px" cellspacing="0" cellpadding="3" style=" background-color: #fff;">
                  <tr>
                    <td>
                      <hr>
                    </td>
                  </tr>
                  <tr>
                    <td align="center" style="padding: 45px;">
                      <a href="{{.reset_url}}">
                        <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/reset_password.png">
                      </a>
                    </td>
                  </tr>
                </table>
              </td>


            </tr>
            <tr align="center">
              <td width="28px" style="padding: 30px;">
                <a href="https://www.github.com/megamsys">
                  <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/git.png">
                </a>
                <a href="https://plus.google.com/105742854503471299978/about">
                  <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/plus.png">
                </a>
                <a href="https://facebook.com/megamsys">
                  <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/fb.png">
                </a>
                <a href="https://twitter.com/megamsys">
                  <img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/twitter.png">
                </a>
              </td>

            </tr>
            <tr align="center">
              <td style="padding: 30px;">
                <p>© Megam 2016 All rights reserved.</p>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</div>
</html>
`))
