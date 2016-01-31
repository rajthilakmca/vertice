package alerts

import (
	"html/template"
)

var balanceTemplate = template.Must(template.New("balance").Parse(`
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Vertice Balance</title>
    <style type="text/css">
        #email .container{
          background-color: #eaf0f2;
        }
        .box1{
            background-color: #fff;
        }
        body{
            font-family: sans-serif;
            font-size: 14px;
            color: #8a9ca2;
        }
        h3{font-size: 24px; font-weight: normal;}
        .bold{font-weight: bold;}
        .logo{
            margin: 30px;
        }
        .f36{font-size: 36px; color: #66719b;}
        .balance{margin: 30px 0px;}
        .pad_b30{padding-bottom: 30px;}
        p{margin: 0px;}
        .pad_45{padding: 45px;}
        .pad_30{padding: 30px;}
    </style>
    <link href='http://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>
</head>
<body id="email">
    <table cellspacing="0" cellpadding="0" border="0" width="100%">
        <tr>
            <td bgcolor="#FFFFFF" align="center">
                <table width="650px" cellspacing="0" cellpadding="3" class="container">
                    <tr>
                        <td align="center">
                            <img class="logo" src="{{.logo}}">
                        </td>
                    </tr>
                    <tr>
                        <td align="center">
                            <h3>Hi <span class="bold">User</span>,<br>Just a friendly reminder on your usage !</h3>

                            <table width="500px" cellspacing="0" cellpadding="3" class="box1">
                                <tr>
                                    <td align="center" width="30%"><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/balance.png"></td>
                                    <td align="left" class="pad_b30">
                                        <h2 class="balance">Balance !</h2>
                                        <p>Here is what left in your account </p>
                                        <p><span class="f36">{{.days}}</span> days <span class="f36">{{.cost}}</span></p>
                                    </td>
                                </tr>
                            </table>
                            <table width="500px" cellspacing="0" cellpadding="3" class="box1">
                                <tr>
                                <td><hr></td>
                                </tr>
                                <tr>
                                    <td align="center" class="pad_45">
                                        <a href=""><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/view_bill.png"></a>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    <tr align="center">
                        <td  width="28px" class="pad_30">
                            <a href="https://www.github.com/megamsys"><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/git.png"></a>
                            <a href="https://plus.google.com/105742854503471299978/about"><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/plus.png"></a>
                            <a href="https://facebook.com/megamsys"><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/fb.png"></a>
                            <a href="https://twitter.com/megamsys"><img src="https://s3-ap-southeast-1.amazonaws.com/megampub/images/mailers/twitter.png"></a>
                        </td>
                    </tr>
                    <tr align="center">
                        <td class="pad_b30">
                            <p>© Megam 2016 All rights reserved.</p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>
`))
