{{define "base"}}
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{block "title" .}}{{end}}</title>
    {{block "style" .}}{{end}}
</head>
<body>
<nav class="navbar navbar-toggleable-md navbar-light">
    <div class="container">
        <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse"
                data-target="#app-navbar-collapse" aria-controls="app-navbar-collapse" aria-expanded="false"
                aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <a class="navbar-brand" href="/template/home.tpl">
            <img src="/static/img/logo/dole-logo.png" height="30" class="d-inline-block align-top" alt="Dole Logo">&nbsp;Applicant
            Lookup
        </a>
        <div class="collapse navbar-collapse" id="app-navbar-collapse">
            <form class="form-inline">
                <input class="form-control" type="text" placeholder="Search">
            </form>
            <ul class="navbar-nav">
                <li class="nav-item"><a href="/template/auth/login.tpl" class="nav-link">Login</a></li>
                <li class="nav-item"><a href="/template/auth/register.tpl" class="nav-link">Register</a></li>
            </ul>
        </div>
    </div>
</nav>
{{template "content" .}}
{{block "script" .}}{{end}}
</body>
</html>
{{end}}